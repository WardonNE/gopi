package builder

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func TestBuilder_Having(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}

	t.Run("builder.Having(column = value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `id` = ?").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Having("id", 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.Having(column In values)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `id` IN (?,?,?)").WithArgs(1, 2, 3).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Having("id", []int{1, 2, 3}).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.Having(column IS Null)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `id` IS NULL").WithoutArgs().WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Having("id", nil).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.Having(column = *gorm.DB)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `id` = (SELECT COUNT(*) FROM `departments`)").WithoutArgs().WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Having("id", mockDB.Table("departments").Select("COUNT(*)")).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.Having(column = *Builder)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `id` = (SELECT COUNT(*) FROM `departments` )").WithoutArgs().WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Having("id", NewBuilder(mockDB).Table("departments").Select(clause.Column{
			Name: "COUNT(*)",
			Raw:  true,
		})).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.Having(column = Callback)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `id` = (SELECT count(*) FROM `departments`)").WithoutArgs().WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Having("id", func(tx *gorm.DB) *gorm.DB {
			var count int64
			return tx.Table("departments").Count(&count)
		}).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.Having(*gorm.DB = value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING (SELECT COUNT(*) FROM `departments`) = ?").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Having(mockDB.Table("departments").Select("COUNT(*)"), 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.Having(*Builder = value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING (SELECT COUNT(*) FROM `departments`) = ?").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Having(
			NewBuilder(mockDB).Table("departments").Select(clause.Column{
				Name: "COUNT(*)",
				Raw:  true,
			}),
			1,
		).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.Having(Callback = value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING (SELECT count(*) FROM `departments`) = ?").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Having(
			func(tx *gorm.DB) *gorm.DB {
				var count int64
				return tx.Table("departments").Count(&count)
			},
			1,
		).Find(&dest)
		assert.Nil(t, err)
	})
}

func TestBuilder_HavingNot(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}

	t.Run("builder.HavingNot(column <> value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `id` <> ?").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").HavingNot("id", 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.HavingNot(column Not In values)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `id` NOT IN (?,?,?)").WithArgs(1, 2, 3).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").HavingNot("id", []int{1, 2, 3}).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.HavingNot(column IS NOT Null)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `id` IS NOT NULL").WithoutArgs().WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").HavingNot("id", nil).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.HavingNot(column <> *gorm.DB)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `id` <> (SELECT COUNT(*) FROM `departments`)").WithoutArgs().WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").HavingNot("id", mockDB.Table("departments").Select("COUNT(*)")).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.HavingNot(column <> *Builder)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `id` <> (SELECT COUNT(*) FROM `departments` )").WithoutArgs().WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").HavingNot("id", NewBuilder(mockDB).Table("departments").Select(clause.Column{
			Name: "COUNT(*)",
			Raw:  true,
		})).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.HavingNot(column <> Callback)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `id` <> (SELECT count(*) FROM `departments`)").WithoutArgs().WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").HavingNot("id", func(tx *gorm.DB) *gorm.DB {
			var count int64
			return tx.Table("departments").Count(&count)
		}).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.HavingNot(*gorm.DB <> value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING (SELECT COUNT(*) FROM `departments`) <> ?").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").HavingNot(mockDB.Table("departments").Select("COUNT(*)"), 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.HavingNot(*Builder = value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING (SELECT COUNT(*) FROM `departments`) <> ?").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").HavingNot(
			NewBuilder(mockDB).Table("departments").Select(clause.Column{
				Name: "COUNT(*)",
				Raw:  true,
			}),
			1,
		).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.HavingNot(Callback <> value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING (SELECT count(*) FROM `departments`) <> ?").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").HavingNot(
			func(tx *gorm.DB) *gorm.DB {
				var count int64
				return tx.Table("departments").Count(&count)
			},
			1,
		).Find(&dest)
		assert.Nil(t, err)
	})
}

func TestBuilder_OrHaving(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}

	t.Run("builder.OrHaving(column = value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `status` = ? OR `id` = ?").WithArgs(1, 1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Having("status", 1).OrHaving("id", 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrHaving(column In values)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `status` = ? OR `id` IN (?,?,?)").WithArgs(1, 1, 2, 3).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Having("status", 1).OrHaving("id", []int{1, 2, 3}).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrHaving(column IS Null)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `status` = ? OR `id` IS NULL").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Having("status", 1).OrHaving("id", nil).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrHaving(column = *gorm.DB)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `status` = ? OR `id` = (SELECT COUNT(*) FROM `departments`)").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Having("status", 1).OrHaving("id", mockDB.Table("departments").Select("COUNT(*)")).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrHaving(column = *Builder)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `status` = ? OR `id` = (SELECT COUNT(*) FROM `departments` )").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Having("status", 1).OrHaving("id", NewBuilder(mockDB).Table("departments").Select(clause.Column{
			Name: "COUNT(*)",
			Raw:  true,
		})).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrHaving(column = Callback)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `status` = ? OR `id` = (SELECT count(*) FROM `departments`)").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Having("status", 1).OrHaving("id", func(tx *gorm.DB) *gorm.DB {
			var count int64
			return tx.Table("departments").Count(&count)
		}).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrHaving(*gorm.DB = value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `status` = ? OR (SELECT COUNT(*) FROM `departments`) = ?").WithArgs(1, 1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Having("status", 1).OrHaving(mockDB.Table("departments").Select("COUNT(*)"), 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrHaving(*Builder = value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `status` = ? OR (SELECT COUNT(*) FROM `departments`) = ?").WithArgs(1, 1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Having("status", 1).OrHaving(
			NewBuilder(mockDB).Table("departments").Select(clause.Column{
				Name: "COUNT(*)",
				Raw:  true,
			}),
			1,
		).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrHaving(Callback = value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `status` = ? OR (SELECT count(*) FROM `departments`) = ?").WithArgs(1, 1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Having("status", 1).OrHaving(
			func(tx *gorm.DB) *gorm.DB {
				var count int64
				return tx.Table("departments").Count(&count)
			},
			1,
		).Find(&dest)
		assert.Nil(t, err)
	})
}

