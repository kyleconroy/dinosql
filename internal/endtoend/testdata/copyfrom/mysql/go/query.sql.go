// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: query.sql

package querytest

import (
	"database/sql"
)

const insertSingleValue = `-- name: InsertSingleValue :copyfrom
INSERT INTO foo (a) VALUES (?)
`

const insertValues = `-- name: InsertValues :copyfrom
INSERT INTO foo (a, b, c, d) VALUES (?, ?, ?, ?)
`

type InsertValuesParams struct {
	A sql.NullString
	B sql.NullInt32
	C sql.NullTime
	D sql.NullTime
}
