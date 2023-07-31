// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package querytest

import (
	"database/sql"
)

type Author struct {
	ID   int64
	Name string
	Age  sql.NullInt64
}

type Book struct {
	ID         int64
	Author     string
	Translator string
	Year       sql.NullInt64
}

type Translator struct {
	ID   int64
	Name string
	Age  sql.NullInt64
}