func TestBuilder_OrHavingNot(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}

	t.Run("builder.OrHavingNot(column <> value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `status` = ? OR `id` <> ?").WithArgs(1, 1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Having("status", 1).OrHavingNot("id", 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrHavingNot(column Not In values)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `status` = ? OR `id` NOT IN (?,?,?)").WithArgs(1, 1, 2, 3).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Having("status", 1).OrHavingNot("id", []int{1, 2, 3}).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrHavingNot(column IS NOT Null)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `status` = ? OR `id` IS NOT NULL").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Having("status", 1).OrHavingNot("id", nil).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrHavingNot(column <> *gorm.DB)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `status` = ? OR `id` <> (SELECT COUNT(*) FROM `departments`)").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Having("status", 1).OrHavingNot("id", mockDB.Table("departments").Select("COUNT(*)")).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrHavingNot(column <> *Builder)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `status` = ? OR `id` <> (SELECT COUNT(*) FROM `departments` )").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Having("status", 1).OrHavingNot("id", NewBuilder(mockDB).Table("departments").Select(clause.Column{
			Name: "COUNT(*)",
			Raw:  true,
		})).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrHavingNot(column <> Callback)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `status` = ? OR `id` <> (SELECT count(*) FROM `departments`)").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Having("status", 1).OrHavingNot("id", func(tx *gorm.DB) *gorm.DB {
			var count int64
			return tx.Table("departments").Count(&count)
		}).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrHavingNot(*gorm.DB <> value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `status` = ? OR (SELECT COUNT(*) FROM `departments`) <> ?").WithArgs(1, 1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Having("status", 1).OrHavingNot(mockDB.Table("departments").Select("COUNT(*)"), 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrHavingNot(*Builder <> value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `status` = ? OR (SELECT COUNT(*) FROM `departments`) <> ?").WithArgs(1, 1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Having("status", 1).OrHavingNot(
			NewBuilder(mockDB).Table("departments").Select(clause.Column{
				Name: "COUNT(*)",
				Raw:  true,
			}),
			1,
		).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrHavingNot(Callback <> value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `status` = ? OR (SELECT count(*) FROM `departments`) <> ?").WithArgs(1, 1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Having("status", 1).OrHavingNot(
			func(tx *gorm.DB) *gorm.DB {
				var count int64
				return tx.Table("departments").Count(&count)
			},
			1,
		).Find(&dest)
		assert.Nil(t, err)
	})
}

func TestBuilder_HavingNull(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}
	t.Run("builder.HavingNull(string IS NULL)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `id` IS NULL").WithoutArgs().WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").HavingNull("id").Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.HavingNull(*gorm.DB IS NULL)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING (SELECT COUNT(*) FROM `departments`) IS NULL").WithoutArgs().WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").HavingNull(mockDB.Table("departments").Select("COUNT(*)")).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.HavingNull(*Builder IS NULL)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING (SELECT COUNT(*) FROM `departments`) IS NULL").WithoutArgs().WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").HavingNull(NewBuilder(mockDB).Table("departments").Select(clause.Column{
			Name: "COUNT(*)",
			Raw:  true,
		})).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.HavingNull(Callback IS NULL)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING (SELECT count(*) FROM `departments`) IS NULL").WithoutArgs().WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").HavingNull(func(tx *gorm.DB) *gorm.DB {
			var count int64
			return tx.Table("departments").Count(&count)
		}).Find(&dest)
		assert.Nil(t, err)
	})
}

func TestBuilder_HavingNotNull(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}
	t.Run("builder.HavingNotNull(string IS NOT NULL)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `id` IS NOT NULL").WithoutArgs().WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").HavingNotNull("id").Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.HavingNotNull(*gorm.DB IS NOT NULL)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING (SELECT COUNT(*) FROM `departments`) IS NOT NULL").WithoutArgs().WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").HavingNotNull(mockDB.Table("departments").Select("COUNT(*)")).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.HavingNotNull(*Builder IS NOT NULL)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING (SELECT COUNT(*) FROM `departments`) IS NOT NULL").WithoutArgs().WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").HavingNotNull(NewBuilder(mockDB).Table("departments").Select(clause.Column{
			Name: "COUNT(*)",
			Raw:  true,
		})).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.HavingNotNull(Callback IS NOT NULL)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING (SELECT count(*) FROM `departments`) IS NOT NULL").WithoutArgs().WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").HavingNotNull(func(tx *gorm.DB) *gorm.DB {
			var count int64
			return tx.Table("departments").Count(&count)
		}).Find(&dest)
		assert.Nil(t, err)
	})
}

func TestBuilder_OrHavingNull(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}
	t.Run("builder.OrHavingNull(string IS NULL)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `status` = ? OR `id` IS NULL").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Having("status", 1).OrHavingNull("id").Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrHavingNull(*gorm.DB IS NULL)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `id` = ? OR (SELECT COUNT(*) FROM `departments`) IS NULL").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Having("id", 1).OrHavingNull(mockDB.Table("departments").Select("COUNT(*)")).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrHavingNull(*Builder IS NULL)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `id` = ? OR (SELECT COUNT(*) FROM `departments`) IS NULL").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Having("id", 1).OrHavingNull(NewBuilder(mockDB).Table("departments").Select(clause.Column{
			Name: "COUNT(*)",
			Raw:  true,
		})).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrHavingNull(Callback IS NULL)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `id` = ? OR (SELECT count(*) FROM `departments`) IS NULL").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Having("id", 1).OrHavingNull(func(tx *gorm.DB) *gorm.DB {
			var count int64
			return tx.Table("departments").Count(&count)
		}).Find(&dest)
		assert.Nil(t, err)
	})
}

