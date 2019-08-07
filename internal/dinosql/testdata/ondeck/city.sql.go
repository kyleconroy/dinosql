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
	Name string
	Slug string
}

func (q *Queries) CreateCity(ctx context.Context, arg CreateCityParams) (City, error) {
	row := q.db.QueryRowContext(ctx, createCity, arg.Name, arg.Slug)
	var i City
	err := row.Scan(&i.Slug, &i.Name)
	return i, err
}

const getCity = `-- name: GetCity :one
SELECT slug, name
FROM city
WHERE slug = $1
`

func (q *Queries) GetCity(ctx context.Context, slug string) (City, error) {
	row := q.db.QueryRowContext(ctx, getCity, slug)
	var i City
	err := row.Scan(&i.Slug, &i.Name)
	return i, err
}

const listCities = `-- name: ListCities :many
SELECT slug, name
FROM city
ORDER BY name
`

func (q *Queries) ListCities(ctx context.Context) ([]City, error) {
	rows, err := q.db.QueryContext(ctx, listCities)
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
	Slug string
	Name string
}

func (q *Queries) UpdateCityName(ctx context.Context, arg UpdateCityNameParams) error {
	_, err := q.db.ExecContext(ctx, updateCityName, arg.Slug, arg.Name)
	return err
}
