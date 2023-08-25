// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package datatype

import (
	"net"
	"net/netip"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type DtCharacter struct {
	A *string
	B *string
	C *string
	D *string
	E *string
}

type DtCharacterNotNull struct {
	A string
	B string
	C string
	D string
	E string
}

type DtDatetime struct {
	A pgtype.Date
	B pgtype.Time
	C pgtype.Time
	D *time.Time
	E pgtype.Timestamp
	F pgtype.Timestamp
	G pgtype.Timestamptz
	H pgtype.Timestamptz
}

type DtDatetimeNotNull struct {
	A pgtype.Date
	B pgtype.Time
	C pgtype.Time
	D time.Time
	E pgtype.Timestamp
	F pgtype.Timestamp
	G pgtype.Timestamptz
	H pgtype.Timestamptz
}

type DtNetType struct {
	A *netip.Addr
	B *netip.Prefix
	C net.HardwareAddr
}

type DtNetTypesNotNull struct {
	A netip.Addr
	B netip.Prefix
	C net.HardwareAddr
}

type DtNumeric struct {
	A *int16
	B *int32
	C *int64
	D pgtype.Numeric
	E pgtype.Numeric
	F *float32
	G *float64
	H *int16
	I *int32
	J *int64
	K *int16
	L *int32
	M *int64
}

type DtNumericNotNull struct {
	A int16
	B int32
	C int64
	D pgtype.Numeric
	E pgtype.Numeric
	F float32
	G float64
	H int16
	I int32
	J int64
	K int16
	L int32
	M int64
}

type DtRange struct {
	A pgtype.Range[pgtype.Int4]
	B pgtype.Range[pgtype.Int8]
	C pgtype.Range[pgtype.Numeric]
	D pgtype.Range[pgtype.Timestamp]
	E pgtype.Range[pgtype.Timestamptz]
	F pgtype.Range[pgtype.Date]
}

type DtRangeNotNull struct {
	A pgtype.Range[pgtype.Int4]
	B pgtype.Range[pgtype.Int8]
	C pgtype.Range[pgtype.Numeric]
	D pgtype.Range[pgtype.Timestamp]
	E pgtype.Range[pgtype.Timestamptz]
	F pgtype.Range[pgtype.Date]
}
