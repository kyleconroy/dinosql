// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: query.sql

package querytest

import (
	"context"
	"database/sql"

	"github.com/jackc/pgconn"
)

const execFoo = `-- name: ExecFoo :exec
INSERT INTO foo (bar) VALUES ('bar')
`

// This function creates a Foo via :exec
func (q *Queries) ExecFoo(ctx context.Context) error {
	_, err := q.db.Exec(ctx, execFoo)
	return err
}

const execResultFoo = `-- name: ExecResultFoo :execresult
INSERT INTO foo (bar) VALUES ('bar')
`

// This function creates a Foo via :execresult
func (q *Queries) ExecResultFoo(ctx context.Context) (pgconn.CommandTag, error) {
	return q.db.Exec(ctx, execResultFoo)
}

const execRowFoo = `-- name: ExecRowFoo :execrows
INSERT INTO foo (bar) VALUES ('bar')
`

// This function creates a Foo via :execrows
func (q *Queries) ExecRowFoo(ctx context.Context) (int64, error) {
	result, err := q.db.Exec(ctx, execRowFoo)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}

const manyFoo = `-- name: ManyFoo :many
SELECT bar FROM foo
`

// This function returns a list of Foos
func (q *Queries) ManyFoo(ctx context.Context) ([]sql.NullString, error) {
	rows, err := q.db.Query(ctx, manyFoo)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []sql.NullString
	for rows.Next() {
		var bar sql.NullString
		if err := rows.Scan(&bar); err != nil {
			return nil, err
		}
		items = append(items, bar)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const oneFoo = `-- name: OneFoo :one
SELECT bar FROM foo
`

// This function returns one Foo
func (q *Queries) OneFoo(ctx context.Context) (sql.NullString, error) {
	row := q.db.QueryRow(ctx, oneFoo)
	var bar sql.NullString
	err := row.Scan(&bar)
	return bar, err
}
