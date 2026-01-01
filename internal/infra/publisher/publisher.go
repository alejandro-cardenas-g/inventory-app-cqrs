package publisher

import (
	"inventory_cqrs/internal/domain/outbox"
)

type IPublisher interface {
	Publish(event *outbox.Event) error
}
