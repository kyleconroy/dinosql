// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
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

func (q *Queries) WithAs(ctx context.Context, aq ...AdditionalQuery) (WithAsRow, error) {
	query := withAs
	queryParams := []interface{}{}

	if len(aq) > 0 {
		query += " " + aq[0].SQL
		queryParams = append(queryParams, aq[0].Args...)
	}

	row := q.db.QueryRowContext(ctx, query, queryParams...)
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

func (q *Queries) WithoutAs(ctx context.Context, aq ...AdditionalQuery) (WithoutAsRow, error) {
	query := withoutAs
	queryParams := []interface{}{}

	if len(aq) > 0 {
		query += " " + aq[0].SQL
		queryParams = append(queryParams, aq[0].Args...)
	}

	row := q.db.QueryRowContext(ctx, query, queryParams...)
	var i WithoutAsRow
	err := row.Scan(&i.X, &i.Y)
	return i, err
}
