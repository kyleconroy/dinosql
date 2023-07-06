// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.0
// source: query.sql

package querytest

import (
	"context"
)

const getNextID = `-- name: GetNextID :one
SELECT pk, pk FROM
 (SELECT nextval('authors_id_seq') as pk) AS alias
`

type GetNextIDRow struct {
	Pk   int64
	Pk_2 int64
}

func (q *Queries) GetNextID(ctx context.Context) (GetNextIDRow, error) {
	row := q.db.QueryRowContext(ctx, getNextID)
	var i GetNextIDRow
	err := row.Scan(&i.Pk, &i.Pk_2)
	return i, err
}