func TestBuilder_OrHavingNotNull(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}
	t.Run("builder.OrHavingNotNull(string IS NOT NULL)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `status` = ? OR `id` IS NOT NULL").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Having("status", 1).OrHavingNotNull("id").Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrHavingNotNull(*gorm.DB IS NOT NULL)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `id` = ? OR (SELECT COUNT(*) FROM `departments`) IS NOT NULL").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Having("id", 1).OrHavingNotNull(mockDB.Table("departments").Select("COUNT(*)")).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrHavingNotNull(*Builder IS NOT NULL)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `id` = ? OR (SELECT COUNT(*) FROM `departments`) IS NOT NULL").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Having("id", 1).OrHavingNotNull(NewBuilder(mockDB).Table("departments").Select(clause.Column{
			Name: "COUNT(*)",
			Raw:  true,
		})).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrHavingNotNull(Callback IS NOT NULL)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `id` = ? OR (SELECT count(*) FROM `departments`) IS NOT NULL").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Having("id", 1).OrHavingNotNull(func(tx *gorm.DB) *gorm.DB {
			var count int64
			return tx.Table("departments").Count(&count)
		}).Find(&dest)
		assert.Nil(t, err)
	})
}

func TestBuilder_HavingEq(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}
	t.Run("builder.HavingEq(string = value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `id` = ?").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").HavingEq("id", 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.HavingEq(*gorm.DB = value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING (SELECT COUNT(*) FROM `departments`) = ?").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").HavingEq(mockDB.Table("departments").Select("COUNT(*)"), 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.HavingEq(*Builder = value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING (SELECT COUNT(*) FROM `departments`) = ?").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").HavingEq(NewBuilder(mockDB).Table("departments").Select(clause.Column{
			Name: "COUNT(*)",
			Raw:  true,
		}), 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.HavingEq(Callback = value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING (SELECT count(*) FROM `departments`) = ?").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").HavingEq(func(tx *gorm.DB) *gorm.DB {
			var count int64
			return tx.Table("departments").Count(&count)
		}, 1).Find(&dest)
		assert.Nil(t, err)
	})
}

func TestBuilder_OrHavingEq(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}
	t.Run("builder.OrHavingEq(string = value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `status` = ? OR `id` = ?").WithArgs(1, 1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Having("status", 1).OrHavingEq("id", 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrHavingEq(*gorm.DB = value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `status` = ? OR (SELECT COUNT(*) FROM `departments`) = ?").WithArgs(1, 1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Having("status", 1).OrHavingEq(mockDB.Table("departments").Select("COUNT(*)"), 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrHavingEq(*Builder = value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `status` = ? OR (SELECT COUNT(*) FROM `departments`) = ?").WithArgs(1, 1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Having("status", 1).OrHavingEq(NewBuilder(mockDB).Table("departments").Select(clause.Column{
			Name: "COUNT(*)",
			Raw:  true,
		}), 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrHavingEq(Callback = value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `status` = ? OR (SELECT count(*) FROM `departments`) = ?").WithArgs(1, 1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Having("status", 1).OrHavingEq(func(tx *gorm.DB) *gorm.DB {
			var count int64
			return tx.Table("departments").Count(&count)
		}, 1).Find(&dest)
		assert.Nil(t, err)
	})
}

func TestBuilder_HavingNeq(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}
	t.Run("builder.HavingNeq(string <> value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `id` <> ?").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").HavingNeq("id", 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.HavingNeq(*gorm.DB <> value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING (SELECT COUNT(*) FROM `departments`) <> ?").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").HavingNeq(mockDB.Table("departments").Select("COUNT(*)"), 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.HavingNeq(*Builder <> value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING (SELECT COUNT(*) FROM `departments`) <> ?").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").HavingNeq(NewBuilder(mockDB).Table("departments").Select(clause.Column{
			Name: "COUNT(*)",
			Raw:  true,
		}), 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.HavingNeq(Callback <> value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING (SELECT count(*) FROM `departments`) <> ?").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").HavingNeq(func(tx *gorm.DB) *gorm.DB {
			var count int64
			return tx.Table("departments").Count(&count)
		}, 1).Find(&dest)
		assert.Nil(t, err)
	})
}

func TestBuilder_OrHavingNeq(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}
	t.Run("builder.OrHavingNeq(string <> value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `status` = ? OR `id` <> ?").WithArgs(1, 1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Having("status", 1).OrHavingNeq("id", 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrHavingNeq(*gorm.DB <> value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `status` = ? OR (SELECT COUNT(*) FROM `departments`) <> ?").WithArgs(1, 1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Having("status", 1).OrHavingNeq(mockDB.Table("departments").Select("COUNT(*)"), 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrHavingNeq(*Builder <> value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `status` = ? OR (SELECT COUNT(*) FROM `departments`) <> ?").WithArgs(1, 1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Having("status", 1).OrHavingNeq(NewBuilder(mockDB).Table("departments").Select(clause.Column{
			Name: "COUNT(*)",
			Raw:  true,
		}), 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrHavingNeq(Callback <> value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `status` = ? OR (SELECT count(*) FROM `departments`) <> ?").WithArgs(1, 1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Having("status", 1).OrHavingNeq(func(tx *gorm.DB) *gorm.DB {
			var count int64
			return tx.Table("departments").Count(&count)
		}, 1).Find(&dest)
		assert.Nil(t, err)
	})
}

