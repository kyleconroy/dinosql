// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: query.sql

package querytest

import (
	"context"
)

const selectFoo = `-- name: SelectFoo :exec
SELECT a, b, c, d FROM foo
`

func (q *Queries) SelectFoo(ctx context.Context) error {
	query := selectFoo
	queryParams := []interface{}{}

	_, err := q.db.ExecContext(ctx, query, queryParams...)
	return err
}
