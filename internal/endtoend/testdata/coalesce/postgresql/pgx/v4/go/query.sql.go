// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: query.sql

package querytest

import (
	"context"
	"database/sql"
)

const coalesceNumeric = `-- name: CoalesceNumeric :many
SELECT coalesce(baz, 0) as login
FROM foo
`

func (q *Queries) CoalesceNumeric(ctx context.Context, aq ...AdditionalQuery) ([]int64, error) {
	query := coalesceNumeric
	queryParams := []interface{}{}

	if len(aq) > 0 {
		query += " " + aq[0].SQL
		queryParams = append(queryParams, aq[0].Args...)
	}

	rows, err := q.db.Query(ctx, query, queryParams...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []int64
	for rows.Next() {
		var login int64
		if err := rows.Scan(&login); err != nil {
			return nil, err
		}
		items = append(items, login)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const coalesceNumericColumns = `-- name: CoalesceNumericColumns :many
SELECT baz, qux, coalesce(baz, qux)
FROM foo
`

type CoalesceNumericColumnsRow struct {
	Baz   sql.NullInt64
	Qux   int64
	Baz_2 int64
}

func (q *Queries) CoalesceNumericColumns(ctx context.Context, aq ...AdditionalQuery) ([]CoalesceNumericColumnsRow, error) {
	query := coalesceNumericColumns
	queryParams := []interface{}{}

	if len(aq) > 0 {
		query += " " + aq[0].SQL
		queryParams = append(queryParams, aq[0].Args...)
	}

	rows, err := q.db.Query(ctx, query, queryParams...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []CoalesceNumericColumnsRow
	for rows.Next() {
		var i CoalesceNumericColumnsRow
		if err := rows.Scan(&i.Baz, &i.Qux, &i.Baz_2); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const coalesceNumericNull = `-- name: CoalesceNumericNull :many
SELECT baz, coalesce(baz)
FROM foo
`

type CoalesceNumericNullRow struct {
	Baz   sql.NullInt64
	Baz_2 sql.NullInt64
}

func (q *Queries) CoalesceNumericNull(ctx context.Context, aq ...AdditionalQuery) ([]CoalesceNumericNullRow, error) {
	query := coalesceNumericNull
	queryParams := []interface{}{}

	if len(aq) > 0 {
		query += " " + aq[0].SQL
		queryParams = append(queryParams, aq[0].Args...)
	}

	rows, err := q.db.Query(ctx, query, queryParams...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []CoalesceNumericNullRow
	for rows.Next() {
		var i CoalesceNumericNullRow
		if err := rows.Scan(&i.Baz, &i.Baz_2); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const coalesceString = `-- name: CoalesceString :many
SELECT coalesce(bar, '') as login
FROM foo
`

func (q *Queries) CoalesceString(ctx context.Context, aq ...AdditionalQuery) ([]string, error) {
	query := coalesceString
	queryParams := []interface{}{}

	if len(aq) > 0 {
		query += " " + aq[0].SQL
		queryParams = append(queryParams, aq[0].Args...)
	}

	rows, err := q.db.Query(ctx, query, queryParams...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var login string
		if err := rows.Scan(&login); err != nil {
			return nil, err
		}
		items = append(items, login)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const coalesceStringColumns = `-- name: CoalesceStringColumns :many
SELECT bar, bat, coalesce(bar, bat)
FROM foo
`

type CoalesceStringColumnsRow struct {
	Bar   sql.NullString
	Bat   string
	Bar_2 string
}

func (q *Queries) CoalesceStringColumns(ctx context.Context, aq ...AdditionalQuery) ([]CoalesceStringColumnsRow, error) {
	query := coalesceStringColumns
	queryParams := []interface{}{}

	if len(aq) > 0 {
		query += " " + aq[0].SQL
		queryParams = append(queryParams, aq[0].Args...)
	}

	rows, err := q.db.Query(ctx, query, queryParams...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []CoalesceStringColumnsRow
	for rows.Next() {
		var i CoalesceStringColumnsRow
		if err := rows.Scan(&i.Bar, &i.Bat, &i.Bar_2); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const coalesceStringNull = `-- name: CoalesceStringNull :many
SELECT bar, coalesce(bar)
FROM foo
`

type CoalesceStringNullRow struct {
	Bar   sql.NullString
	Bar_2 sql.NullString
}

func (q *Queries) CoalesceStringNull(ctx context.Context, aq ...AdditionalQuery) ([]CoalesceStringNullRow, error) {
	query := coalesceStringNull
	queryParams := []interface{}{}

	if len(aq) > 0 {
		query += " " + aq[0].SQL
		queryParams = append(queryParams, aq[0].Args...)
	}

	rows, err := q.db.Query(ctx, query, queryParams...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []CoalesceStringNullRow
	for rows.Next() {
		var i CoalesceStringNullRow
		if err := rows.Scan(&i.Bar, &i.Bar_2); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
