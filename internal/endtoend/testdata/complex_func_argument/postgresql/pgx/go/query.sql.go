// Code generated by sqlc. DO NOT EDIT.
// source: query.sql

package querytest

import (
	"context"
)

const listFoos = `-- name: ListFoos :one
SELECT id FROM foo WHERE id = get_id(CASE WHEN $1 = 100 THEN $2 ELSE 'baz' END)
`

func (q *Queries) ListFoos(ctx context.Context, dollar_1 interface{}) (string, error) {
	row := q.db.QueryRow(ctx, listFoos, dollar_1)
	var id string
	err := row.Scan(&id)
	return id, err
}
