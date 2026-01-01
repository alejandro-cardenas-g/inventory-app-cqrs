package outbox

import (
	"fmt"
	bs "inventory_cqrs/internal/bootstrap"
	"inventory_cqrs/internal/config"
	"inventory_cqrs/internal/domain/outbox"
	"sync"
	"time"

	"go.uber.org/zap"
)

type OutboxWorker struct {
	container    *bs.WorkerContainer
	logger       *zap.Logger
	cfg          config.WorkerConfig
	processingCh chan *outbox.Event
	wg           *sync.WaitGroup
}

func NewOutboxWorker(
	container *bs.WorkerContainer,
	logger *zap.Logger,
	cfg config.WorkerConfig) *OutboxWorker {
	return &OutboxWorker{container: container, logger: logger, cfg: cfg, processingCh: make(chan *outbox.Event, cfg.BatchSize), wg: &sync.WaitGroup{}}
}

func (w *OutboxWorker) Run() {
	var interval time.Duration = w.cfg.PollInterval
	var maxInterval time.Duration = w.cfg.PollInterval * 8

	var count int = 0

	for {
		fmt.Println("Polling for events", interval)

		processedEvents, _ := w.container.DispatcherService.ManageEvents(w.wg)

		if processedEvents < 1 {
			count++
		} else {
			interval = w.cfg.PollInterval
		}

		time.Sleep(interval)

		if count > 10 {
			count = 0
			interval = interval * 2
			if interval > maxInterval {
				interval = maxInterval
			}
		}
	}
}

func (w *OutboxWorker) Shutdown() {
	w.logger.Info("Shutting down outbox worker")
	w.wg.Wait()
	close(w.processingCh)
}
