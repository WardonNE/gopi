package builder

// Delete deletes all matched records
func (builder *Builder) Delete() error {
	builder.onExecutionFinished = true
	return builder.DB().Delete(nil).Error
}
