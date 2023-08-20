// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: query.sql

package querytest

import (
	"context"
	"database/sql"
)

const identOnNonNullable = `-- name: IdentOnNonNullable :many
SELECT bar FROM foo WHERE bar = $1
`

func (q *Queries) IdentOnNonNullable(ctx context.Context, bar sql.NullString, aq ...AdditionalQuery) ([]string, error) {
	query := identOnNonNullable
	queryParams := []interface{}{bar}

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
		var bar string
		if err := rows.Scan(&bar); err != nil {
			return nil, err
		}
		items = append(items, bar)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const identOnNullable = `-- name: IdentOnNullable :many
SELECT maybe_bar FROM foo WHERE maybe_bar = $1
`

func (q *Queries) IdentOnNullable(ctx context.Context, maybeBar sql.NullString, aq ...AdditionalQuery) ([]sql.NullString, error) {
	query := identOnNullable
	queryParams := []interface{}{maybeBar}

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
		var maybe_bar sql.NullString
		if err := rows.Scan(&maybe_bar); err != nil {
			return nil, err
		}
		items = append(items, maybe_bar)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const stringOnNonNullable = `-- name: StringOnNonNullable :many
SELECT bar FROM foo WHERE bar = $1
`

func (q *Queries) StringOnNonNullable(ctx context.Context, bar sql.NullString, aq ...AdditionalQuery) ([]string, error) {
	query := stringOnNonNullable
	queryParams := []interface{}{bar}

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
		var bar string
		if err := rows.Scan(&bar); err != nil {
			return nil, err
		}
		items = append(items, bar)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const stringOnNullable = `-- name: StringOnNullable :many
SELECT maybe_bar FROM foo WHERE maybe_bar = $1
`

func (q *Queries) StringOnNullable(ctx context.Context, maybeBar sql.NullString, aq ...AdditionalQuery) ([]sql.NullString, error) {
	query := stringOnNullable
	queryParams := []interface{}{maybeBar}

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
		var maybe_bar sql.NullString
		if err := rows.Scan(&maybe_bar); err != nil {
			return nil, err
		}
		items = append(items, maybe_bar)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
