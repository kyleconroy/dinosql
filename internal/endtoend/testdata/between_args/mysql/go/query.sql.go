// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.0
// source: query.sql

package querytest

import (
	"context"
)

const getBetweenPrices = `-- name: GetBetweenPrices :many
SELECT  id, name, price
FROM    products
WHERE   price BETWEEN ? AND ?
`

type GetBetweenPricesParams struct {
	FromPrice int32
	ToPrice   int32
}

func (q *Queries) GetBetweenPrices(ctx context.Context, arg GetBetweenPricesParams) ([]Product, error) {
	rows, err := q.db.QueryContext(ctx, getBetweenPrices, arg.FromPrice, arg.ToPrice)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Product
	for rows.Next() {
		var i Product
		if err := rows.Scan(&i.ID, &i.Name, &i.Price); err != nil {
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

const getBetweenPricesNamed = `-- name: GetBetweenPricesNamed :many
SELECT  id, name, price
FROM    products
WHERE   price BETWEEN ? AND ?
`

type GetBetweenPricesNamedParams struct {
	MinPrice int32
	MaxPrice int32
}

func (q *Queries) GetBetweenPricesNamed(ctx context.Context, arg GetBetweenPricesNamedParams) ([]Product, error) {
	rows, err := q.db.QueryContext(ctx, getBetweenPricesNamed, arg.MinPrice, arg.MaxPrice)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Product
	for rows.Next() {
		var i Product
		if err := rows.Scan(&i.ID, &i.Name, &i.Price); err != nil {
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

const getBetweenPricesTable = `-- name: GetBetweenPricesTable :many
SELECT  id, name, price
FROM    products
WHERE   products.price BETWEEN ? AND ?
`

type GetBetweenPricesTableParams struct {
	FromPrice int32
	ToPrice   int32
}

func (q *Queries) GetBetweenPricesTable(ctx context.Context, arg GetBetweenPricesTableParams) ([]Product, error) {
	rows, err := q.db.QueryContext(ctx, getBetweenPricesTable, arg.FromPrice, arg.ToPrice)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Product
	for rows.Next() {
		var i Product
		if err := rows.Scan(&i.ID, &i.Name, &i.Price); err != nil {
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

const getBetweenPricesTableAlias = `-- name: GetBetweenPricesTableAlias :many
SELECT  id, name, price
FROM    products as p
WHERE   p.price BETWEEN ? AND ?
`

type GetBetweenPricesTableAliasParams struct {
	FromPrice int32
	ToPrice   int32
}

func (q *Queries) GetBetweenPricesTableAlias(ctx context.Context, arg GetBetweenPricesTableAliasParams) ([]Product, error) {
	rows, err := q.db.QueryContext(ctx, getBetweenPricesTableAlias, arg.FromPrice, arg.ToPrice)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Product
	for rows.Next() {
		var i Product
		if err := rows.Scan(&i.ID, &i.Name, &i.Price); err != nil {
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
