package compiler

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
	"github.com/kyleconroy/sqlc/internal/sql/ast/pg"
	"github.com/kyleconroy/sqlc/internal/sql/catalog"
)

type QueryCatalog struct {
	catalog *catalog.Catalog
	ctes    map[string]*Table
}

func buildQueryCatalog(c *catalog.Catalog, node ast.Node) (*QueryCatalog, error) {
	var with *pg.WithClause
	switch n := node.(type) {
	case *pg.InsertStmt:
		with = n.WithClause
	case *pg.UpdateStmt:
		with = n.WithClause
	case *pg.SelectStmt:
		with = n.WithClause
	default:
		with = nil
	}
	qc := &QueryCatalog{catalog: c, ctes: map[string]*Table{}}
	if with != nil {
		for _, item := range with.Ctes.Items {
			if cte, ok := item.(*pg.CommonTableExpr); ok {
				cols, err := outputColumns(qc, cte.Ctequery)
				if err != nil {
					return nil, err
				}
				qc.ctes[*cte.Ctename] = &Table{
					Rel: &ast.TableName{
						Name: *cte.Ctename,
					},
					Columns: cols,
				}
			}
		}
	}
	return qc, nil
}

func (qc QueryCatalog) GetTable(rel *ast.TableName) (*Table, error) {
	cte, exists := qc.ctes[rel.Name]
	if exists {
		return cte, nil
	}
	src, err := qc.catalog.GetTable(rel)
	if err != nil {
		return nil, err
	}
	var cols []*Column
	for _, c := range src.Columns {
		cols = append(cols, &Column{
			Table:    rel,
			Name:     c.Name,
			DataType: dataType(&c.Type),
			NotNull:  c.IsNotNull,
			IsArray:  c.IsArray,
		})
	}
	return &Table{Rel: rel, Columns: cols}, nil
}
