package builder

import (
	"database/sql"
	"fmt"

	"gorm.io/gorm"
)

// Begin transaction begin
func (builder *Builder) Begin(opts ...*sql.TxOptions) {
	if builder.transactionLevel == 0 {
		builder.db = builder.db.Begin(opts...)
		builder.onTransaction = true
		builder.transactionLevel = 1
	} else {
		builder.db = builder.db.SavePoint(fmt.Sprintf("sp%d", builder.transactionLevel))
		builder.transactionLevel++
	}
}

// Commit transaction commit
func (builder *Builder) Commit() error {
	builder.onTransaction = false
	return builder.db.Commit().Error
}

// Rollback transaction rollback
func (builder *Builder) Rollback() error {
	if builder.transactionLevel == 1 {
		builder.onTransaction = false
		return builder.db.Rollback().Error
	}
	builder.db = builder.db.RollbackTo(fmt.Sprintf("sp%d", builder.transactionLevel-1))
	builder.transactionLevel--
	return builder.db.Error
}

// Transaction execute with transaction
func (builder *Builder) Transaction(callback func(tx *gorm.DB) error, opts ...*sql.TxOptions) error {
	return builder.db.Transaction(callback, opts...)
}
