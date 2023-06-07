// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
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
	Price   int32
	Price_2 int32
}

func (q *Queries) GetBetweenPrices(ctx context.Context, arg GetBetweenPricesParams) ([]Product, error) {
	rows, err := q.db.QueryContext(ctx, getBetweenPrices, arg.Price, arg.Price_2)
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
	Price   int32
	Price_2 int32
}

func (q *Queries) GetBetweenPricesTable(ctx context.Context, arg GetBetweenPricesTableParams) ([]Product, error) {
	rows, err := q.db.QueryContext(ctx, getBetweenPricesTable, arg.Price, arg.Price_2)
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
	Price   int32
	Price_2 int32
}

func (q *Queries) GetBetweenPricesTableAlias(ctx context.Context, arg GetBetweenPricesTableAliasParams) ([]Product, error) {
	rows, err := q.db.QueryContext(ctx, getBetweenPricesTableAlias, arg.Price, arg.Price_2)
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
