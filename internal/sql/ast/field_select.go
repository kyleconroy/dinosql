package ast

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type FieldSelect struct {
	Xpr          Node
	Arg          Node
	Fieldnum     AttrNumber
	Resulttype   Oid
	Resulttypmod int32
	Resultcollid Oid
}

func (n *FieldSelect) Pos() int {
	return 0
}
