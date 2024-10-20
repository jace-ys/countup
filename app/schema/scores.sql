-- name: ListScores :many
SELECT *
FROM scores;

-- name: InsertScore :exec
INSERT INTO scores (user_email, score)
VALUES ($1, $2);