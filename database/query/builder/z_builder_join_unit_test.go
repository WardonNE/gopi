package builder

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm/clause"
)

func TestBuilder_LeftJoin(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}
	mock.ExpectQuery("SELECT * FROM `users` LEFT JOIN departments ON departments.id = users.department_id AND departments.status = ?").WithArgs(1).WillReturnRows(result)
	var dest = make([]map[string]any, 0)
	err := NewBuilder(mockDB).Table("users").LeftJoin("departments", "departments.id = users.department_id AND departments.status = ?", 1).Find(&dest)
	assert.Nil(t, err)
}

func TestBuilder_LeftJoinSub(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}
	t.Run("Bulder.LeftJoinSub(*Builder)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `departments` LEFT JOIN (SELECT `department_id`,COUNT(`user_id`) AS `user_number` FROM `users` GROUP BY `department_id`) t ON t.department_id = departments.id").WithoutArgs().WillReturnRows(result)
		var dest = make([]map[string]any, 0)
		var subQuery = NewBuilder(mockDB).Table("users").Select("department_id", clause.Expr{
			SQL: "COUNT(`user_id`) AS `user_number`",
		}).Group("department_id")
		err := NewBuilder(mockDB).Table("departments").LeftJoinSub(subQuery, "t", "t.department_id = departments.id").Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("Builder.LeftJoinSub(*gorm.DB)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `departments` LEFT JOIN (SELECT department_id,COUNT(`user_id`) AS `user_number` FROM `users` GROUP BY `department_id`) t ON t.department_id = departments.id").WithoutArgs().WillReturnRows(result)
		var dest = make([]map[string]any, 0)
		var subQuery = mockDB.Table("users").Select("department_id", "COUNT(`user_id`) AS `user_number`").Group("department_id")
		err := NewBuilder(mockDB).Table("departments").LeftJoinSub(subQuery, "t", "t.department_id = departments.id").Find(&dest)
		assert.Nil(t, err)
	})
}

func TestBuilder_RightJoin(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}
	mock.ExpectQuery("SELECT * FROM `users` LEFT JOIN departments ON departments.id = users.department_id AND departments.status = ?").WithArgs(1).WillReturnRows(result)
	var dest = make([]map[string]any, 0)
	err := NewBuilder(mockDB).Table("users").LeftJoin("departments", "departments.id = users.department_id AND departments.status = ?", 1).Find(&dest)
	assert.Nil(t, err)
}

func TestBuilder_RightJoinSub(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}
	t.Run("Bulder.RightJoinSub(*Builder)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `departments` RIGHT JOIN (SELECT `department_id`,COUNT(`user_id`) AS `user_number` FROM `users` GROUP BY `department_id`) t ON t.department_id = departments.id").WithoutArgs().WillReturnRows(result)
		var dest = make([]map[string]any, 0)
		var subQuery = NewBuilder(mockDB).Table("users").Select("department_id", clause.Expr{
			SQL: "COUNT(`user_id`) AS `user_number`",
		}).Group("department_id")
		err := NewBuilder(mockDB).Table("departments").RightJoinSub(subQuery, "t", "t.department_id = departments.id").Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("Builder.RightJoinSub(*gorm.DB)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `departments` RIGHT JOIN (SELECT department_id,COUNT(`user_id`) AS `user_number` FROM `users` GROUP BY `department_id`) t ON t.department_id = departments.id").WithoutArgs().WillReturnRows(result)
		var dest = make([]map[string]any, 0)
		var subQuery = mockDB.Table("users").Select("department_id", "COUNT(`user_id`) AS `user_number`").Group("department_id")
		err := NewBuilder(mockDB).Table("departments").RightJoinSub(subQuery, "t", "t.department_id = departments.id").Find(&dest)
		assert.Nil(t, err)
	})
}

func TestBuilder_InnerJoin(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}
	mock.ExpectQuery("SELECT * FROM `users` INNER JOIN departments ON departments.id = users.department_id AND departments.status = ?").WithArgs(1).WillReturnRows(result)
	var dest = make([]map[string]any, 0)
	err := NewBuilder(mockDB).Table("users").InnerJoin("departments", "departments.id = users.department_id AND departments.status = ?", 1).Find(&dest)
	assert.Nil(t, err)
}

