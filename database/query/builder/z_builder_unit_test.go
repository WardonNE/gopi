package builder

import (
	"database/sql"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	mock      sqlmock.Sqlmock
	mockDB    *gorm.DB
	mockUsers = []map[string]any{
		{"id": 1, "name": "user1"},
		{"id": 2, "name": "user2"},
		{"id": 3, "name": "user3"},
		{"id": 4, "name": "user4"},
		{"id": 5, "name": "user5"},
		{"id": 6, "name": "user6"},
		{"id": 7, "name": "user7"},
		{"id": 8, "name": "user8"},
		{"id": 9, "name": "user9"},
		{"id": 10, "name": "user10"},
		{"id": 11, "name": "user11"},
		{"id": 12, "name": "user12"},
		{"id": 13, "name": "user13"},
	}
)

type stringer struct {
	Value string
}

func (s *stringer) String() string {
	return s.Value
}

func TestMain(m *testing.M) {
	var (
		err error
		db  *sql.DB
	)
	db, mock, err = sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		panic(err)
	}
	mockDB, err = gorm.Open(mysql.New(mysql.Config{Conn: db, SkipInitializeWithVersion: true}), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(logger.Info),
	})
	os.Exit(m.Run())
}
