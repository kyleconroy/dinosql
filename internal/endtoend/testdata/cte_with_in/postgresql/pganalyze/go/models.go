// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0

package querytest

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type L struct {
	ID       int64
	ParentID pgtype.Int4
}

type T struct {
	ID  int64
	LID pgtype.Int4
	F   pgtype.Text
}
