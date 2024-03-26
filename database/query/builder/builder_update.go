package builder

// Update update all matched records
func (builder *Builder) Update(values any) error {
	builder.db = builder.DB().Updates(values)
	return builder.db.Error
}
