package ast

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type InferenceElem struct {
	Xpr          Node
	Expr         Node
	Infercollid  Oid
	Inferopclass Oid
}

func (n *InferenceElem) Pos() int {
	return 0
}
