// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: query.sql

package querytest

import (
	"context"
)

const complicated = `-- name: Complicated :many
WITH names AS (SELECT name from foo WHERE foo.name = ?)
SELECT name FROM names WHERE name IN (SELECT name FROM foo WHERE foo.name = ?)
`

type ComplicatedParams struct {
	Slug string
}

func (q *Queries) Complicated(ctx context.Context, arg ComplicatedParams) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, complicated, arg.Slug, arg.Slug)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		items = append(items, name)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const funcParamIdent = `-- name: FuncParamIdent :many
SELECT name FROM foo WHERE name = ?
`

func (q *Queries) FuncParamIdent(ctx context.Context, slug string) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, funcParamIdent, slug)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		items = append(items, name)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const funcParamString = `-- name: FuncParamString :many
SELECT name FROM foo WHERE name = ?
`

func (q *Queries) FuncParamString(ctx context.Context, slug string) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, funcParamString, slug)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		items = append(items, name)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
