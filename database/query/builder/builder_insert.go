package builder

import (
	"fmt"

	"github.com/wardonne/gopi/database/exception"
	"gorm.io/gorm/clause"
)

// FirstOrCreate gets the first record, if not found, create it
func (builder *Builder) FirstOrCreate(dest any) error {
	builder.onExecutionFinished = true
	return builder.DB().FirstOrCreate(dest).Error
}

// FirstOrInit gets the firsst record, if not found, return an inited instance
func (builder *Builder) FirstOrInit(dest any) error {
	builder.onExecutionFinished = true
	return builder.DB().FirstOrInit(dest).Error
}

// Create create a new record
func (builder *Builder) Create(value any) error {
	builder.onExecutionFinished = true
	return builder.DB().Create(value).Error
}

// CreateInBatches inserts values in batches of batchSize
func (builder *Builder) CreateInBatches(values []any, batchSize int) error {
	builder.onExecutionFinished = true
	return builder.DB().CreateInBatches(values, batchSize).Error
}

// Upsert insert new records if unique key is conflict then update it
func (builder *Builder) Upsert(values any, keys []any, updates ...string) error {
	builder.onExecutionFinished = true
	expr := clause.OnConflict{}
	for _, key := range keys {
		switch value := key.(type) {
		case string:
			expr.Columns = append(expr.Columns, clause.Column{Name: value})
		case fmt.Stringer:
			expr.Columns = append(expr.Columns, clause.Column{Name: value.String()})
		case clause.Column:
			expr.Columns = append(expr.Columns, value)
		case clause.Expr:
			expr.Columns = append(expr.Columns, clause.Column{Name: value.SQL, Raw: true})
		default:
			exception.ThrowInvalidParamTypeErr("upsert.key", key)
		}
	}
	if len(updates) == 0 {
		expr.UpdateAll = true
	} else {
		expr.DoUpdates = clause.AssignmentColumns(updates)
	}
	return builder.DB().Clauses(expr).Create(values).Error
}
