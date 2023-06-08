// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: query.sql

package querytest

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const findAccount = `-- name: FindAccount :one
SELECT
    a.id, a.state,
    ua.name
    -- other fields
FROM
    accounts a
    INNER JOIN users_accounts ua ON a.id = ua.id2
WHERE
    a.id = $1
`

type FindAccountRow struct {
	ID    uuid.UUID `sometagtype:"some_value"`
	State sql.NullString
	Name  sql.NullString
}

func (q *Queries) FindAccount(ctx context.Context, accountID uuid.UUID) (FindAccountRow, error) {
	row := q.db.QueryRowContext(ctx, findAccount, accountID)
	var i FindAccountRow
	err := row.Scan(&i.ID, &i.State, &i.Name)
	return i, err
}
