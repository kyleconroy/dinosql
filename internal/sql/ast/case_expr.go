package ast

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type CaseExpr struct {
	Xpr        Node
	Casetype   Oid
	Casecollid Oid
	Arg        Node
	Args       *List
	Defresult  Node
	Location   int
}

func (n *CaseExpr) Pos() int {
	return n.Location
}
