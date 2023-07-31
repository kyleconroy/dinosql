// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: query.sql

package querytest

import (
	"context"
)

const test = `-- name: test :one
SELECT id, !!!nobody,_,-would-believe---this-...?!, parent        id from bar limit 1
`

func (q *Queries) test(ctx context.Context) (Bar, error) {
	row := q.db.QueryRowContext(ctx, test)
	var i Bar
	err := row.Scan(&i.ID, &i.NobodyWouldBelieveThis, &i.ParentID)
	return i, err
}
