// Code generated by sqlc. DO NOT EDIT.

package querytest

import (
	"database/sql"
)

type Bar struct {
	ID    int32
	Title sql.NullString
}

type Foo struct {
	ID int32
}
