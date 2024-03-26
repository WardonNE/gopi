package builder

// QuoteField returns quoted field
func (builder *Builder) QuoteField(field any) string {
	return QuoteField(builder.conn, field)
}

// FormatValue returns formatted value
func (builder *Builder) FormatValue(value any) any {
	return BuildValue(builder.conn, value)
}

// FormatValues returns formatted values
func (builder *Builder) FormatValues(values ...any) []any {
	args := make([]any, 0, len(values))
	for _, value := range values {
		args = append(args, builder.FormatValue(value))
	}
	return args
}
