// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0

package querytest

import (
	"context"
)

type Querier interface {
	CreateAuthors(ctx context.Context, authorID []int32) (int64, error)
}

var _ Querier = (*Queries)(nil)
