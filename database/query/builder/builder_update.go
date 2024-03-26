package builder

// Update update all matched records
func (builder *Builder) Update(values any) error {
	return builder.DB().Updates(values).Error
}
