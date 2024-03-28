package builder

// Exec executes insert/update/delete raw sql
func (builder *Builder) Exec(sql string, values ...any) (int64, error) {
	builder.onExecutionFinished = true
	tx := builder.db.Exec(sql, values...)
	return tx.RowsAffected, tx.Error
}

// Fetch executes select raw sql
func (builder *Builder) Fetch(dest any, sql string, values ...any) error {
	builder.onExecutionFinished = true
	tx := builder.db.Raw(sql, values...)
	if tx.Error != nil {
		return tx.Error
	}
	return tx.Scan(dest).Error
}
