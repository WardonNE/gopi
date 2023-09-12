package migration

type Migrator interface {
	Commit()
	Rollback()
}
