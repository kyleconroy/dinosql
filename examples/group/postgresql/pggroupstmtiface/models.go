// Code generated by sqlc. DO NOT EDIT.

package pggroupstmtiface

import (
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

type Author struct {
	AuthorID int32
	Name     string
}

type Book struct {
	BookID    int32
	AuthorID  int32
	Isbn      string
	BookType  BookType
	Title     string
	Year      int32
	Available time.Time
	Tags      []string
}
