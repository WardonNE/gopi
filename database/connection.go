package database

import (
	"gorm.io/gorm"
)

type Connection struct {
	db *gorm.DB
}
