// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: copyfrom.go

package querytest

import (
	"context"
	"fmt"
	"io"
	"sync/atomic"

	"github.com/go-sql-driver/mysql"
	"github.com/hexon/mysqltsv"
)

var readerHandlerSequenceForBulkInsert uint32 = 1

func convertRowsForBulkInsert(w *io.PipeWriter, arg []BulkInsertParams) {
	e := mysqltsv.NewEncoder(w, 2, nil)
	for _, row := range arg {
		e.AppendBytes(row.A)
		e.AppendBytes(row.B)
	}
	w.CloseWithError(e.Close())
}

// BulkInsert uses MySQL's LOAD DATA LOCAL INFILE and is not atomic.
//
// Errors and duplicate keys are treated as warnings and insertion will
// continue, even without an error for some cases.  Use this in a transaction
// and use SHOW WARNINGS to check for any problems and roll back if you want to.
//
// Check the documentation for more information:
// https://dev.mysql.com/doc/refman/8.0/en/load-data.html#load-data-error-handling
func (q *Queries) BulkInsert(ctx context.Context, arg []BulkInsertParams) (int64, error) {
	pr, pw := io.Pipe()
	defer pr.Close()
	rh := fmt.Sprintf("BulkInsert_%d", atomic.AddUint32(&readerHandlerSequenceForBulkInsert, 1))
	mysql.RegisterReaderHandler(rh, func() io.Reader { return pr })
	defer mysql.DeregisterReaderHandler(rh)
	go convertRowsForBulkInsert(pw, arg)
	// The string interpolation is necessary because LOAD DATA INFILE requires
	// the file name to be given as a literal string.
	result, err := q.db.ExecContext(ctx, fmt.Sprintf("LOAD DATA LOCAL INFILE '%s' INTO TABLE `foo` %s (a, b)", "Reader::"+rh, mysqltsv.Escaping))
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
