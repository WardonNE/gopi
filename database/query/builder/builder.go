package builder

import (
	"context"
	"strings"

	"github.com/wardonne/gopi/support/collection/list"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

// Builder query builder
type Builder struct {
	db                    *gorm.DB
	conn                  *gorm.DB
	distinct              bool
	selects               *list.ArrayList[clause.Column]
	joins                 *list.ArrayList[clause.Join]
	having                *list.ArrayList[clause.Expression]
	transactionLevel      uint
	onTransaction         bool
	onTransactionBuilding bool
	onExecutionFinished   bool
}

// NewBuilder create a new query builder
//
//	builder := NewBuilder(db)
func NewBuilder(db *gorm.DB) *Builder {
	return &Builder{
		db:      db,
		conn:    db,
		selects: list.NewArrayList[clause.Column](),
		joins:   list.NewArrayList[clause.Join](),
		having:  list.NewArrayList[clause.Expression](),
	}
}

func (builder *Builder) instance() *Builder {
	if builder.onTransaction {
		if builder.onExecutionFinished {
			return &Builder{
				db:               builder.db.Session(&gorm.Session{NewDB: true}),
				conn:             builder.conn,
				selects:          list.NewArrayList[clause.Column](),
				joins:            list.NewArrayList[clause.Join](),
				having:           list.NewArrayList[clause.Expression](),
				transactionLevel: builder.transactionLevel,
				onTransaction:    true,
			}
		} else if builder.onTransactionBuilding {
			return builder
		} else {
			builder.onTransactionBuilding = true
			return builder
		}
	}
	return builder
}

func (builder *Builder) addClauses() *Builder {
	var selectClause = clause.Select{
		Distinct: builder.distinct,
		Columns:  builder.selects.ToArray(),
	}
	var joinClause = clause.From{
		Joins: builder.joins.ToArray(),
	}
	var groupClause = clause.GroupBy{
		Having: builder.having.ToArray(),
	}
	builder.db = builder.db.Clauses(selectClause, joinClause, groupClause)
	return builder
}

// DB returns *gorm.DB
func (builder *Builder) DB() *gorm.DB {
	return builder.addClauses().db
}

// Debug enable debug mode
func (builder *Builder) Debug() *Builder {
	builder.db = builder.db.Session(&gorm.Session{
		Logger: builder.db.Logger.LogMode(logger.Info),
	})
	return builder
}

// DryRun enable dry run mode
func (builder *Builder) DryRun() *Builder {
	builder.db = builder.db.Session(&gorm.Session{
		DryRun:                 true,
		SkipDefaultTransaction: true,
	})
	return builder
}

// WithContext bind context to builder
func (builder *Builder) WithContext(ctx context.Context) *Builder {
	builder.db = builder.db.Session(&gorm.Session{
		Context: ctx,
	})
	return builder
}

// Assign assign attributes
func (builder *Builder) Assign(attrs ...any) *Builder {
	builder = builder.instance()
	builder.db = builder.db.Assign(attrs...)
	return builder
}

// Attrs sets init attributes
func (builder *Builder) Attrs(attrs ...any) *Builder {
	builder = builder.instance()
	builder.db = builder.db.Attrs(attrs...)
	return builder
}

// Clone clone a new [Builder]
func (builder *Builder) Clone() *Builder {
	return NewBuilder(builder.conn)
}

// ToSQL returns sql
func (builder *Builder) ToSQL() string {
	builder.onExecutionFinished = true
	return strings.TrimSpace(builder.DB().ToSQL(func(tx *gorm.DB) *gorm.DB {
		dest := make([]map[string]any, 0)
		return tx.Find(&dest)
	}))
}

// Exec executes insert/update/delete raw sql
func (builder *Builder) Exec(sql string, values ...any) error {
	builder.onExecutionFinished = true
	builder.db = builder.db.Exec(sql, values...)
	return builder.db.Error
}

// Fetch executes select raw sql
func (builder *Builder) Fetch(dest any, sql string, values ...any) error {
	builder.onExecutionFinished = true
	builder.db = builder.db.Exec(sql).Scan(dest)
	return builder.db.Error
}
