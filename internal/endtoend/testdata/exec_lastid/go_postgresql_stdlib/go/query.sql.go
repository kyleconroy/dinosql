// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package querytest

import (
	"context"
)

const insertBar = `-- name: InsertBar :execlastid
INSERT INTO bar () VALUES ()
`

func (q *Queries) InsertBar(ctx context.Context) (int64, error) {
	result, err := q.db.ExecContext(ctx, insertBar)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}
