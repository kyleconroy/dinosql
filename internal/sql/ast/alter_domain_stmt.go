package ast

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type AlterDomainStmt struct {
	Subtype   byte
	TypeName  *List
	Name      *string
	Def       Node
	Behavior  DropBehavior
	MissingOk bool
}

func (n *AlterDomainStmt) Pos() int {
	return 0
}
