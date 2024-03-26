package builder

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func TestBuilder_Select(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}

	t.Run("Builder.Select(string)", func(t *testing.T) {
		mock.ExpectQuery("SELECT `id`,`name` FROM `users`").WithoutArgs().WillReturnRows(result)
		var dest = make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Select("id").Select("name").Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("Builder.Select(fmt.Stringer)", func(t *testing.T) {
		mock.ExpectQuery("SELECT `id` FROM `users`").WithoutArgs().WillReturnRows(result)
		var dest = make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Select(&stringer{"id"}).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("Builder.Select(clause.Column)", func(t *testing.T) {
		t.Run("clause.Column with Alias", func(t *testing.T) {
			mock.ExpectQuery("SELECT `id` AS `userId`,`name` FROM `users`").WithoutArgs().WillReturnRows(result)
			var dest = make([]map[string]any, 0)
			err := NewBuilder(mockDB).Table("users").Select(clause.Column{
				Name:  "id",
				Alias: "userId",
			}, clause.Column{
				Name: "name",
			}).Find(&dest)
			assert.Nil(t, err)
		})

		t.Run("clause.Column without Alias", func(t *testing.T) {
			mock.ExpectQuery("SELECT `id` FROM `users`").WithoutArgs().WillReturnRows(result)
			var dest = make([]map[string]any, 0)
			err := NewBuilder(mockDB).Table("users").Select(clause.Column{
				Name: "id",
			}).Find(&dest)
			assert.Nil(t, err)
		})

		t.Run("clause.Column with Raw SQL", func(t *testing.T) {
			mock.ExpectQuery("SELECT `id` AS `userId` FROM `users`").WithoutArgs().WillReturnRows(result)
			var dest = make([]map[string]any, 0)
			err := NewBuilder(mockDB).Table("users").Select(clause.Column{
				Name: "`id` AS `userId`",
				Raw:  true,
			}).Find(&dest)
			assert.Nil(t, err)
		})

		t.Run("clause.Column with Table", func(t *testing.T) {
			mock.ExpectQuery("SELECT `users`.`id` FROM `users`").WithoutArgs().WillReturnRows(result)
			var dest = make([]map[string]any, 0)
			err := NewBuilder(mockDB).Table("users").Select(clause.Column{
				Table: "users",
				Name:  "id",
			}).Find(&dest)
			assert.Nil(t, err)
		})
	})

	t.Run("Builder.Select(clause.Expr)", func(t *testing.T) {
		t.Run("clause.Expr without Vars", func(t *testing.T) {
			mock.ExpectQuery("SELECT `users`.`id` AS `userId` FROM `users`").WithoutArgs().WillReturnRows(result)
			var dest = make([]map[string]any, 0)
			err := NewBuilder(mockDB).Table("users").Select(clause.Expr{
				SQL: "`users`.`id` AS `userId`",
			}).Find(&dest)
			assert.Nil(t, err)
		})

		t.Run("clause.Expr with Vars", func(t *testing.T) {
			mock.ExpectQuery("SELECT (SELECT COUNT(*) FROM departments WHERE status = (SELECT id FROM `departments` LIMIT 1)) AS departmentCount FROM `users`").WithoutArgs().WillReturnRows(result)
			var dest = make([]map[string]any, 0)
			err := NewBuilder(mockDB).Table("users").Select(clause.Expr{
				SQL:  "(SELECT COUNT(*) FROM departments WHERE status = (?)) AS departmentCount",
				Vars: []any{mockDB.Table("departments").Select("id").Limit(1)},
			}).Find(&dest)
			assert.Nil(t, err)
		})

		t.Run("clause.Expr multi fields", func(t *testing.T) {
			mock.ExpectQuery("SELECT `users`.`id` AS `userId`,(SELECT COUNT(*) FROM departments WHERE status = 1) AS departmentCount FROM `users`").WithoutArgs().WillReturnRows(result)
			var dest = make([]map[string]any, 0)
			err := NewBuilder(mockDB).Table("users").Select(
				clause.Expr{
					SQL: "`users`.`id` AS `userId`",
				},
				clause.Expr{
					SQL:  "(SELECT COUNT(*) FROM departments WHERE status = ?) AS departmentCount",
					Vars: []any{1},
				},
			).Find(&dest)
			assert.Nil(t, err)
		})
	})

	t.Run("Builder.Select(clause.NamedExpr)", func(t *testing.T) {
		t.Run("clause.NamedExpr without Vars", func(t *testing.T) {
			mock.ExpectQuery("SELECT `users`.`id` AS `userId` FROM `users`").WithoutArgs().WillReturnRows(result)
			var dest = make([]map[string]any, 0)
			err := NewBuilder(mockDB).Table("users").Select(clause.NamedExpr{
				SQL: "`users`.`id` AS `userId`",
			}).Find(&dest)
			assert.Nil(t, err)
		})
		t.Run("clause.NamedExpr with Vars", func(t *testing.T) {
			mock.ExpectQuery("SELECT (SELECT COUNT(*) FROM departments WHERE status = 1) AS departmentCount FROM `users`").WithoutArgs().WillReturnRows(result)
			var dest = make([]map[string]any, 0)
			err := NewBuilder(mockDB).Table("users").Select(clause.NamedExpr{
				SQL:  "(SELECT COUNT(*) FROM departments WHERE status = @status) AS departmentCount",
				Vars: []any{sql.Named("status", 1)},
			}).Find(&dest)
			assert.Nil(t, err)
		})
	})

	t.Run("Builder.Select(*gorm.DB)", func(t *testing.T) {
		mock.ExpectQuery("SELECT (SELECT COUNT(*) FROM `departments`) FROM `users`").WithoutArgs().WillReturnRows(result)
		var dest = make([]map[string]any, 0)
		subquery := mockDB.Table("departments").Select("COUNT(*)")
		err := NewBuilder(mockDB).Table("users").Select(subquery).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("Builder.Select(*Builder)", func(t *testing.T) {
		mock.ExpectQuery("SELECT (SELECT COUNT(*) FROM `departments`) FROM `users`").WithoutArgs().WillReturnRows(result)
		var dest = make([]map[string]any, 0)
		subquery := NewBuilder(mockDB).Table("departments").Select(clause.Column{
			Name: "COUNT(*)",
			Raw:  true,
		})
		err := NewBuilder(mockDB).Table("users").Select(subquery).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("Builder.Select(Callback)", func(t *testing.T) {
		mock.ExpectQuery("SELECT (SELECT count(*) FROM `departments`) FROM `users`").WithoutArgs().WillReturnRows(result)
		var dest = make([]map[string]any, 0)
		callback := func(tx *gorm.DB) *gorm.DB {
			var count int64
			return tx.Table("departments").Count(&count)
		}
		err := NewBuilder(mockDB).Table("users").Select(callback).Find(&dest)
		assert.Nil(t, err)
	})
}

func TestBuilder_Distinct(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}

	mock.ExpectQuery("SELECT DISTINCT `id` FROM `users`").WithoutArgs().WillReturnRows(result)
	var dest = make([]map[string]any, 0)
	err := NewBuilder(mockDB).Table("users").Distinct("id").Find(&dest)
	assert.Nil(t, err)
}
