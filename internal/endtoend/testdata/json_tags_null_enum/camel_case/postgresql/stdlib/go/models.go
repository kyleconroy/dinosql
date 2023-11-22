// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package db

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
)

type JobPostLocationType string

const (
	JobPostLocationTypeRemote   JobPostLocationType = "remote"
	JobPostLocationTypeInOffice JobPostLocationType = "in_office"
	JobPostLocationTypeHybrid   JobPostLocationType = "hybrid"
)

func (e *JobPostLocationType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = JobPostLocationType(s)
	case string:
		*e = JobPostLocationType(s)
	default:
		return fmt.Errorf("unsupported scan type for JobPostLocationType: %T", src)
	}
	return nil
}

type NullJobPostLocationType struct {
	JobPostLocationType JobPostLocationType `json:"jobPostLocationType"`
	Valid               bool                `json:"valid"` // Valid is true if JobPostLocationType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullJobPostLocationType) Scan(value interface{}) error {
	if value == nil {
		ns.JobPostLocationType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.JobPostLocationType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullJobPostLocationType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.JobPostLocationType), nil
}

type Author struct {
	ID   int64                   `json:"id"`
	Type NullJobPostLocationType `json:"type"`
	Name string                  `json:"name"`
	Bio  sql.NullString          `json:"bio"`
}
