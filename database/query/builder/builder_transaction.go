package builder

import (
	"database/sql"

	"github.com/wardonne/gopi/support/collection/list"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Begin transaction begin
func (builder *Builder) Begin(opts ...*sql.TxOptions) *Builder {
	return &Builder{
		db:       builder.db.Begin(opts...),
		conn:     builder.db,
		distinct: false,
		selects:  list.NewArrayList[clause.Column](),
		joins:    list.NewArrayList[clause.Join](),
		having:   list.NewArrayList[clause.Expression](),
		stage:    transaction,
	}
}

// Commit transaction commit
func (builder *Builder) Commit() *Builder {
	builder.db = builder.db.Commit()
	return builder
}

// Rollback transaction rollback
func (builder *Builder) Rollback() *Builder {
	builder.db = builder.db.Rollback()
	return builder
}

// Transaction execute with transaction
func (builder *Builder) Transaction(callback func(tx *gorm.DB) error, opts ...*sql.TxOptions) error {

	return builder.db.Transaction(callback, opts...)
}
