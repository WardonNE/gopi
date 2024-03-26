package builder

// Limit set limit
func (builder *Builder) Limit(limit int) *Builder {

	builder.db = builder.db.Limit(limit)
	return builder
}

// Offset set offset
func (builder *Builder) Offset(offset int) *Builder {

	builder.db = builder.db.Offset(offset)
	return builder
}
