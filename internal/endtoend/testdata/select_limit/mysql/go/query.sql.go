// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: query.sql

package querytest

import (
	"context"
	"database/sql"
)

const fooLimit = `-- name: FooLimit :many
SELECT a FROM foo
LIMIT ?
`

func (q *Queries) FooLimit(ctx context.Context, limit int32, aq ...AdditionalQuery) ([]sql.NullString, error) {
	query := fooLimit
	queryParams := []interface{}{limit}

	if len(aq) > 0 {
		query += " " + aq[0].SQL
		queryParams = append(queryParams, aq[0].Args...)
	}

	rows, err := q.db.QueryContext(ctx, query, queryParams...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []sql.NullString
	for rows.Next() {
		var a sql.NullString
		if err := rows.Scan(&a); err != nil {
			return nil, err
		}
		items = append(items, a)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const fooLimitOffset = `-- name: FooLimitOffset :many
SELECT a FROM foo
LIMIT ? OFFSET ?
`

type FooLimitOffsetParams struct {
	Limit  int32
	Offset int32
}

func (q *Queries) FooLimitOffset(ctx context.Context, arg FooLimitOffsetParams, aq ...AdditionalQuery) ([]sql.NullString, error) {
	query := fooLimitOffset
	queryParams := []interface{}{arg.Limit, arg.Offset}

	if len(aq) > 0 {
		query += " " + aq[0].SQL
		queryParams = append(queryParams, aq[0].Args...)
	}

	rows, err := q.db.QueryContext(ctx, query, queryParams...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []sql.NullString
	for rows.Next() {
		var a sql.NullString
		if err := rows.Scan(&a); err != nil {
			return nil, err
		}
		items = append(items, a)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
