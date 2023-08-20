// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: query.sql

package querytest

import (
	"context"
)

const deleteAuthor = `-- name: DeleteAuthor :exec
UPDATE
  authors,
  books
SET
  authors.deleted_at = now(),
  books.deleted_at = now()
WHERE
  books.is_amazing = 1
  AND authors.name = ?
`

func (q *Queries) DeleteAuthor(ctx context.Context, name string) error {
	query := deleteAuthor
	queryParams := []interface{}{name}

	_, err := q.db.ExecContext(ctx, query, queryParams...)
	return err
}
