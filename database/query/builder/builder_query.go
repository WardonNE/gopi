package builder

import (
	"gorm.io/gorm"
)

// Count counts matched records
func (builder *Builder) Count(dest *int64) error {
	builder.selects.Clear()
	return builder.DB().Count(dest).Error
}

// Take gets the first matched record without specific order
func (builder *Builder) Take(dest any) error {
	return builder.DB().Take(dest).Error
}

// First gets the first matched record order by primary key asc
func (builder *Builder) First(dest any) error {
	return builder.DB().First(dest).Error
}

// Last gets the last matched record order by primary key desc
func (builder *Builder) Last(dest any) error {
	return builder.DB().Last(dest).Error
}

// Find find all matched records
func (builder *Builder) Find(dest any) error {
	return builder.DB().Find(dest).Error
}

// Pluck gets single column from results
func (builder *Builder) Pluck(column any, dest any) error {
	builder.selects.Clear()
	builder = builder.Select(column)
	return builder.DB().Pluck(builder.selects.First().Name, dest).Error
}

// Chunk find all matched records in batches of batchSize
func (builder *Builder) Chunk(dest any, batchSize int, callback func(tx *gorm.DB, batch int) error) error {
	return builder.DB().FindInBatches(dest, batchSize, callback).Error
}

// Cursor iteration
func (builder *Builder) Cursor(dest any, callback func() error) error {
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
