// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0

package querytest

import (
	"database/sql"
)

type Author struct {
	ID       int64
	Username sql.NullString
	Email    sql.NullString
	Name     string
	Bio      sql.NullString
}
