// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package datatype

import (
	"time"

	"github.com/google/uuid"
)

type Foo struct {
	Bar *time.Time
	Baz *uuid.UUID
}
