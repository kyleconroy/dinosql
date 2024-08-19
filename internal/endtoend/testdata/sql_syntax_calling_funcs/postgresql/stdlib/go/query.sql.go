// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package querytest

import (
	"context"
)

const mixedNotation = `-- name: MixedNotation :one
SELECT concat_lower_or_upper('Hello', 'World', uppercase => true)
`

func (q *Queries) MixedNotation(ctx context.Context) (string, error) {
	row := q.db.QueryRowContext(ctx, mixedNotation)
	var concat_lower_or_upper string
	err := row.Scan(&concat_lower_or_upper)
	return concat_lower_or_upper, err
}

const namedAnyOrder = `-- name: NamedAnyOrder :one
SELECT concat_lower_or_upper(a => 'Hello', b => 'World', uppercase => true)
`

func (q *Queries) NamedAnyOrder(ctx context.Context) (string, error) {
	row := q.db.QueryRowContext(ctx, namedAnyOrder)
	var concat_lower_or_upper string
	err := row.Scan(&concat_lower_or_upper)
	return concat_lower_or_upper, err
}

const namedNotation = `-- name: NamedNotation :one
SELECT concat_lower_or_upper(a => 'Hello', b => 'World')
`

func (q *Queries) NamedNotation(ctx context.Context) (string, error) {
	row := q.db.QueryRowContext(ctx, namedNotation)
	var concat_lower_or_upper string
	err := row.Scan(&concat_lower_or_upper)
	return concat_lower_or_upper, err
}

const namedOtherOrder = `-- name: NamedOtherOrder :one
SELECT concat_lower_or_upper(a => 'Hello', uppercase => true, b => 'World')
`

func (q *Queries) NamedOtherOrder(ctx context.Context) (string, error) {
	row := q.db.QueryRowContext(ctx, namedOtherOrder)
	var concat_lower_or_upper string
	err := row.Scan(&concat_lower_or_upper)
	return concat_lower_or_upper, err
}

const positionalNoDefaault = `-- name: PositionalNoDefaault :one
SELECT concat_lower_or_upper('Hello', 'World')
`

func (q *Queries) PositionalNoDefaault(ctx context.Context) (string, error) {
	row := q.db.QueryRowContext(ctx, positionalNoDefaault)
	var concat_lower_or_upper string
	err := row.Scan(&concat_lower_or_upper)
	return concat_lower_or_upper, err
}

const positionalNotation = `-- name: PositionalNotation :one
SELECT concat_lower_or_upper('Hello', 'World', true)
`

func (q *Queries) PositionalNotation(ctx context.Context) (string, error) {
	row := q.db.QueryRowContext(ctx, positionalNotation)
	var concat_lower_or_upper string
	err := row.Scan(&concat_lower_or_upper)
	return concat_lower_or_upper, err
}
