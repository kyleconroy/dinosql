// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package querytest

import (
	"context"
	"database/sql"
	"time"
)

const updateCode = `-- name: UpdateCode :one
WITH cc AS (
            UPDATE td3.codes
            SET
                created_by = $1,
                updated_by  = $1,
                code = $2,
                hash = $3,
                is_private = false
            RETURNING hash
)
UPDATE td3.test_codes
SET
    created_by = $1,
    updated_by  = $1,
    test_id = $4,
    code_hash = cc.hash
    FROM cc
RETURNING hash, id, ts_created, ts_updated, created_by, updated_by, test_id, code_hash
`

type UpdateCodeParams struct {
	CreatedBy string
	Code      sql.NullString
	Hash      sql.NullString
	TestID    int32
}

type UpdateCodeRow struct {
	Hash      sql.NullString
	ID        int32
	TsCreated time.Time
	TsUpdated time.Time
	CreatedBy string
	UpdatedBy string
	TestID    int32
	CodeHash  string
}

func (q *Queries) UpdateCode(ctx context.Context, arg UpdateCodeParams) (UpdateCodeRow, error) {
	row := q.db.QueryRowContext(ctx, updateCode,
		arg.CreatedBy,
		arg.Code,
		arg.Hash,
		arg.TestID,
	)
	var i UpdateCodeRow
	err := row.Scan(
		&i.Hash,
		&i.ID,
		&i.TsCreated,
		&i.TsUpdated,
		&i.CreatedBy,
		&i.UpdatedBy,
		&i.TestID,
		&i.CodeHash,
	)
	return i, err
}
