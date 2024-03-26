package builder

// With alias of [Builder.Preload]
func (builder *Builder) With(relation string, args ...any) *Builder {
	builder = builder.instance()
	builder.db = builder.db.Preload(relation, args...)
	return builder
}

// Preload preload relations
func (builder *Builder) Preload(relation string, args ...any) *Builder {
	builder = builder.instance()
	builder.db = builder.db.Preload(relation, args...)
	return builder
}
