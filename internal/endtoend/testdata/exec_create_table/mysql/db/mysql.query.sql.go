// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: mysql.query.sql

package db

import (
	"context"
)

const createTable = `-- name: CreateTable :exec
CREATE TABLE test (id INTEGER NOT NULL)
`

func (q *Queries) CreateTable(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, createTable)
	return err
}
