package builder

import (
	"fmt"

	"github.com/wardonne/gopi/database/exception"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Group add groupby clause
//
//	builder.GroupBy(
//		"id",
//		clause.Expr{SQL: "users.name"},
//		clause.Column{Name: "department_id"},
//	)
func (builder *Builder) Group(columns ...any) *Builder {
	builder = builder.instance()
	for _, column := range columns {
		switch value := column.(type) {
		case string, clause.Column, fmt.Stringer, *Builder, *gorm.DB:
			builder.db = builder.db.Group(builder.QuoteField(value))
		case clause.Expr:
			builder.db = builder.db.Group(value.SQL)
		default:
			exception.ThrowInvalidParamTypeErr("group by", value)
		}
	}
	return builder
}
