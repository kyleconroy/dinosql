// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package querytest

import (
	"github.com/jackc/pgtype"
)

type TestTable struct {
	VCidNull pgtype.CID
	VOidNull pgtype.OID
	VTidNull pgtype.TID
	VXidNull pgtype.XID
	VCid     pgtype.CID
	VOid     pgtype.OID
	VTid     pgtype.TID
	VXid     pgtype.XID
}
