// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package querytest

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const insertSelect = `-- name: InsertSelect :exec
INSERT INTO authors (id, name, bio) 
SELECT $1, name, bio
FROM authors
WHERE name = $2
`

type InsertSelectParams struct {
	ID   pgtype.Int8
	Name pgtype.Text
}

func (q *Queries) InsertSelect(ctx context.Context, arg InsertSelectParams) error {
	_, err := q.db.Exec(ctx, insertSelect, arg.ID, arg.Name)
	return err
}
