// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: query.sql

package querytest

import (
	"context"
)

const caseStatementText = `-- name: CaseStatementText :many
SELECT CASE 
  WHEN id = $1 THEN 'foo'
  ELSE 'bar'
END is_one
FROM foo
`

func (q *Queries) CaseStatementText(ctx context.Context, id string) ([]string, error) {
	rows, err := q.db.Query(ctx, caseStatementText, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var is_one string
		if err := rows.Scan(&is_one); err != nil {
			return nil, err
		}
		items = append(items, is_one)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
