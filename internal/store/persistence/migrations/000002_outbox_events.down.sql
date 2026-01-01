DROP INDEX IF EXISTS idx_outbox_aggregate;
DROP INDEX IF EXISTS idx_outbox_unprocessed;
DROP INDEX IF EXISTS uq_outbox_event_id;

DROP TABLE IF EXISTS outbox_events;