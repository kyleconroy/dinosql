// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package querytest

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Memory struct {
	ID        uuid.UUID
	VampireID uuid.UUID
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

type Vampire struct {
	ID uuid.UUID
}
