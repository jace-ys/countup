-- name: GetCounter :one
SELECT *
FROM counter
WHERE id = 1;

-- name: IncrementCounter :exec
UPDATE counter
SET
  count = count + 1,
  last_increment_by = $1,
  last_increment_at = $2,
  next_finalize_at = NULL
WHERE id = 1;

-- name: ResetCounter :exec
UPDATE counter
SET
  count = 0,
  last_increment_by = NULL,
  last_increment_at = NULL,
  next_finalize_at = NULL
WHERE id = 1;

-- name: UpdateCounterFinalizeTime :exec
UPDATE counter
SET
  next_finalize_at = $1
WHERE id = 1;

-- name: ListIncrementRequests :many
SELECT *
FROM increment_requests;

-- name: InsertIncrementRequest :one
WITH inserted AS (
  INSERT INTO increment_requests (requested_by, requested_at)
  VALUES ($1, $2)
  RETURNING *
)
SELECT COUNT(*) AS num_requests
FROM increment_requests;

-- name: DeleteIncrementRequests :exec
TRUNCATE TABLE increment_requests;