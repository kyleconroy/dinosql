// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: query.sql

package querytest

import (
	"context"
)

const placeholder = `-- name: Placeholder :exec
SELECT 1
`

func (q *Queries) Placeholder(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, placeholder)
	return err
}
