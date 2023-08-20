// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: city.sql

package ondeck

import (
	"context"
)

const createCity = `-- name: CreateCity :one
INSERT INTO city (
    name,
    slug
) VALUES (
    $1,
    $2
) RETURNING slug, name
`

type CreateCityParams struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

// Create a new city. The slug must be unique.
// This is the second line of the comment
// This is the third line
func (q *Queries) CreateCity(ctx context.Context, arg CreateCityParams, aq ...AdditionalQuery) (City, error) {
	query := createCity
	queryParams := []interface{}{arg.Name, arg.Slug}

	if len(aq) > 0 {
		query += " " + aq[0].SQL
		queryParams = append(queryParams, aq[0].Args...)
	}

	row := q.queryRow(ctx, q.createCityStmt, query, queryParams...)
	var i City
	err := row.Scan(&i.Slug, &i.Name)
	return i, err
}

const getCity = `-- name: GetCity :one
SELECT slug, name
FROM city
WHERE slug = $1
`

func (q *Queries) GetCity(ctx context.Context, slug string, aq ...AdditionalQuery) (City, error) {
	query := getCity
	queryParams := []interface{}{slug}

	if len(aq) > 0 {
		query += " " + aq[0].SQL
		queryParams = append(queryParams, aq[0].Args...)
	}

	row := q.queryRow(ctx, q.getCityStmt, query, queryParams...)
	var i City
	err := row.Scan(&i.Slug, &i.Name)
	return i, err
}

const listCities = `-- name: ListCities :many
SELECT slug, name
FROM city
ORDER BY name
`

func (q *Queries) ListCities(ctx context.Context, aq ...AdditionalQuery) ([]City, error) {
	query := listCities
	queryParams := []interface{}{}

	if len(aq) > 0 {
		query += " " + aq[0].SQL
		queryParams = append(queryParams, aq[0].Args...)
	}

	rows, err := q.query(ctx, q.listCitiesStmt, query, queryParams...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []City
	for rows.Next() {
		var i City
		if err := rows.Scan(&i.Slug, &i.Name); err != nil {
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

const updateCityName = `-- name: UpdateCityName :exec
UPDATE city
SET name = $2
WHERE slug = $1
`

type UpdateCityNameParams struct {
	Slug string `json:"slug"`
	Name string `json:"name"`
}

func (q *Queries) UpdateCityName(ctx context.Context, arg UpdateCityNameParams) error {
	query := updateCityName
	queryParams := []interface{}{arg.Slug, arg.Name}

	_, err := q.exec(ctx, q.updateCityNameStmt, query, queryParams...)
	return err
}
