// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: query.sql

package querytest

import (
	"context"
)

const mixedNotation = `-- name: MixedNotation :one
SELECT concat_lower_or_upper('Hello', 'World', uppercase => true)
`

func (q *Queries) MixedNotation(ctx context.Context, aq ...AdditionalQuery) (string, error) {
	query := mixedNotation
	queryParams := []interface{}{}

	if len(aq) > 0 {
		query += " " + aq[0].SQL
		queryParams = append(queryParams, aq[0].Args...)
	}

	row := q.db.QueryRowContext(ctx, query, queryParams...)
	var concat_lower_or_upper string
	err := row.Scan(&concat_lower_or_upper)
	return concat_lower_or_upper, err
}

const namedAnyOrder = `-- name: NamedAnyOrder :one
SELECT concat_lower_or_upper(a => 'Hello', b => 'World', uppercase => true)
`

func (q *Queries) NamedAnyOrder(ctx context.Context, aq ...AdditionalQuery) (string, error) {
	query := namedAnyOrder
	queryParams := []interface{}{}

	if len(aq) > 0 {
		query += " " + aq[0].SQL
		queryParams = append(queryParams, aq[0].Args...)
	}

	row := q.db.QueryRowContext(ctx, query, queryParams...)
	var concat_lower_or_upper string
	err := row.Scan(&concat_lower_or_upper)
	return concat_lower_or_upper, err
}

const namedNotation = `-- name: NamedNotation :one
SELECT concat_lower_or_upper(a => 'Hello', b => 'World')
`

func (q *Queries) NamedNotation(ctx context.Context, aq ...AdditionalQuery) (string, error) {
	query := namedNotation
	queryParams := []interface{}{}

	if len(aq) > 0 {
		query += " " + aq[0].SQL
		queryParams = append(queryParams, aq[0].Args...)
	}

	row := q.db.QueryRowContext(ctx, query, queryParams...)
	var concat_lower_or_upper string
	err := row.Scan(&concat_lower_or_upper)
	return concat_lower_or_upper, err
}

const namedOtherOrder = `-- name: NamedOtherOrder :one
SELECT concat_lower_or_upper(a => 'Hello', uppercase => true, b => 'World')
`

func (q *Queries) NamedOtherOrder(ctx context.Context, aq ...AdditionalQuery) (string, error) {
	query := namedOtherOrder
	queryParams := []interface{}{}

	if len(aq) > 0 {
		query += " " + aq[0].SQL
		queryParams = append(queryParams, aq[0].Args...)
	}

	row := q.db.QueryRowContext(ctx, query, queryParams...)
	var concat_lower_or_upper string
	err := row.Scan(&concat_lower_or_upper)
	return concat_lower_or_upper, err
}

const positionalNoDefaault = `-- name: PositionalNoDefaault :one
SELECT concat_lower_or_upper('Hello', 'World')
`

func (q *Queries) PositionalNoDefaault(ctx context.Context, aq ...AdditionalQuery) (string, error) {
	query := positionalNoDefaault
	queryParams := []interface{}{}

	if len(aq) > 0 {
		query += " " + aq[0].SQL
		queryParams = append(queryParams, aq[0].Args...)
	}

	row := q.db.QueryRowContext(ctx, query, queryParams...)
	var concat_lower_or_upper string
	err := row.Scan(&concat_lower_or_upper)
	return concat_lower_or_upper, err
}

const positionalNotation = `-- name: PositionalNotation :one
SELECT concat_lower_or_upper('Hello', 'World', true)
`

func (q *Queries) PositionalNotation(ctx context.Context, aq ...AdditionalQuery) (string, error) {
	query := positionalNotation
	queryParams := []interface{}{}

	if len(aq) > 0 {
		query += " " + aq[0].SQL
		queryParams = append(queryParams, aq[0].Args...)
	}

	row := q.db.QueryRowContext(ctx, query, queryParams...)
	var concat_lower_or_upper string
	err := row.Scan(&concat_lower_or_upper)
	return concat_lower_or_upper, err
}
