package builder

import (
	"gorm.io/gorm"
)

// Exists select exists
func (builder *Builder) Exists() (bool, error) {
	builder.onExecutionFinished = true
	var dest = new(struct {
		Result bool
	})
	err := builder.conn.Raw("SELECT EXISTS (?) AS `result`", builder.DB()).Scan(dest).Error
	return dest.Result, err
}

// Take gets the first matched record without specific order
func (builder *Builder) Take(dest any) error {
	builder.onExecutionFinished = true
	return builder.DB().Take(dest).Error
}

// First gets the first matched record order by primary key asc
func (builder *Builder) First(dest any) error {
	builder.onExecutionFinished = true
	return builder.DB().First(dest).Error
}

// Last gets the last matched record order by primary key desc
func (builder *Builder) Last(dest any) error {
	builder.onExecutionFinished = true
	return builder.DB().Last(dest).Error
}

// Find find all matched records
func (builder *Builder) Find(dest any) error {
	builder.onExecutionFinished = true
	return builder.DB().Find(dest).Error
}

// Pluck gets single column from results
func (builder *Builder) Pluck(column any, dest any) error {
	builder.selects.Clear()
	builder = builder.Select(column)
	builder.onExecutionFinished = true
	return builder.DB().Pluck(builder.selects.First().Name, dest).Error
}

// Chunk find all matched records in batches of batchSize
func (builder *Builder) Chunk(dest any, batchSize int, callback func(tx *gorm.DB, batch int) error) error {
	builder.onExecutionFinished = true
	return builder.DB().FindInBatches(dest, batchSize, callback).Error
}

// Cursor iteration
func (builder *Builder) Cursor(dest any, callback func() error) error {
	builder.onExecutionFinished = true
	rows, err := builder.DB().Rows()
	if err != nil {
		return err
	}
	for rows.Next() {
		err = builder.db.ScanRows(rows, dest)
		if err != nil {
			return err
		}
		err = callback()
		if err != nil {
			return err
		}
	}
	return nil
}
