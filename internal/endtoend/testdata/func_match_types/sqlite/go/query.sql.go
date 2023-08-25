// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: query.sql

package querytest

import (
	"context"
	"database/sql"
)

const authorPages = `-- name: AuthorPages :many
select author, count(title) as num_books, sum(pages) as total_pages
from books
group by author
`

type AuthorPagesRow struct {
	Author     string
	NumBooks   int64
	TotalPages sql.NullFloat64
}

func (q *Queries) AuthorPages(ctx context.Context) ([]AuthorPagesRow, error) {
	rows, err := q.db.QueryContext(ctx, authorPages)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []AuthorPagesRow
	for rows.Next() {
		var i AuthorPagesRow
		if err := rows.Scan(&i.Author, &i.NumBooks, &i.TotalPages); err != nil {
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
