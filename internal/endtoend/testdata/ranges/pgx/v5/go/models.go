// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1

package querytest

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type TestTable struct {
	VDaterangeNull      pgtype.Range[pgtype.Date]
	VDatemultirangeNull pgtype.Multirange[pgtype.Range[pgtype.Date]]
	VTsrangeNull        pgtype.Range[pgtype.Timestamp]
	VTsmultirangeNull   pgtype.Multirange[pgtype.Range[pgtype.Timestamp]]
	VTstzrangeNull      pgtype.Range[pgtype.Timestamptz]
	VTstzmultirangeNull pgtype.Multirange[pgtype.Range[pgtype.Timestamptz]]
	VNumrangeNull       pgtype.Range[pgtype.Numeric]
	VNummultirangeNull  pgtype.Multirange[pgtype.Range[pgtype.Numeric]]
	VInt4rangeNull      pgtype.Range[pgtype.Int4]
	VInt4multirangeNull pgtype.Multirange[pgtype.Range[pgtype.Int4]]
	VInt8rangeNull      pgtype.Range[pgtype.Int8]
	VInt8multirangeNull pgtype.Multirange[pgtype.Range[pgtype.Int8]]
	VDaterange          pgtype.Range[pgtype.Date]
	VDatemultirange     pgtype.Multirange[pgtype.Range[pgtype.Date]]
	VTsrange            pgtype.Range[pgtype.Timestamp]
	VTsmultirange       pgtype.Multirange[pgtype.Range[pgtype.Timestamp]]
	VTstzrange          pgtype.Range[pgtype.Timestamptz]
	VTstzmultirange     pgtype.Multirange[pgtype.Range[pgtype.Timestamptz]]
	VNumrange           pgtype.Range[pgtype.Numeric]
	VNummultirange      pgtype.Multirange[pgtype.Range[pgtype.Numeric]]
	VInt4range          pgtype.Range[pgtype.Int4]
	VInt4multirange     pgtype.Multirange[pgtype.Range[pgtype.Int4]]
	VInt8range          pgtype.Range[pgtype.Int8]
	VInt8multirange     pgtype.Multirange[pgtype.Range[pgtype.Int8]]
}
