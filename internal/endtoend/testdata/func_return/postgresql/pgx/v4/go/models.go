// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1

package querytest

import (
	"database/sql"
)

type User struct {
	ID        sql.NullInt32
	FirstName string
}
