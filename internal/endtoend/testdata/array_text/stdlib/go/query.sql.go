// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package querytest

import (
	"context"

	"github.com/lib/pq"
)

const textArray = `-- name: TextArray :many
SELECT tags FROM bar
`

func (q *Queries) TextArray(ctx context.Context) ([][]string, error) {
	rows, err := q.db.QueryContext(ctx, textArray)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items [][]string
	for rows.Next() {
		var tags []string
		if err := rows.Scan(pq.Array(&tags)); err != nil {
			return nil, err
		}
		items = append(items, tags)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
