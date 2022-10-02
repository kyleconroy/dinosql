// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: copyfrom.go

package querytest

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

// iteratorForInsertSingleValue implements pgx.CopyFromSource.
type iteratorForInsertSingleValue struct {
	rows                 []pgtype.Text
	skippedFirstNextCall bool
}

func (r *iteratorForInsertSingleValue) Next() bool {
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

func (r iteratorForInsertSingleValue) Values() ([]interface{}, error) {
	return []interface{}{
		r.rows[0],
	}, nil
}

func (r iteratorForInsertSingleValue) Err() error {
	return nil
}

// InsertSingleValue inserts a single value using copy.
func (q *Queries) InsertSingleValue(ctx context.Context, a []pgtype.Text) (int64, error) {
	return q.db.CopyFrom(ctx, []string{"myschema", "foo"}, []string{"a"}, &iteratorForInsertSingleValue{rows: a})
}

// iteratorForInsertValues implements pgx.CopyFromSource.
type iteratorForInsertValues struct {
	rows                 []InsertValuesParams
	skippedFirstNextCall bool
}

func (r *iteratorForInsertValues) Next() bool {
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

func (r iteratorForInsertValues) Values() ([]interface{}, error) {
	return []interface{}{
		r.rows[0].A,
		r.rows[0].B,
	}, nil
}

func (r iteratorForInsertValues) Err() error {
	return nil
}

// InsertValues inserts multiple values using copy.
func (q *Queries) InsertValues(ctx context.Context, arg []InsertValuesParams) (int64, error) {
	return q.db.CopyFrom(ctx, []string{"myschema", "foo"}, []string{"a", "b"}, &iteratorForInsertValues{rows: arg})
}
