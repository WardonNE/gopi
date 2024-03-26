package builder

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/wardonne/gopi/database/exception"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func TestBuilder_Table(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}
	t.Run("builder.Table(string)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users`").WithoutArgs().WillReturnRows(result)
		var dest = make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.Table(clause.Table)", func(t *testing.T) {
		t.Run("clause.Table without Alias", func(t *testing.T) {
			mock.ExpectQuery("SELECT * FROM `users`").WithoutArgs().WillReturnRows(result)
			var dest = make([]map[string]any, 0)
			err := NewBuilder(mockDB).Table(clause.Table{Name: "users"}).Find(&dest)
			assert.Nil(t, err)
		})

		t.Run("clause.Table with Alias", func(t *testing.T) {
			mock.ExpectQuery("SELECT * FROM `users` `u`").WithoutArgs().WillReturnRows(result)
			var dest = make([]map[string]any, 0)
			err := NewBuilder(mockDB).Table(clause.Table{Name: "users", Alias: "u"}).Find(&dest)
			assert.Nil(t, err)
		})

		t.Run("clause.Table with Raw SQL", func(t *testing.T) {
			mock.ExpectQuery("SELECT * FROM users AS `u1`").WithoutArgs().WillReturnRows(result)
			var dest = make([]map[string]any, 0)
			err := NewBuilder(mockDB).Table(clause.Table{Name: "users AS `u1`", Raw: true}).Find(&dest)
			assert.Nil(t, err)
		})
	})

	t.Run("builder.Table(clause.Expr)", func(t *testing.T) {
		t.Run("clause.Expr without Vars", func(t *testing.T) {
			mock.ExpectQuery("SELECT * FROM `users` `u2`").WithoutArgs().WillReturnRows(result)
			var dest = make([]map[string]any, 0)
			err := NewBuilder(mockDB).Table(clause.Expr{SQL: "`users` `u2`"}).Find(&dest)
			assert.Nil(t, err)
		})

		t.Run("clause.Expr with Vars", func(t *testing.T) {
			mock.ExpectQuery("SELECT * FROM (SELECT * FROM departments WHERE status = ?) AS t").WithArgs(1).WillReturnRows(result)
			var dest = make([]map[string]any, 0)
			err := NewBuilder(mockDB).Table(clause.Expr{
				SQL:  "(SELECT * FROM departments WHERE status = ?) AS t",
				Vars: []any{1},
			}).Find(&dest)
			assert.Nil(t, err)
		})
	})

	t.Run("builder.Table(clause.NamedExpr)", func(t *testing.T) {
		t.Run("clause.NamedExpr without Vars", func(t *testing.T) {
			mock.ExpectQuery("SELECT * FROM `users` `u2`").WithoutArgs().WillReturnRows(result)
			var dest = make([]map[string]any, 0)
			err := NewBuilder(mockDB).Table(clause.NamedExpr{SQL: "`users` `u2`"}).Find(&dest)
			assert.Nil(t, err)
		})

		t.Run("clause.NamedExpr with Vars", func(t *testing.T) {
			mock.ExpectQuery("SELECT * FROM (SELECT * FROM departments WHERE status = @status) AS t").WithArgs(1).WillReturnRows(result)
			var dest = make([]map[string]any, 0)
			err := NewBuilder(mockDB).Table(clause.Expr{
				SQL:  "(SELECT * FROM departments WHERE status = @status) AS t",
				Vars: []any{1},
			}).Find(&dest)
			assert.Nil(t, err)
		})
	})

	t.Run("builder.Table(fmt.Stringer)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users`").WithoutArgs().WillReturnRows(result)
		var dest = make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table(&stringer{"users"}).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.Table(*gorm.DB)", func(t *testing.T) {
		t.Run("*gorm.DB without Alias", func(t *testing.T) {
			var dest = make([]map[string]any, 0)
			assert.PanicsWithError(t, exception.NewEmptySubQueryAliasErr().Error(), func() {
				NewBuilder(mockDB).Table(mockDB.Table("departments").Where("status = ?", 1)).Find(&dest)
			})
		})

		t.Run("*gorm.DB with Invalid Type Alias", func(t *testing.T) {
			var dest = make([]map[string]any, 0)
			assert.PanicsWithError(t, exception.NewInvalidParamTypeErr("Alias", 123).Error(), func() {
				NewBuilder(mockDB).Table(mockDB.Table("departments").Where("status = ?", 1), 123).Find(&dest)
			})
		})

		t.Run("*gorm.DB with Alias", func(t *testing.T) {
			mock.ExpectQuery("SELECT * FROM (SELECT * FROM `departments` WHERE status = 1) AS `t`").WithoutArgs().WillReturnRows(result)
			var dest = make([]map[string]any, 0)
			err := NewBuilder(mockDB).Table(mockDB.Table("departments").Where("status = ?", 1), "t").Find(&dest)
			assert.Nil(t, err)
		})
	})

	t.Run("builder.Table(*Builder)", func(t *testing.T) {
		t.Run("*Builder without Alias", func(t *testing.T) {
			assert.PanicsWithError(t, exception.NewEmptySubQueryAliasErr().Error(), func() {
				var dest = make([]map[string]any, 0)
				NewBuilder(mockDB).Table(NewBuilder(mockDB).Table("departments")).Find(&dest)
			})
		})

		t.Run("*Builder with Invalid Type Alias", func(t *testing.T) {
			assert.PanicsWithError(t, exception.NewInvalidParamTypeErr("Alias", 123).Error(), func() {
				var dest = make([]map[string]any, 0)
				NewBuilder(mockDB).Table(NewBuilder(mockDB).Table("departments"), 123).Find(&dest)
			})
		})

		t.Run("*Builder with Alias", func(t *testing.T) {
			mock.ExpectQuery("SELECT * FROM (SELECT * FROM `departments` WHERE `status` = 1) AS `t`").WithoutArgs().WillReturnRows(result)
			var dest = make([]map[string]any, 0)
			err := NewBuilder(mockDB).Table(NewBuilder(mockDB).Table("departments").Where("status", 1), "t").Find(&dest)
			assert.Nil(t, err)
		})
	})

	t.Run("builder.Table(Callback)", func(t *testing.T) {
		t.Run("Callback without Alias", func(t *testing.T) {
			assert.PanicsWithError(t, exception.NewEmptySubQueryAliasErr().Error(), func() {
				var dest = make([]map[string]any, 0)
				NewBuilder(mockDB).Table(func(tx *gorm.DB) *gorm.DB {
					return tx.Table("departments")
				}).Find(&dest)
			})
		})

		t.Run("Callback with Invalid Type Alias", func(t *testing.T) {
			assert.PanicsWithError(t, exception.NewInvalidParamTypeErr("Alias", 123).Error(), func() {
				var dest = make([]map[string]any, 0)
				NewBuilder(mockDB).Table(func(tx *gorm.DB) *gorm.DB {
					return tx.Table("departments")
				}, 123).Find(&dest)
			})
		})

		t.Run("Callback with Alias", func(t *testing.T) {
			mock.ExpectQuery("SELECT * FROM (SELECT * FROM `departments` WHERE status = 1) AS `t`").WithoutArgs().WillReturnRows(result)
			var dest = make([]map[string]any, 0)
			err := NewBuilder(mockDB).Table(func(tx *gorm.DB) *gorm.DB {
				var dest = make([]map[string]any, 0)
				return tx.Table("departments").Where("status = ?", 1).Find(&dest)
			}, "t").Find(&dest)
			assert.Nil(t, err)
		})
	})
}
