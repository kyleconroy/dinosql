// Code generated by sqlc. DO NOT EDIT.
// source: books.sql

package mygroupstmt

import (
	"context"
	"database/sql"
	"time"
)

type Books struct {
	db DBTX
}

func NewBooks(db DBTX) *Books {
	return &Books{db: db}
}

const booksCreate = `-- name: Create :execresult
INSERT INTO books (
    author_id,
    isbn,
    book_type,
    title,
    yr,
    available,
    tags
) VALUES (
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?
)
`

type BooksCreateParams struct {
	AuthorID  int32
	Isbn      string
	BookType  BooksBookType
	Title     string
	Yr        int32
	Available time.Time
	Tags      string
}

func (q *Books) Create(ctx context.Context, arg BooksCreateParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, booksCreate,
		arg.AuthorID,
		arg.Isbn,
		arg.BookType,
		arg.Title,
		arg.Yr,
		arg.Available,
		arg.Tags,
	)
}

type BooksCreateStmt struct {
	stmt *sql.Stmt
}

func (s *BooksCreateStmt) Close() error {
	return s.stmt.Close()
}

func (s *BooksCreateStmt) JoinTx(ctx context.Context, tx *sql.Tx) *BooksCreateStmt {
	stmt := tx.StmtContext(ctx, s.stmt)
	return &BooksCreateStmt{stmt: stmt}
}

func (s *BooksCreateStmt) Exec(ctx context.Context, arg BooksCreateParams) (sql.Result, error) {
	return s.stmt.ExecContext(ctx,
		arg.AuthorID,
		arg.Isbn,
		arg.BookType,
		arg.Title,
		arg.Yr,
		arg.Available,
		arg.Tags,
	)
}

func (q *Books) PrepareCreate(ctx context.Context) (*BooksCreateStmt, error) {
	stmt, err := q.db.PrepareContext(ctx, booksCreate)
	if err != nil {
		return nil, err
	}
	return &BooksCreateStmt{stmt: stmt}, nil
}

const booksDelete = `-- name: Delete :exec
DELETE FROM books
WHERE book_id = ?
`

func (q *Books) Delete(ctx context.Context, bookID int32) error {
	_, err := q.db.ExecContext(ctx, booksDelete, bookID)
	return err
}

type BooksDeleteStmt struct {
	stmt *sql.Stmt
}

func (s *BooksDeleteStmt) Close() error {
	return s.stmt.Close()
}

func (s *BooksDeleteStmt) JoinTx(ctx context.Context, tx *sql.Tx) *BooksDeleteStmt {
	stmt := tx.StmtContext(ctx, s.stmt)
	return &BooksDeleteStmt{stmt: stmt}
}

func (s *BooksDeleteStmt) Exec(ctx context.Context, bookID int32) error {
	_, err := s.stmt.ExecContext(ctx, bookID)
	return err
}

func (q *Books) PrepareDelete(ctx context.Context) (*BooksDeleteStmt, error) {
	stmt, err := q.db.PrepareContext(ctx, booksDelete)
	if err != nil {
		return nil, err
	}
	return &BooksDeleteStmt{stmt: stmt}, nil
}

const booksDeleteAuthorBeforeYear = `-- name: DeleteAuthorBeforeYear :exec
DELETE FROM books
WHERE yr < ? AND author_id = ?
`

type BooksDeleteAuthorBeforeYearParams struct {
	Yr       int32
	AuthorID int32
}

func (q *Books) DeleteAuthorBeforeYear(ctx context.Context, arg BooksDeleteAuthorBeforeYearParams) error {
	_, err := q.db.ExecContext(ctx, booksDeleteAuthorBeforeYear, arg.Yr, arg.AuthorID)
	return err
}

type BooksDeleteAuthorBeforeYearStmt struct {
	stmt *sql.Stmt
}

func (s *BooksDeleteAuthorBeforeYearStmt) Close() error {
	return s.stmt.Close()
}

func (s *BooksDeleteAuthorBeforeYearStmt) JoinTx(ctx context.Context, tx *sql.Tx) *BooksDeleteAuthorBeforeYearStmt {
	stmt := tx.StmtContext(ctx, s.stmt)
	return &BooksDeleteAuthorBeforeYearStmt{stmt: stmt}
}

func (s *BooksDeleteAuthorBeforeYearStmt) Exec(ctx context.Context, arg BooksDeleteAuthorBeforeYearParams) error {
	_, err := s.stmt.ExecContext(ctx, arg.Yr, arg.AuthorID)
	return err
}

