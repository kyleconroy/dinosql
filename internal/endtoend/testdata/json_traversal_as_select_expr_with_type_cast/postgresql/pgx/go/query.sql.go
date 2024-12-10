// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package querytest

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const myGet = `-- name: MyGet :many
SELECT id, myjson, (mt.myjson->'thing1'->'thing2')::text
FROM mytable mt
`

type MyGetRow struct {
	ID      int64       `json:"id"`
	Myjson  []byte      `json:"myjson"`
	Column3 pgtype.Text `json:"column_3"`
}

func (q *Queries) MyGet(ctx context.Context) ([]MyGetRow, error) {
	rows, err := q.db.Query(ctx, myGet)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []MyGetRow
	for rows.Next() {
		var i MyGetRow
		if err := rows.Scan(&i.ID, &i.Myjson, &i.Column3); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}