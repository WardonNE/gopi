package builder

import (
	"fmt"

	"gorm.io/gorm/clause"
)

// Count counts matched records
func (builder *Builder) Count() (int64, error) {
	builder.selects.Clear()
	builder.onExecutionFinished = true
	var dest int64
	err := builder.DB().Count(&dest).Error
	return dest, err
}

// Sum select sum
func (builder *Builder) Sum(column any) (float64, error) {
	builder.onExecutionFinished = true
	builder.selects.Clear()
	if builder.distinct {
		builder = builder.Select(clause.Expr{
			SQL: fmt.Sprintf("SUM(DISTINCT %s) AS `sum`", builder.QuoteField(column)),
		})
	} else {
		builder = builder.Select(clause.Expr{
			SQL: fmt.Sprintf("SUM(%s) AS `sum`", builder.QuoteField(column)),
		})
	}
	var dest float64
	err := builder.DB().Find(&dest).Error
	return dest, err
}

// Avg select avg
func (builder *Builder) Avg(column any) (float64, error) {
	builder.onExecutionFinished = true
	builder.selects.Clear()
	if builder.distinct {
		builder = builder.Select(clause.Expr{
			SQL: fmt.Sprintf("AVG(DISTINCT %s) AS `avg`", builder.QuoteField(column)),
		})
	} else {
		builder = builder.Select(clause.Expr{
			SQL: fmt.Sprintf("AVG(%s) AS `avg`", builder.QuoteField(column)),
		})
	}
	var dest float64
	err := builder.DB().Find(&dest).Error
	return dest, err
}

// Max select max
func (builder *Builder) Max(column any) (float64, error) {
	builder.onExecutionFinished = true
	builder.selects.Clear()
	if builder.distinct {
		builder = builder.Select(clause.Expr{
			SQL: fmt.Sprintf("MAX(DISTINCT %s) AS `max`", builder.QuoteField(column)),
		})
	} else {
		builder = builder.Select(clause.Expr{
			SQL: fmt.Sprintf("MAX(%s) AS `max`", builder.QuoteField(column)),
		})
	}
	var dest float64
	err := builder.DB().Find(&dest).Error
	return dest, err
}

// Min select min
func (builder *Builder) Min(column any) (float64, error) {
	builder.onExecutionFinished = true
	builder.selects.Clear()
	if builder.distinct {
		builder = builder.Select(clause.Expr{
			SQL: fmt.Sprintf("MIN(DISTINCT %s) AS `min`", builder.QuoteField(column)),
		})
	} else {
		builder = builder.Select(clause.Expr{
			SQL: fmt.Sprintf("MIN(%s) AS `min`", builder.QuoteField(column)),
		})
	}
	var dest float64
	err := builder.DB().Find(&dest).Error
	return dest, err
}
