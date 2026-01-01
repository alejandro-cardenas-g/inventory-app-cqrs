-- name: SaveOutboxEvent :exec
INSERT INTO outbox_events (
    event_id,
    event_type,
    aggregate_type,
    aggregate_id,
    payload,
    occurred_at,
    processed_at,
    retry_count,
    correlation_id,
    causation_id,
    status
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, 'pending');

-- name: GetUnprocessedOutboxEvents :many
SELECT * FROM outbox_events WHERE status = 'pending' LIMIT $1 FOR UPDATE SKIP LOCKED;

-- name: MarkEventsAsProcessing :exec
UPDATE outbox_events SET status = 'processing' WHERE id = ANY($1::bigint[]);

-- name: MarkProcessedOutboxEvent :exec
UPDATE outbox_events SET status = 'done', processed_at = now() WHERE id = $1;

-- name: IncrementRetryCount :exec
UPDATE outbox_events SET retry_count = retry_count + 1 WHERE id = $1 and status = 'pending';