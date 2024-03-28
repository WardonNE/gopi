package builder

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/wardonne/gopi/database/exception"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Callback query callback
type Callback = func(tx *gorm.DB) *gorm.DB

// Clause builder clause
type Clause = func(builder *Builder) *Builder

// QuoteField returns quoted field
func QuoteField(db *gorm.DB, field any) string {
	switch value := any(field).(type) {
	case fmt.Stringer:
		return db.Statement.Quote(value.String())
	case Callback:
		return fmt.Sprintf("(%s)", db.ToSQL(value))
	case *Builder:
		return fmt.Sprintf("(%s)", value.ToSQL())
	case *gorm.DB:
		callback := func(tx *gorm.DB) *gorm.DB {
			dest := make([]map[string]any, 0)
			return tx.Find(&dest)
		}
		return fmt.Sprintf("(%s)", value.ToSQL(callback))
	case *datatypes.JSONQueryExpression:
		return db.Clauses().ToSQL(func(tx *gorm.DB) *gorm.DB {
			stmt := tx.Statement
			value.Build(stmt)
			return tx
		})
	default:
		return db.Statement.Quote(value)
	}
}

// BuildValue returns formatted value
func BuildValue(db *gorm.DB, value any) any {
	switch v := any(value).(type) {
	case *Builder:
		return v.addClauses().db
	case Callback:
		return v(NewBuilder(db).DryRun().db)
	case sql.NamedArg:
		return sql.NamedArg{Name: v.Name, Value: BuildValue(db, v.Value)}
	default:
		return v
	}
}

// BuildTable returns quoted table
func BuildTable(db *gorm.DB, table any, alias string, values ...any) string {
	alias = strings.TrimSpace(alias)
	switch value := table.(type) {
	case Callback:
		if alias == "" {
			exception.ThrowEmptySubQueryAliasErr()
		}
		return fmt.Sprintf("(%s) AS %s", db.ToSQL(value), QuoteField(db, alias))
	case *Builder:
		if alias == "" {
			exception.ThrowEmptySubQueryAliasErr()
		}
		callback := func(tx *gorm.DB) *gorm.DB {
			var dest = make([]map[string]any, 0)
			return tx.Find(&dest)
		}
		return fmt.Sprintf("(%s) AS %s", value.db.ToSQL(callback), QuoteField(db, alias))
	case *gorm.DB:
		if alias == "" {
			exception.ThrowEmptySubQueryAliasErr()
		}
		callback := func(tx *gorm.DB) *gorm.DB {
			var dest = make([]map[string]any, 0)
			return tx.Find(&dest)
		}
		return fmt.Sprintf("(%s) AS %s", value.ToSQL(callback), QuoteField(db, alias))
	case string:
		if alias == "" {
			return QuoteField(db, value)
		}
		return fmt.Sprintf("%s AS %s", QuoteField(db, value), QuoteField(db, alias))
	case fmt.Stringer:
		if alias == "" {
			return QuoteField(db, value.String())
		}
		return fmt.Sprintf("%s AS %s", QuoteField(db, value.String()), QuoteField(db, alias))
	case clause.Table:
		if alias != "" {
			value.Alias = alias
		}
		return QuoteField(db, value)
	case clause.Expr:
		for _, v := range values {
			value.Vars = append(value.Vars, BuildValue(db, v))
		}
		return QuoteField(db, value)
	default:
		exception.ThrowInvalidParamTypeErr("Table", value)
	}
	return ""
}
