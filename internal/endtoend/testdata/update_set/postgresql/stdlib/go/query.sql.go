// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: query.sql

package querytest

import (
	"context"
)

const updateSet = `-- name: UpdateSet :exec
UPDATE foo SET name = $2 WHERE slug = $1
`

type UpdateSetParams struct {
	Slug string
	Name string
}

func (q *Queries) UpdateSet(ctx context.Context, arg UpdateSetParams) error {
	_, err := q.db.ExecContext(ctx, updateSet, arg.Slug, arg.Name)
	return err
}
