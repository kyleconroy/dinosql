// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0

package db

import (
	"config/global"
)

type Override struct {
	ONE   global.Text
	Two   int64
	Three global.ShouldSeeThis
	Four  []byte
	Five  []global.Text
}