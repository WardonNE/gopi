package builder

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm/clause"
)

func TestBuilder_Group(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}
	t.Run("Builder.Group(string)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` GROUP BY `department_id`").WithoutArgs().WillReturnRows(result)
		var dest = make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Group("department_id").Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("Builder.Group(clause.Column)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` GROUP BY `department_id`").WithoutArgs().WillReturnRows(result)
		var dest = make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Group(clause.Column{
			Name: "department_id",
		}).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("Builder.Group(fmt.Stringer)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` GROUP BY `department_id`").WithoutArgs().WillReturnRows(result)
		var dest = make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Group(&stringer{Value: "department_id"}).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("Builder.Group(*Builder)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` GROUP BY (SELECT `id` FROM `departments` LIMIT 1)").WithoutArgs().WillReturnRows(result)
		var dest = make([]map[string]any, 0)
		var subQuery = NewBuilder(mockDB).Select("id").Table("departments").Limit(1)
		err := NewBuilder(mockDB).Table("users").Group(subQuery).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("Builder.Group(*gorm.DB)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` GROUP BY (SELECT id FROM `departments` LIMIT 1)").WithoutArgs().WillReturnRows(result)
		var dest = make([]map[string]any, 0)
		var subQuery = mockDB.Select("id").Table("departments").Limit(1)
		err := NewBuilder(mockDB).Table("users").Group(subQuery).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("Builder.Group(clause.Expr)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` GROUP BY (SELECT id FROM departments LIMIT 1)").WithoutArgs().WillReturnRows(result)
		var dest = make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Group(clause.Expr{
			SQL: "(SELECT id FROM departments LIMIT 1)",
		}).Find(&dest)
		assert.Nil(t, err)
	})
}
