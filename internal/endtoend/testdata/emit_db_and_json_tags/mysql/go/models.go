// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0

package querytest

import (
	"database/sql"
)

type User struct {
	ID        int32          `db:"id" json:"id"`
	FirstName string         `db:"first_name" json:"first_name"`
	LastName  sql.NullString `db:"last_name" json:"last_name"`
	Age       int32          `db:"age" json:"age"`
}
