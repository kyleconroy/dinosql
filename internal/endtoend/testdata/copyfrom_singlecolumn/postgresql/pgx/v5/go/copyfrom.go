// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: copyfrom.go

package querytest

import (
	"context"
)

// iteratorForCreateAuthors implements pgx.CopyFromSource.
type iteratorForCreateAuthors struct {
	rows                 []int32
	skippedFirstNextCall bool
}

func (r *iteratorForCreateAuthors) Next() bool {
	if len(r.rows) == 0 {
		return false
	}
	if !r.skippedFirstNextCall {
		r.skippedFirstNextCall = true
		return true
	}
	r.rows = r.rows[1:]
	return len(r.rows) > 0
}

func (r iteratorForCreateAuthors) Values() ([]interface{}, error) {
	return []interface{}{
		r.rows[0],
	}, nil
}

func (r iteratorForCreateAuthors) Err() error {
	return nil
}

func (q *Queries) CreateAuthors(ctx context.Context, authorID []int32) (int64, error) {
	return q.db.CopyFrom(ctx, []string{"authors"}, []string{"author_id"}, &iteratorForCreateAuthors{rows: authorID})
}
