// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package authors

import (
	"context"
	"database/sql"
	"iter"
)

const createAuthor = `-- name: CreateAuthor :execresult
INSERT INTO authors (
  name, bio
) VALUES (
  ?, ? 
)
`

type CreateAuthorParams struct {
	Name string
	Bio  sql.NullString
}

func (q *Queries) CreateAuthor(ctx context.Context, arg CreateAuthorParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createAuthor, arg.Name, arg.Bio)
}

const deleteAuthor = `-- name: DeleteAuthor :exec
DELETE FROM authors
WHERE id = ?
`

func (q *Queries) DeleteAuthor(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteAuthor, id)
	return err
}

const getAuthor = `-- name: GetAuthor :one
SELECT id, name, bio FROM authors
WHERE id = ? LIMIT 1
`

func (q *Queries) GetAuthor(ctx context.Context, id int64) (Author, error) {
	row := q.db.QueryRowContext(ctx, getAuthor, id)
	var i Author
	err := row.Scan(&i.ID, &i.Name, &i.Bio)
	return i, err
}

const iterAuthors = `-- name: IterAuthors :iter
SELECT id, name, bio FROM authors
ORDER BY name
`

func (q *Queries) IterAuthors(ctx context.Context) IterAuthorsRows {
	rows, err := q.db.QueryContext(ctx, iterAuthors)
	if err != nil {
		return IterAuthorsRows{err: err}
	}
	return IterAuthorsRows{rows: rows}
}

type IterAuthorsRows struct {
	rows *sql.Rows
	err  error
}

func (r *IterAuthorsRows) Iterate() iter.Seq[Author] {
	if r.rows == nil {
		return func(yield func(Author) bool) {}
	}

	return func(yield func(Author) bool) {
		defer r.rows.Close()

		for r.rows.Next() {
			var i Author
			err := r.rows.Scan(&i.ID, &i.Name, &i.Bio)
			if err != nil {
				r.err = err
				return
			}

			if !yield(i) {
				r.err = r.rows.Close()
				return
			}
		}
	}
}

func (r *IterAuthorsRows) Close() error {
	return r.rows.Close()
}

func (r *IterAuthorsRows) Err() error {
	if r.err != nil {
		return r.err
	}
	return r.rows.Err()
}

const listAuthors = `-- name: ListAuthors :many
SELECT id, name, bio FROM authors
ORDER BY name
`

func (q *Queries) ListAuthors(ctx context.Context) ([]Author, error) {
	rows, err := q.db.QueryContext(ctx, listAuthors)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Author
	for rows.Next() {
		var i Author
		if err := rows.Scan(&i.ID, &i.Name, &i.Bio); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
