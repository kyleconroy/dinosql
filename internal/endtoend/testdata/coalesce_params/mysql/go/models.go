// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package querytest

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"
)

type CalendarMaincalendar string

const (
	CalendarMaincalendarTrue  CalendarMaincalendar = "true"
	CalendarMaincalendarFalse CalendarMaincalendar = "false"
)

func (e *CalendarMaincalendar) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = CalendarMaincalendar(s)
	case string:
		*e = CalendarMaincalendar(s)
	default:
		return fmt.Errorf("unsupported scan type for CalendarMaincalendar: %T", src)
	}
	return nil
}

type NullCalendarMaincalendar struct {
	CalendarMaincalendar CalendarMaincalendar
	Valid                bool // Valid is true if CalendarMaincalendar is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullCalendarMaincalendar) Scan(value interface{}) error {
	if value == nil {
		ns.CalendarMaincalendar, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.CalendarMaincalendar.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullCalendarMaincalendar) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.CalendarMaincalendar), nil
}

type Author struct {
	ID      int64
	Address string
	Name    string
	Bio     string
}

type Calendar struct {
	ID           uint64
	Relation     uint64
	Calendarname []byte
	Title        []byte
	Description  []byte
	Timezone     string
	Uniquekey    string
	Idkey        string
	Maincalendar CalendarMaincalendar
	Createdate   time.Time
	Modifydate   time.Time
}

type Event struct {
	ID                uint64
	Relation          uint64
	Calendarreference uint64
	Uniquekey         string
	Eventname         []byte
	Description       []byte
	Location          string
	Timezone          string
	Idkey             sql.NullString
}
