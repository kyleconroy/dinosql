// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package db

import (
	"context"
	"strings"
)

const getFullNames = `-- name: GetFullNames :many
SELECT
 full_name
FROM
  (
    SELECT
  	 concat(first_name, ' ', last_name) as full_name
    FROM
      customers
  ) subquery
WHERE
  full_name IN ($1)
`

func (q *Queries) GetFullNames(ctx context.Context, fullNames []interface{}) ([]interface{}, error) {
	query := getFullNames
	var queryParams []interface{}
	if len(fullNames) > 0 {
		for _, v := range fullNames {
			queryParams = append(queryParams, v)
		}
		query = strings.Replace(query, "/*SLICE:full_names*/?", strings.Repeat(",?", len(fullNames))[1:], 1)
	} else {
		query = strings.Replace(query, "/*SLICE:full_names*/?", "NULL", 1)
	}
	rows, err := q.db.QueryContext(ctx, query, queryParams...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []interface{}
	for rows.Next() {
		var full_name interface{}
		if err := rows.Scan(&full_name); err != nil {
			return nil, err
		}
		items = append(items, full_name)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getFullNames2 = `-- name: GetFullNames2 :many
WITH subquery AS (
    SELECT
  	 concat(first_name, ' ', last_name) as full_name
    FROM
      customers
  )
SELECT
 full_name
FROM
  subquery
WHERE
  full_name IN ($1)
`

func (q *Queries) GetFullNames2(ctx context.Context, fullNames []interface{}) ([]interface{}, error) {
	query := getFullNames2
	var queryParams []interface{}
	if len(fullNames) > 0 {
		for _, v := range fullNames {
			queryParams = append(queryParams, v)
		}
		query = strings.Replace(query, "/*SLICE:full_names*/?", strings.Repeat(",?", len(fullNames))[1:], 1)
	} else {
		query = strings.Replace(query, "/*SLICE:full_names*/?", "NULL", 1)
	}
	rows, err := q.db.QueryContext(ctx, query, queryParams...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []interface{}
	for rows.Next() {
		var full_name interface{}
		if err := rows.Scan(&full_name); err != nil {
			return nil, err
		}
		items = append(items, full_name)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
