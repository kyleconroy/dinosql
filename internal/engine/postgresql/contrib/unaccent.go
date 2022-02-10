// Code generated by sqlc-pg-gen. DO NOT EDIT.

package contrib

import (
	"github.com/egtann/sqlc/internal/sql/ast"
	"github.com/egtann/sqlc/internal/sql/catalog"
)

func Unaccent() *catalog.Schema {
	s := &catalog.Schema{Name: "pg_catalog"}
	s.Funcs = []*catalog.Function{
		{
			Name: "unaccent",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "regdictionary"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "unaccent",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
	}
	return s
}