func TestBuilder_HavingGt(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}
	t.Run("builder.HavingGt(string > value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `id` > ?").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").HavingGt("id", 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.HavingGt(*gorm.DB > value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING (SELECT COUNT(*) FROM `departments`) > ?").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").HavingGt(mockDB.Table("departments").Select("COUNT(*)"), 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.HavingGt(*Builder > value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING (SELECT COUNT(*) FROM `departments`) > ?").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").HavingGt(NewBuilder(mockDB).Table("departments").Select(clause.Column{
			Name: "COUNT(*)",
			Raw:  true,
		}), 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.HavingGt(Callback > value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING (SELECT count(*) FROM `departments`) > ?").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").HavingGt(func(tx *gorm.DB) *gorm.DB {
			var count int64
			return tx.Table("departments").Count(&count)
		}, 1).Find(&dest)
		assert.Nil(t, err)
	})
}

func TestBuilder_OrHavingGt(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}
	t.Run("builder.OrHavingGt(string > value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `status` = ? OR `id` > ?").WithArgs(1, 1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Having("status", 1).OrHavingGt("id", 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrHavingGt(*gorm.DB > value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `status` = ? OR (SELECT COUNT(*) FROM `departments`) > ?").WithArgs(1, 1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Having("status", 1).OrHavingGt(mockDB.Table("departments").Select("COUNT(*)"), 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrHavingGt(*Builder > value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `status` = ? OR (SELECT COUNT(*) FROM `departments`) > ?").WithArgs(1, 1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Having("status", 1).OrHavingGt(NewBuilder(mockDB).Table("departments").Select(clause.Column{
			Name: "COUNT(*)",
			Raw:  true,
		}), 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrHavingGt(Callback > value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `status` = ? OR (SELECT count(*) FROM `departments`) > ?").WithArgs(1, 1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Having("status", 1).OrHavingGt(func(tx *gorm.DB) *gorm.DB {
			var count int64
			return tx.Table("departments").Count(&count)
		}, 1).Find(&dest)
		assert.Nil(t, err)
	})
}

func TestBuilder_HavingGte(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}
	t.Run("builder.HavingGte(string >= value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `id` >= ?").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").HavingGte("id", 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.HavingGte(*gorm.DB >= value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING (SELECT COUNT(*) FROM `departments`) >= ?").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").HavingGte(mockDB.Table("departments").Select("COUNT(*)"), 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.HavingGte(*Builder >= value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING (SELECT COUNT(*) FROM `departments`) >= ?").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").HavingGte(NewBuilder(mockDB).Table("departments").Select(clause.Column{
			Name: "COUNT(*)",
			Raw:  true,
		}), 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.HavingGte(Callback >= value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING (SELECT count(*) FROM `departments`) >= ?").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").HavingGte(func(tx *gorm.DB) *gorm.DB {
			var count int64
			return tx.Table("departments").Count(&count)
		}, 1).Find(&dest)
		assert.Nil(t, err)
	})
}

func TestBuilder_OrHavingGte(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}
	t.Run("builder.OrHavingGte(string >= value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `status` = ? OR `id` >= ?").WithArgs(1, 1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Having("status", 1).OrHavingGte("id", 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrHavingGte(*gorm.DB >= value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `status` = ? OR (SELECT COUNT(*) FROM `departments`) >= ?").WithArgs(1, 1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Having("status", 1).OrHavingGte(mockDB.Table("departments").Select("COUNT(*)"), 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrHavingGte(*Builder >= value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `status` = ? OR (SELECT COUNT(*) FROM `departments`) >= ?").WithArgs(1, 1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Having("status", 1).OrHavingGte(NewBuilder(mockDB).Table("departments").Select(clause.Column{
			Name: "COUNT(*)",
			Raw:  true,
		}), 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrHavingGte(Callback >= value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `status` = ? OR (SELECT count(*) FROM `departments`) >= ?").WithArgs(1, 1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Having("status", 1).OrHavingGte(func(tx *gorm.DB) *gorm.DB {
			var count int64
			return tx.Table("departments").Count(&count)
		}, 1).Find(&dest)
		assert.Nil(t, err)
	})
}

func TestBuilder_HavingLt(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}
	t.Run("builder.HavingLt(string < value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `id` < ?").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").HavingLt("id", 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.HavingLt(*gorm.DB < value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING (SELECT COUNT(*) FROM `departments`) < ?").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").HavingLt(mockDB.Table("departments").Select("COUNT(*)"), 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.HavingLt(*Builder < value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING (SELECT COUNT(*) FROM `departments`) < ?").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").HavingLt(NewBuilder(mockDB).Table("departments").Select(clause.Column{
			Name: "COUNT(*)",
			Raw:  true,
		}), 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.HavingLt(Callback < value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING (SELECT count(*) FROM `departments`) < ?").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").HavingLt(func(tx *gorm.DB) *gorm.DB {
			var count int64
			return tx.Table("departments").Count(&count)
		}, 1).Find(&dest)
		assert.Nil(t, err)
	})
}

func TestBuilder_OrHavingLt(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}
	t.Run("builder.OrHavingLt(string < value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `status` = ? OR `id` < ?").WithArgs(1, 1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Having("status", 1).OrHavingLt("id", 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrHavingLt(*gorm.DB < value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `status` = ? OR (SELECT COUNT(*) FROM `departments`) < ?").WithArgs(1, 1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Having("status", 1).OrHavingLt(mockDB.Table("departments").Select("COUNT(*)"), 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrHavingLt(*Builder < value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `status` = ? OR (SELECT COUNT(*) FROM `departments`) < ?").WithArgs(1, 1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Having("status", 1).OrHavingLt(NewBuilder(mockDB).Table("departments").Select(clause.Column{
			Name: "COUNT(*)",
			Raw:  true,
		}), 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrHavingLt(Callback < value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `status` = ? OR (SELECT count(*) FROM `departments`) < ?").WithArgs(1, 1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Having("status", 1).OrHavingLt(func(tx *gorm.DB) *gorm.DB {
			var count int64
			return tx.Table("departments").Count(&count)
		}, 1).Find(&dest)
		assert.Nil(t, err)
	})
}

