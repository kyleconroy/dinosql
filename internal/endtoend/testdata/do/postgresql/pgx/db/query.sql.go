// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: query.sql

package db

import (
	"context"
)

const doStuff = `-- name: DoStuff :exec
DO $$
    BEGIN
        ALTER TABLE authors
        ADD COLUMN marked_for_processing bool;
    END
$$
`

func (q *Queries) DoStuff(ctx context.Context) error {
	_, err := q.db.Exec(ctx, doStuff)
	return err
}
