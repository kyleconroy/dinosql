// Code generated by sqlc. DO NOT EDIT.

package querytest

import (
	"fmt"
)

type FooFoobar string

const (
	FooFoobarFooA FooFoobar = "foo-a"
	FooFoobarFooB FooFoobar = "foo_b"
	FooFoobarFooC FooFoobar = "foo:c"
	FooFoobarFooD FooFoobar = "foo/d"
	FooFoobarFooe FooFoobar = "foo@e"
	FooFoobarFoof FooFoobar = "foo+f"
	FooFoobarFoog FooFoobar = "foo!g"
)

func (e *FooFoobar) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = FooFoobar(s)
	case string:
		*e = FooFoobar(s)
	default:
		return fmt.Errorf("unsupported scan type for FooFoobar: %T", src)
	}
	return nil
}

type Foo struct {
	Foobar FooFoobar
}
