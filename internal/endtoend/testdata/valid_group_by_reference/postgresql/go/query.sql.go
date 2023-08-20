// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: query.sql

package querytest

import (
	"context"
	"database/sql"
)

const listAuthors = `-- name: ListAuthors :many
SELECT   id, name as name, bio
FROM     authors
GROUP BY name
`

func (q *Queries) ListAuthors(ctx context.Context, aq ...AdditionalQuery) ([]Author, error) {
	query := listAuthors
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
	var items []Author
	for rows.Next() {
		var i Author
		if err := rows.Scan(&i.ID, &i.Name, &i.Bio); err != nil {
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

const listAuthorsIdenticalAlias = `-- name: ListAuthorsIdenticalAlias :many
SELECT   id, name as name, bio
FROM     authors
GROUP BY name
`

func (q *Queries) ListAuthorsIdenticalAlias(ctx context.Context, aq ...AdditionalQuery) ([]Author, error) {
	query := listAuthorsIdenticalAlias
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
	var items []Author
	for rows.Next() {
		var i Author
		if err := rows.Scan(&i.ID, &i.Name, &i.Bio); err != nil {
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

const listMetrics = `-- name: ListMetrics :many
SELECT time_bucket('15 days', time) AS bucket, city_name, AVG(temp_c)
FROM weather_metrics
WHERE time > NOW() - (6 * INTERVAL '1 month')
GROUP BY bucket, city_name
ORDER BY bucket DESC
`

type ListMetricsRow struct {
	Bucket   interface{}
	CityName sql.NullString
	Avg      float64
}

func (q *Queries) ListMetrics(ctx context.Context, aq ...AdditionalQuery) ([]ListMetricsRow, error) {
	query := listMetrics
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
	var items []ListMetricsRow
	for rows.Next() {
		var i ListMetricsRow
		if err := rows.Scan(&i.Bucket, &i.CityName, &i.Avg); err != nil {
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
