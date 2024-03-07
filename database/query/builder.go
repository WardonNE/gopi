package query

import (
	"fmt"
	"strings"

	"github.com/wardonne/gopi/database/exception"
	"github.com/wardonne/gopi/support/collection/list"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Builder query builder
type Builder[T any] struct {
	db         *gorm.DB
	selects    *list.ArrayList[clause.Column]
	table      *clause.Table
	joins      *list.ArrayList[clause.Join]
	conditions *list.ArrayList[clause.Expression]
	groups     *list.ArrayList[clause.Column]
	havings    *list.ArrayList[clause.Expression]
	limit      *int
	offset     int
}

// NewBuilder create a new query builder
//
//	builder := NewBuilder[any](db)
func NewBuilder[T any](db *gorm.DB) *Builder[T] {
	return &Builder[T]{
		db:         db.Session(&gorm.Session{Initialized: true}),
		selects:    list.NewArrayList[clause.Column](),
		joins:      list.NewArrayList[clause.Join](),
		conditions: list.NewArrayList[clause.Expression](),
		groups:     list.NewArrayList[clause.Column](),
		havings:    list.NewArrayList[clause.Expression](),
	}
}

func (builder *Builder[T]) build() {
	// add select clause
	if builder.selects.IsEmpty() {
		builder.db.Statement.AddClauseIfNotExists(clause.Select{
			Columns: []clause.Column{{Name: "*", Raw: true}},
		})
	} else {
		builder.db.Statement.AddClauseIfNotExists(clause.Select{Columns: builder.selects.ToArray()})
	}
	// add from clause
	if builder.table == nil {
		builder.db.Statement.AddClauseIfNotExists(clause.From{
			Tables: []clause.Table{{Name: clause.CurrentTable}},
			Joins:  builder.joins.ToArray(),
		})
	} else {
		builder.db.Statement.AddClauseIfNotExists(clause.From{
			Tables: []clause.Table{*builder.table},
			Joins:  builder.joins.ToArray(),
		})
	}
	// add where clause
	builder.db.Statement.AddClauseIfNotExists(clause.Where{
		Exprs: builder.conditions.ToArray(),
	})
	// add group by and having clause
	builder.db.Statement.AddClauseIfNotExists(clause.GroupBy{
		Columns: builder.groups.ToArray(),
		Having:  builder.havings.ToArray(),
	})
	// add limit offset clause
	builder.db.Statement.AddClauseIfNotExists(clause.Limit{
		Limit:  builder.limit,
		Offset: builder.offset,
	})
	builder.db.Statement.Build("SELECT", "FROM", "WHERE", "GROUP BY", "ORDER BY", "LIMIT", "FOR")
}

func (builder *Builder[T]) buildConditionValue(value any) any {
	switch v := value.(type) {
	case string:
		return v
		// case int, uint, int8, uint8, int16, uint16, int32, uint32,
	}
	return value
}

// DB returns *gorm.DB
func (builder *Builder[T]) DB() *gorm.DB {
	builder.build()
	return builder.db
}

// Select add select clause
//
//	subquery := NewBuilder[any]().Where("id", 1)
//	builder.Select(
//		"id",
//		clause.Expr{SQL: "max(id) AS maxId"},
//		clause.Column{Table: "users", Name: "name", Alias: "username"},
//		subquery,
//	)
func (builder *Builder[T]) Select(columns ...any) *Builder[T] {
	for _, column := range columns {
		switch value := column.(type) {
		case string:
			builder.selects.Add(clause.Column{
				Name: value,
			})
		case clause.Expr:
			builder.selects.Add(clause.Column{
				Name: value.SQL,
				Raw:  true,
			})
		case clause.Column:
			builder.selects.Add(value)
		case fmt.Stringer:
			builder.selects.Add(clause.Column{
				Name: value.String(),
			})
		case *Builder[any]:
			builder.selects.Add(clause.Column{
				Name: fmt.Sprintf("(%s)", strings.TrimSpace(value.ToSQL())),
				Raw:  true,
			})
		default:
			exception.ThrowInvalidParamTypeErr("select", value)
		}
	}
	return builder
}

// From add from clause
//
//	builder.From("users")
//	builder.From(clause.Expr{SQL: "users AS u"})
//	builder.From(clause.Table{Name: "users", Alias: "u"})
func (builder *Builder[T]) From(table any) *Builder[T] {
	switch value := table.(type) {
	case string:
		builder.table = &clause.Table{
			Name: value,
		}
	case clause.Expr:
		builder.table = &clause.Table{
			Name: value.SQL,
			Raw:  true,
		}
	case clause.Table:
		builder.table = &value
	case fmt.Stringer:
		builder.table = &clause.Table{
			Name: value.String(),
		}
	default:
		exception.ThrowInvalidParamTypeErr("from", value)
	}
	return builder
}

// FromRaw add from clause by raw sql
//
//	builder.FromRaw("users as u")
//	builder.FromRaw("(SELECT * FROM users WHERE status = 1) AS u")
func (builder *Builder[T]) FromRaw(sql string) *Builder[T] {
	builder.table = &clause.Table{
		Name: sql,
		Raw:  true,
	}
	return builder
}

// FromQuery add from subquery clause
//
//	query := NewBuilder[any](db).From("users").Where("status", 1)
//	builder.FromQuery(query)
func (builder *Builder[T]) FromQuery(query *Builder[any], alias string) *Builder[T] {
	builder.table = &clause.Table{
		Name:  fmt.Sprintf("(%s)", strings.TrimSpace(query.ToSQL())),
		Alias: alias,
		Raw:   true,
	}
	return builder
}

// Join add a join clause
//
//	builder.Join("emails", "users.id = emails.id")
//	builder.Join("emails", "users.id = emails.id AND emails.status = ?", 1)
//	builder.Join(clause.Table{Name: "emails"}, "users.id = emails.id")
//	builder.Join(clause.Expr{SQL: "emails"}, "users.id = emails.id")
//	builder.Join(
//		clause.Join{
//			Table: clause.Table{
//				Name: "emails", Alias: "e",
//			},
//			On: clause.Where{
//				Exprs:[]clause.Expression{
//					clause.Expr{SQL:"users.id = e.user_id"},
//				}
//			}
//		}
//	)
//
//	type CustomTable struct {
//		Name string
//	}
//
//	func (c CustomTable) String() string {
//		return c.Name
//	}
//
//	builder.Join(CustomTable{Name: "emails"}, "users.id = emails.user_id")
//
//	builder.Join("emails", clause.Expr{SQL: "users.id = emails.user_id"})
//	builder.Join("emails", clause.Where{
//		Exprs: []clause.Expression{clause.Expr{SQL: "users.id = emails.user_id"}}
//	})
func (builder *Builder[T]) Join(table any, condition any, values ...any) *Builder[T] {
	var buildOn = func(condition any) clause.Where {
		var on clause.Where
		switch value := condition.(type) {
		case string:
			on = clause.Where{Exprs: []clause.Expression{
				clause.Expr{SQL: value, Vars: values},
			}}
		case fmt.Stringer:
			on = clause.Where{Exprs: []clause.Expression{
				clause.Expr{SQL: value.String(), Vars: values},
			}}
		case clause.Expression:
			on = clause.Where{Exprs: []clause.Expression{value}}
		case clause.Where:
			on = value
		}
		return on
	}
	switch value := table.(type) {
	case string:
		builder.joins.Add(clause.Join{
			Table: clause.Table{Name: value},
			ON:    buildOn(condition),
		})
	case clause.Expr:
		builder.joins.Add(clause.Join{
			Table: clause.Table{Name: value.SQL, Raw: true},
			ON:    buildOn(condition),
		})
	case clause.Table:
		builder.joins.Add(clause.Join{
			Table: value,
			ON:    buildOn(condition),
		})
	case clause.Join:
		builder.joins.Add(value)
	case fmt.Stringer:
		builder.joins.Add(clause.Join{
			Table: clause.Table{Name: value.String()},
			ON:    buildOn(condition),
		})
	}
	return builder
}

// JoinRaw add join clause by raw sql
//
//	builder.JoinRaw("emails AS e ON users.id = e.user_id")
//	builder.JoinRaw("emails AS e ON users.id = e.user_id AND e.status = ?", 1)
func (builder *Builder[T]) JoinRaw(sql string, values ...any) *Builder[T] {
	builder.joins.Add(clause.Join{
		Expression: clause.Expr{SQL: sql, Vars: values},
	})
	return builder
}

func (builder *Builder[T]) JoinQuery(query *Builder[any], alias, condition string, values ...any) *Builder[T] {
	builder.joins.Add(clause.Join{
		Table: clause.Table{Name: fmt.Sprintf("(%s)", strings.TrimSpace(query.ToSQL())), Alias: alias, Raw: true},
		ON:    clause.Where{Exprs: []clause.Expression{clause.Expr{SQL: condition, Vars: values}}},
	})
	return builder
}

func (builder *Builder[T]) CrossJoin(tableName string, condition string, values ...any) *Builder[T] {
	builder.joins.Add(clause.Join{
		Type:  clause.CrossJoin,
		Table: clause.Table{Name: tableName},
		ON: clause.Where{Exprs: []clause.Expression{
			clause.Expr{SQL: condition, Vars: values},
		}},
	})
	return builder
}

func (builder *Builder[T]) CrossJoinQuery(query *Builder[T], alias, condition string, values ...any) *Builder[T] {
	builder.joins.Add(clause.Join{
		Type:  clause.CrossJoin,
		Table: clause.Table{Name: fmt.Sprintf("(%s)", strings.TrimSpace(query.ToSQL())), Alias: alias, Raw: true},
		ON:    clause.Where{Exprs: []clause.Expression{clause.Expr{SQL: condition, Vars: values}}},
	})
	return builder
}

func (builder *Builder[T]) InnerJoin(tableName string, condition string, values ...any) *Builder[T] {
	builder.joins.Add(clause.Join{
		Type:  clause.InnerJoin,
		Table: clause.Table{Name: tableName},
		ON: clause.Where{Exprs: []clause.Expression{
			clause.Expr{SQL: condition, Vars: values},
		}},
	})
	return builder
}

func (builder *Builder[T]) InnerJoinQuery(query *Builder[T], alias, condition string, values ...any) *Builder[T] {
	builder.joins.Add(clause.Join{
		Type:  clause.InnerJoin,
		Table: clause.Table{Name: fmt.Sprintf("(%s)", strings.TrimSpace(query.ToSQL())), Alias: alias, Raw: true},
		ON:    clause.Where{Exprs: []clause.Expression{clause.Expr{SQL: condition, Vars: values}}},
	})
	return builder
}

func (builder *Builder[T]) LeftJoin(tableName string, condition string, values ...any) *Builder[T] {
	builder.joins.Add(clause.Join{
		Type:  clause.LeftJoin,
		Table: clause.Table{Name: tableName},
		ON: clause.Where{Exprs: []clause.Expression{
			clause.Expr{SQL: condition, Vars: values},
		}},
	})
	return builder
}

func (builder *Builder[T]) LeftJoinQuery(query *Builder[T], alias, condition string, values ...any) *Builder[T] {
	builder.joins.Add(clause.Join{
		Type:  clause.InnerJoin,
		Table: clause.Table{Name: fmt.Sprintf("(%s)", strings.TrimSpace(query.ToSQL())), Alias: alias, Raw: true},
		ON:    clause.Where{Exprs: []clause.Expression{clause.Expr{SQL: condition, Vars: values}}},
	})
	return builder
}

func (builder *Builder[T]) RightJoin(tableName string, condition string, values ...any) *Builder[T] {
	builder.joins.Add(clause.Join{
		Type:  clause.RightJoin,
		Table: clause.Table{Name: tableName},
		ON: clause.Where{Exprs: []clause.Expression{
			clause.Expr{SQL: condition, Vars: values},
		}},
	})
	return builder
}

func (builder *Builder[T]) RightJoinQuery(query *Builder[T], alias, condition string, values ...any) *Builder[T] {
	builder.joins.Add(clause.Join{
		Type:  clause.RightJoin,
		Table: clause.Table{Name: fmt.Sprintf("(%s)", strings.TrimSpace(query.ToSQL())), Alias: alias, Raw: true},
		ON:    clause.Where{Exprs: []clause.Expression{clause.Expr{SQL: condition, Vars: values}}},
	})
	return builder
}

func (builder *Builder[T]) Where(column string, value any) *Builder[T] {
	switch v := value.(type) {
	case *Builder[any]:
		builder.conditions.Add(clause.Expr{
			SQL: fmt.Sprintf("`%s` = (%s)", column, strings.TrimSpace(v.ToSQL())),
		})
	case *gorm.DB:
		builder.conditions.Add(clause.Expr{
			SQL: fmt.Sprintf("`%s` = (%s)", column, strings.TrimSpace(v.ToSQL(func(tx *gorm.DB) *gorm.DB {
				return tx
			}))),
		})
	default:
		builder.conditions.Add(clause.Eq{
			Column: column,
			Value:  value,
		})
	}
	return builder
}

func (builder *Builder[T]) WhereNot(column string, value any) *Builder[T] {
	builder.conditions.Add(clause.Not(clause.Eq{
		Column: column,
		Value:  value,
	}))
	return builder
}

func (builder *Builder[T]) OrWhere(column string, value any) *Builder[T] {
	builder.conditions.Add(clause.Or(clause.Eq{
		Column: column,
		Value:  value,
	}))
	return builder
}

func (builder *Builder[T]) OrWhereNot(column string, value any) *Builder[T] {
	builder.conditions.Add(clause.Or(clause.Not(clause.Eq{
		Column: column,
		Value:  value,
	})))
	return builder
}

func (builder *Builder[T]) WhereRaw(sql string, args ...any) *Builder[T] {
	builder.conditions.Add(clause.Expr{
		SQL:  sql,
		Vars: args,
	})
	return builder
}

func (builder *Builder[T]) WhereRawNot(sql string, args ...any) *Builder[T] {
	builder.conditions.Add(clause.Expr{
		SQL:  sql,
		Vars: args,
	})
	return builder
}

func (builder *Builder[T]) OrWhereRaw(sql string, args ...any) *Builder[T] {
	builder.conditions.Add(clause.Or(clause.Expr{
		SQL:  sql,
		Vars: args,
	}))
	return builder
}

func (builder *Builder[T]) OrWhereRawNot(sql string, args ...any) *Builder[T] {
	builder.conditions.Add(clause.Expr{
		SQL:  sql,
		Vars: args,
	})
	return builder
}

func (builder *Builder[T]) WhereQuery(query *Builder[any]) *Builder[T] {
	builder.conditions.Add(clause.And(
		query.conditions.ToArray()...,
	))
	return builder
}

func (builder *Builder[T]) OrWhereQuery(query *Builder[any]) *Builder[T] {
	builder.conditions.Add(clause.Or(
		query.conditions.ToArray()...,
	))
	return builder
}

func (builder *Builder[T]) WhereNotQuery(query *Builder[any]) *Builder[T] {
	builder.conditions.Add(clause.Not(clause.And(query.conditions.ToArray()...)))
	return builder
}

func (builder *Builder[T]) OrWhereNotQuery(query *Builder[any]) *Builder[T] {
	builder.conditions.Add(clause.Or(clause.Not(query.conditions.ToArray()...)))
	return builder
}

func (builder *Builder[T]) WhereNull(column string) *Builder[T] {
	builder.conditions.Add(clause.Eq{
		Column: column,
	})
	return builder
}

func (builder *Builder[T]) OrWhereNull(column string) *Builder[T] {
	builder.conditions.Add(clause.Or(clause.Eq{
		Column: column,
	}))
	return builder
}

func (builder *Builder[T]) WhereNotNull(column string) *Builder[T] {
	builder.conditions.Add(clause.Not(clause.Eq{
		Column: column,
	}))
	return builder
}

func (builder *Builder[T]) OrWhereNotNull(column string) *Builder[T] {
	builder.conditions.Add(clause.Or(clause.Not(clause.Eq{
		Column: column,
	})))
	return builder
}

func (builder *Builder[T]) WhereEq(column string, value any) *Builder[T] {
	builder.conditions.Add(clause.Eq{
		Column: column,
		Value:  value,
	})
	return builder
}

func (builder *Builder[T]) WhereNeq(column string, value any) *Builder[T] {
	builder.conditions.Add(clause.Neq{
		Column: column,
		Value:  value,
	})
	return builder
}

func (builder *Builder[T]) OrWhereEq(column string, value any) *Builder[T] {
	builder.conditions.Add(clause.Or(clause.Eq{
		Column: column,
		Value:  value,
	}))
	return builder
}

func (builder *Builder[T]) OrWhereNeq(column string, value any) *Builder[T] {
	builder.conditions.Add(clause.Or(clause.Not(clause.Neq{
		Column: column,
		Value:  value,
	})))
	return builder
}

func (builder *Builder[T]) WhereGt(column string, value any) *Builder[T] {
	builder.conditions.Add(clause.Gt{
		Column: column,
		Value:  value,
	})
	return builder
}

func (builder *Builder[T]) OrWhereGt(column string, value any) *Builder[T] {
	builder.conditions.Add(clause.Or(clause.Gt{
		Column: column,
		Value:  value,
	}))
	return builder
}

func (builder *Builder[T]) WhereGte(column string, value any) *Builder[T] {
	builder.conditions.Add(clause.Gte{
		Column: column,
		Value:  value,
	})
	return builder
}

func (builder *Builder[T]) OrWhereGte(column string, value any) *Builder[T] {
	builder.conditions.Add(clause.Or(clause.Gte{
		Column: column,
		Value:  value,
	}))
	return builder
}

func (builder *Builder[T]) WhereLt(column string, value any) *Builder[T] {
	builder.conditions.Add(clause.Lt{
		Column: column,
		Value:  value,
	})
	return builder
}

func (builder *Builder[T]) OrWhereLt(column string, value any) *Builder[T] {
	builder.conditions.Add(clause.Or(clause.Lt{
		Column: column,
		Value:  value,
	}))
	return builder
}

func (builder *Builder[T]) WhereLte(column string, value any) *Builder[T] {
	builder.conditions.Add(clause.Lte{
		Column: column,
		Value:  value,
	})
	return builder
}

func (builder *Builder[T]) OrWhereLte(column string, value any) *Builder[T] {
	builder.conditions.Add(clause.Or(clause.Lte{
		Column: column,
		Value:  value,
	}))
	return builder
}

func (builder *Builder[T]) WhereIn(column string, values ...any) *Builder[T] {
	builder.conditions.Add(clause.IN{
		Column: column,
		Values: values,
	})
	return builder
}

func (builder *Builder[T]) OrWhereIn(column string, values ...any) *Builder[T] {
	builder.conditions.Add(clause.Or(clause.IN{
		Column: column,
		Values: values,
	}))
	return builder
}

func (builder *Builder[T]) WhereNotIn(column string, values ...any) *Builder[T] {
	builder.conditions.Add(clause.Not(clause.IN{
		Column: column,
		Values: values,
	}))
	return builder
}

func (builder *Builder[T]) OrWhereNotIn(column string, values ...any) *Builder[T] {
	builder.conditions.Add(clause.Or(clause.Not(clause.IN{
		Column: column,
		Values: values,
	})))
	return builder
}

func (builder *Builder[T]) WhereBetween(column string, values [2]any) *Builder[T] {
	builder.conditions.Add(clause.Expr{
		SQL:  fmt.Sprintf("`%s` BETWEEN ? AND ?", column),
		Vars: values[:],
	})
	return builder
}

func (builder *Builder[T]) OrWhereBetween(column string, values [2]any) *Builder[T] {
	builder.conditions.Add(clause.Or(clause.Expr{
		SQL:  fmt.Sprintf("`%s` BETWEEN ? AND ?", column),
		Vars: values[:],
	}))
	return builder
}

func (builder *Builder[T]) WhereNotBetween(column string, values [2]any) *Builder[T] {
	builder.conditions.Add(clause.Expr{
		SQL:  fmt.Sprintf("`%s` NOT BETWEEN ? AND ?", column),
		Vars: values[:],
	})
	return builder
}

func (builder *Builder[T]) OrWhereNotBetween(column string, values [2]any) *Builder[T] {
	builder.conditions.Add(clause.Or(clause.Expr{
		SQL:  fmt.Sprintf("`%s` NOT BETWEEN ? AND ?", column),
		Vars: values[:],
	}))
	return builder
}

func (builder *Builder[T]) WhereLike(column string, value string) *Builder[T] {
	builder.conditions.Add(clause.Like{
		Column: column,
		Value:  value,
	})
	return builder
}

func (builder *Builder[T]) WhereNotLike(column string, value string) *Builder[T] {
	builder.conditions.Add(clause.Not(clause.Like{
		Column: column,
		Value:  value,
	}))
	return builder
}

func (builder *Builder[T]) OrWhereLike(column string, value string) *Builder[T] {
	builder.conditions.Add(clause.Or(clause.Like{
		Column: column,
		Value:  value,
	}))
	return builder
}

func (builder *Builder[T]) OrWhereNotLike(column string, value string) *Builder[T] {
	builder.conditions.Add(clause.Not(clause.Or(clause.Like{
		Column: column,
		Value:  value,
	})))
	return builder
}

func (builder *Builder[T]) GroupBy(columns ...any) *Builder[T] {
	for _, column := range columns {
		switch value := column.(type) {
		case string:
			builder.groups.Add(clause.Column{
				Name: value,
			})
		case clause.Expr:
			builder.groups.Add(clause.Column{
				Name: value.SQL,
				Raw:  true,
			})
		case clause.Column:
			builder.groups.Add(value)
		case fmt.Stringer:
			builder.groups.Add(clause.Column{
				Name: value.String(),
			})
		default:
			exception.ThrowInvalidParamTypeErr("group by", value)
		}
	}
	return builder
}

func (builder *Builder[T]) Having(column string, value any) *Builder[T] {
	builder.havings.Add(clause.Eq{
		Column: column,
		Value:  value,
	})
	return builder
}

func (builder *Builder[T]) HavingNot(column string, value any) *Builder[T] {
	builder.havings.Add(clause.Not(clause.Eq{
		Column: column,
		Value:  value,
	}))
	return builder
}

func (builder *Builder[T]) OrHaving(column string, value any) *Builder[T] {
	builder.havings.Add(clause.Or(clause.Eq{
		Column: column,
		Value:  value,
	}))
	return builder
}

func (builder *Builder[T]) OrHavingNot(column string, value any) *Builder[T] {
	builder.havings.Add(clause.Or(clause.Not(clause.Eq{
		Column: column,
		Value:  value,
	})))
	return builder
}

func (builder *Builder[T]) HavingRaw(sql string, args ...any) *Builder[T] {
	builder.havings.Add(clause.Expr{
		SQL:  sql,
		Vars: args,
	})
	return builder
}

func (builder *Builder[T]) OrHavingRaw(sql string, args ...any) *Builder[T] {
	builder.havings.Add(clause.Or(clause.Expr{
		SQL:  sql,
		Vars: args,
	}))
	return builder
}

func (builder *Builder[T]) OrHavingRawNot(sql string, args ...any) *Builder[T] {
	builder.havings.Add(clause.Or(clause.Not(clause.Expr{
		SQL:  sql,
		Vars: args,
	})))
	return builder
}

func (builder *Builder[T]) HavingQuery(query *Builder[any]) *Builder[T] {
	builder.havings.Add(clause.And(query.havings.ToArray()...))
	return builder
}

func (builder *Builder[T]) OrHavingQuery(query *Builder[any]) *Builder[T] {
	builder.havings.Add(clause.Or(query.havings.ToArray()...))
	return builder
}

func (builder *Builder[T]) HavingNotQuery(query *Builder[any]) *Builder[T] {
	builder.havings.Add(clause.Not(clause.And(query.havings.ToArray()...)))
	return builder
}

func (builder *Builder[T]) OrHavingNotQuery(query *Builder[any]) *Builder[T] {
	builder.havings.Add(clause.Or(clause.Or(query.havings.ToArray()...)))
	return builder
}

func (builder *Builder[T]) HavingNull(column string) *Builder[T] {
	builder.havings.Add(clause.Eq{
		Column: column,
	})
	return builder
}

func (builder *Builder[T]) OrHavingNull(column string) *Builder[T] {
	builder.havings.Add(clause.Or(clause.Eq{
		Column: column,
	}))
	return builder
}

func (builder *Builder[T]) HavingNotNull(column string) *Builder[T] {
	builder.havings.Add(clause.Not(clause.Eq{
		Column: column,
	}))
	return builder
}

func (builder *Builder[T]) OrHavingNotNull(column string) *Builder[T] {
	builder.havings.Add(clause.Or(clause.Not(clause.Eq{
		Column: column,
	})))
	return builder
}

func (builder *Builder[T]) HavingEq(column string, value any) *Builder[T] {
	builder.havings.Add(clause.Eq{
		Column: column,
		Value:  value,
	})
	return builder
}

func (builder *Builder[T]) OrHavingEq(column string, value any) *Builder[T] {
	builder.havings.Add(clause.Or(clause.Eq{
		Column: column,
		Value:  value,
	}))
	return builder
}

func (builder *Builder[T]) HavingNeq(column string, value any) *Builder[T] {
	builder.havings.Add(clause.Not(clause.Eq{
		Column: column,
		Value:  value,
	}))
	return builder
}

func (builder *Builder[T]) OrHavingNeq(column string, value any) *Builder[T] {
	builder.havings.Add(clause.Or(clause.Not(clause.Eq{
		Column: column,
		Value:  value,
	})))
	return builder
}

func (builder *Builder[T]) HavingGt(column string, value any) *Builder[T] {
	builder.havings.Add(clause.Gt{
		Column: column,
		Value:  value,
	})
	return builder
}

func (builder *Builder[T]) OrHavingGt(column string, value any) *Builder[T] {
	builder.havings.Add(clause.Or(clause.Gt{
		Column: column,
		Value:  value,
	}))
	return builder
}

func (builder *Builder[T]) HavingGte(column string, value any) *Builder[T] {
	builder.havings.Add(clause.Gte{
		Column: column,
		Value:  value,
	})
	return builder
}

func (builder *Builder[T]) OrHavingGte(column string, value any) *Builder[T] {
	builder.havings.Add(clause.Or(clause.Gte{
		Column: column,
		Value:  value,
	}))
	return builder
}

func (builder *Builder[T]) HavingLt(column string, value any) *Builder[T] {
	builder.havings.Add(clause.Lt{
		Column: column,
		Value:  value,
	})
	return builder
}

func (builder *Builder[T]) OrHavingLt(column string, value any) *Builder[T] {
	builder.havings.Add(clause.Or(clause.Lt{
		Column: column,
		Value:  value,
	}))
	return builder
}

func (builder *Builder[T]) HavingLte(column string, value any) *Builder[T] {
	builder.havings.Add(clause.Lte{
		Column: column,
		Value:  value,
	})
	return builder
}

func (builder *Builder[T]) OrHavingLte(column string, value any) *Builder[T] {
	builder.havings.Add(clause.Or(clause.Lte{
		Column: column,
		Value:  value,
	}))
	return builder
}

func (builder *Builder[T]) HavingIn(column string, values ...any) *Builder[T] {
	builder.havings.Add(clause.IN{
		Column: column,
		Values: values,
	})
	return builder
}

func (builder *Builder[T]) OrHavingIn(column string, values ...any) *Builder[T] {
	builder.havings.Add(clause.Or(clause.IN{
		Column: column,
		Values: values,
	}))
	return builder
}

func (builder *Builder[T]) HavingNotIn(column string, values ...any) *Builder[T] {
	builder.havings.Add(clause.Not(clause.IN{
		Column: column,
		Values: values,
	}))
	return builder
}

func (builder *Builder[T]) OrHavingNotIn(column string, values ...any) *Builder[T] {
	builder.havings.Add(clause.Or(clause.Not(clause.IN{
		Column: column,
		Values: values,
	})))
	return builder
}

func (builder *Builder[T]) HavingBetween(column string, values [2]any) *Builder[T] {
	builder.havings.Add(clause.Expr{
		SQL:  fmt.Sprintf("`%s` BETWEEN ? AND ?", column),
		Vars: values[:],
	})
	return builder
}

func (builder *Builder[T]) OrHavingBetween(column string, values [2]any) *Builder[T] {
	builder.havings.Add(clause.Or(clause.Expr{
		SQL:  fmt.Sprintf("`%s` BETWEEN ? AND ?", column),
		Vars: values[:],
	}))
	return builder
}

func (builder *Builder[T]) HavingNotBetween(column string, values [2]any) *Builder[T] {
	builder.havings.Add(clause.Expr{
		SQL:  fmt.Sprintf("`%s` NOT BETWEEN ? AND ?", column),
		Vars: values[:],
	})
	return builder
}

func (builder *Builder[T]) OrHavingNotBetween(column string, values [2]any) *Builder[T] {
	builder.havings.Add(clause.Or(clause.Expr{
		SQL:  fmt.Sprintf("`%s` NOT BETWEEN ? AND ?", column),
		Vars: values[:],
	}))
	return builder
}

func (builder *Builder[T]) HavingLike(column string, value string) *Builder[T] {
	builder.conditions.Add(clause.Like{
		Column: column,
		Value:  value,
	})
	return builder
}

func (builder *Builder[T]) HavingNotLike(column string, value string) *Builder[T] {
	builder.conditions.Add(clause.Not(clause.Like{
		Column: column,
		Value:  value,
	}))
	return builder
}

func (builder *Builder[T]) OrHavingLike(column string, value string) *Builder[T] {
	builder.conditions.Add(clause.Or(clause.Like{
		Column: column,
		Value:  value,
	}))
	return builder
}

func (builder *Builder[T]) OrHavingNotLike(column string, value string) *Builder[T] {
	builder.conditions.Add(clause.Not(clause.Or(clause.Like{
		Column: column,
		Value:  value,
	})))
	return builder
}

func (builder *Builder[T]) Clone() *Builder[T] {
	return &Builder[T]{
		db:         builder.db,
		selects:    builder.selects,
		table:      builder.table,
		joins:      builder.joins,
		conditions: builder.conditions,
		groups:     builder.groups,
		havings:    builder.havings,
		limit:      builder.limit,
		offset:     builder.offset,
	}
}

func (builder *Builder[T]) ToAnyBuilder() *Builder[any] {
	return &Builder[any]{
		db:         builder.db,
		selects:    builder.selects,
		table:      builder.table,
		joins:      builder.joins,
		conditions: builder.conditions,
		groups:     builder.groups,
		havings:    builder.havings,
		limit:      builder.limit,
		offset:     builder.offset,
	}
}

func (builder *Builder[T]) ToSQL() string {
	builder.build()
	return builder.db.Statement.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx
	})
}
