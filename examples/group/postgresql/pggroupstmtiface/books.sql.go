// Code generated by sqlc. DO NOT EDIT.
// source: books.sql

package pggroupstmtiface

import (
	"context"
	"database/sql"
	"time"

	"github.com/lib/pq"
)

type Books interface {
	Create(ctx context.Context, arg BooksCreateParams) (Book, error)
	PrepareCreate(ctx context.Context) (BooksCreateStmt, error)
	Delete(ctx context.Context, bookID int32) error
	PrepareDelete(ctx context.Context) (BooksDeleteStmt, error)
	Get(ctx context.Context, bookID int32) (Book, error)
	PrepareGet(ctx context.Context) (BooksGetStmt, error)
	ListByTags(ctx context.Context, dollar_1 []string) ([]BooksListByTagsRow, error)
	PrepareListByTags(ctx context.Context) (BooksListByTagsStmt, error)
	ListByTitleYear(ctx context.Context, arg BooksListByTitleYearParams) ([]Book, error)
	PrepareListByTitleYear(ctx context.Context) (BooksListByTitleYearStmt, error)
	Update(ctx context.Context, arg BooksUpdateParams) error
	PrepareUpdate(ctx context.Context) (BooksUpdateStmt, error)
	UpdateISBN(ctx context.Context, arg BooksUpdateISBNParams) error
	PrepareUpdateISBN(ctx context.Context) (BooksUpdateISBNStmt, error)
}

func NewBooks(db DBTX) Books {
	return &books{db: db}
}

type books struct {
	db DBTX
}

const booksCreate = `-- name: Create :one
INSERT INTO books (
    author_id,
    isbn,
    book_type,
    title,
    year,
    available,
    tags
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7
)
RETURNING book_id, author_id, isbn, book_type, title, year, available, tags
`

type BooksCreateParams struct {
	AuthorID  int32
	Isbn      string
	BookType  BookType
	Title     string
	Year      int32
	Available time.Time
	Tags      []string
}

func (q *books) Create(ctx context.Context, arg BooksCreateParams) (Book, error) {
	row := q.db.QueryRowContext(ctx, booksCreate,
		arg.AuthorID,
		arg.Isbn,
		arg.BookType,
		arg.Title,
		arg.Year,
		arg.Available,
		pq.Array(arg.Tags),
	)
	var i Book
	err := row.Scan(
		&i.BookID,
		&i.AuthorID,
		&i.Isbn,
		&i.BookType,
		&i.Title,
		&i.Year,
		&i.Available,
		pq.Array(&i.Tags),
	)
	return i, err
}

type BooksCreateStmt interface {
	Close() error
	JoinTx(ctx context.Context, tx *sql.Tx) BooksCreateStmt
	Exec(ctx context.Context, arg BooksCreateParams) (Book, error)
}

type booksCreateStmt struct {
	stmt *sql.Stmt
}

func (s *booksCreateStmt) Close() error {
	return s.stmt.Close()
}

func (s *booksCreateStmt) JoinTx(ctx context.Context, tx *sql.Tx) BooksCreateStmt {
	stmt := tx.StmtContext(ctx, s.stmt)
	return &booksCreateStmt{stmt: stmt}
}

func (s *booksCreateStmt) Exec(ctx context.Context, arg BooksCreateParams) (Book, error) {
	row := s.stmt.QueryRowContext(ctx,
		arg.AuthorID,
		arg.Isbn,
		arg.BookType,
		arg.Title,
		arg.Year,
		arg.Available,
		pq.Array(arg.Tags),
	)
	var i Book
	err := row.Scan(
		&i.BookID,
		&i.AuthorID,
		&i.Isbn,
		&i.BookType,
		&i.Title,
		&i.Year,
		&i.Available,
		pq.Array(&i.Tags),
	)
	return i, err
}

func (q *books) PrepareCreate(ctx context.Context) (BooksCreateStmt, error) {
	stmt, err := q.db.PrepareContext(ctx, booksCreate)
	if err != nil {
		return nil, err
	}
	return &booksCreateStmt{stmt: stmt}, nil
}

const booksDelete = `-- name: Delete :exec
DELETE FROM books
WHERE book_id = $1
`

func (q *books) Delete(ctx context.Context, bookID int32) error {
	_, err := q.db.ExecContext(ctx, booksDelete, bookID)
	return err
}

type BooksDeleteStmt interface {
	Close() error
	JoinTx(ctx context.Context, tx *sql.Tx) BooksDeleteStmt
	Exec(ctx context.Context, bookID int32) error
}

type booksDeleteStmt struct {
	stmt *sql.Stmt
}

func (s *booksDeleteStmt) Close() error {
	return s.stmt.Close()
}

