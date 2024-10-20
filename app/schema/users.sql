-- name: GetUser :one
SELECT *
FROM users
WHERE id = $1;

-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE email = $1;

-- name: InsertUserIfNotExists :exec
INSERT INTO users (id, email)
VALUES ($1, $2)
ON CONFLICT(email) DO NOTHING;