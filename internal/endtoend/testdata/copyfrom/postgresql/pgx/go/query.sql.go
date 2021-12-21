// Code generated by sqlc. DO NOT EDIT.
// source: query.sql

package querytest

import (
	"context"
	"database/sql"
)

// iteratorForInsertSingleValue implements pgx.CopyFromSource.
type iteratorForInsertSingleValue struct {
	rows                 []sql.NullString
	skippedFirstNextCall bool
}

func (r *iteratorForInsertSingleValue) Next() bool {
	if len(r.rows) == 0 {
		return false
	}
	if !r.skippedFirstNextCall {
		r.skippedFirstNextCall = true
	} else {
		r.rows = r.rows[1:]
	}
	return true
}

func (r iteratorForInsertSingleValue) Values() ([]interface{}, error) {
	return []interface{}{
		r.rows[0],
	}, nil
}

func (r iteratorForInsertSingleValue) Err() error {
	return nil
}

func (q *Queries) InsertSingleValue(ctx context.Context, a []sql.NullString) (int64, error) {
	return q.db.CopyFrom(ctx, []string{"foo"}, []string{"a"}, &iteratorForInsertSingleValue{rows: a})
}

type InsertValuesParams struct {
	A sql.NullString
	B sql.NullInt32
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
	} else {
		r.rows = r.rows[1:]
	}
	return true
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

func (q *Queries) InsertValues(ctx context.Context, arg []InsertValuesParams) (int64, error) {
	return q.db.CopyFrom(ctx, []string{"foo"}, []string{"a", "b"}, &iteratorForInsertValues{rows: arg})
}
