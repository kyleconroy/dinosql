// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package querytest

import (
	"database/sql"
	"time"
)

type Author struct {
	ID   int64
	Name string
	Bio  sql.NullString
}

type WeatherMetric struct {
	Time            time.Time
	TimezoneShift   sql.NullInt32
	CityName        sql.NullString
	TempC           sql.NullFloat64
	FeelsLikeC      sql.NullFloat64
	TempMinC        sql.NullFloat64
	TempMaxC        sql.NullFloat64
	PressureHpa     sql.NullFloat64
	HumidityPercent sql.NullFloat64
	WindSpeedMs     sql.NullFloat64
	WindDeg         sql.NullInt32
	Rain1hMm        sql.NullFloat64
	Rain3hMm        sql.NullFloat64
	Snow1hMm        sql.NullFloat64
	Snow3hMm        sql.NullFloat64
	CloudsPercent   sql.NullInt32
	WeatherTypeID   sql.NullInt32
}
