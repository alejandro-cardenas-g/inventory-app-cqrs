package outbox

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	id            int64
	eventID       uuid.UUID
	eventType     string
	aggregateType string
	aggregateID   int64
	payload       []byte
	occurredAt    time.Time
	processedAt   *time.Time
	retryCount    int
	correlationID *uuid.UUID
	causationID   *uuid.UUID
}

func New(
	eventType string,
	aggregateType string,
	aggregateID int64,
	payload []byte,
	correlationID *uuid.UUID,
	causationID *uuid.UUID,
) *Event {
	eventId := uuid.New()

	return &Event{
		eventID:       eventId,
		eventType:     eventType,
		aggregateType: aggregateType,
		aggregateID:   aggregateID,
		payload:       payload,
		correlationID: correlationID,
		causationID:   causationID,
	}
}

func (e *Event) SetID(id int64) {
	e.id = id
}

func (e *Event) MarkProcessed(t time.Time) {
	e.processedAt = &t
}

func (e *Event) IncrementRetry() {
	e.retryCount++
}

func (e *Event) GetID() int64 {
	return e.id
}

func (e *Event) GetEventID() uuid.UUID {
	return e.eventID
}

func (e *Event) GetEventType() string {
	return e.eventType
}

func (e *Event) GetAggregateType() string {
	return e.aggregateType
}

func (e *Event) GetAggregateID() int64 {
	return e.aggregateID
}

func (e *Event) GetPayload() []byte {
	return e.payload
}

func (e *Event) GetOccurredAt() time.Time {
	return e.occurredAt
}

func (e *Event) GetProcessedAt() *time.Time {
	return e.processedAt
}

func (e *Event) GetRetryCount() int {
	return e.retryCount
}

func (e *Event) GetCorrelationID() *uuid.UUID {
	return e.correlationID
}

func (e *Event) GetCausationID() *uuid.UUID {
	return e.causationID
}