func (q *Books) PrepareDeleteAuthorBeforeYear(ctx context.Context) (*BooksDeleteAuthorBeforeYearStmt, error) {
	stmt, err := q.db.PrepareContext(ctx, booksDeleteAuthorBeforeYear)
	if err != nil {
		return nil, err
	}
	return &BooksDeleteAuthorBeforeYearStmt{stmt: stmt}, nil
}

const booksGet = `-- name: Get :one
SELECT book_id, author_id, isbn, book_type, title, yr, available, tags FROM books
WHERE book_id = ?
`

func (q *Books) Get(ctx context.Context, bookID int32) (Book, error) {
	row := q.db.QueryRowContext(ctx, booksGet, bookID)
	var i Book
	err := row.Scan(
		&i.BookID,
		&i.AuthorID,
		&i.Isbn,
		&i.BookType,
		&i.Title,
		&i.Yr,
		&i.Available,
		&i.Tags,
	)
	return i, err
}

type BooksGetStmt struct {
	stmt *sql.Stmt
}

func (s *BooksGetStmt) Close() error {
	return s.stmt.Close()
}

func (s *BooksGetStmt) JoinTx(ctx context.Context, tx *sql.Tx) *BooksGetStmt {
	stmt := tx.StmtContext(ctx, s.stmt)
	return &BooksGetStmt{stmt: stmt}
}

func (s *BooksGetStmt) Exec(ctx context.Context, bookID int32) (Book, error) {
	row := s.stmt.QueryRowContext(ctx, bookID)
	var i Book
	err := row.Scan(
		&i.BookID,
		&i.AuthorID,
		&i.Isbn,
		&i.BookType,
		&i.Title,
		&i.Yr,
		&i.Available,
		&i.Tags,
	)
	return i, err
}

func (q *Books) PrepareGet(ctx context.Context) (*BooksGetStmt, error) {
	stmt, err := q.db.PrepareContext(ctx, booksGet)
	if err != nil {
		return nil, err
	}
	return &BooksGetStmt{stmt: stmt}, nil
}

const booksListByTags = `-- name: ListByTags :many
SELECT
  book_id,
  title,
  name,
  isbn,
  tags
FROM books
LEFT JOIN authors ON books.author_id = authors.author_id
WHERE tags = ?
`

type BooksListByTagsRow struct {
	BookID int32
	Title  string
	Name   string
	Isbn   string
	Tags   string
}

