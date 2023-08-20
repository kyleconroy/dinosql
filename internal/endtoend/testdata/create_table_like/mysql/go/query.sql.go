// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: query.sql

package querytest

import (
	"context"
)

const getAll = `-- name: GetAll :many
SELECT id, first_name, last_name, date_of_birth FROM super_users
`

func (q *Queries) GetAll(ctx context.Context, aq ...AdditionalQuery) ([]SuperUser, error) {
	query := getAll
	queryParams := []interface{}{}

	if len(aq) > 0 {
		query += " " + aq[0].SQL
		queryParams = append(queryParams, aq[0].Args...)
	}

	rows, err := q.db.QueryContext(ctx, query, queryParams...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SuperUser
	for rows.Next() {
		var i SuperUser
		if err := rows.Scan(
			&i.ID,
			&i.FirstName,
			&i.LastName,
			&i.DateOfBirth,
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
