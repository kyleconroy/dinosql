// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: query.sql

package db

import (
	"context"
)

const deleteJoin = `-- name: DeleteJoin :exec
DELETE jt.*,
pt.*
FROM
        join_table as jt
        JOIN primary_table as pt ON jt.primary_table_id = pt.id
WHERE
        jt.id = ?
        AND pt.user_id = ?
`

type DeleteJoinParams struct {
	ID     uint64
	UserID uint64
}

func (q *Queries) DeleteJoin(ctx context.Context, arg DeleteJoinParams) error {
	query := deleteJoin
	queryParams := []interface{}{arg.ID, arg.UserID}

	_, err := q.db.ExecContext(ctx, query, queryParams...)
	return err
}

const deleteLeftJoin = `-- name: DeleteLeftJoin :exec
DELETE jt.*,
pt.*
FROM
        join_table as jt
        LEFT JOIN primary_table as pt ON jt.primary_table_id = pt.id
WHERE
        jt.id = ?
        AND pt.user_id = ?
`

type DeleteLeftJoinParams struct {
	ID     uint64
	UserID uint64
}

func (q *Queries) DeleteLeftJoin(ctx context.Context, arg DeleteLeftJoinParams) error {
	query := deleteLeftJoin
	queryParams := []interface{}{arg.ID, arg.UserID}

	_, err := q.db.ExecContext(ctx, query, queryParams...)
	return err
}

const deleteRightJoin = `-- name: DeleteRightJoin :exec
DELETE jt.*,
pt.*
FROM
        join_table as jt
        RIGHT JOIN primary_table as pt ON jt.primary_table_id = pt.id
WHERE
        jt.id = ?
        AND pt.user_id = ?
`

type DeleteRightJoinParams struct {
	ID     uint64
	UserID uint64
}

func (q *Queries) DeleteRightJoin(ctx context.Context, arg DeleteRightJoinParams) error {
	query := deleteRightJoin
	queryParams := []interface{}{arg.ID, arg.UserID}

	_, err := q.db.ExecContext(ctx, query, queryParams...)
	return err
}
