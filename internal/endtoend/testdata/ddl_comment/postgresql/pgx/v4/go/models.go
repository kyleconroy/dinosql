// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package querytest

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
)

// Enum comment
type FooBat string

const (
	FooBatBat FooBat = "bat"
)

func (e *FooBat) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = FooBat(s)
	case string:
		*e = FooBat(s)
	default:
		return fmt.Errorf("unsupported scan type for FooBat: %T", src)
	}
	return nil
}

type NullFooBat struct {
	FooBat FooBat
	Valid  bool // Valid is true if FooBat is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullFooBat) Scan(value interface{}) error {
	if value == nil {
		ns.FooBat, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.FooBat.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullFooBat) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.FooBat), nil
}

// Table comment
type FooBar struct {
	// Column comment
	Baz sql.NullString
}
