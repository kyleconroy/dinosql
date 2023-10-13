// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: query.sql

package querytest

import (
	"context"
)

const test2 = `-- name: Test2 :one
select txt from Demo
where txt like '%' || $1::text || '%'
`

func (q *Queries) Test2(ctx context.Context, val string) (string, error) {
	row := q.db.QueryRow(ctx, test2, val)
	var txt string
	err := row.Scan(&txt)
	return txt, err
}

const test3 = `-- name: Test3 :one
select txt from Demo
where txt like concat('%', $1::text, '%')
`

func (q *Queries) Test3(ctx context.Context, val string) (string, error) {
	row := q.db.QueryRow(ctx, test3, val)
	var txt string
	err := row.Scan(&txt)
	return txt, err
}
