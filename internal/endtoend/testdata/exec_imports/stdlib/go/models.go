// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package querytest

import (
	"database/sql"
)

type Foo struct {
	Bar  sql.NullInt32
	Bars int32
}
