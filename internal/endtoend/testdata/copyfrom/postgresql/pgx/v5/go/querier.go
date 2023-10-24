// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0

package querytest

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type Querier interface {
	// InsertSingleValue inserts a single value using copy.
	InsertSingleValue(ctx context.Context, a []pgtype.Text) (int64, error)
	// InsertValues inserts multiple values using copy.
	InsertValues(ctx context.Context, arg []InsertValuesParams) (int64, error)
}

var _ Querier = (*Queries)(nil)