func (s *booksDeleteStmt) JoinTx(ctx context.Context, tx *sql.Tx) BooksDeleteStmt {
	stmt := tx.StmtContext(ctx, s.stmt)
	return &booksDeleteStmt{stmt: stmt}
}

func (s *booksDeleteStmt) Exec(ctx context.Context, bookID int32) error {
	_, err := s.stmt.ExecContext(ctx, bookID)
	return err
}

func (q *books) PrepareDelete(ctx context.Context) (BooksDeleteStmt, error) {
	stmt, err := q.db.PrepareContext(ctx, booksDelete)
	if err != nil {
		return nil, err
	}
	return &booksDeleteStmt{stmt: stmt}, nil
}

const booksGet = `-- name: Get :one
SELECT book_id, author_id, isbn, book_type, title, year, available, tags FROM books
WHERE book_id = $1
`

func (q *books) Get(ctx context.Context, bookID int32) (Book, error) {
	row := q.db.QueryRowContext(ctx, booksGet, bookID)
	var i Book
	err := row.Scan(
		&i.BookID,
		&i.AuthorID,
		&i.Isbn,
		&i.BookType,
		&i.Title,
		&i.Year,
		&i.Available,
		pq.Array(&i.Tags),
	)
	return i, err
}

type BooksGetStmt interface {
	Close() error
	JoinTx(ctx context.Context, tx *sql.Tx) BooksGetStmt
	Exec(ctx context.Context, bookID int32) (Book, error)
}

type booksGetStmt struct {
	stmt *sql.Stmt
}

func (s *booksGetStmt) Close() error {
	return s.stmt.Close()
}

func (s *booksGetStmt) JoinTx(ctx context.Context, tx *sql.Tx) BooksGetStmt {
	stmt := tx.StmtContext(ctx, s.stmt)
	return &booksGetStmt{stmt: stmt}
}

func (s *booksGetStmt) Exec(ctx context.Context, bookID int32) (Book, error) {
	row := s.stmt.QueryRowContext(ctx, bookID)
	var i Book
	err := row.Scan(
		&i.BookID,
		&i.AuthorID,
		&i.Isbn,
		&i.BookType,
		&i.Title,
		&i.Year,
		&i.Available,
		pq.Array(&i.Tags),
	)
	return i, err
}

