// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: users.sql

package userstore

import (
	"context"
)

const getUser = `-- name: GetUser :one
SELECT id, email
FROM users
WHERE id = $1
`

func (q *Queries) GetUser(ctx context.Context, db DBTX, id string) (User, error) {
	row := db.QueryRow(ctx, getUser, id)
	var i User
	err := row.Scan(&i.ID, &i.Email)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, email
FROM users
WHERE email = $1
`

func (q *Queries) GetUserByEmail(ctx context.Context, db DBTX, email string) (User, error) {
	row := db.QueryRow(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(&i.ID, &i.Email)
	return i, err
}

const insertUserIfNotExists = `-- name: InsertUserIfNotExists :exec
INSERT INTO users (id, email)
VALUES ($1, $2)
ON CONFLICT(email) DO NOTHING
`

type InsertUserIfNotExistsParams struct {
	ID    string
	Email string
}

func (q *Queries) InsertUserIfNotExists(ctx context.Context, db DBTX, arg InsertUserIfNotExistsParams) error {
	_, err := db.Exec(ctx, insertUserIfNotExists, arg.ID, arg.Email)
	return err
}