func TestBuilder_HavingLte(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}
	t.Run("builder.HavingLte(string <= value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `id` <= ?").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").HavingLte("id", 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.HavingLte(*gorm.DB <= value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING (SELECT COUNT(*) FROM `departments`) <= ?").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").HavingLte(mockDB.Table("departments").Select("COUNT(*)"), 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.HavingLte(*Builder <= value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING (SELECT COUNT(*) FROM `departments`) <= ?").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").HavingLte(NewBuilder(mockDB).Table("departments").Select(clause.Column{
			Name: "COUNT(*)",
			Raw:  true,
		}), 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.HavingLte(Callback <= value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING (SELECT count(*) FROM `departments`) <= ?").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").HavingLte(func(tx *gorm.DB) *gorm.DB {
			var count int64
			return tx.Table("departments").Count(&count)
		}, 1).Find(&dest)
		assert.Nil(t, err)
	})
}

func TestBuilder_OrHavingLte(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}
	t.Run("builder.OrHavingLte(string <= value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `status` = ? OR `id` <= ?").WithArgs(1, 1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Having("status", 1).OrHavingLte("id", 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrHavingLte(*gorm.DB <= value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `status` = ? OR (SELECT COUNT(*) FROM `departments`) <= ?").WithArgs(1, 1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Having("status", 1).OrHavingLte(mockDB.Table("departments").Select("COUNT(*)"), 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrHavingLte(*Builder <= value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `status` = ? OR (SELECT COUNT(*) FROM `departments`) <= ?").WithArgs(1, 1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Having("status", 1).OrHavingLte(NewBuilder(mockDB).Table("departments").Select(clause.Column{
			Name: "COUNT(*)",
			Raw:  true,
		}), 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrHavingLte(Callback <= value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `status` = ? OR (SELECT count(*) FROM `departments`) <= ?").WithArgs(1, 1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Having("status", 1).OrHavingLte(func(tx *gorm.DB) *gorm.DB {
			var count int64
			return tx.Table("departments").Count(&count)
		}, 1).Find(&dest)
		assert.Nil(t, err)
	})
}

func TestBuilder_HavingIn(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}
	t.Run("builder.HavingIn(string IN values)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `id` IN (?,?,?)").WithArgs(1, 2, 3).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").HavingIn("id", 1, 2, 3).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.HavingIn(string IN *gorm.DB)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `id` IN (SELECT user_id FROM `vip_users`)").WithoutArgs().WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").HavingIn(
			"id",
			mockDB.Table("vip_users").Select("user_id"),
		).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.HavingIn(string IN *Builder)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `id` IN (SELECT `user_id` FROM `vip_users` )").WithoutArgs().WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").HavingIn(
			"id",
			NewBuilder(mockDB).Table("vip_users").Select("user_id"),
		).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.HavingIn(string IN Callback)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `id` IN (SELECT user_id FROM `vip_users`)").WithoutArgs().WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").HavingIn(
			"id",
			func(tx *gorm.DB) *gorm.DB {
				dest := make([]map[string]any, 0)
				return tx.Table("vip_users").Select("user_id").Find(&dest)
			},
		).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.HavingIn(*gorm.DB IN values)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING (SELECT COUNT(*) FROM `departments`) IN (?,?,?)").WithArgs(1, 2, 3).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").HavingIn(mockDB.Table("departments").Select("COUNT(*)"), 1, 2, 3).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.HavingIn(*Builder IN values)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING (SELECT COUNT(*) FROM `departments`) IN (?,?,?)").WithArgs(1, 2, 3).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").HavingIn(NewBuilder(mockDB).Table("departments").Select(clause.Column{
			Name: "COUNT(*)",
			Raw:  true,
		}), 1, 2, 3).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.HavingIn(Callback IN values)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING (SELECT count(*) FROM `departments`) IN (?,?,?)").WithArgs(1, 2, 3).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").HavingIn(func(tx *gorm.DB) *gorm.DB {
			var count int64
			return tx.Table("departments").Count(&count)
		}, 1, 2, 3).Find(&dest)
		assert.Nil(t, err)
	})
}

func TestBuilder_OrHavingIn(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}
	t.Run("builder.OrHavingIn(string IN values)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `status` = ? OR `id` IN (?,?,?)").WithArgs(1, 1, 2, 3).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Having("status", 1).OrHavingIn("id", 1, 2, 3).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrHavingIn(*gorm.DB IN values)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `status` = ? OR (SELECT COUNT(*) FROM `departments`) IN (?,?,?)").WithArgs(1, 1, 2, 3).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Having("status", 1).OrHavingIn(mockDB.Table("departments").Select("COUNT(*)"), 1, 2, 3).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrHavingIn(*Builder IN values)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `status` = ? OR (SELECT COUNT(*) FROM `departments`) IN (?,?,?)").WithArgs(1, 1, 2, 3).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Having("status", 1).OrHavingIn(NewBuilder(mockDB).Table("departments").Select(clause.Column{
			Name: "COUNT(*)",
			Raw:  true,
		}), 1, 2, 3).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrHavingIn(Callback IN values)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `status` = ? OR (SELECT count(*) FROM `departments`) IN (?,?,?)").WithArgs(1, 1, 2, 3).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Having("status", 1).OrHavingIn(func(tx *gorm.DB) *gorm.DB {
			var count int64
			return tx.Table("departments").Count(&count)
		}, 1, 2, 3).Find(&dest)
		assert.Nil(t, err)
	})
}

