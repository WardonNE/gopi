package builder

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type selectColumn struct {
	Column clause.Column
	Vars   []any
}

// Select add select clause
//
//	builder.Select(
//		"id",
//		clause.Column{Table: "users", Name: "name", Alias: "username"},
//		NewBuilder(db).Table("departments").Where("status", 1).Select(clause.Expr{SQL: "COUNT(*) AS count"}),
//		db.Table("tasks").Where("status", 1).Select("COUNT(*)"),
//		func(tx *gorm.DB) *gorm.DB {
//			var count int64
//			return tx.Table("orders").Where("status", 1).Count(&count)
//		},
//	)
func (builder *Builder) Select(columns ...any) *Builder {

	for _, column := range columns {
		switch value := column.(type) {
		case clause.Expr:
			rawCol := builder.conn.Clauses().ToSQL(func(tx *gorm.DB) *gorm.DB {
				stmt := tx.Statement
				value.Build(stmt)
				return tx
			})
			builder.selects.Add(clause.Column{
				Name: rawCol,
				Raw:  true,
			})
		case clause.NamedExpr:
			rawCol := builder.conn.Clauses().ToSQL(func(tx *gorm.DB) *gorm.DB {
				stmt := tx.Statement
				value.Build(stmt)
				return tx
			})
			builder.selects.Add(clause.Column{
				Name: rawCol,
				Raw:  true,
			})
		default:
			builder.selects.Add(clause.Column{
				Name: builder.QuoteField(column),
				Raw:  true,
			})
		}
	}
	return builder
}

// Distinct distinct
//
//	builder.Distinct()
//	builder.Distinct(
//		"id",
//		clause.Expr{SQL: "max(id) AS maxId"},
//		clause.Column{Table: "users", Name: "name", Alias: "username"},
//		NewBuilder(db).Table("departments").Where("status", 1).Select(clause.Expr{SQL: "COUNT(*) AS count"}),
//		db.Table("tasks").Where("status", 1).Select("COUNT(*)"),
//		func(tx *gorm.DB) *gorm.DB {
//			var count int64
//			return tx.Table("orders").Where("status = ?", 1).Count(&count)
//		},
//	)
func (builder *Builder) Distinct(columns ...any) *Builder {

	if len(columns) > 0 {
		builder = builder.Select(columns...)
	}
	builder.distinct = true
	return builder
}
