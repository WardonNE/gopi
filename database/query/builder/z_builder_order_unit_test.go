package builder

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm/clause"
)

func TestBuilder_OrderAsc(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}
	t.Run("Builder.OrderAsc(string)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` ORDER BY `id`").WithoutArgs().WillReturnRows(result)
		var dest = make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").OrderAsc("id").Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("Builder.OrderAsc(clause.Column)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` ORDER BY `id`").WithoutArgs().WillReturnRows(result)
		var dest = make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").OrderAsc(clause.Column{
			Name: "id",
		}).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("Builder.OrderAsc(clause.Expr)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` ORDER BY `id`").WithoutArgs().WillReturnRows(result)
		var dest = make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").OrderAsc(clause.Expr{
			SQL: "`id`",
		}).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("Builder.OrderAsc(*Builder)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` ORDER BY (SELECT `id` FROM `departments` LIMIT 1)").WithoutArgs().WillReturnRows(result)
		var dest = make([]map[string]any, 0)
		var subQuery = NewBuilder(mockDB).Table("departments").Select("id").Limit(1)
		err := NewBuilder(mockDB).Table("users").OrderAsc(subQuery).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("Builder.OrderAsc(*gorm.DB)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` ORDER BY (SELECT id FROM `departments` LIMIT 1)").WithoutArgs().WillReturnRows(result)
		var dest = make([]map[string]any, 0)
		var subQuery = mockDB.Table("departments").Select("id").Limit(1)
		err := NewBuilder(mockDB).Table("users").OrderAsc(subQuery).Find(&dest)
		assert.Nil(t, err)
	})
}

func TestBuilder_OrderDesc(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}
	t.Run("Builder.OrderAsc(string)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` ORDER BY `id` DESC").WithoutArgs().WillReturnRows(result)
		var dest = make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").OrderDesc("id").Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("Builder.OrderAsc(clause.Column)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` ORDER BY `id` DESC").WithoutArgs().WillReturnRows(result)
		var dest = make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").OrderDesc(clause.Column{
			Name: "id",
		}).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("Builder.OrderAsc(clause.Expr)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` ORDER BY `id` DESC").WithoutArgs().WillReturnRows(result)
		var dest = make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").OrderDesc(clause.Expr{
			SQL: "`id`",
		}).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("Builder.OrderAsc(*Builder)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` ORDER BY (SELECT `id` FROM `departments` LIMIT 1) DESC").WithoutArgs().WillReturnRows(result)
		var dest = make([]map[string]any, 0)
		var subQuery = NewBuilder(mockDB).Table("departments").Select("id").Limit(1)
		err := NewBuilder(mockDB).Table("users").OrderDesc(subQuery).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("Builder.OrderAsc(*gorm.DB)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` ORDER BY (SELECT id FROM `departments` LIMIT 1) DESC").WithoutArgs().WillReturnRows(result)
		var dest = make([]map[string]any, 0)
		var subQuery = mockDB.Table("departments").Select("id").Limit(1)
		err := NewBuilder(mockDB).Table("users").OrderDesc(subQuery).Find(&dest)
		assert.Nil(t, err)
	})
}
