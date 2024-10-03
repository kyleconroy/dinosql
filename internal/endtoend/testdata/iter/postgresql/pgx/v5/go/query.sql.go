// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package querytest

import (
	"context"
	"iter"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

const iterValues = `-- name: IterValues :iter
SELECT a, b
FROM myschema.foo
WHERE b = $1
`

func (q *Queries) IterValues(ctx context.Context, b pgtype.Int4) IterValuesRows {
	rows, err := q.db.Query(ctx, iterValues, b)
	if err != nil {
		return IterValuesRows{err: err}
	}
	return IterValuesRows{rows: rows}
}

type IterValuesRows struct {
	rows pgx.Rows
	err  error
}

func (r *IterValuesRows) Iterate() iter.Seq[MyschemaFoo] {
	if r.rows == nil {
		return func(yield func(MyschemaFoo) bool) {}
	}

	return func(yield func(MyschemaFoo) bool) {
		for r.rows.Next() {
			var i MyschemaFoo
			err := r.rows.Scan(&i.A, &i.B)
			if err != nil {
				r.err = err
				return
			}

			if !yield(i) {
				r.rows.Close()
				return
			}
		}
	}
}

func (r *IterValuesRows) Close() {
	r.rows.Close()
}

func (r *IterValuesRows) Err() error {
	if r.err != nil {
		return r.err
	}
	return r.rows.Err()
}
