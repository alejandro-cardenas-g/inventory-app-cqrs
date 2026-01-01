package elastic

import (
	"fmt"
	"inventory_cqrs/internal/domain/outbox"
	"inventory_cqrs/internal/infra/publisher"
	"inventory_cqrs/internal/store/persistence"
	"time"
)

type CreateProductPublisher struct {
	readModelRepository *persistence.ProductViewRepository
}

func NewCreateProductPublisher(readModelRepository *persistence.ProductViewRepository) publisher.IPublisher {
	return &CreateProductPublisher{}
}

func (p *CreateProductPublisher) Publish(event *outbox.Event) error {
	fmt.Println("handling event", event.GetEventType(), event.GetEventID())
	time.Sleep(10 * time.Second)
	return nil
}
