// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0

package batch

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type BookType string

const (
	BookTypeFICTION    BookType = "FICTION"
	BookTypeNONFICTION BookType = "NONFICTION"
)

func (e *BookType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = BookType(s)
	case string:
		*e = BookType(s)
	default:
		return fmt.Errorf("unsupported scan type for BookType: %T", src)
	}
	return nil
}

// NullBookType is the nullable version of BookType.
type NullBookType struct {
	BookType BookType
	Valid    bool
}

func (e *NullBookType) Scan(src interface{}) error {
	if src == nil {
		e.Valid = false
		return nil
	}
	switch s := src.(type) {
	case []byte:
		e.BookType = BookType(s)
	case string:
		e.BookType = BookType(s)
	default:
		return fmt.Errorf("unsupported scan type for NullBookType: %T", src)
	}
	e.Valid = len(e.BookType) > 0
	return nil
}

func (e *NullBookType) Value() (driver.Value, error) {
	if !e.Valid {
		return nil, nil
	}
	return string(e.BookType), nil
}

type Author struct {
	AuthorID int32  `json:"author_id"`
	Name     string `json:"name"`
}

type Book struct {
	BookID    int32     `json:"book_id"`
	AuthorID  int32     `json:"author_id"`
	Isbn      string    `json:"isbn"`
	BookType  BookType  `json:"book_type"`
	Title     string    `json:"title"`
	Year      int32     `json:"year"`
	Available time.Time `json:"available"`
	Tags      []string  `json:"tags"`
}
