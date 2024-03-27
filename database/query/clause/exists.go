package clause

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Exists exists for where
type Exists struct {
	Query any
	Vars  []any
}

// Build build exists
func (b Exists) Build(builder clause.Builder) {
	builder.WriteString(" EXISTS ")
	switch expr := b.Query.(type) {
	case *gorm.DB:
		builder.WriteByte('(')
		builder.AddVar(builder, expr)
		builder.WriteByte(')')
	case string:
		builder.WriteByte('(')
		(clause.Expr{SQL: expr, Vars: b.Vars}).Build(builder)
		builder.WriteByte(')')
	case fmt.Stringer:
		builder.WriteByte('(')
		(clause.Expr{SQL: expr.String(), Vars: b.Vars}).Build(builder)
		builder.WriteByte(')')
	case clause.Expression:
		expr.Build(builder)
	default:
		builder.WriteByte('(')
		(clause.Expr{SQL: fmt.Sprintf("%s", expr), Vars: b.Vars}).Build(builder)
		builder.WriteByte(')')
	}
}

// NegationBuild build not exists
func (b Exists) NegationBuild(builder clause.Builder) {
	builder.WriteString(" NOT EXISTS ")
	switch expr := b.Query.(type) {
	case *gorm.DB:
		builder.WriteByte('(')
		builder.AddVar(builder, expr)
		builder.WriteByte(')')
	case string:
		builder.WriteByte('(')
		(clause.Expr{SQL: expr, Vars: b.Vars}).Build(builder)
		builder.WriteByte(')')
	case fmt.Stringer:
		builder.WriteByte('(')
		(clause.Expr{SQL: expr.String(), Vars: b.Vars}).Build(builder)
		builder.WriteByte(')')
	case clause.Expression:
		expr.Build(builder)
	default:
		builder.WriteByte('(')
		(clause.Expr{SQL: fmt.Sprintf("%s", expr), Vars: b.Vars}).Build(builder)
		builder.WriteByte(')')
	}
}
