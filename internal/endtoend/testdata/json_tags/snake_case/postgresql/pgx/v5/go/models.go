// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package querytest

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type User struct {
	FirstName pgtype.Text `json:"first_name"`
	LastName  pgtype.Text `json:"last_name"`
	Age       pgtype.Int2 `json:"age"`
}
