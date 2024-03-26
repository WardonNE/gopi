package builder

// Delete deletes all matched records
func (builder *Builder) Delete() error {
	builder.db = builder.DB().Delete(nil)
	return builder.db.Error
}
