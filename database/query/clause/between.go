package clause

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Between between for where
type Between struct {
	Column any
	Values [2]any
}

// Build build between
func (b Between) Build(builder clause.Builder) {
	builder.WriteQuoted(b.Column)
	builder.WriteString(" BETWEEN ")
	if _, ok := b.Values[0].(*gorm.DB); ok {
		builder.WriteByte('(')
		builder.AddVar(builder, b.Values[0])
		builder.WriteByte(')')
	} else {
		builder.AddVar(builder, b.Values[0])
	}
	builder.WriteString(" AND ")
	if _, ok := b.Values[1].(*gorm.DB); ok {
		builder.WriteByte('(')
		builder.AddVar(builder, b.Values[1])
		builder.WriteByte(')')
	} else {
		builder.AddVar(builder, b.Values[1])
	}
}

// NegationBuild build not between
func (b Between) NegationBuild(builder clause.Builder) {
	builder.WriteQuoted(b.Column)
	builder.WriteString(" NOT BETWEEN ")
	if _, ok := b.Values[0].(*gorm.DB); ok {
		builder.WriteByte('(')
		builder.AddVar(builder, b.Values[0])
		builder.WriteByte(')')
	} else {
		builder.AddVar(builder, b.Values[0])
	}
	builder.WriteString(" AND ")
	if _, ok := b.Values[1].(*gorm.DB); ok {
		builder.WriteByte('(')
		builder.AddVar(builder, b.Values[1])
		builder.WriteByte(')')
	} else {
		builder.AddVar(builder, b.Values[1])
	}
}
