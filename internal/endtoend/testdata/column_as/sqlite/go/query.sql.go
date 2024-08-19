// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package querytest

import (
	"context"
)

const withAs = `-- name: WithAs :one
SELECT 1 AS x, 2 AS y
`

type WithAsRow struct {
	X int64
	Y int64
}

func (q *Queries) WithAs(ctx context.Context) (WithAsRow, error) {
	row := q.db.QueryRowContext(ctx, withAs)
	var i WithAsRow
	err := row.Scan(&i.X, &i.Y)
	return i, err
}

const withoutAs = `-- name: WithoutAs :one
SELECT 1 x, 2 y
`

type WithoutAsRow struct {
	X int64
	Y int64
}

func (q *Queries) WithoutAs(ctx context.Context) (WithoutAsRow, error) {
	row := q.db.QueryRowContext(ctx, withoutAs)
	var i WithoutAsRow
	err := row.Scan(&i.X, &i.Y)
	return i, err
}
