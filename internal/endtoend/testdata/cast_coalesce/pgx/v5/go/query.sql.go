// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package querytest

import (
	"context"
)

const castCoalesce = `-- name: CastCoalesce :many
SELECT coalesce(bar, '')::text as login
FROM foo
`

func (q *Queries) CastCoalesce(ctx context.Context) ([]string, error) {
	rows, err := q.db.Query(ctx, castCoalesce)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var login string
		if err := rows.Scan(&login); err != nil {
			return nil, err
		}
		items = append(items, login)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
