// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: query.sql

package querytest

import (
	"context"
	"database/sql"
	"strings"
)

const selectUsers = `-- name: SelectUsers :many
SELECT first_name, last_name FROM users WHERE age > ?
`

type SelectUsersRow struct {
	FirstName string
	LastName  sql.NullString
}

func (q *Queries) SelectUsers(ctx context.Context, age int32) ([]SelectUsersRow, error) {
	rows, err := q.db.QueryContext(ctx, selectUsers, age)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SelectUsersRow
	for rows.Next() {
		var i SelectUsersRow
		if err := rows.Scan(&i.FirstName, &i.LastName); err != nil {
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

const selectUsersDynamic = `-- name: SelectUsersDynamic :many
SELECT first_name, last_name FROM users WHERE age > ? AND /*DYNAMIC:dynamic*/?
`

type SelectUsersDynamicParams struct {
	Age     int32
	Dynamic DynamicSql
}

type SelectUsersDynamicRow struct {
	FirstName string
	LastName  sql.NullString
}

func (q *Queries) SelectUsersDynamic(ctx context.Context, arg SelectUsersDynamicParams) ([]SelectUsersDynamicRow, error) {
	query := selectUsersDynamic
	var queryParams []interface{}
	curNumb := 2
	queryParams = append(queryParams, arg.Age)
	replaceText, args := arg.Dynamic.ToSql(curNumb)
	curNumb += len(args)
	query = strings.ReplaceAll(query, "/*DYNAMIC:dynamic*/?", "replaceText", 1)
	queryParams = append(queryParams, v)
	rows, err := q.db.QueryContext(ctx, query, queryParams...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SelectUsersDynamicRow
	for rows.Next() {
		var i SelectUsersDynamicRow
		if err := rows.Scan(&i.FirstName, &i.LastName); err != nil {
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

const selectUsersDynamic2 = `-- name: SelectUsersDynamic2 :many
SELECT first_name, last_name
FROM users
WHERE age > ? AND
    job_status = ? AND
    /*DYNAMIC:dynamic*/?
`

type SelectUsersDynamic2Params struct {
	Age     int32
	Status  string
	Dynamic DynamicSql
}

type SelectUsersDynamic2Row struct {
	FirstName string
	LastName  sql.NullString
}

func (q *Queries) SelectUsersDynamic2(ctx context.Context, arg SelectUsersDynamic2Params) ([]SelectUsersDynamic2Row, error) {
	query := selectUsersDynamic2
	var queryParams []interface{}
	curNumb := 3
	queryParams = append(queryParams, arg.Age)
	queryParams = append(queryParams, arg.Status)
	replaceText, args := arg.Dynamic.ToSql(curNumb)
	curNumb += len(args)
	query = strings.ReplaceAll(query, "/*DYNAMIC:dynamic*/?", "replaceText", 1)
	queryParams = append(queryParams, v)
	rows, err := q.db.QueryContext(ctx, query, queryParams...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SelectUsersDynamic2Row
	for rows.Next() {
		var i SelectUsersDynamic2Row
		if err := rows.Scan(&i.FirstName, &i.LastName); err != nil {
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
