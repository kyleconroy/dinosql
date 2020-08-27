// Code generated by sqlc. DO NOT EDIT.
// source: query.sql

package querytest

import (
	"context"
	"database/sql"
)

const doubleDash = `-- name: DoubleDash :one
SELECT bar FROM foo WHERE bar = $1
`

func (q *Queries) DoubleDash(ctx context.Context, bar sql.NullString) (sql.NullString, error) {
	row := q.db.QueryRowContext(ctx, doubleDash, bar)
	var bar sql.NullString
	err := row.Scan(&bar)
	return bar, err
}

const slashStar = `-- name: SlashStar :one
SELECT bar FROM foo WHERE bar = $1
`

func (q *Queries) SlashStar(ctx context.Context, bar sql.NullString) (sql.NullString, error) {
	row := q.db.QueryRowContext(ctx, slashStar, bar)
	var bar sql.NullString
	err := row.Scan(&bar)
	return bar, err
}
