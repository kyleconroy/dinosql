// Code generated by sqlc. DO NOT EDIT.

package querytest

import (
	"fmt"
)

type Foobar string

const (
	FoobarFooA Foobar = "foo-a"
	FoobarFooB Foobar = "foo_b"
	FoobarFooC Foobar = "foo:c"
	FoobarFooD Foobar = "foo/d"
	FoobarFooe Foobar = "foo@e"
	FoobarFoof Foobar = "foo+f"
	FoobarFoog Foobar = "foo!g"
)

func (e *Foobar) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Foobar(s)
	case string:
		*e = Foobar(s)
	default:
		return fmt.Errorf("unsupported scan type for Foobar: %T", src)
	}
	return nil
}

type Foo struct {
	Foobar Foobar
}