func TestBuilder_HavingNotIn(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}
	t.Run("builder.HavingNotIn(string NOT IN values)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `id` NOT IN (?,?,?)").WithArgs(1, 2, 3).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").HavingNotIn("id", 1, 2, 3).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.HavingNotIn(*gorm.DB NOT IN values)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING (SELECT COUNT(*) FROM `departments`) NOT IN (?,?,?)").WithArgs(1, 2, 3).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").HavingNotIn(mockDB.Table("departments").Select("COUNT(*)"), 1, 2, 3).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.HavingNotIn(*Builder NOT IN values)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING (SELECT COUNT(*) FROM `departments`) NOT IN (?,?,?)").WithArgs(1, 2, 3).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").HavingNotIn(NewBuilder(mockDB).Table("departments").Select(clause.Column{
			Name: "COUNT(*)",
			Raw:  true,
		}), 1, 2, 3).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.HavingNotIn(Callback NOT IN values)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING (SELECT count(*) FROM `departments`) NOT IN (?,?,?)").WithArgs(1, 2, 3).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").HavingNotIn(func(tx *gorm.DB) *gorm.DB {
			var count int64
			return tx.Table("departments").Count(&count)
		}, 1, 2, 3).Find(&dest)
		assert.Nil(t, err)
	})
}

func TestBuilder_OrHavingNotIn(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}
	t.Run("builder.OrHavingNotIn(string NOT IN values)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `status` = ? OR `id` NOT IN (?,?,?)").WithArgs(1, 1, 2, 3).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Having("status", 1).OrHavingNotIn("id", 1, 2, 3).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrHavingNotIn(*gorm.DB NOT IN values)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `status` = ? OR (SELECT COUNT(*) FROM `departments`) NOT IN (?,?,?)").WithArgs(1, 1, 2, 3).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Having("status", 1).OrHavingNotIn(mockDB.Table("departments").Select("COUNT(*)"), 1, 2, 3).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrHavingNotIn(*Builder NOT IN values)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `status` = ? OR (SELECT COUNT(*) FROM `departments`) NOT IN (?,?,?)").WithArgs(1, 1, 2, 3).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Having("status", 1).OrHavingNotIn(NewBuilder(mockDB).Table("departments").Select(clause.Column{
			Name: "COUNT(*)",
			Raw:  true,
		}), 1, 2, 3).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrHavingNotIn(Callback NOT IN values)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `status` = ? OR (SELECT count(*) FROM `departments`) NOT IN (?,?,?)").WithArgs(1, 1, 2, 3).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Having("status", 1).OrHavingNotIn(func(tx *gorm.DB) *gorm.DB {
			var count int64
			return tx.Table("departments").Count(&count)
		}, 1, 2, 3).Find(&dest)
		assert.Nil(t, err)
	})
}

func TestBuilder_HavingBetween(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}
	t.Run("builder.HavingBetween(string BETWEEN values)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `id` BETWEEN ? AND ?").WithArgs(1, 10).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").HavingBetween("id", 1, 10).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.HavingBetween(string BETWEEN *gorm.DB AND *gorm.DB)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `id` BETWEEN (SELECT MIN(`id`) FROM `vip_users`) AND (SELECT MAX(`id`) FROM `vip_users`)").WithoutArgs().WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").HavingBetween(
			"id",
			mockDB.Table("vip_users").Select("MIN(`id`)"),
			mockDB.Table("vip_users").Select("MAX(`id`)"),
		).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.HavingBetween(string BETWEEN *Builder AND *Builder)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `id` BETWEEN (SELECT MIN(`id`) FROM `vip_users` ) AND (SELECT MAX(`id`) FROM `vip_users` )").WithoutArgs().WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").HavingBetween(
			"id",
			NewBuilder(mockDB).Table("vip_users").Select(clause.Column{Name: "MIN(`id`)", Raw: true}),
			NewBuilder(mockDB).Table("vip_users").Select(clause.Column{Name: "MAX(`id`)", Raw: true}),
		).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.HavingBetween(string BETWEEN Callback AND Callback)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `id` BETWEEN (SELECT MIN(`id`) FROM `vip_users`) AND (SELECT MAX(`id`) FROM `vip_users`)").WithoutArgs().WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").HavingBetween(
			"id",
			func(tx *gorm.DB) *gorm.DB {
				dest := make([]map[string]any, 0)
				return tx.Table("vip_users").Select("MIN(`id`)").Find(&dest)
			},
			func(tx *gorm.DB) *gorm.DB {
				dest := make([]map[string]any, 0)
				return tx.Table("vip_users").Select("MAX(`id`)").Find(&dest)
			},
		).Find(&dest)
		assert.Nil(t, err)
	})
}

