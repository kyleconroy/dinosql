// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: query.sql

package querytest

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const starExpansionCTE = `-- name: StarExpansionCTE :many
WITH cte AS (SELECT a, b FROM foo) SELECT a, b FROM cte
`

type StarExpansionCTERow struct {
	A pgtype.Text
	B pgtype.Text
}

func (q *Queries) StarExpansionCTE(ctx context.Context) ([]StarExpansionCTERow, error) {
	rows, err := q.db.Query(ctx, starExpansionCTE)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []StarExpansionCTERow
	for rows.Next() {
		var i StarExpansionCTERow
		if err := rows.Scan(&i.A, &i.B); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
