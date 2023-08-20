// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: query.sql

package querytest

import (
	"context"
	"database/sql"
)

const test = `-- name: Test :one
select txt from Demo
where txt ~~ '%' || $1 || '%'
`

func (q *Queries) Test(ctx context.Context, val string, aq ...AdditionalQuery) (string, error) {
	query := test
	queryParams := []interface{}{val}

	if len(aq) > 0 {
		query += " " + aq[0].SQL
		queryParams = append(queryParams, aq[0].Args...)
	}

	row := q.db.QueryRowContext(ctx, query, queryParams...)
	var txt string
	err := row.Scan(&txt)
	return txt, err
}

const test2 = `-- name: Test2 :one
select txt from Demo
where txt like '%' || $1 || '%'
`

func (q *Queries) Test2(ctx context.Context, val sql.NullString, aq ...AdditionalQuery) (string, error) {
	query := test2
	queryParams := []interface{}{val}

	if len(aq) > 0 {
		query += " " + aq[0].SQL
		queryParams = append(queryParams, aq[0].Args...)
	}

	row := q.db.QueryRowContext(ctx, query, queryParams...)
	var txt string
	err := row.Scan(&txt)
	return txt, err
}

const test3 = `-- name: Test3 :one
select txt from Demo
where txt like concat('%', $1, '%')
`

func (q *Queries) Test3(ctx context.Context, val interface{}, aq ...AdditionalQuery) (string, error) {
	query := test3
	queryParams := []interface{}{val}

	if len(aq) > 0 {
		query += " " + aq[0].SQL
		queryParams = append(queryParams, aq[0].Args...)
	}

	row := q.db.QueryRowContext(ctx, query, queryParams...)
	var txt string
	err := row.Scan(&txt)
	return txt, err
}