func TestBuilder_HavingNotBetween(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}
	t.Run("builder.HavingNotBetween(string NOT BETWEEN values)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `id` NOT BETWEEN ? AND ?").WithArgs(1, 10).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").HavingNotBetween("id", 1, 10).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.HavingNotBetween(string NOT BETWEEN *gorm.DB AND *gorm.DB)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `id` NOT BETWEEN (SELECT MIN(`id`) FROM `vip_users`) AND (SELECT MAX(`id`) FROM `vip_users`)").WithoutArgs().WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").HavingNotBetween(
			"id",
			mockDB.Table("vip_users").Select("MIN(`id`)"),
			mockDB.Table("vip_users").Select("MAX(`id`)"),
		).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.HavingNotBetween(string NOT BETWEEN *Builder AND *Builder)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `id` NOT BETWEEN (SELECT MIN(`id`) FROM `vip_users` ) AND (SELECT MAX(`id`) FROM `vip_users` )").WithoutArgs().WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").HavingNotBetween(
			"id",
			NewBuilder(mockDB).Table("vip_users").Select(clause.Column{Name: "MIN(`id`)", Raw: true}),
			NewBuilder(mockDB).Table("vip_users").Select(clause.Column{Name: "MAX(`id`)", Raw: true}),
		).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.HavingNotBetween(string NOT BETWEEN Callback AND Callback)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `id` NOT BETWEEN (SELECT MIN(`id`) FROM `vip_users`) AND (SELECT MAX(`id`) FROM `vip_users`)").WithoutArgs().WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").HavingNotBetween(
			"id",
			func(tx *gorm.DB) *gorm.DB {
				dest := make([]map[string]any, 0)
				return tx.Table("vip_users").Select("MIN(`id`)").Find(&dest)
			},
			func(tx *gorm.DB) *gorm.DB {
				dest := make([]map[string]any, 0)
				return tx.Table("vip_users").Select("MAX(`id`)").Find(&dest)
			},
		).Find(&dest)
		assert.Nil(t, err)
	})
}

func TestBuilder_OrHavingBetween(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}
	t.Run("builder.OrHavingBetween(string BETWEEN values)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `status` = ? OR `id` BETWEEN ? AND ?").WithArgs(1, 1, 10).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Having("status", 1).OrHavingBetween("id", 1, 10).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrHavingBetween(string BETWEEN *gorm.DB AND *gorm.DB)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `status` = ? OR `id` BETWEEN (SELECT MIN(`id`) FROM `vip_users`) AND (SELECT MAX(`id`) FROM `vip_users`)").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Having("status", 1).OrHavingBetween(
			"id",
			mockDB.Table("vip_users").Select("MIN(`id`)"),
			mockDB.Table("vip_users").Select("MAX(`id`)"),
		).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrHavingBetween(string BETWEEN *Builder AND *Builder)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `status` = ? OR `id` BETWEEN (SELECT MIN(`id`) FROM `vip_users` ) AND (SELECT MAX(`id`) FROM `vip_users` )").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Having("status", 1).OrHavingBetween(
			"id",
			NewBuilder(mockDB).Table("vip_users").Select(clause.Column{Name: "MIN(`id`)", Raw: true}),
			NewBuilder(mockDB).Table("vip_users").Select(clause.Column{Name: "MAX(`id`)", Raw: true}),
		).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrHavingBetween(string BETWEEN Callback AND Callback)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `status` = ? OR `id` BETWEEN (SELECT MIN(`id`) FROM `vip_users`) AND (SELECT MAX(`id`) FROM `vip_users`)").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Having("status", 1).OrHavingBetween(
			"id",
			func(tx *gorm.DB) *gorm.DB {
				dest := make([]map[string]any, 0)
				return tx.Table("vip_users").Select("MIN(`id`)").Find(&dest)
			},
			func(tx *gorm.DB) *gorm.DB {
				dest := make([]map[string]any, 0)
				return tx.Table("vip_users").Select("MAX(`id`)").Find(&dest)
			},
		).Find(&dest)
		assert.Nil(t, err)
	})
}

func TestBuilder_OrHavingNotBetween(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}
	t.Run("builder.OrWhereNotBetween(string BETWEEN values)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `status` = ? OR `id` NOT BETWEEN ? AND ?").WithArgs(1, 1, 10).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Having("status", 1).OrHavingNotBetween("id", 1, 10).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrWhereNotBetween(string BETWEEN *gorm.DB AND *gorm.DB)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `status` = ? OR `id` NOT BETWEEN (SELECT MIN(`id`) FROM `vip_users`) AND (SELECT MAX(`id`) FROM `vip_users`)").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Having("status", 1).OrHavingNotBetween(
			"id",
			mockDB.Table("vip_users").Select("MIN(`id`)"),
			mockDB.Table("vip_users").Select("MAX(`id`)"),
		).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrWhereNotBetween(string BETWEEN *Builder AND *Builder)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `status` = ? OR `id` NOT BETWEEN (SELECT MIN(`id`) FROM `vip_users` ) AND (SELECT MAX(`id`) FROM `vip_users` )").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Having("status", 1).OrHavingNotBetween(
			"id",
			NewBuilder(mockDB).Table("vip_users").Select(clause.Column{Name: "MIN(`id`)", Raw: true}),
			NewBuilder(mockDB).Table("vip_users").Select(clause.Column{Name: "MAX(`id`)", Raw: true}),
		).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrWhereNotBetween(string BETWEEN Callback AND Callback)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `status` = ? OR `id` NOT BETWEEN (SELECT MIN(`id`) FROM `vip_users`) AND (SELECT MAX(`id`) FROM `vip_users`)").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Having("status", 1).OrHavingNotBetween(
			"id",
			func(tx *gorm.DB) *gorm.DB {
				dest := make([]map[string]any, 0)
				return tx.Table("vip_users").Select("MIN(`id`)").Find(&dest)
			},
			func(tx *gorm.DB) *gorm.DB {
				dest := make([]map[string]any, 0)
				return tx.Table("vip_users").Select("MAX(`id`)").Find(&dest)
			},
		).Find(&dest)
		assert.Nil(t, err)
	})
}

func TestBuilder_HavingLike(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}

	mock.ExpectQuery("SELECT * FROM `users` HAVING `username` LIKE ?").WithArgs("%wardonne%").WillReturnRows(result)
	dest := make([]map[string]any, 0)
	err := NewBuilder(mockDB).Table("users").HavingLike("username", "%wardonne%").Find(&dest)
	assert.Nil(t, err)
}

