// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: query.sql

package querytest

import (
	"context"
	"database/sql"
	"strings"
)

const funcNullable = `-- name: FuncNullable :many
SELECT bar FROM foo
WHERE id IN (/*SLICE:favourites*/?)
`

func (q *Queries) FuncNullable(ctx context.Context, favourites []int32) ([]sql.NullString, error) {
	query := funcNullable
	var queryParams []interface{}
	if len(favourites) > 0 {
		for _, v := range favourites {
			queryParams = append(queryParams, v)
		}
		query = strings.Replace(query, "/*SLICE:favourites*/?", strings.Repeat(",?", len(favourites))[1:], 1)
	} else {
		query = strings.Replace(query, "/*SLICE:favourites*/?", "NULL", 1)
	}
	rows, err := q.db.QueryContext(ctx, query, queryParams...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []sql.NullString
	for rows.Next() {
		var bar sql.NullString
		if err := rows.Scan(&bar); err != nil {
			return nil, err
		}
		items = append(items, bar)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const funcParamIdent = `-- name: FuncParamIdent :many
SELECT name FROM foo
WHERE name = ?
  AND id IN (/*SLICE:favourites*/?)
`

type FuncParamIdentParams struct {
	Slug       string
	Favourites []int32
}

func (q *Queries) FuncParamIdent(ctx context.Context, arg FuncParamIdentParams) ([]string, error) {
	query := funcParamIdent
	var queryParams []interface{}
	queryParams = append(queryParams, arg.Slug)
	if len(arg.Favourites) > 0 {
		for _, v := range arg.Favourites {
			queryParams = append(queryParams, v)
		}
		query = strings.Replace(query, "/*SLICE:favourites*/?", strings.Repeat(",?", len(arg.Favourites))[1:], 1)
	} else {
		query = strings.Replace(query, "/*SLICE:favourites*/?", "NULL", 1)
	}
	rows, err := q.db.QueryContext(ctx, query, queryParams...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		items = append(items, name)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const funcParamSoloArg = `-- name: FuncParamSoloArg :many
SELECT name FROM foo
WHERE id IN (/*SLICE:favourites*/?)
`

func (q *Queries) FuncParamSoloArg(ctx context.Context, favourites []int32) ([]string, error) {
	query := funcParamSoloArg
	var queryParams []interface{}
	if len(favourites) > 0 {
		for _, v := range favourites {
			queryParams = append(queryParams, v)
		}
		query = strings.Replace(query, "/*SLICE:favourites*/?", strings.Repeat(",?", len(favourites))[1:], 1)
	} else {
		query = strings.Replace(query, "/*SLICE:favourites*/?", "NULL", 1)
	}
	rows, err := q.db.QueryContext(ctx, query, queryParams...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		items = append(items, name)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const funcParamString = `-- name: FuncParamString :many
SELECT name FROM foo
WHERE name = ?
  AND id IN (/*SLICE:favourites*/?)
`

type FuncParamStringParams struct {
	Slug       string
	Favourites []int32
}

func (q *Queries) FuncParamString(ctx context.Context, arg FuncParamStringParams) ([]string, error) {
	query := funcParamString
	var queryParams []interface{}
	queryParams = append(queryParams, arg.Slug)
	if len(arg.Favourites) > 0 {
		for _, v := range arg.Favourites {
			queryParams = append(queryParams, v)
		}
		query = strings.Replace(query, "/*SLICE:favourites*/?", strings.Repeat(",?", len(arg.Favourites))[1:], 1)
	} else {
		query = strings.Replace(query, "/*SLICE:favourites*/?", "NULL", 1)
	}
	rows, err := q.db.QueryContext(ctx, query, queryParams...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		items = append(items, name)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const sliceExec = `-- name: SliceExec :exec
UPDATE foo SET name = ?
WHERE id IN (/*SLICE:favourites*/?)
`

type SliceExecParams struct {
	Slug       string
	Favourites []int32
}

func (q *Queries) SliceExec(ctx context.Context, arg SliceExecParams) error {
	query := sliceExec
	var queryParams []interface{}
	queryParams = append(queryParams, arg.Slug)
	if len(arg.Favourites) > 0 {
		for _, v := range arg.Favourites {
			queryParams = append(queryParams, v)
		}
		query = strings.Replace(query, "/*SLICE:favourites*/?", strings.Repeat(",?", len(arg.Favourites))[1:], 1)
	} else {
		query = strings.Replace(query, "/*SLICE:favourites*/?", "NULL", 1)
	}
	_, err := q.db.ExecContext(ctx, query, queryParams...)
	return err
}
