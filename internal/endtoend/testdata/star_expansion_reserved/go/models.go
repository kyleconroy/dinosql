// Code generated by sqlc. DO NOT EDIT.

package querytest

import (
	"database/sql"
)

type Foo struct {
	Group sql.NullString
}

func (t *Foo) GetGroup() sql.NullString {
	return t.Group
}
