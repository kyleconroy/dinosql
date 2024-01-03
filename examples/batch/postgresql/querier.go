// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package batch

import (
	"context"

	"github.com/jackc/pgx/v5/pgconn"
)

type Querier interface {
	BooksByYear(ctx context.Context, year []int32) *BooksByYearBatchResults
	CreateAuthor(ctx context.Context, name string) (Author, error)
	CreateBook(ctx context.Context, arg []CreateBookParams) *CreateBookBatchResults
	DeleteBook(ctx context.Context, bookID []int32) *DeleteBookBatchResults
	DeleteBookExecResult(ctx context.Context, bookID int32) (pgconn.CommandTag, error)
	DeleteBookNamedFunc(ctx context.Context, bookID []int32) *DeleteBookNamedFuncBatchResults
	DeleteBookNamedSign(ctx context.Context, bookID []int32) *DeleteBookNamedSignBatchResults
	GetAuthor(ctx context.Context, authorID int32) (Author, error)
	GetBiography(ctx context.Context, authorID []int32) *GetBiographyBatchResults
	UpdateBook(ctx context.Context, arg []UpdateBookParams) *UpdateBookBatchResults
}

var _ Querier = (*Queries)(nil)
