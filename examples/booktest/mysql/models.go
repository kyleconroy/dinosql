// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0

package booktest

import (
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
