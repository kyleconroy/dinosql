// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: query.sql

package querytest

import (
	"context"
)

const barExists = `-- name: BarExists :one
SELECT
    EXISTS (
        SELECT
            1
        FROM
            bar
        where
            id = $1
    )
`

func (q *Queries) BarExists(ctx context.Context, id int32, aq ...AdditionalQuery) (bool, error) {
	query := barExists
	queryParams := []interface{}{id}
	row := q.db.QueryRow(ctx, query, queryParams...)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}
