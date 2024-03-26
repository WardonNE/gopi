package builder

import "gorm.io/gorm"

// Scopes add scopes
func (builder *Builder) Scopes(scopes ...func(tx *gorm.DB) *gorm.DB) *Builder {

	builder.db = builder.db.Scopes(scopes...)
	return builder
}
