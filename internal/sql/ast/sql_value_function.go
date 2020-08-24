package ast

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type SQLValueFunction struct {
	Xpr      Node
	Op       SQLValueFunctionOp
	Type     Oid
	Typmod   int32
	Location int
}

func (n *SQLValueFunction) Pos() int {
	return n.Location
}
