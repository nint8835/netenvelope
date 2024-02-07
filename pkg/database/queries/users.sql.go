// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: users.sql

package queries

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (username, password_hash)
VALUES (?, ?)
RETURNING id, username, password_hash
`

type CreateUserParams struct {
	Username     string
	PasswordHash []byte
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Username, arg.PasswordHash)
	var i User
	err := row.Scan(&i.ID, &i.Username, &i.PasswordHash)
	return i, err
}
