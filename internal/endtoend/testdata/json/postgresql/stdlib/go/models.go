// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package querytest

import (
	"encoding/json"

	"github.com/tabbed/pqtype"
)

type Foo struct {
	A json.RawMessage
	B json.RawMessage
	C pqtype.NullRawMessage
	D pqtype.NullRawMessage
}
