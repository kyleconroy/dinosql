// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: query.sql

package querytest

import (
	"context"
)

const callInsertData = `-- name: CallInsertData :exec
CALL insert_data($1, $2)
`

type CallInsertDataParams struct {
	A int32
	B int32
}

func (q *Queries) CallInsertData(ctx context.Context, arg CallInsertDataParams) error {
	query := callInsertData
	queryParams := []interface{}{arg.A, arg.B}

	_, err := q.db.ExecContext(ctx, query, queryParams...)
	return err
}

const callInsertDataNamed = `-- name: CallInsertDataNamed :exec
CALL insert_data(b => $1, a => $2)
`

type CallInsertDataNamedParams struct {
	B int32
	A int32
}

func (q *Queries) CallInsertDataNamed(ctx context.Context, arg CallInsertDataNamedParams) error {
	query := callInsertDataNamed
	queryParams := []interface{}{arg.B, arg.A}

	_, err := q.db.ExecContext(ctx, query, queryParams...)
	return err
}

const callInsertDataNoArgs = `-- name: CallInsertDataNoArgs :exec
CALL insert_data(1, 2)
`

func (q *Queries) CallInsertDataNoArgs(ctx context.Context) error {
	query := callInsertDataNoArgs
	queryParams := []interface{}{}

	_, err := q.db.ExecContext(ctx, query, queryParams...)
	return err
}

const callInsertDataSqlcArgs = `-- name: CallInsertDataSqlcArgs :exec
CALL insert_data($1, $2)
`

type CallInsertDataSqlcArgsParams struct {
	Foo int32
	Bar int32
}

func (q *Queries) CallInsertDataSqlcArgs(ctx context.Context, arg CallInsertDataSqlcArgsParams) error {
	query := callInsertDataSqlcArgs
	queryParams := []interface{}{arg.Foo, arg.Bar}

	_, err := q.db.ExecContext(ctx, query, queryParams...)
	return err
}