func (q *books) PrepareGet(ctx context.Context) (BooksGetStmt, error) {
	stmt, err := q.db.PrepareContext(ctx, booksGet)
	if err != nil {
		return nil, err
	}
	return &booksGetStmt{stmt: stmt}, nil
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
WHERE tags && $1::varchar[]
`

type BooksListByTagsRow struct {
	BookID int32
	Title  string
	Name   string
	Isbn   string
	Tags   []string
}

func (q *books) ListByTags(ctx context.Context, dollar_1 []string) ([]BooksListByTagsRow, error) {
	rows, err := q.db.QueryContext(ctx, booksListByTags, pq.Array(dollar_1))
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
			pq.Array(&i.Tags),
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

type BooksListByTagsStmt interface {
	Close() error
	JoinTx(ctx context.Context, tx *sql.Tx) BooksListByTagsStmt
	Exec(ctx context.Context, dollar_1 []string) ([]BooksListByTagsRow, error)
}

type booksListByTagsStmt struct {
	stmt *sql.Stmt
}

func (s *booksListByTagsStmt) Close() error {
	return s.stmt.Close()
}

func (s *booksListByTagsStmt) JoinTx(ctx context.Context, tx *sql.Tx) BooksListByTagsStmt {
	stmt := tx.StmtContext(ctx, s.stmt)
	return &booksListByTagsStmt{stmt: stmt}
}

func (s *booksListByTagsStmt) Exec(ctx context.Context, dollar_1 []string) ([]BooksListByTagsRow, error) {
	rows, err := s.stmt.QueryContext(ctx, pq.Array(dollar_1))
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
			pq.Array(&i.Tags),
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

func (q *books) PrepareListByTags(ctx context.Context) (BooksListByTagsStmt, error) {
	stmt, err := q.db.PrepareContext(ctx, booksListByTags)
	if err != nil {
		return nil, err
	}
	return &booksListByTagsStmt{stmt: stmt}, nil
}

const booksListByTitleYear = `-- name: ListByTitleYear :many
SELECT book_id, author_id, isbn, book_type, title, year, available, tags FROM books
WHERE title = $1 AND year = $2
`

type BooksListByTitleYearParams struct {
	Title string
	Year  int32
}

func (q *books) ListByTitleYear(ctx context.Context, arg BooksListByTitleYearParams) ([]Book, error) {
	rows, err := q.db.QueryContext(ctx, booksListByTitleYear, arg.Title, arg.Year)
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
			&i.Year,
			&i.Available,
			pq.Array(&i.Tags),
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

type BooksListByTitleYearStmt interface {
	Close() error
	JoinTx(ctx context.Context, tx *sql.Tx) BooksListByTitleYearStmt
	Exec(ctx context.Context, arg BooksListByTitleYearParams) ([]Book, error)
}

type booksListByTitleYearStmt struct {
	stmt *sql.Stmt
}

func (s *booksListByTitleYearStmt) Close() error {
	return s.stmt.Close()
}

func (s *booksListByTitleYearStmt) JoinTx(ctx context.Context, tx *sql.Tx) BooksListByTitleYearStmt {
	stmt := tx.StmtContext(ctx, s.stmt)
	return &booksListByTitleYearStmt{stmt: stmt}
}

func (s *booksListByTitleYearStmt) Exec(ctx context.Context, arg BooksListByTitleYearParams) ([]Book, error) {
	rows, err := s.stmt.QueryContext(ctx, arg.Title, arg.Year)
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
			&i.Year,
			&i.Available,
			pq.Array(&i.Tags),
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

func (q *books) PrepareListByTitleYear(ctx context.Context) (BooksListByTitleYearStmt, error) {
	stmt, err := q.db.PrepareContext(ctx, booksListByTitleYear)
	if err != nil {
		return nil, err
	}
	return &booksListByTitleYearStmt{stmt: stmt}, nil
}

const booksUpdate = `-- name: Update :exec
UPDATE books
SET title = $1, tags = $2
WHERE book_id = $3
`

type BooksUpdateParams struct {
	Title  string
	Tags   []string
	BookID int32
}

func (q *books) Update(ctx context.Context, arg BooksUpdateParams) error {
	_, err := q.db.ExecContext(ctx, booksUpdate, arg.Title, pq.Array(arg.Tags), arg.BookID)
	return err
}

type BooksUpdateStmt interface {
	Close() error
	JoinTx(ctx context.Context, tx *sql.Tx) BooksUpdateStmt
	Exec(ctx context.Context, arg BooksUpdateParams) error
}

type booksUpdateStmt struct {
	stmt *sql.Stmt
}

func (s *booksUpdateStmt) Close() error {
	return s.stmt.Close()
}

func (s *booksUpdateStmt) JoinTx(ctx context.Context, tx *sql.Tx) BooksUpdateStmt {
	stmt := tx.StmtContext(ctx, s.stmt)
	return &booksUpdateStmt{stmt: stmt}
}

func (s *booksUpdateStmt) Exec(ctx context.Context, arg BooksUpdateParams) error {
	_, err := s.stmt.ExecContext(ctx, arg.Title, pq.Array(arg.Tags), arg.BookID)
	return err
}

func (q *books) PrepareUpdate(ctx context.Context) (BooksUpdateStmt, error) {
	stmt, err := q.db.PrepareContext(ctx, booksUpdate)
	if err != nil {
		return nil, err
	}
	return &booksUpdateStmt{stmt: stmt}, nil
}

const booksUpdateISBN = `-- name: UpdateISBN :exec
UPDATE books
SET title = $1, tags = $2, isbn = $4
WHERE book_id = $3
`

type BooksUpdateISBNParams struct {
	Title  string
	Tags   []string
	BookID int32
	Isbn   string
}

func (q *books) UpdateISBN(ctx context.Context, arg BooksUpdateISBNParams) error {
	_, err := q.db.ExecContext(ctx, booksUpdateISBN,
		arg.Title,
		pq.Array(arg.Tags),
		arg.BookID,
		arg.Isbn,
	)
	return err
}

type BooksUpdateISBNStmt interface {
	Close() error
	JoinTx(ctx context.Context, tx *sql.Tx) BooksUpdateISBNStmt
	Exec(ctx context.Context, arg BooksUpdateISBNParams) error
}

type booksUpdateISBNStmt struct {
	stmt *sql.Stmt
}

func (s *booksUpdateISBNStmt) Close() error {
	return s.stmt.Close()
}

func (s *booksUpdateISBNStmt) JoinTx(ctx context.Context, tx *sql.Tx) BooksUpdateISBNStmt {
	stmt := tx.StmtContext(ctx, s.stmt)
	return &booksUpdateISBNStmt{stmt: stmt}
}

func (s *booksUpdateISBNStmt) Exec(ctx context.Context, arg BooksUpdateISBNParams) error {
	_, err := s.stmt.ExecContext(ctx,
		arg.Title,
		pq.Array(arg.Tags),
		arg.BookID,
		arg.Isbn,
	)
	return err
}

func (q *books) PrepareUpdateISBN(ctx context.Context) (BooksUpdateISBNStmt, error) {
	stmt, err := q.db.PrepareContext(ctx, booksUpdateISBN)
	if err != nil {
		return nil, err
	}
	return &booksUpdateISBNStmt{stmt: stmt}, nil
}
