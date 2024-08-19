// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package querytest

import (
	"context"
)

const caseStatementBoolean = `-- name: CaseStatementBoolean :many
SELECT CASE 
  WHEN id = $1 THEN true
  ELSE false
END is_one
FROM foo
`

func (q *Queries) CaseStatementBoolean(ctx context.Context, id string) ([]bool, error) {
	rows, err := q.db.Query(ctx, caseStatementBoolean, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []bool
	for rows.Next() {
		var is_one bool
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
