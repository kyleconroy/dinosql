// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: query.sql

package querytest

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createMemories = `-- name: CreateMemories :many
INSERT INTO memories (vampire_id)
SELECT
    unnest($1::uuid[]) AS vampire_id
RETURNING
    id, vampire_id, created_at, updated_at
`

func (q *Queries) CreateMemories(ctx context.Context, vampireID pgtype.Array[pgtype.UUID]) ([]Memory, error) {
	rows, err := q.db.Query(ctx, createMemories, vampireID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Memory
	for rows.Next() {
		var i Memory
		if err := rows.Scan(
			&i.ID,
			&i.VampireID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
