// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: query.sql

package override

import (
	"context"

	t "github.com/jackc/pgtype"
)

const test = `-- name: test :exec
UPDATE foo SET langs = $1
`

func (q *Queries) test(ctx context.Context, langs *t.Text) error {
	_, err := q.db.ExecContext(ctx, test, langs)
	return err
}
