// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.0

package booktest

import (
	"time"
)

type Author struct {
	AuthorID int64
	Name     string
}

type Book struct {
	BookID    int64
	AuthorID  int64
	Isbn      string
	BookType  string
	Title     string
	Yr        int64
	Available time.Time
	Tag       string
}
