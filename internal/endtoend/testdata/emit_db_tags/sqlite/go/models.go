// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package querytest

import (
	"database/sql"
)

type User struct {
	ID        int64          `db:"id"`
	FirstName string         `db:"first_name"`
	LastName  sql.NullString `db:"last_name"`
	Age       int64          `db:"age"`
}
