// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package querytest

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.funcParamIdentStmt, err = db.PrepareContext(ctx, funcParamIdent); err != nil {
		return nil, fmt.Errorf("error preparing query FuncParamIdent: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.funcParamIdentStmt != nil {
		if cerr := q.funcParamIdentStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing funcParamIdentStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                 DBTX
	tx                 *sql.Tx
	funcParamIdentStmt *sql.Stmt
}

func (q *Queries) PrepareFuncParamIdent(ctx context.Context) error {
	var err error
	if q.funcParamIdentStmt, err = q.db.PrepareContext(ctx, funcParamIdent); err != nil {
		return fmt.Errorf("error preparing query FuncParamIdent: %w", err)
	}
	return nil
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                 tx,
		tx:                 tx,
		funcParamIdentStmt: q.funcParamIdentStmt,
	}
}
