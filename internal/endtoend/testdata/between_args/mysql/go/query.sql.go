// Code generated by sqlc. DO NOT EDIT.
// source: query.sql

package querytest

import (
	"context"
)

const getBetweenPrices = `-- name: GetBetweenPrices :many
SELECT	id, name, price
FROM	products
WHERE	price BETWEEN ? AND ?
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