func (q *Books) ListByTags(ctx context.Context, tags string) ([]BooksListByTagsRow, error) {
	rows, err := q.db.QueryContext(ctx, booksListByTags, tags)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []BooksListByTagsRow{}
	for rows.Next() {
		var i BooksListByTagsRow
		if err := rows.Scan(
			&i.BookID,
			&i.Title,
			&i.Name,
			&i.Isbn,
			&i.Tags,
		); err != nil {
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

type BooksListByTagsStmt struct {
	stmt *sql.Stmt
}

func (s *BooksListByTagsStmt) Close() error {
	return s.stmt.Close()
}

func (s *BooksListByTagsStmt) JoinTx(ctx context.Context, tx *sql.Tx) *BooksListByTagsStmt {
	stmt := tx.StmtContext(ctx, s.stmt)
	return &BooksListByTagsStmt{stmt: stmt}
}

func (s *BooksListByTagsStmt) Exec(ctx context.Context, tags string) ([]BooksListByTagsRow, error) {
	rows, err := s.stmt.QueryContext(ctx, tags)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []BooksListByTagsRow{}
	for rows.Next() {
		var i BooksListByTagsRow
		if err := rows.Scan(
			&i.BookID,
			&i.Title,
			&i.Name,
			&i.Isbn,
			&i.Tags,
		); err != nil {
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

func (q *Books) PrepareListByTags(ctx context.Context) (*BooksListByTagsStmt, error) {
	stmt, err := q.db.PrepareContext(ctx, booksListByTags)
	if err != nil {
		return nil, err
	}
	return &BooksListByTagsStmt{stmt: stmt}, nil
}

const booksListByTitleYear = `-- name: ListByTitleYear :many
SELECT book_id, author_id, isbn, book_type, title, yr, available, tags FROM books
WHERE title = ? AND yr = ?
`

type BooksListByTitleYearParams struct {
	Title string
	Yr    int32
}

func (q *Books) ListByTitleYear(ctx context.Context, arg BooksListByTitleYearParams) ([]Book, error) {
	rows, err := q.db.QueryContext(ctx, booksListByTitleYear, arg.Title, arg.Yr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Book{}
	for rows.Next() {
		var i Book
		if err := rows.Scan(
			&i.BookID,
			&i.AuthorID,
			&i.Isbn,
			&i.BookType,
			&i.Title,
			&i.Yr,
			&i.Available,
			&i.Tags,
		); err != nil {
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

type BooksListByTitleYearStmt struct {
	stmt *sql.Stmt
}

func (s *BooksListByTitleYearStmt) Close() error {
	return s.stmt.Close()
}

func (s *BooksListByTitleYearStmt) JoinTx(ctx context.Context, tx *sql.Tx) *BooksListByTitleYearStmt {
	stmt := tx.StmtContext(ctx, s.stmt)
	return &BooksListByTitleYearStmt{stmt: stmt}
}

func (s *BooksListByTitleYearStmt) Exec(ctx context.Context, arg BooksListByTitleYearParams) ([]Book, error) {
	rows, err := s.stmt.QueryContext(ctx, arg.Title, arg.Yr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Book{}
	for rows.Next() {
		var i Book
		if err := rows.Scan(
			&i.BookID,
			&i.AuthorID,
			&i.Isbn,
			&i.BookType,
			&i.Title,
			&i.Yr,
			&i.Available,
			&i.Tags,
		); err != nil {
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

func (q *Books) PrepareListByTitleYear(ctx context.Context) (*BooksListByTitleYearStmt, error) {
	stmt, err := q.db.PrepareContext(ctx, booksListByTitleYear)
	if err != nil {
		return nil, err
	}
	return &BooksListByTitleYearStmt{stmt: stmt}, nil
}

const booksUpdate = `-- name: Update :exec
UPDATE books
SET title = ?, tags = ?
WHERE book_id = ?
`

type BooksUpdateParams struct {
	Title  string
	Tags   string
	BookID int32
}

func (q *Books) Update(ctx context.Context, arg BooksUpdateParams) error {
	_, err := q.db.ExecContext(ctx, booksUpdate, arg.Title, arg.Tags, arg.BookID)
	return err
}

type BooksUpdateStmt struct {
	stmt *sql.Stmt
}

func (s *BooksUpdateStmt) Close() error {
	return s.stmt.Close()
}

func (s *BooksUpdateStmt) JoinTx(ctx context.Context, tx *sql.Tx) *BooksUpdateStmt {
	stmt := tx.StmtContext(ctx, s.stmt)
	return &BooksUpdateStmt{stmt: stmt}
}

func (s *BooksUpdateStmt) Exec(ctx context.Context, arg BooksUpdateParams) error {
	_, err := s.stmt.ExecContext(ctx, arg.Title, arg.Tags, arg.BookID)
	return err
}

func (q *Books) PrepareUpdate(ctx context.Context) (*BooksUpdateStmt, error) {
	stmt, err := q.db.PrepareContext(ctx, booksUpdate)
	if err != nil {
		return nil, err
	}
	return &BooksUpdateStmt{stmt: stmt}, nil
}

const booksUpdateISBN = `-- name: UpdateISBN :exec
UPDATE books
SET title = ?, tags = ?, isbn = ?
WHERE book_id = ?
`

type BooksUpdateISBNParams struct {
	Title  string
	Tags   string
	Isbn   string
	BookID int32
}

func (q *Books) UpdateISBN(ctx context.Context, arg BooksUpdateISBNParams) error {
	_, err := q.db.ExecContext(ctx, booksUpdateISBN,
		arg.Title,
		arg.Tags,
		arg.Isbn,
		arg.BookID,
	)
	return err
}

type BooksUpdateISBNStmt struct {
	stmt *sql.Stmt
}

func (s *BooksUpdateISBNStmt) Close() error {
	return s.stmt.Close()
}

func (s *BooksUpdateISBNStmt) JoinTx(ctx context.Context, tx *sql.Tx) *BooksUpdateISBNStmt {
	stmt := tx.StmtContext(ctx, s.stmt)
	return &BooksUpdateISBNStmt{stmt: stmt}
}

func (s *BooksUpdateISBNStmt) Exec(ctx context.Context, arg BooksUpdateISBNParams) error {
	_, err := s.stmt.ExecContext(ctx,
		arg.Title,
		arg.Tags,
		arg.Isbn,
		arg.BookID,
	)
	return err
}

func (q *Books) PrepareUpdateISBN(ctx context.Context) (*BooksUpdateISBNStmt, error) {
	stmt, err := q.db.PrepareContext(ctx, booksUpdateISBN)
	if err != nil {
		return nil, err
	}
	return &BooksUpdateISBNStmt{stmt: stmt}, nil
}
