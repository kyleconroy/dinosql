// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package querytest

import (
	"database/sql"
)

type Author struct {
	ID       int32
	Name     string
	ParentID sql.NullInt32
}
