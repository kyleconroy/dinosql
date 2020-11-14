// Code generated by sqlc. DO NOT EDIT.
// source: query.sql

package override

import (
	"context"

	"github.com/gofrs/uuid"
)

const loadFoo = `-- name: LoadFoo :many
SELECT id, about FROM foo WHERE id = $1
`

func (q *Queries) LoadFoo(ctx context.Context, id uuid.UUID) ([]Foo, error) {
	rows, err := q.db.QueryContext(ctx, loadFoo, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Foo
	for rows.Next() {
		var i Foo
		if err := rows.Scan(&i.ID, &i.About); err != nil {
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
