// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: query.sql

package querytest

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const nestedSelect = `-- name: NestedSelect :one
SELECT latest.id, t.count
FROM (
  SELECT id, max(update_time) AS update_time
  FROM test
  WHERE id = ANY ($1::bigint[])
    -- ERROR HERE on update_time
    AND update_time >= $2
  GROUP BY id
) latest
INNER JOIN test t USING (id, update_time)
`

type NestedSelectParams struct {
	IDs       []int64
	StartTime pgtype.Int8
}

type NestedSelectRow struct {
	ID    int64
	Count int64
}

func (q *Queries) NestedSelect(ctx context.Context, arg NestedSelectParams) (NestedSelectRow, error) {
	row := q.db.QueryRow(ctx, nestedSelect, arg.IDs, arg.StartTime)
	var i NestedSelectRow
	err := row.Scan(&i.ID, &i.Count)
	return i, err
}
