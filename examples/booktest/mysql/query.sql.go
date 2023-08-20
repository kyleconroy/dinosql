// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: query.sql

package booktest

import (
	"context"
	"database/sql"
	"time"
)

const booksByTags = `-- name: BooksByTags :many
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

type BooksByTagsRow struct {
	BookID int32
	Title  string
	Name   sql.NullString
	Isbn   string
	Tags   string
}

func (q *Queries) BooksByTags(ctx context.Context, tags string, aq ...AdditionalQuery) ([]BooksByTagsRow, error) {
	query := booksByTags
	queryParams := []interface{}{tags}

	if len(aq) > 0 {
		query += " " + aq[0].SQL
		queryParams = append(queryParams, aq[0].Args...)
	}

	rows, err := q.db.QueryContext(ctx, query, queryParams...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []BooksByTagsRow
	for rows.Next() {
		var i BooksByTagsRow
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

const booksByTitleYear = `-- name: BooksByTitleYear :many
SELECT book_id, author_id, isbn, book_type, title, yr, available, tags FROM books
WHERE title = ? AND yr = ?
`

type BooksByTitleYearParams struct {
	Title string
	Yr    int32
}

func (q *Queries) BooksByTitleYear(ctx context.Context, arg BooksByTitleYearParams, aq ...AdditionalQuery) ([]Book, error) {
	query := booksByTitleYear
	queryParams := []interface{}{arg.Title, arg.Yr}

	if len(aq) > 0 {
		query += " " + aq[0].SQL
		queryParams = append(queryParams, aq[0].Args...)
	}

	rows, err := q.db.QueryContext(ctx, query, queryParams...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Book
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

const createAuthor = `-- name: CreateAuthor :execresult
INSERT INTO authors (name) VALUES (?)
`

func (q *Queries) CreateAuthor(ctx context.Context, name string) (sql.Result, error) {
	query := createAuthor
	queryParams := []interface{}{name}

	return q.db.ExecContext(ctx, query, queryParams...)
}

const createBook = `-- name: CreateBook :execresult
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

type CreateBookParams struct {
	AuthorID  int32
	Isbn      string
	BookType  BooksBookType
	Title     string
	Yr        int32
	Available time.Time
	Tags      string
}

func (q *Queries) CreateBook(ctx context.Context, arg CreateBookParams) (sql.Result, error) {
	query := createBook
	queryParams := []interface{}{
		arg.AuthorID,
		arg.Isbn,
		arg.BookType,
		arg.Title,
		arg.Yr,
		arg.Available,
		arg.Tags,
	}

	return q.db.ExecContext(ctx, query, queryParams...)
}

const deleteAuthorBeforeYear = `-- name: DeleteAuthorBeforeYear :exec
DELETE FROM books
WHERE yr < ? AND author_id = ?
`

type DeleteAuthorBeforeYearParams struct {
	Yr       int32
	AuthorID int32
}

func (q *Queries) DeleteAuthorBeforeYear(ctx context.Context, arg DeleteAuthorBeforeYearParams) error {
	query := deleteAuthorBeforeYear
	queryParams := []interface{}{arg.Yr, arg.AuthorID}

	_, err := q.db.ExecContext(ctx, query, queryParams...)
	return err
}

const deleteBook = `-- name: DeleteBook :exec
DELETE FROM books
WHERE book_id = ?
`

func (q *Queries) DeleteBook(ctx context.Context, bookID int32) error {
	query := deleteBook
	queryParams := []interface{}{bookID}

	_, err := q.db.ExecContext(ctx, query, queryParams...)
	return err
}

const getAuthor = `-- name: GetAuthor :one
SELECT author_id, name FROM authors
WHERE author_id = ?
`

func (q *Queries) GetAuthor(ctx context.Context, authorID int32, aq ...AdditionalQuery) (Author, error) {
	query := getAuthor
	queryParams := []interface{}{authorID}

	if len(aq) > 0 {
		query += " " + aq[0].SQL
		queryParams = append(queryParams, aq[0].Args...)
	}

	row := q.db.QueryRowContext(ctx, query, queryParams...)
	var i Author
	err := row.Scan(&i.AuthorID, &i.Name)
	return i, err
}

const getBook = `-- name: GetBook :one
SELECT book_id, author_id, isbn, book_type, title, yr, available, tags FROM books
WHERE book_id = ?
`

func (q *Queries) GetBook(ctx context.Context, bookID int32, aq ...AdditionalQuery) (Book, error) {
	query := getBook
	queryParams := []interface{}{bookID}

	if len(aq) > 0 {
		query += " " + aq[0].SQL
		queryParams = append(queryParams, aq[0].Args...)
	}

	row := q.db.QueryRowContext(ctx, query, queryParams...)
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

const updateBook = `-- name: UpdateBook :exec
UPDATE books
SET title = ?, tags = ?
WHERE book_id = ?
`

type UpdateBookParams struct {
	Title  string
	Tags   string
	BookID int32
}

func (q *Queries) UpdateBook(ctx context.Context, arg UpdateBookParams) error {
	query := updateBook
	queryParams := []interface{}{arg.Title, arg.Tags, arg.BookID}

	_, err := q.db.ExecContext(ctx, query, queryParams...)
	return err
}

const updateBookISBN = `-- name: UpdateBookISBN :exec
UPDATE books
SET title = ?, tags = ?, isbn = ?
WHERE book_id = ?
`

type UpdateBookISBNParams struct {
	Title  string
	Tags   string
	Isbn   string
	BookID int32
}

func (q *Queries) UpdateBookISBN(ctx context.Context, arg UpdateBookISBNParams) error {
	query := updateBookISBN
	queryParams := []interface{}{
		arg.Title,
		arg.Tags,
		arg.Isbn,
		arg.BookID,
	}

	_, err := q.db.ExecContext(ctx, query, queryParams...)
	return err
}
