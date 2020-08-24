package ast

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type FieldStore struct {
	Xpr        Node
	Arg        Node
	Newvals    *List
	Fieldnums  *List
	Resulttype Oid
}

func (n *FieldStore) Pos() int {
	return 0
}