func TestBuilder_HavingNotLike(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}

	mock.ExpectQuery("SELECT * FROM `users` HAVING `username` NOT LIKE ?").WithArgs("%wardonne%").WillReturnRows(result)
	dest := make([]map[string]any, 0)
	err := NewBuilder(mockDB).Table("users").HavingNotLike("username", "%wardonne%").Find(&dest)
	assert.Nil(t, err)
}

func TestBuilder_OrHavingLike(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}

	mock.ExpectQuery("SELECT * FROM `users` HAVING `status` = ? OR `username` LIKE ?").WithArgs(1, "%wardonne%").WillReturnRows(result)
	dest := make([]map[string]any, 0)
	err := NewBuilder(mockDB).Table("users").Having("status", 1).OrHavingLike("username", "%wardonne%").Find(&dest)
	assert.Nil(t, err)
}

func TestBuilder_OrHavingNotLike(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}

	mock.ExpectQuery("SELECT * FROM `users` HAVING `status` = ? OR `username` NOT LIKE ?").WithArgs(1, "%wardonne%").WillReturnRows(result)
	dest := make([]map[string]any, 0)
	err := NewBuilder(mockDB).Table("users").Having("status", 1).OrHavingNotLike("username", "%wardonne%").Find(&dest)
	assert.Nil(t, err)
}

func TestBuilder_HavingBuilder(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}
	mock.ExpectQuery("SELECT * FROM `users` HAVING `status` = ? AND (`id` = ? AND `department_id` = ?)").WithArgs(1, 1, 1).WillReturnRows(result)
	var dest = make([]map[string]any, 0)
	err := NewBuilder(mockDB).
		Table("users").
		Having("status", 1).
		HavingBuilder(NewBuilder(mockDB).Having("id", 1).Having("department_id", 1)).
		Find(&dest)
	assert.Nil(t, err)
}

func TestBuilder_HavingNotBuilder(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}
	mock.ExpectQuery("SELECT * FROM `users` HAVING `status` = ? AND (`id` <> ? AND `department_id` <> ?)").WithArgs(1, 1, 1).WillReturnRows(result)
	var dest = make([]map[string]any, 0)
	err := NewBuilder(mockDB).
		Table("users").
		Having("status", 1).
		HavingNotBuilder(NewBuilder(mockDB).Having("id", 1).Having("department_id", 1)).
		Find(&dest)
	assert.Nil(t, err)
}

func TestBuilder_OrHavingBuilder(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}
	mock.ExpectQuery("SELECT * FROM `users` HAVING `status` = ? OR (`id` = ? AND `department_id` = ?)").WithArgs(1, 1, 1).WillReturnRows(result)
	var dest = make([]map[string]any, 0)
	err := NewBuilder(mockDB).
		Table("users").
		Having("status", 1).
		OrHavingBuilder(NewBuilder(mockDB).Having("id", 1).Having("department_id", 1)).
		Find(&dest)
	assert.Nil(t, err)
}
func TestBuilder_OrHavingNotBuilder(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}
	mock.ExpectQuery("SELECT * FROM `users` HAVING `status` = ? OR (`id` <> ? AND `department_id` <> ?)").WithArgs(1, 1, 1).WillReturnRows(result)
	var dest = make([]map[string]any, 0)
	err := NewBuilder(mockDB).
		Table("users").
		Having("status", 1).
		OrHavingNotBuilder(NewBuilder(mockDB).Having("id", 1).Having("department_id", 1)).
		Find(&dest)
	assert.Nil(t, err)
}

func TestBuilder_HavingCallback(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}
	mock.ExpectQuery("SELECT * FROM `users` HAVING `status` = ? AND (`id` = ? AND `department_id` = ?)").WithArgs(1, 1, 1).WillReturnRows(result)
	var dest = make([]map[string]any, 0)
	err := NewBuilder(mockDB).
		Table("users").
		Having("status", 1).
		HavingCallback(func(builder *Builder) *Builder {
			return builder.Having("id", 1).Having("department_id", 1)
		}).
		Find(&dest)
	assert.Nil(t, err)
}

func TestBuilder_HavingNotCallback(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}
	mock.ExpectQuery("SELECT * FROM `users` HAVING `status` = ? AND (`id` <> ? AND `department_id` <> ?)").WithArgs(1, 1, 1).WillReturnRows(result)
	var dest = make([]map[string]any, 0)
	err := NewBuilder(mockDB).
		Table("users").
		Having("status", 1).
		HavingNotCallback(func(builder *Builder) *Builder {
			return builder.Having("id", 1).Having("department_id", 1)
		}).
		Find(&dest)
	assert.Nil(t, err)
}

func TestBuilder_OrHavingCallback(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}
	mock.ExpectQuery("SELECT * FROM `users` HAVING `status` = ? OR (`id` = ? AND `department_id` = ?)").WithArgs(1, 1, 1).WillReturnRows(result)
	var dest = make([]map[string]any, 0)
	err := NewBuilder(mockDB).
		Table("users").
		Having("status", 1).
		OrHavingCallback(func(builder *Builder) *Builder {
			return builder.Having("id", 1).Having("department_id", 1)
		}).
		Find(&dest)
	assert.Nil(t, err)
}

func TestBuilder_OrHavingNotCallback(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}
	mock.ExpectQuery("SELECT * FROM `users` HAVING `status` = ? OR (`id` <> ? AND `department_id` <> ?)").WithArgs(1, 1, 1).WillReturnRows(result)
	var dest = make([]map[string]any, 0)
	err := NewBuilder(mockDB).
		Table("users").
		Having("status", 1).
		OrHavingNotCallback(func(builder *Builder) *Builder {
			return builder.Having("id", 1).Having("department_id", 1)
		}).
		Find(&dest)
	assert.Nil(t, err)
}
