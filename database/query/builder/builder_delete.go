package builder

// Delete deletes all matched records
func (builder *Builder) Delete() error {
	return builder.DB().Delete(nil).Error
}
