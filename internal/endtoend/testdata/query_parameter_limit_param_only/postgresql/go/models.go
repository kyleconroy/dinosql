// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package querytest

import (
	"database/sql"
	"time"
)

type Notice struct {
	ID        int32
	Cnt       int32
	Status    string
	NoticeAt  sql.NullTime
	CreatedAt time.Time
}
