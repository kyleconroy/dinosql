// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: query.sql

package querytest

import (
	"context"
	"database/sql"
)

const duplicate = `-- name: Duplicate :one
SELECT users.id, users.name, users.age, users.id, users.name, users.age FROM users
`

type DuplicateRow struct {
	User   User
	User_2 User
}

func (q *Queries) Duplicate(ctx context.Context, aq ...AdditionalQuery) (DuplicateRow, error) {
	query := duplicate
	queryParams := []interface{}{}

	if len(aq) > 0 {
		query += " " + aq[0].SQL
		queryParams = append(queryParams, aq[0].Args...)
	}

	row := q.db.QueryRowContext(ctx, query, queryParams...)
	var i DuplicateRow
	err := row.Scan(
		&i.User.ID,
		&i.User.Name,
		&i.User.Age,
		&i.User_2.ID,
		&i.User_2.Name,
		&i.User_2.Age,
	)
	return i, err
}

const join = `-- name: Join :one
SELECT users.id, users.name, users.age, posts.id, posts.user_id FROM posts
INNER JOIN users ON posts.user_id = users.id
`

type JoinRow struct {
	User User
	Post Post
}

func (q *Queries) Join(ctx context.Context, aq ...AdditionalQuery) (JoinRow, error) {
	query := join
	queryParams := []interface{}{}

	if len(aq) > 0 {
		query += " " + aq[0].SQL
		queryParams = append(queryParams, aq[0].Args...)
	}

	row := q.db.QueryRowContext(ctx, query, queryParams...)
	var i JoinRow
	err := row.Scan(
		&i.User.ID,
		&i.User.Name,
		&i.User.Age,
		&i.Post.ID,
		&i.Post.UserID,
	)
	return i, err
}

const only = `-- name: Only :one
SELECT users.id, users.name, users.age FROM users
`

type OnlyRow struct {
	User User
}

func (q *Queries) Only(ctx context.Context, aq ...AdditionalQuery) (OnlyRow, error) {
	query := only
	queryParams := []interface{}{}

	if len(aq) > 0 {
		query += " " + aq[0].SQL
		queryParams = append(queryParams, aq[0].Args...)
	}

	row := q.db.QueryRowContext(ctx, query, queryParams...)
	var i OnlyRow
	err := row.Scan(&i.User.ID, &i.User.Name, &i.User.Age)
	return i, err
}

const withAlias = `-- name: WithAlias :one
SELECT u.id, u.name, u.age FROM users u
`

type WithAliasRow struct {
	User User
}

func (q *Queries) WithAlias(ctx context.Context, aq ...AdditionalQuery) (WithAliasRow, error) {
	query := withAlias
	queryParams := []interface{}{}

	if len(aq) > 0 {
		query += " " + aq[0].SQL
		queryParams = append(queryParams, aq[0].Args...)
	}

	row := q.db.QueryRowContext(ctx, query, queryParams...)
	var i WithAliasRow
	err := row.Scan(&i.User.ID, &i.User.Name, &i.User.Age)
	return i, err
}

const withAsterisk = `-- name: WithAsterisk :one
SELECT users.id, users.name, users.age, id, name, age FROM users
`

type WithAsteriskRow struct {
	User User
	ID   int32
	Name string
	Age  sql.NullInt32
}

func (q *Queries) WithAsterisk(ctx context.Context, aq ...AdditionalQuery) (WithAsteriskRow, error) {
	query := withAsterisk
	queryParams := []interface{}{}

	if len(aq) > 0 {
		query += " " + aq[0].SQL
		queryParams = append(queryParams, aq[0].Args...)
	}

	row := q.db.QueryRowContext(ctx, query, queryParams...)
	var i WithAsteriskRow
	err := row.Scan(
		&i.User.ID,
		&i.User.Name,
		&i.User.Age,
		&i.ID,
		&i.Name,
		&i.Age,
	)
	return i, err
}

const withCrossSchema = `-- name: WithCrossSchema :many
SELECT users.id, users.name, users.age, bu.id, bu.name FROM users
INNER JOIN baz.users bu ON users.id = bu.id
`

type WithCrossSchemaRow struct {
	User    User
	BazUser BazUser
}

func (q *Queries) WithCrossSchema(ctx context.Context, aq ...AdditionalQuery) ([]WithCrossSchemaRow, error) {
	query := withCrossSchema
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
	var items []WithCrossSchemaRow
	for rows.Next() {
		var i WithCrossSchemaRow
		if err := rows.Scan(
			&i.User.ID,
			&i.User.Name,
			&i.User.Age,
			&i.BazUser.ID,
			&i.BazUser.Name,
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

const withSchema = `-- name: WithSchema :one
SELECT bu.id, bu.name FROM baz.users bu
`

type WithSchemaRow struct {
	BazUser BazUser
}

func (q *Queries) WithSchema(ctx context.Context, aq ...AdditionalQuery) (WithSchemaRow, error) {
	query := withSchema
	queryParams := []interface{}{}

	if len(aq) > 0 {
		query += " " + aq[0].SQL
		queryParams = append(queryParams, aq[0].Args...)
	}

	row := q.db.QueryRowContext(ctx, query, queryParams...)
	var i WithSchemaRow
	err := row.Scan(&i.BazUser.ID, &i.BazUser.Name)
	return i, err
}

const withSubquery = `-- name: WithSubquery :many
SELECT users.id, users.name, users.age, (SELECT count(*) FROM users) AS total_count FROM users
`

type WithSubqueryRow struct {
	User       User
	TotalCount int64
}

func (q *Queries) WithSubquery(ctx context.Context, aq ...AdditionalQuery) ([]WithSubqueryRow, error) {
	query := withSubquery
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
	var items []WithSubqueryRow
	for rows.Next() {
		var i WithSubqueryRow
		if err := rows.Scan(
			&i.User.ID,
			&i.User.Name,
			&i.User.Age,
			&i.TotalCount,
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
