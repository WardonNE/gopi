package builder

import (
	"fmt"
	"time"

	"gorm.io/gorm/clause"
	"gorm.io/hints"
)

// Clauses add clauses
func (builder *Builder) Clauses(clauses ...clause.Expression) *Builder {
	builder = builder.instance()
	builder.db = builder.db.Clauses(clauses...)
	return builder
}

// Hint sets hints
func (builder *Builder) Hint(content string) *Builder {
	builder = builder.instance()
	return builder.Clauses(hints.New(content))
}

// MaxExecutionTime sets max execution time hint
func (builder *Builder) MaxExecutionTime(value time.Duration) *Builder {
	builder = builder.instance()
	return builder.Hint(fmt.Sprintf("MAX_EXECUTION_TIME(%d)", value.Milliseconds()))
}

// UseIndex set use index
func (builder *Builder) UseIndex(names ...string) *Builder {
	builder = builder.instance()
	return builder.Clauses(hints.UseIndex(names...))
}

// IgnoreIndex sets ignore index
func (builder *Builder) IgnoreIndex(names ...string) *Builder {
	builder = builder.instance()
	return builder.Clauses(hints.IgnoreIndex(names...))
}

// ForceIndex sets force index
func (builder *Builder) ForceIndex(names ...string) *Builder {
	builder = builder.instance()
	return builder.Clauses(hints.ForceIndex(names...))
}

// ForceIndexForJoin sets force index for join
func (builder *Builder) ForceIndexForJoin(names ...string) *Builder {
	builder = builder.instance()
	return builder.Clauses(hints.ForceIndex(names...).ForJoin())
}

// ForceIndexForOrderBy sets force index for order by
func (builder *Builder) ForceIndexForOrderBy(names ...string) *Builder {
	builder = builder.instance()
	return builder.Clauses(hints.ForceIndex(names...).ForOrderBy())
}

// ForceIndexForGroupBy sets force index for group by
func (builder *Builder) ForceIndexForGroupBy(names ...string) *Builder {
	builder = builder.instance()
	return builder.Clauses(hints.ForceIndex(names...).ForGroupBy())
}
