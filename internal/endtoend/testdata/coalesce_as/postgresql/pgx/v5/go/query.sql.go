// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: query.sql

package querytest

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const sumBaz = `-- name: SumBaz :many
SELECT bar, coalesce(sum(baz), 0) as quantity
FROM foo
GROUP BY 1
`

type SumBazRow struct {
	Bar      pgtype.Text
	Quantity interface{}
}

func (q *Queries) SumBaz(ctx context.Context) ([]SumBazRow, error) {
	rows, err := q.db.Query(ctx, sumBaz)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SumBazRow
	for rows.Next() {
		var i SumBazRow
		if err := rows.Scan(&i.Bar, &i.Quantity); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
