// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: placeholder.sql

package queries

import (
	"context"
)

const getOne = `-- name: GetOne :one
SELECT 1
`

func (q *Queries) GetOne(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, getOne)
	var column_1 int64
	err := row.Scan(&column_1)
	return column_1, err
}
