// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: query.sql

package querytest

import (
	"context"
)

const cTEFilter = `-- name: CTEFilter :many
WITH filter_count AS (
	SELECT count(*) FROM bar WHERE ready = $1
)
SELECT filter_count.count
FROM filter_count
`

func (q *Queries) CTEFilter(ctx context.Context, ready bool) ([]int64, error) {
	rows, err := q.db.Query(ctx, cTEFilter, ready)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []int64
	for rows.Next() {
		var count int64
		if err := rows.Scan(&count); err != nil {
			return nil, err
		}
		items = append(items, count)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
