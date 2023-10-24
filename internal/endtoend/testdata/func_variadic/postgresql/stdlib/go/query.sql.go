// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: query.sql

package querytest

import (
	"context"
	"database/sql"
)

const updateJ = `-- name: UpdateJ :exec
UPDATE
    test
SET
    j = jsonb_build_object($1::text, $2::text, $3::text, $4::text)
WHERE
    id = $5
`

type UpdateJParams struct {
	Column1 string
	Column2 string
	Column3 string
	Column4 string
	ID      sql.NullInt32
}

func (q *Queries) UpdateJ(ctx context.Context, arg UpdateJParams) error {
	_, err := q.db.ExecContext(ctx, updateJ,
		arg.Column1,
		arg.Column2,
		arg.Column3,
		arg.Column4,
		arg.ID,
	)
	return err
}
