// Code generated by sqlc. DO NOT EDIT.
// source: authors.sql

package pggroupiface

import (
	"context"
)

type Authors interface {
	Create(ctx context.Context, name string) (Author, error)
	Get(ctx context.Context, authorID int32) (Author, error)
}

func NewAuthors(db DBTX) Authors {
	return &authors{db: db}
}

type authors struct {
	db DBTX
}

const authorsCreate = `-- name: Create :one
INSERT INTO authors (name) VALUES ($1)
RETURNING author_id, name
`

func (q *authors) Create(ctx context.Context, name string) (Author, error) {
	row := q.db.QueryRowContext(ctx, authorsCreate, name)
	var i Author
	err := row.Scan(&i.AuthorID, &i.Name)
	return i, err
}

const authorsGet = `-- name: Get :one
SELECT author_id, name FROM authors
WHERE author_id = $1
`

func (q *authors) Get(ctx context.Context, authorID int32) (Author, error) {
	row := q.db.QueryRowContext(ctx, authorsGet, authorID)
	var i Author
	err := row.Scan(&i.AuthorID, &i.Name)
	return i, err
}
