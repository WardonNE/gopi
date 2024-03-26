package builder

import (
	"github.com/wardonne/gopi/database/exception"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// WithJoin add a join by a defined relationship
func (builder *Builder) WithJoin(relation string) *Builder {

	builder.db = builder.db.Joins(relation)
	return builder
}

// LeftJoin add a left join
//
//	builder.LeftJoin("departments", "users.department_id = departments.id")
//	builder.LeftJoin("departments", "users.department_id = departments.id and departments.status = ?", 1)
func (builder *Builder) LeftJoin(table string, condition string, values ...any) *Builder {

	builder.joins.Add(clause.Join{
		Type:  clause.LeftJoin,
		Table: clause.Table{Name: table, Raw: true},
		ON: clause.Where{
			Exprs: []clause.Expression{
				clause.Expr{
					SQL:  condition,
					Vars: values,
				},
			},
		},
	})
	return builder
}

// LeftJoinSub add a subquery left join
//
//	builder.LeftJoinSub(NewBuilder(db).Table("departments"), "d", "d.id = users.department_id")
//	builder.LeftJoinSub(db.Table("departments"), "d", "d.id = users.department_id")
func (builder *Builder) LeftJoinSub(query any, alias string, condition string, values ...any) *Builder {

	var rawSubQuery string
	switch q := query.(type) {
	case *Builder, *gorm.DB:
		rawSubQuery = builder.QuoteField(q)
	default:
		exception.ThrowInvalidParamTypeErr("LeftJoinSub", q)
	}
	builder.joins.Add(clause.Join{
		Type:  clause.LeftJoin,
		Table: clause.Table{Name: rawSubQuery, Alias: alias, Raw: true},
		ON: clause.Where{
			Exprs: []clause.Expression{
				clause.Expr{
					SQL:  condition,
					Vars: values,
				},
			},
		},
	})
	return builder
}

// RightJoin add a right join
//
//	builder.RightJoin("departments", "users.department_id = departments.id")
//	builder.RightJoin("departments", "users.department_id = departments.id and departments.status = ?", 1)
func (builder *Builder) RightJoin(table string, condition string, values ...any) *Builder {

	builder.joins.Add(clause.Join{
		Type:  clause.RightJoin,
		Table: clause.Table{Name: table, Raw: true},
		ON: clause.Where{
			Exprs: []clause.Expression{
				clause.Expr{
					SQL:  condition,
					Vars: values,
				},
			},
		},
	})
	return builder
}

// RightJoinSub add a subquery right join
//
//	builder.RightJoinSub(NewBuilder(db).Table("departments"), "d", "d.id = users.department_id")
//	builder.RightJoinSub(db.Table("departments"), "d", "d.id = users.department_id")
func (builder *Builder) RightJoinSub(query any, alias string, condition string, values ...any) *Builder {

	var rawSubQuery string
	switch q := query.(type) {
	case *Builder, *gorm.DB:
		rawSubQuery = builder.QuoteField(q)
	default:
		exception.ThrowInvalidParamTypeErr("RightJoinSub", q)
	}
	builder.joins.Add(clause.Join{
		Type:  clause.RightJoin,
		Table: clause.Table{Name: rawSubQuery, Alias: alias, Raw: true},
		ON: clause.Where{
			Exprs: []clause.Expression{
				clause.Expr{
					SQL:  condition,
					Vars: values,
				},
			},
		},
	})
	return builder
}

// InnerJoin add an inner join
//
//	builder.InnerJoin("departments", "users.department_id = departments.id")
//	builder.InnerJoin("departments", "users.department_id = departments.id and departments.status = ?", 1)
func (builder *Builder) InnerJoin(table string, condition string, values ...any) *Builder {

	builder.joins.Add(clause.Join{
		Type:  clause.InnerJoin,
		Table: clause.Table{Name: table, Raw: true},
		ON: clause.Where{
			Exprs: []clause.Expression{
				clause.Expr{
					SQL:  condition,
					Vars: values,
				},
			},
		},
	})
	return builder
}

// InnerJoinSub add a subquery inner join
//
//	builder.InnerJoinSub(NewBuilder(db).Table("departments"), "d", "d.id = users.department_id")
//	builder.InnerJoinSub(db.Table("departments"), "d", "d.id = users.department_id")
func (builder *Builder) InnerJoinSub(query any, alias string, condition string, values ...any) *Builder {

	var rawSubQuery string
	switch q := query.(type) {
	case *Builder, *gorm.DB:
		rawSubQuery = builder.QuoteField(q)
	default:
		exception.ThrowInvalidParamTypeErr("InnerJoinSub", q)
	}
	builder.joins.Add(clause.Join{
		Type:  clause.InnerJoin,
		Table: clause.Table{Name: rawSubQuery, Alias: alias, Raw: true},
		ON: clause.Where{
			Exprs: []clause.Expression{
				clause.Expr{
					SQL:  condition,
					Vars: values,
				},
			},
		},
	})
	return builder
}

// CrossJoin add a cross join
//
//	builder.CrossJoin("departments", "users.department_id = departments.id")
//	builder.CrossJoin("departments", "users.department_id = departments.id and departments.status = ?", 1)
func (builder *Builder) CrossJoin(table string, condition string, values ...any) *Builder {

	builder.joins.Add(clause.Join{
		Type:  clause.CrossJoin,
		Table: clause.Table{Name: table, Raw: true},
		ON: clause.Where{
			Exprs: []clause.Expression{
				clause.Expr{
					SQL:  condition,
					Vars: values,
				},
			},
		},
	})
	return builder
}

// CrossJoinSub add a subquery cross join
//
//	builder.CrossJoinSub(NewBuilder(db).Table("departments"), "d", "d.id = users.department_id")
//	builder.CrossJoinSub(db.Table("departments"), "d", "d.id = users.department_id")
func (builder *Builder) CrossJoinSub(query any, alias string, condition string, values ...any) *Builder {

	var rawSubQuery string
	switch q := query.(type) {
	case *Builder, *gorm.DB:
		rawSubQuery = builder.QuoteField(q)
	default:
		exception.ThrowInvalidParamTypeErr("CrossJoinSub", q)
	}
	builder.joins.Add(clause.Join{
		Type:  clause.CrossJoin,
		Table: clause.Table{Name: rawSubQuery, Alias: alias, Raw: true},
		ON: clause.Where{
			Exprs: []clause.Expression{
				clause.Expr{
					SQL:  condition,
					Vars: values,
				},
			},
		},
	})
	return builder
}