func TestBuilder_InnerJoinSub(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}
	t.Run("Bulder.InnerJoinSub(*Builder)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `departments` INNER JOIN (SELECT `department_id`,COUNT(`user_id`) AS `user_number` FROM `users` GROUP BY `department_id`) t ON t.department_id = departments.id").WithoutArgs().WillReturnRows(result)
		var dest = make([]map[string]any, 0)
		var subQuery = NewBuilder(mockDB).Table("users").Select("department_id", clause.Expr{
			SQL: "COUNT(`user_id`) AS `user_number`",
		}).Group("department_id")
		err := NewBuilder(mockDB).Table("departments").InnerJoinSub(subQuery, "t", "t.department_id = departments.id").Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("Builder.InnerJoinSub(*gorm.DB)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `departments` INNER JOIN (SELECT department_id,COUNT(`user_id`) AS `user_number` FROM `users` GROUP BY `department_id`) t ON t.department_id = departments.id").WithoutArgs().WillReturnRows(result)
		var dest = make([]map[string]any, 0)
		var subQuery = mockDB.Table("users").Select("department_id", "COUNT(`user_id`) AS `user_number`").Group("department_id")
		err := NewBuilder(mockDB).Table("departments").InnerJoinSub(subQuery, "t", "t.department_id = departments.id").Find(&dest)
		assert.Nil(t, err)
	})
}

func TestBuilder_CrossJoin(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}
	mock.ExpectQuery("SELECT * FROM `users` CROSS JOIN departments ON departments.id = users.department_id AND departments.status = ?").WithArgs(1).WillReturnRows(result)
	var dest = make([]map[string]any, 0)
	err := NewBuilder(mockDB).Table("users").CrossJoin("departments", "departments.id = users.department_id AND departments.status = ?", 1).Find(&dest)
	assert.Nil(t, err)
}

func TestBuilder_CrossJoinSub(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}
	t.Run("Bulder.CrossJoinSub(*Builder)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `departments` CROSS JOIN (SELECT `department_id`,COUNT(`user_id`) AS `user_number` FROM `users` GROUP BY `department_id`) t ON t.department_id = departments.id").WithoutArgs().WillReturnRows(result)
		var dest = make([]map[string]any, 0)
		var subQuery = NewBuilder(mockDB).Table("users").Select("department_id", clause.Expr{
			SQL: "COUNT(`user_id`) AS `user_number`",
		}).Group("department_id")
		err := NewBuilder(mockDB).Table("departments").CrossJoinSub(subQuery, "t", "t.department_id = departments.id").Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("Builder.CrossJoinSub(*gorm.DB)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `departments` CROSS JOIN (SELECT department_id,COUNT(`user_id`) AS `user_number` FROM `users` GROUP BY `department_id`) t ON t.department_id = departments.id").WithoutArgs().WillReturnRows(result)
		var dest = make([]map[string]any, 0)
		var subQuery = mockDB.Table("users").Select("department_id", "COUNT(`user_id`) AS `user_number`").Group("department_id")
		err := NewBuilder(mockDB).Table("departments").CrossJoinSub(subQuery, "t", "t.department_id = departments.id").Find(&dest)
		assert.Nil(t, err)
	})
}

func TestBuilder_WithJoin(t *testing.T) {
	type Department struct {
		ID   uint64
		Name string
	}

	type User struct {
		ID           uint64
		DepartmentID uint64
		Department   *Department
	}
	var result = mock.NewRows([]string{"id", "name"}).AddRow(1, "wardonne")
	mock.ExpectQuery("SELECT * FROM `users` LEFT JOIN `departments` `Department` ON `users`.`department_id` = `Department`.`id`").WithoutArgs().WillReturnRows(result)
	var dest = make([]map[string]any, 0)
	err := NewBuilder(mockDB).Model(new(User)).WithJoin("Department").Find(&dest)
	assert.Nil(t, err)
}
