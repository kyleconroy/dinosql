// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: query.sql

package querytest

import (
	"context"
)

const createFoo = `-- name: CreateFoo :exec
INSERT INTO foo (id) VALUES (?)
`

func (q *Queries) CreateFoo(ctx context.Context, id uint32) error {
	_, err := q.db.ExecContext(ctx, createFoo, id)
	return err
}
