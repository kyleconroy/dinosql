// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package querytest

import (
	"database/sql"
)

type BazUser struct {
	ID   int64
	Name string
}

type Post struct {
	ID     int64
	UserID int64
}

type User struct {
	ID   int64
	Name string
	Age  sql.NullInt64
}
