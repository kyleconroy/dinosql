// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package querytest

import (
	"database/sql"
)

type Bar struct {
	ID    uint64
	Title sql.NullString
}

type Foo struct {
	ID uint64
}
