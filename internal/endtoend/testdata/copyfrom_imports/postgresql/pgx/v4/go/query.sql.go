// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: query.sql

package querytest

import (
	"context"
	"database/sql"

	"github.com/jackc/pgconn"
)

const deleteValues = `-- name: DeleteValues :execresult
DELETE
FROM myschema.foo
`

func (q *Queries) DeleteValues(ctx context.Context) (pgconn.CommandTag, error) {
	return q.db.Exec(ctx, deleteValues)
}

const insertSingleValue = `-- name: InsertSingleValue :exec
INSERT INTO myschema.foo (a) VALUES ($1)
`

func (q *Queries) InsertSingleValue(ctx context.Context, a sql.NullString) error {
	_, err := q.db.Exec(ctx, insertSingleValue, a)
	return err
}

type InsertValuesParams struct {
	A sql.NullString
	B sql.NullInt32
}
