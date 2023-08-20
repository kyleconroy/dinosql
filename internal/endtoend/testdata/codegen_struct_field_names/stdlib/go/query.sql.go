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

func (q *Queries) test(ctx context.Context, aq ...AdditionalQuery) (Bar, error) {
	query := test
	queryParams := []interface{}{}

	if len(aq) > 0 {
		query += " " + aq[0].SQL
		queryParams = append(queryParams, aq[0].Args...)
	}

	row := q.db.QueryRowContext(ctx, query, queryParams...)
	var i Bar
	err := row.Scan(&i.ID, &i.NobodyWouldBelieveThis, &i.ParentID)
	return i, err
}
