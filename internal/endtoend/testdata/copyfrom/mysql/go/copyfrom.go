// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: copyfrom.go

package querytest

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"sync/atomic"

	"github.com/go-sql-driver/mysql"
	"github.com/hexon/mysqltsv"
)

var readerHandlerSequenceForInsertSingleValue uint32 = 1

func convertRowsForInsertSingleValue(w *io.PipeWriter, a []sql.NullString) {
	e := mysqltsv.NewEncoder(w, 1, nil)
	for _, row := range a {
		e.AppendValue(row)
	}
	w.CloseWithError(e.Close())
}

// InsertSingleValue uses MySQL's LOAD DATA LOCAL INFILE and is not atomic. Errors and duplicate keys are treated as warnings and insertion will continue, even without an error for some cases.
// Use this in a transaction and use SHOW WARNINGS to check for any problems and roll back if you want to.
// This is a MySQL limitation, not sqlc. Check the documentation for more information: https://dev.mysql.com/doc/refman/8.0/en/load-data.html#load-data-error-handling
func (q *Queries) InsertSingleValue(ctx context.Context, a []sql.NullString) (int64, error) {
	pr, pw := io.Pipe()
	defer pr.Close()
	rh := fmt.Sprintf("InsertSingleValue_%d", atomic.AddUint32(&readerHandlerSequenceForInsertSingleValue, 1))
	mysql.RegisterReaderHandler(rh, func() io.Reader { return pr })
	defer mysql.DeregisterReaderHandler(rh)
	go convertRowsForInsertSingleValue(pw, a)
	result, err := q.db.ExecContext(ctx, fmt.Sprintf("LOAD DATA LOCAL INFILE '%s' INTO TABLE `foo` %s (a)", "Reader::"+rh, mysqltsv.Escaping))
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

var readerHandlerSequenceForInsertValues uint32 = 1

func convertRowsForInsertValues(w *io.PipeWriter, arg []InsertValuesParams) {
	e := mysqltsv.NewEncoder(w, 4, nil)
	for _, row := range arg {
		e.AppendValue(row.A)
		e.AppendValue(row.B)
		e.AppendValue(row.C)
		e.AppendValue(row.D)
	}
	w.CloseWithError(e.Close())
}

// InsertValues uses MySQL's LOAD DATA LOCAL INFILE and is not atomic. Errors and duplicate keys are treated as warnings and insertion will continue, even without an error for some cases.
// Use this in a transaction and use SHOW WARNINGS to check for any problems and roll back if you want to.
// This is a MySQL limitation, not sqlc. Check the documentation for more information: https://dev.mysql.com/doc/refman/8.0/en/load-data.html#load-data-error-handling
func (q *Queries) InsertValues(ctx context.Context, arg []InsertValuesParams) (int64, error) {
	pr, pw := io.Pipe()
	defer pr.Close()
	rh := fmt.Sprintf("InsertValues_%d", atomic.AddUint32(&readerHandlerSequenceForInsertValues, 1))
	mysql.RegisterReaderHandler(rh, func() io.Reader { return pr })
	defer mysql.DeregisterReaderHandler(rh)
	go convertRowsForInsertValues(pw, arg)
	result, err := q.db.ExecContext(ctx, fmt.Sprintf("LOAD DATA LOCAL INFILE '%s' INTO TABLE `foo` %s (a, b, c, d)", "Reader::"+rh, mysqltsv.Escaping))
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
