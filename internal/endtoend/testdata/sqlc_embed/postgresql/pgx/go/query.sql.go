// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
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
	User   User `db:"user" json:"user"`
	User_2 User `db:"user_2" json:"user_2"`
}

func (q *Queries) Duplicate(ctx context.Context) (DuplicateRow, error) {
	row := q.db.QueryRow(ctx, duplicate)
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
	User User `db:"user" json:"user"`
	Post Post `db:"post" json:"post"`
}

func (q *Queries) Join(ctx context.Context) (JoinRow, error) {
	row := q.db.QueryRow(ctx, join)
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

const listUserLink = `-- name: ListUserLink :many
SELECT
    owner.id, owner.name, owner.age,
    consumer.id, consumer.name, consumer.age
FROM
    user_links
    INNER JOIN users AS owner ON owner.id = user_links.owner_id
    INNER JOIN users AS consumer ON consumer.id = user_links.consumer_id
`

type ListUserLinkRow struct {
	User   User `db:"user" json:"user"`
	User_2 User `db:"user_2" json:"user_2"`
}

func (q *Queries) ListUserLink(ctx context.Context) ([]ListUserLinkRow, error) {
	rows, err := q.db.Query(ctx, listUserLink)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListUserLinkRow
	for rows.Next() {
		var i ListUserLinkRow
		if err := rows.Scan(
			&i.User.ID,
			&i.User.Name,
			&i.User.Age,
			&i.User_2.ID,
			&i.User_2.Name,
			&i.User_2.Age,
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

const only = `-- name: Only :one
SELECT users.id, users.name, users.age FROM users
`

type OnlyRow struct {
	User User `db:"user" json:"user"`
}

func (q *Queries) Only(ctx context.Context) (OnlyRow, error) {
	row := q.db.QueryRow(ctx, only)
	var i OnlyRow
	err := row.Scan(&i.User.ID, &i.User.Name, &i.User.Age)
	return i, err
}

const withAlias = `-- name: WithAlias :one
SELECT u.id, u.name, u.age FROM users u
`

type WithAliasRow struct {
	User User `db:"user" json:"user"`
}

func (q *Queries) WithAlias(ctx context.Context) (WithAliasRow, error) {
	row := q.db.QueryRow(ctx, withAlias)
	var i WithAliasRow
	err := row.Scan(&i.User.ID, &i.User.Name, &i.User.Age)
	return i, err
}

const withAsterisk = `-- name: WithAsterisk :one
SELECT users.id, users.name, users.age, id, name, age FROM users
`

type WithAsteriskRow struct {
	User User          `db:"user" json:"user"`
	ID   int32         `db:"id" json:"id"`
	Name string        `db:"name" json:"name"`
	Age  sql.NullInt32 `db:"age" json:"age"`
}

func (q *Queries) WithAsterisk(ctx context.Context) (WithAsteriskRow, error) {
	row := q.db.QueryRow(ctx, withAsterisk)
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
	User    User    `db:"user" json:"user"`
	BazUser BazUser `db:"baz_user" json:"baz_user"`
}

func (q *Queries) WithCrossSchema(ctx context.Context) ([]WithCrossSchemaRow, error) {
	rows, err := q.db.Query(ctx, withCrossSchema)
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
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const withSchema = `-- name: WithSchema :one
SELECT bu.id, bu.name FROM baz.users bu
`

type WithSchemaRow struct {
	BazUser BazUser `db:"baz_user" json:"baz_user"`
}

func (q *Queries) WithSchema(ctx context.Context) (WithSchemaRow, error) {
	row := q.db.QueryRow(ctx, withSchema)
	var i WithSchemaRow
	err := row.Scan(&i.BazUser.ID, &i.BazUser.Name)
	return i, err
}

const withSubquery = `-- name: WithSubquery :many
SELECT users.id, users.name, users.age, (SELECT count(*) FROM users) AS total_count FROM users
`

type WithSubqueryRow struct {
	User       User  `db:"user" json:"user"`
	TotalCount int64 `db:"total_count" json:"total_count"`
}

func (q *Queries) WithSubquery(ctx context.Context) ([]WithSubqueryRow, error) {
	rows, err := q.db.Query(ctx, withSubquery)
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
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
