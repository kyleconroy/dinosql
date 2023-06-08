// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: query.sql

package querytest

import (
	"context"
	"database/sql"
)

const insertValues = `-- name: InsertValues :exec
INSERT INTO foo (a, b) VALUES ($1, $2)
`

type InsertValuesParams struct {
	A sql.NullString
	B sql.NullInt32
}

func (q *Queries) InsertValues(ctx context.Context, arg InsertValuesParams) error {
	_, err := q.db.ExecContext(ctx, insertValues, arg.A, arg.B)
	return err
}
