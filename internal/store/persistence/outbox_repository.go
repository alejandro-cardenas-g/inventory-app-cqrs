package persistence

import (
	"context"
	"inventory_cqrs/internal/domain/outbox"
	db "inventory_cqrs/internal/store/persistence/sqlc"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type OutboxRepository struct {
	queries *db.Queries
}

func NewOutboxRepository(tx DBTX) *OutboxRepository {
	return &OutboxRepository{queries: db.New(tx)}
}

func (r *OutboxRepository) UseTX(tx DBTX) *OutboxRepository {
	return &OutboxRepository{queries: db.New(tx)}
}

func (r *OutboxRepository) Save(ctx context.Context, e *outbox.Event) error {

	var eventID pgtype.UUID

	eBytes, err := e.GetEventID().MarshalBinary()

	if err != nil {
		return ErrorPerformingOperation
	}

	eventID = pgtype.UUID{Bytes: [16]byte(eBytes), Valid: true}

	params := db.SaveOutboxEventParams{
		EventID:       eventID,
		EventType:     e.GetEventType(),
		AggregateType: e.GetAggregateType(),
		AggregateID:   e.GetAggregateID(),
		Payload:       e.GetPayload(),
		OccurredAt:    pgtype.Timestamptz{Time: e.GetOccurredAt(), Valid: true},
	}

	if e.GetCorrelationID() != nil {
		cBytes, err := e.GetCorrelationID().MarshalBinary()
		if err == nil {
			params.CorrelationID = pgtype.UUID{Bytes: [16]byte(cBytes), Valid: true}
		}
	}

	if e.GetCausationID() != nil {
		cBytes, err := e.GetCausationID().MarshalBinary()
		if err == nil {
			params.CausationID = pgtype.UUID{Bytes: [16]byte(cBytes), Valid: true}
		}
	}

	err = r.queries.SaveOutboxEvent(ctx, params)

	if err != nil {
		return ErrorPerformingOperation
	}

	return nil
}

func (r *OutboxRepository) GetUnprocessed(ctx context.Context, limit int) ([]*outbox.Event, error) {
	eventsDB, err := r.queries.GetUnprocessedOutboxEvents(ctx, int32(limit))

	if err != nil {
		return nil, err
	}

	events := make([]*outbox.Event, len(eventsDB))

	for i, e := range eventsDB {

		var correlationID *uuid.UUID
		var causationID *uuid.UUID

		if v, err := uuid.Parse(e.CorrelationID.String()); err == nil {
			correlationID = &v
		}

		if v, err := uuid.Parse(e.CausationID.String()); err == nil {
			causationID = &v
		}

		events[i] = outbox.New(
			e.EventType,
			e.AggregateType,
			e.AggregateID,
			e.Payload,
			correlationID,
			causationID,
		)

		events[i].SetID(e.ID)
	}
	return events, nil
}

func (r *OutboxRepository) MarkProcessed(ctx context.Context, id int64) error {
	return r.queries.MarkProcessedOutboxEvent(ctx, id)
}

func (r *OutboxRepository) MarkEventsAsProcessing(ctx context.Context, ids []int64) error {
	return r.queries.MarkEventsAsProcessing(ctx, ids)
}

func (r *OutboxRepository) IncrementRetryCount(ctx context.Context, id int64) error {
	return r.queries.IncrementRetryCount(ctx, id)
}
