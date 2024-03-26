package builder

import (
	"github.com/wardonne/gopi/database/exception"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Table add from clause
//
//	builder.Table("users", "u")
//	builder.Table(clause.Expr{SQL: "users AS u"})
//	builder.Table(clause.Table{Name: "users", Alias: "u"})
//	builder.Table(db.Table("users").Where("status", 1), u)
//	builder.Table(NewBuilder(db).Table("users").Where("status", 1), u)
//	builder.Table(func(tx *gorm.DB) *gorm.DB {
//		var dest = make([]map[string]any, 0)
//		return tx.Table("users").Where("status = ?", 1).Find(&dest)
//	})
func (builder *Builder) Table(table any, values ...any) *Builder {

	switch value := table.(type) {
	case *Builder:
		if len(values) == 0 {
			exception.ThrowEmptySubQueryAliasErr()
		}
		var alias string
		var args []any
		switch (values[0]).(type) {
		case string:
			alias = values[0].(string)
			if len(values) > 1 {
				args = values[1:]
			}
		default:
			exception.ThrowInvalidParamTypeErr("Alias", values[0])
		}
		builder.db = builder.db.Table(BuildTable(builder.db, value, alias, args...))
	case *gorm.DB:
		if len(values) == 0 {
			exception.ThrowEmptySubQueryAliasErr()
		}
		var alias string
		var args []any
		switch (values[0]).(type) {
		case string:
			alias = values[0].(string)
			if len(values) > 1 {
				args = values[1:]
			}
		default:
			exception.ThrowInvalidParamTypeErr("Alias", values[0])
		}
		builder.db = builder.db.Table(BuildTable(builder.db, value, alias, args...))
	case Callback:
		if len(values) == 0 {
			exception.ThrowEmptySubQueryAliasErr()
		}
		var alias string
		var args []any
		switch (values[0]).(type) {
		case string:
			alias = values[0].(string)
			if len(values) > 1 {
				args = values[1:]
			}
		default:
			exception.ThrowInvalidParamTypeErr("Alias", values[0])
		}
		builder.db = builder.db.Table(BuildTable(builder.db, value, alias, args...))
	case clause.Expr:
		vars := make([]any, 0, len(value.Vars)+len(values))
		vars = append(vars, value.Vars...)
		vars = append(vars, values...)
		builder.db = builder.db.Table(value.SQL, builder.FormatValues(vars...)...)
	case clause.NamedExpr:
		vars := make([]any, 0, len(value.Vars)+len(values))
		vars = append(vars, value.Vars...)
		vars = append(vars, values...)
		builder.db = builder.db.Table(value.SQL, builder.FormatValues(vars...)...)
	default:
		builder.db = builder.db.Table(BuildTable(builder.db, value, "", values...))
	}
	return builder
}

// Model set table by model instance
//
//	builder.Model(new(User))
//	buulder.Model(&user)
func (builder *Builder) Model(value any) *Builder {

	builder.db = builder.db.Model(value)
	return builder
}
