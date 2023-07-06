// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.0
// source: query.sql

package db

import (
	"context"
)

const getRepro = `-- name: GetRepro :one
select id, name, seq from repro where id = ? limit 1
`

func (q *Queries) GetRepro(ctx context.Context, id interface{}) (Repro, error) {
	row := q.db.QueryRowContext(ctx, getRepro, id)
	var i Repro
	err := row.Scan(&i.ID, &i.Name, &i.Seq)
	return i, err
}
