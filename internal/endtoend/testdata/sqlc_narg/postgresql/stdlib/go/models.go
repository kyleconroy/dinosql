// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package querytest

import (
	"database/sql"
)

type Foo struct {
	Bar      string
	MaybeBar sql.NullString
}
