package ast

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type CollateExpr struct {
	Xpr      Node
	Arg      Node
	CollOid  Oid
	Location int
}

func (n *CollateExpr) Pos() int {
	return n.Location
}
