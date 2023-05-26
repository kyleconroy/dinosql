// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: query.sql

package querytest

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const getUserByID = `-- name: GetUserByID :one
SELECT first_name, id, last_name FROM users WHERE id = $1
`

type GetUserByIDRow struct {
	FirstName string
	ID        int32
	LastName  pgtype.Text
}

func (q *Queries) GetUserByID(ctx context.Context, targetID int32) (GetUserByIDRow, error) {
	row := q.db.QueryRow(ctx, getUserByID, targetID)
	var i GetUserByIDRow
	err := row.Scan(&i.FirstName, &i.ID, &i.LastName)
	return i, err
}

const insertNewUser = `-- name: InsertNewUser :exec
INSERT INTO users (first_name, last_name) VALUES ($1, $2)
`

type InsertNewUserParams struct {
	FirstName string
	LastName  pgtype.Text
}

func (q *Queries) InsertNewUser(ctx context.Context, arg InsertNewUserParams) error {
	_, err := q.db.Exec(ctx, insertNewUser, arg.FirstName, arg.LastName)
	return err
}

const limitSQLCArg = `-- name: LimitSQLCArg :many
select first_name, id FROM users LIMIT $1
`

type LimitSQLCArgRow struct {
	FirstName string
	ID        int32
}

func (q *Queries) LimitSQLCArg(ctx context.Context, limit int32) ([]LimitSQLCArgRow, error) {
	rows, err := q.db.Query(ctx, limitSQLCArg, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []LimitSQLCArgRow
	for rows.Next() {
		var i LimitSQLCArgRow
		if err := rows.Scan(&i.FirstName, &i.ID); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listUserOrders = `-- name: ListUserOrders :many
SELECT
	users.id,
	users.first_name,
	orders.price
FROM
	orders
LEFT JOIN users ON orders.user_id = users.id
WHERE orders.price > $1
`

type ListUserOrdersRow struct {
	ID        pgtype.Int4
	FirstName pgtype.Text
	Price     pgtype.Numeric
}

func (q *Queries) ListUserOrders(ctx context.Context, minPrice pgtype.Numeric) ([]ListUserOrdersRow, error) {
	rows, err := q.db.Query(ctx, listUserOrders, minPrice)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListUserOrdersRow
	for rows.Next() {
		var i ListUserOrdersRow
		if err := rows.Scan(&i.ID, &i.FirstName, &i.Price); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listUserParenExpr = `-- name: ListUserParenExpr :many
SELECT id, first_name, last_name, age, job_status FROM users WHERE (job_status = 'APPLIED' OR job_status = 'PENDING')
AND id > $1
ORDER BY id
LIMIT $2
`

type ListUserParenExprParams struct {
	ID    int32
	Limit int32
}

func (q *Queries) ListUserParenExpr(ctx context.Context, arg ListUserParenExprParams) ([]User, error) {
	rows, err := q.db.Query(ctx, listUserParenExpr, arg.ID, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.FirstName,
			&i.LastName,
			&i.Age,
			&i.JobStatus,
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

const listUsersByFamily = `-- name: ListUsersByFamily :many
SELECT first_name, last_name FROM users WHERE age < $1 AND last_name = $2
`

type ListUsersByFamilyParams struct {
	MaxAge   int32
	InFamily pgtype.Text
}

type ListUsersByFamilyRow struct {
	FirstName string
	LastName  pgtype.Text
}

func (q *Queries) ListUsersByFamily(ctx context.Context, arg ListUsersByFamilyParams) ([]ListUsersByFamilyRow, error) {
	rows, err := q.db.Query(ctx, listUsersByFamily, arg.MaxAge, arg.InFamily)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListUsersByFamilyRow
	for rows.Next() {
		var i ListUsersByFamilyRow
		if err := rows.Scan(&i.FirstName, &i.LastName); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listUsersByID = `-- name: ListUsersByID :many
SELECT first_name, id, last_name FROM users WHERE id < $1
`

type ListUsersByIDRow struct {
	FirstName string
	ID        int32
	LastName  pgtype.Text
}

func (q *Queries) ListUsersByID(ctx context.Context, id int32) ([]ListUsersByIDRow, error) {
	rows, err := q.db.Query(ctx, listUsersByID, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListUsersByIDRow
	for rows.Next() {
		var i ListUsersByIDRow
		if err := rows.Scan(&i.FirstName, &i.ID, &i.LastName); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listUsersWithLimit = `-- name: ListUsersWithLimit :many
SELECT first_name, last_name FROM users LIMIT $1
`

type ListUsersWithLimitRow struct {
	FirstName string
	LastName  pgtype.Text
}

func (q *Queries) ListUsersWithLimit(ctx context.Context, limit int32) ([]ListUsersWithLimitRow, error) {
	rows, err := q.db.Query(ctx, listUsersWithLimit, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListUsersWithLimitRow
	for rows.Next() {
		var i ListUsersWithLimitRow
		if err := rows.Scan(&i.FirstName, &i.LastName); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
