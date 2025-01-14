// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package querytest

import (
	"context"
)

const authorPages = `-- name: AuthorPages :many
select author, count(title) as num_books, SUM(pages) as total_pages, SUM(score) AS sum_score, SUM(price) AS sum_price, SUM(avg_word_length) as sum_avg_length
from books
group by author
`

type AuthorPagesRow struct {
	Author       string
	NumBooks     int64
	TotalPages   int64
	SumScore     int64
	SumPrice     int64
	SumAvgLength int64
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
		if err := rows.Scan(
			&i.Author,
			&i.NumBooks,
			&i.TotalPages,
			&i.SumScore,
			&i.SumPrice,
			&i.SumAvgLength,
		); err != nil {
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
