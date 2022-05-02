// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0

package booktest

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type BooksBookType string

const (
	BooksBookTypeFICTION    BooksBookType = "FICTION"
	BooksBookTypeNONFICTION BooksBookType = "NONFICTION"
)

func (e *BooksBookType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = BooksBookType(s)
	case string:
		*e = BooksBookType(s)
	default:
		return fmt.Errorf("unsupported scan type for BooksBookType: %T", src)
	}
	return nil
}

// NullBooksBookType is the nullable version of BooksBookType.
type NullBooksBookType struct {
	BooksBookType BooksBookType
	Valid         bool
}

func (e *NullBooksBookType) Scan(src interface{}) error {
	if src == nil {
		e.Valid = false
		return nil
	}
	switch s := src.(type) {
	case []byte:
		e.BooksBookType = BooksBookType(s)
	case string:
		e.BooksBookType = BooksBookType(s)
	default:
		return fmt.Errorf("unsupported scan type for NullBooksBookType: %T", src)
	}
	e.Valid = len(e.BooksBookType) > 0
	return nil
}

func (e *NullBooksBookType) Value() (driver.Value, error) {
	if !e.Valid {
		return nil, nil
	}
	return string(e.BooksBookType), nil
}

type Author struct {
	AuthorID int32
	Name     string
}

type Book struct {
	BookID    int32
	AuthorID  int32
	Isbn      string
	BookType  BooksBookType
	Title     string
	Yr        int32
	Available time.Time
	Tags      string
}
