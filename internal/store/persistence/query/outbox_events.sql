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
    causation_id
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10);

-- name: GetUnprocessedOutboxEvents :many
SELECT * FROM outbox_events WHERE processed_at IS NULL LIMIT $1;

-- name: MarkProcessedOutboxEvent :exec
UPDATE outbox_events SET processed_at = $1 WHERE id = $2;