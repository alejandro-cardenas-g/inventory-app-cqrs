CREATE TABLE outbox_events (
    id BIGSERIAL PRIMARY KEY,

    event_id UUID NOT NULL,
    event_type VARCHAR(100) NOT NULL,

    aggregate_type VARCHAR(50) NOT NULL,
    aggregate_id BIGINT NOT NULL,

    payload JSONB NOT NULL,

    occurred_at TIMESTAMPTZ NOT NULL DEFAULT now(),

    processed_at TIMESTAMPTZ NULL,
    retry_count INT NOT NULL DEFAULT 0,

    correlation_id UUID NULL,
    causation_id UUID NULL
);

CREATE UNIQUE INDEX uq_outbox_event_id
    ON outbox_events(event_id);

CREATE INDEX idx_outbox_unprocessed
    ON outbox_events(processed_at)
    WHERE processed_at IS NULL;

CREATE INDEX idx_outbox_aggregate
    ON outbox_events(aggregate_type, aggregate_id);