package database

import (
	"sync"

	"gorm.io/gorm"
)

var databases = &databaseContainer{
	mu:  new(sync.RWMutex),
	dbs: make(map[string]*gorm.DB),
}

type databaseContainer struct {
	mu  *sync.RWMutex
	dbs map[string]*gorm.DB
}

func (dc *databaseContainer) Store(key string, db *gorm.DB) {
	databases.dbs[key] = db
}

func (dc *databaseContainer) Load(key string) *gorm.DB {
	return dc.dbs[key]
}

func RegisterDatabase(alias string, dialector gorm.Dialector, opts ...gorm.Option) error {
	if db, err := gorm.Open(dialector, opts...); err != nil {
		return err
	} else {
		databases.mu.Lock()
		defer databases.mu.Unlock()
		databases.Store(alias, db)
		return nil
	}
}

func DB(alias ...string) *gorm.DB {
	key := "default"
	if len(alias) > 0 {
		key = alias[0]
	}
	databases.mu.RLock()
	defer databases.mu.RUnlock()
	return databases.Load(key)
}
