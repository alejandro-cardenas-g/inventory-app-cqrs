package publisher

import (
	"context"
	"inventory_cqrs/internal/config"
	"inventory_cqrs/internal/constants"
	"inventory_cqrs/internal/domain/outbox"
	"inventory_cqrs/internal/store/persistence"
	"sync"

	"go.uber.org/zap"
)

type DispatcherService struct {
	publishersMap map[constants.Event]IPublisher
	txManager     *persistence.TxManager
	obRepository  *persistence.OutboxRepository
	logger        *zap.Logger
	cfg           config.WorkerConfig
}

func NewDispatcherService(publishersMap map[constants.Event]IPublisher, txManager *persistence.TxManager, obRepository *persistence.OutboxRepository, logger *zap.Logger, cfg config.WorkerConfig) *DispatcherService {
	return &DispatcherService{publishersMap: publishersMap, txManager: txManager, obRepository: obRepository, logger: logger, cfg: cfg}
}

func (d *DispatcherService) ManageEvents(wg *sync.WaitGroup) (int, error) {

	var events []*outbox.Event = []*outbox.Event{}

	err := d.txManager.WithTx(context.Background(), func(tx persistence.DBTX) error {
		var err error
		ob := d.obRepository.UseTX(tx)

		events, err = ob.GetUnprocessed(context.Background(), d.cfg.BatchSize)

		if err != nil {
			return err
		}

		if len(events) == 0 {
			return nil
		}

		eventIDs := make([]int64, len(events))
		for i, event := range events {
			eventIDs[i] = event.GetID()
		}
		err = ob.MarkEventsAsProcessing(context.Background(), eventIDs)

		if err != nil {
			return err
		}

		return nil
	})

	for _, event := range events {
		wg.Add(1)
		go func() {
			defer wg.Done()
			d.dispatch(event)
		}()
	}

	d.logger.Info("Processing events")

	wg.Wait()

	return len(events), err
}

func (d *DispatcherService) dispatch(event *outbox.Event) {
	publisher, ok := d.publishersMap[constants.Event(event.GetEventType())]

	if !ok {
		d.logger.Error("publisher not found for event type", zap.String("event_type", event.GetEventType()))
		d.obRepository.MarkProcessed(context.Background(), event.GetID())
		return
	}

	err := publisher.Publish(event)
	if err != nil {
		if event.GetRetryCount()+1 > d.cfg.MaxRetries {
			d.logger.Error("max retries reached for event", zap.String("event_type", event.GetEventType()))
			d.obRepository.MarkProcessed(context.Background(), event.GetID())
		} else {
			d.obRepository.IncrementRetryCount(context.Background(), event.GetID())
		}
		return
	}

	d.obRepository.MarkProcessed(context.Background(), event.GetID())
}
