package ast

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type PartitionRangeDatum struct {
	Kind     PartitionRangeDatumKind
	Value    Node
	Location int
}

func (n *PartitionRangeDatum) Pos() int {
	return n.Location
}
