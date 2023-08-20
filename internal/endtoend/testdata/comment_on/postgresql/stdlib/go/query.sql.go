// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: query.sql

package querytest

import (
	"context"
)

const listBar = `-- name: ListBar :many
SELECT baz FROM foo.bar
`

func (q *Queries) ListBar(ctx context.Context, aq ...AdditionalQuery) ([]string, error) {
	query := listBar
	queryParams := []interface{}{}

	if len(aq) > 0 {
		query += " " + aq[0].SQL
		queryParams = append(queryParams, aq[0].Args...)
	}

	rows, err := q.db.QueryContext(ctx, query, queryParams...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var baz string
		if err := rows.Scan(&baz); err != nil {
			return nil, err
		}
		items = append(items, baz)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listBat = `-- name: ListBat :many
SELECT baz FROM foo.bat
`

func (q *Queries) ListBat(ctx context.Context, aq ...AdditionalQuery) ([]string, error) {
	query := listBat
	queryParams := []interface{}{}

	if len(aq) > 0 {
		query += " " + aq[0].SQL
		queryParams = append(queryParams, aq[0].Args...)
	}

	rows, err := q.db.QueryContext(ctx, query, queryParams...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var baz string
		if err := rows.Scan(&baz); err != nil {
			return nil, err
		}
		items = append(items, baz)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
