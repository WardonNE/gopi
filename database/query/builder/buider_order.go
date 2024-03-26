package builder

import (
	"fmt"

	"github.com/wardonne/gopi/database/exception"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Order add order by clause
//
//	builder.Order("id") // ORDER BY id ASC
//	builder.Order("id", true) // ORDER BY id DESC
func (builder *Builder) Order(column any, desc bool) *Builder {

	switch value := column.(type) {
	case string, fmt.Stringer, clause.Column, *Builder, *gorm.DB:
		builder.db = builder.db.Order(clause.OrderByColumn{
			Column: clause.Column{
				Name: builder.QuoteField(value),
				Raw:  true,
			},
			Desc: desc,
		})
	case clause.Expr:
		builder.db = builder.db.Order(clause.OrderByColumn{
			Column: clause.Column{
				Name: value.SQL,
				Raw:  true,
			},
			Desc: desc,
		})
	default:
		exception.ThrowInvalidParamTypeErr("order by", value)
	}
	return builder
}

// OrderAsc add order by asc clause
//
//	builder.OrderAsc("id") // ORDER BY id ASC
func (builder *Builder) OrderAsc(column any) *Builder {

	return builder.Order(column, false)
}

// OrderDesc add order by desc clause
//
//	builder.OrderDesc("id") // ORDER BY id DESC
func (builder *Builder) OrderDesc(column any) *Builder {

	return builder.Order(column, true)
}
