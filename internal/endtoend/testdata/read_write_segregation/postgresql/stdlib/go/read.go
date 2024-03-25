// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package querytest

import (
	"context"
	"database/sql"
	"fmt"
)

func PrepareReadOnly(ctx context.Context, db DBTX) (*ReadQueries, error) {
	q := ReadQueries{db: db}
	var err error
	if q.getUserByIDStmt, err = db.PrepareContext(ctx, getUserByID); err != nil {
		return nil, fmt.Errorf("error preparing query GetUserByID: %w", err)
	}
	if q.listUsersStmt, err = db.PrepareContext(ctx, listUsers); err != nil {
		return nil, fmt.Errorf("error preparing query ListUsers: %w", err)
	}
	return &q, nil
}

type ReadQueries struct {
	db              DBTX
	getUserByIDStmt *sql.Stmt
	listUsersStmt   *sql.Stmt
}

func (q *ReadQueries) Close() error {
	var err error
	if q.getUserByIDStmt != nil {
		if cerr := q.getUserByIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserByIDStmt: %w", cerr)
		}
	}
	if q.listUsersStmt != nil {
		if cerr := q.listUsersStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listUsersStmt: %w", cerr)
		}
	}
	return err
}

func (q *ReadQueries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *ReadQueries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

func (q *ReadQueries) GetUserByID(ctx context.Context, targetID int32) (GetUserByIDRow, error) {
	row := q.queryRow(ctx, q.getUserByIDStmt, getUserByID, targetID)
	var i GetUserByIDRow
	err := row.Scan(&i.FirstName, &i.ID, &i.LastName)
	return i, err
}

func (q *ReadQueries) ListUsers(ctx context.Context) ([]ListUsersRow, error) {
	rows, err := q.query(ctx, q.listUsersStmt, listUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListUsersRow
	for rows.Next() {
		var i ListUsersRow
		if err := rows.Scan(&i.FirstName, &i.LastName); err != nil {
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
