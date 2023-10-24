// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: query.sql

package querytest

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const findWallets = `-- name: FindWallets :many
select id, address, balance, total_balance from 
(
	select id, address, balance,
	  sum(balance) over (order by balance desc rows between unbounded preceding and current row) as total_balance,
	  sum(balance) over (order by balance desc rows between unbounded preceding and current row) - balance as last_balance
	from wallets
	where type=$1
) amounts
where amounts.last_balance < $2
`

type FindWalletsParams struct {
	Column1 pgtype.Text
	Column2 pgtype.Numeric
}

type FindWalletsRow struct {
	ID           int64
	Address      string
	Balance      pgtype.Numeric
	TotalBalance pgtype.Numeric
}

func (q *Queries) FindWallets(ctx context.Context, arg FindWalletsParams) ([]FindWalletsRow, error) {
	rows, err := q.db.Query(ctx, findWallets, arg.Column1, arg.Column2)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FindWalletsRow
	for rows.Next() {
		var i FindWalletsRow
		if err := rows.Scan(
			&i.ID,
			&i.Address,
			&i.Balance,
			&i.TotalBalance,
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
