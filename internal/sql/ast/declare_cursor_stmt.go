package ast

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type DeclareCursorStmt struct {
	Portalname *string
	Options    int
	Query      Node
}

func (n *DeclareCursorStmt) Pos() int {
	return 0
}
