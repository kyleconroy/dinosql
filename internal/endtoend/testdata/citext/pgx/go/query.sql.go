// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package querytest

import (
	"context"
)

const getCitexts = `-- name: GetCitexts :many
SELECT bar, bat
FROM foo
`

func (q *Queries) GetCitexts(ctx context.Context) ([]Foo, error) {
	rows, err := q.db.Query(ctx, getCitexts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Foo
	for rows.Next() {
		var i Foo
		if err := rows.Scan(&i.Bar, &i.Bat); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
