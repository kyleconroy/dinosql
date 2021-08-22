// Code generated by sqlc. DO NOT EDIT.
// source: query.sql

package datatype

import (
	"context"
	"database/sql"
)

const createExample = `-- name: CreateExample :one
INSERT INTO examples (value) VALUES ($1)
RETURNING example_id, value, create_time, update_time
`

func (q *Queries) CreateExample(ctx context.Context, value sql.NullString) (Example, error) {
	row := q.db.QueryRow(ctx, createExample, value)
	var i Example
	err := row.Scan(
		&i.ExampleID,
		&i.Value,
		&i.CreateTime,
		&i.UpdateTime,
	)
	return i, err
}
