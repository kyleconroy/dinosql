package ast

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type CurrentOfExpr struct {
	Xpr         Node
	Cvarno      Index
	CursorName  *string
	CursorParam int
}

func (n *CurrentOfExpr) Pos() int {
	return 0
}
