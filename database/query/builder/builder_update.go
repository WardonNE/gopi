package builder

// Update update all matched records
func (builder *Builder) Update(values any) error {
	builder.onExecutionFinished = true
	return builder.DB().Updates(values).Error
}
