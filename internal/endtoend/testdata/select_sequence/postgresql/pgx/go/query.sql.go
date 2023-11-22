// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: query.sql

package querytest

import (
	"context"
)

const getLastValue = `-- name: GetLastValue :one
SELECT last_value FROM my_sequence
`

func (q *Queries) GetLastValue(ctx context.Context) (int64, error) {
	row := q.db.QueryRow(ctx, getLastValue)
	var last_value int64
	err := row.Scan(&last_value)
	return last_value, err
}
