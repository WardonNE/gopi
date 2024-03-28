package builder

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func TestBuilder_Where(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}

	t.Run("builder.Where(column = value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `id` = ?").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Where("id", 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.Where(column In values)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `id` IN (?,?,?)").WithArgs(1, 2, 3).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Where("id", []int{1, 2, 3}).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.Where(column IS Null)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `id` IS NULL").WithoutArgs().WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Where("id", nil).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.Where(column = *gorm.DB)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `id` = (SELECT COUNT(*) FROM `departments`)").WithoutArgs().WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Where("id", mockDB.Table("departments").Select("COUNT(*)")).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.Where(column = *Builder)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `id` = (SELECT COUNT(*) FROM `departments` )").WithoutArgs().WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Where("id", NewBuilder(mockDB).Table("departments").Select(clause.Column{
			Name: "COUNT(*)",
			Raw:  true,
		})).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.Where(column = Callback)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `id` = (SELECT count(*) FROM `departments`)").WithoutArgs().WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Where("id", func(tx *gorm.DB) *gorm.DB {
			var count int64
			return tx.Table("departments").Count(&count)
		}).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.Where(*gorm.DB = value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE (SELECT COUNT(*) FROM `departments`) = ?").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Where(mockDB.Table("departments").Select("COUNT(*)"), 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.Where(*Builder = value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE (SELECT COUNT(*) FROM `departments`) = ?").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Where(
			NewBuilder(mockDB).Table("departments").Select(clause.Column{
				Name: "COUNT(*)",
				Raw:  true,
			}),
			1,
		).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.Where(Callback = value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE (SELECT count(*) FROM `departments`) = ?").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Where(
			func(tx *gorm.DB) *gorm.DB {
				var count int64
				return tx.Table("departments").Count(&count)
			},
			1,
		).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.Where(JSON_EXTRACT)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE JSON_EXTRACT(`meta`,'$.department.status') = ?").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Where(datatypes.JSONQuery("meta").Extract("department.status"), 1).Find(&dest)
		assert.Nil(t, err)
	})
}

func TestBuilder_WhereNot(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}

	t.Run("builder.WhereNot(column <> value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `id` <> ?").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").WhereNot("id", 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.WhereNot(column Not In values)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `id` NOT IN (?,?,?)").WithArgs(1, 2, 3).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").WhereNot("id", []int{1, 2, 3}).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.WhereNot(column IS NOT Null)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `id` IS NOT NULL").WithoutArgs().WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").WhereNot("id", nil).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.WhereNot(column <> *gorm.DB)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `id` <> (SELECT COUNT(*) FROM `departments`)").WithoutArgs().WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").WhereNot("id", mockDB.Table("departments").Select("COUNT(*)")).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.WhereNot(column <> *Builder)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `id` <> (SELECT COUNT(*) FROM `departments` )").WithoutArgs().WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").WhereNot("id", NewBuilder(mockDB).Table("departments").Select(clause.Column{
			Name: "COUNT(*)",
			Raw:  true,
		})).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.WhereNot(column <> Callback)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `id` <> (SELECT count(*) FROM `departments`)").WithoutArgs().WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").WhereNot("id", func(tx *gorm.DB) *gorm.DB {
			var count int64
			return tx.Table("departments").Count(&count)
		}).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.WhereNot(*gorm.DB <> value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE (SELECT COUNT(*) FROM `departments`) <> ?").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").WhereNot(mockDB.Table("departments").Select("COUNT(*)"), 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.WhereNot(*Builder = value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE (SELECT COUNT(*) FROM `departments`) <> ?").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").WhereNot(
			NewBuilder(mockDB).Table("departments").Select(clause.Column{
				Name: "COUNT(*)",
				Raw:  true,
			}),
			1,
		).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.WhereNot(Callback <> value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE (SELECT count(*) FROM `departments`) <> ?").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").WhereNot(
			func(tx *gorm.DB) *gorm.DB {
				var count int64
				return tx.Table("departments").Count(&count)
			},
			1,
		).Find(&dest)
		assert.Nil(t, err)
	})
}

func TestBuilder_OrWhere(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}

	t.Run("builder.OrWhere(column = value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? OR `id` = ?").WithArgs(1, 1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Where("status", 1).OrWhere("id", 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrWhere(column In values)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? OR `id` IN (?,?,?)").WithArgs(1, 1, 2, 3).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Where("status", 1).OrWhere("id", []int{1, 2, 3}).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrWhere(column IS Null)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? OR `id` IS NULL").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Where("status", 1).OrWhere("id", nil).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrWhere(column = *gorm.DB)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? OR `id` = (SELECT COUNT(*) FROM `departments`)").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Where("status", 1).OrWhere("id", mockDB.Table("departments").Select("COUNT(*)")).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrWhere(column = *Builder)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? OR `id` = (SELECT COUNT(*) FROM `departments` )").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Where("status", 1).OrWhere("id", NewBuilder(mockDB).Table("departments").Select(clause.Column{
			Name: "COUNT(*)",
			Raw:  true,
		})).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrWhere(column = Callback)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? OR `id` = (SELECT count(*) FROM `departments`)").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Where("status", 1).OrWhere("id", func(tx *gorm.DB) *gorm.DB {
			var count int64
			return tx.Table("departments").Count(&count)
		}).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrWhere(*gorm.DB = value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? OR (SELECT COUNT(*) FROM `departments`) = ?").WithArgs(1, 1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Where("status", 1).OrWhere(mockDB.Table("departments").Select("COUNT(*)"), 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrWhere(*Builder = value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? OR (SELECT COUNT(*) FROM `departments`) = ?").WithArgs(1, 1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Where("status", 1).OrWhere(
			NewBuilder(mockDB).Table("departments").Select(clause.Column{
				Name: "COUNT(*)",
				Raw:  true,
			}),
			1,
		).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrWhere(Callback = value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? OR (SELECT count(*) FROM `departments`) = ?").WithArgs(1, 1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Where("status", 1).OrWhere(
			func(tx *gorm.DB) *gorm.DB {
				var count int64
				return tx.Table("departments").Count(&count)
			},
			1,
		).Find(&dest)
		assert.Nil(t, err)
	})
}

func TestBuilder_OrWhereNot(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}

	t.Run("builder.OrWhereNot(column <> value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? OR `id` <> ?").WithArgs(1, 1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Where("status", 1).OrWhereNot("id", 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrWhereNot(column Not In values)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? OR `id` NOT IN (?,?,?)").WithArgs(1, 1, 2, 3).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Where("status", 1).OrWhereNot("id", []int{1, 2, 3}).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrWhereNot(column IS NOT Null)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? OR `id` IS NOT NULL").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Where("status", 1).OrWhereNot("id", nil).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrWhereNot(column <> *gorm.DB)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? OR `id` <> (SELECT COUNT(*) FROM `departments`)").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Where("status", 1).OrWhereNot("id", mockDB.Table("departments").Select("COUNT(*)")).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrWhereNot(column <> *Builder)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? OR `id` <> (SELECT COUNT(*) FROM `departments` )").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Where("status", 1).OrWhereNot("id", NewBuilder(mockDB).Table("departments").Select(clause.Column{
			Name: "COUNT(*)",
			Raw:  true,
		})).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrWhereNot(column <> Callback)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? OR `id` <> (SELECT count(*) FROM `departments`)").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Where("status", 1).OrWhereNot("id", func(tx *gorm.DB) *gorm.DB {
			var count int64
			return tx.Table("departments").Count(&count)
		}).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrWhereNot(*gorm.DB <> value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? OR (SELECT COUNT(*) FROM `departments`) <> ?").WithArgs(1, 1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Where("status", 1).OrWhereNot(mockDB.Table("departments").Select("COUNT(*)"), 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrWhereNot(*Builder <> value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? OR (SELECT COUNT(*) FROM `departments`) <> ?").WithArgs(1, 1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Where("status", 1).OrWhereNot(
			NewBuilder(mockDB).Table("departments").Select(clause.Column{
				Name: "COUNT(*)",
				Raw:  true,
			}),
			1,
		).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrWhereNot(Callback <> value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? OR (SELECT count(*) FROM `departments`) <> ?").WithArgs(1, 1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Where("status", 1).OrWhereNot(
			func(tx *gorm.DB) *gorm.DB {
				var count int64
				return tx.Table("departments").Count(&count)
			},
			1,
		).Find(&dest)
		assert.Nil(t, err)
	})
}

func TestBuilder_WhereNull(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}
	t.Run("builder.WhereNull(string IS NULL)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `id` IS NULL").WithoutArgs().WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").WhereNull("id").Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.WhereNull(*gorm.DB IS NULL)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE (SELECT COUNT(*) FROM `departments`) IS NULL").WithoutArgs().WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").WhereNull(mockDB.Table("departments").Select("COUNT(*)")).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.WhereNull(*Builder IS NULL)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE (SELECT COUNT(*) FROM `departments`) IS NULL").WithoutArgs().WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").WhereNull(NewBuilder(mockDB).Table("departments").Select(clause.Column{
			Name: "COUNT(*)",
			Raw:  true,
		})).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.WhereNull(Callback IS NULL)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE (SELECT count(*) FROM `departments`) IS NULL").WithoutArgs().WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").WhereNull(func(tx *gorm.DB) *gorm.DB {
			var count int64
			return tx.Table("departments").Count(&count)
		}).Find(&dest)
		assert.Nil(t, err)
	})
}

func TestBuilder_WhereNotNull(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}
	t.Run("builder.WhereNotNull(string IS NOT NULL)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `id` IS NOT NULL").WithoutArgs().WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").WhereNotNull("id").Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.WhereNotNull(*gorm.DB IS NOT NULL)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE (SELECT COUNT(*) FROM `departments`) IS NOT NULL").WithoutArgs().WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").WhereNotNull(mockDB.Table("departments").Select("COUNT(*)")).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.WhereNotNull(*Builder IS NOT NULL)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE (SELECT COUNT(*) FROM `departments`) IS NOT NULL").WithoutArgs().WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").WhereNotNull(NewBuilder(mockDB).Table("departments").Select(clause.Column{
			Name: "COUNT(*)",
			Raw:  true,
		})).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.WhereNotNull(Callback IS NOT NULL)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE (SELECT count(*) FROM `departments`) IS NOT NULL").WithoutArgs().WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").WhereNotNull(func(tx *gorm.DB) *gorm.DB {
			var count int64
			return tx.Table("departments").Count(&count)
		}).Find(&dest)
		assert.Nil(t, err)
	})
}

func TestBuilder_OrWhereNull(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}
	t.Run("builder.OrWhereNull(string IS NULL)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? OR `id` IS NULL").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Where("status", 1).OrWhereNull("id").Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrWhereNull(*gorm.DB IS NULL)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `id` = ? OR (SELECT COUNT(*) FROM `departments`) IS NULL").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Where("id", 1).OrWhereNull(mockDB.Table("departments").Select("COUNT(*)")).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrWhereNull(*Builder IS NULL)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `id` = ? OR (SELECT COUNT(*) FROM `departments`) IS NULL").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Where("id", 1).OrWhereNull(NewBuilder(mockDB).Table("departments").Select(clause.Column{
			Name: "COUNT(*)",
			Raw:  true,
		})).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrWhereNull(Callback IS NULL)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `id` = ? OR (SELECT count(*) FROM `departments`) IS NULL").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Where("id", 1).OrWhereNull(func(tx *gorm.DB) *gorm.DB {
			var count int64
			return tx.Table("departments").Count(&count)
		}).Find(&dest)
		assert.Nil(t, err)
	})
}

func TestBuilder_OrWhereNotNull(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}
	t.Run("builder.OrWhereNotNull(string IS NOT NULL)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? OR `id` IS NOT NULL").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Where("status", 1).OrWhereNotNull("id").Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrWhereNotNull(*gorm.DB IS NOT NULL)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `id` = ? OR (SELECT COUNT(*) FROM `departments`) IS NOT NULL").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Where("id", 1).OrWhereNotNull(mockDB.Table("departments").Select("COUNT(*)")).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrWhereNotNull(*Builder IS NOT NULL)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `id` = ? OR (SELECT COUNT(*) FROM `departments`) IS NOT NULL").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Where("id", 1).OrWhereNotNull(NewBuilder(mockDB).Table("departments").Select(clause.Column{
			Name: "COUNT(*)",
			Raw:  true,
		})).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrWhereNotNull(Callback IS NOT NULL)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `id` = ? OR (SELECT count(*) FROM `departments`) IS NOT NULL").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Where("id", 1).OrWhereNotNull(func(tx *gorm.DB) *gorm.DB {
			var count int64
			return tx.Table("departments").Count(&count)
		}).Find(&dest)
		assert.Nil(t, err)
	})
}

func TestBuilder_WhereEq(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}
	t.Run("builder.WhereEq(string = value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `id` = ?").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").WhereEq("id", 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.WhereEq(*gorm.DB = value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE (SELECT COUNT(*) FROM `departments`) = ?").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").WhereEq(mockDB.Table("departments").Select("COUNT(*)"), 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.WhereEq(*Builder = value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE (SELECT COUNT(*) FROM `departments`) = ?").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").WhereEq(NewBuilder(mockDB).Table("departments").Select(clause.Column{
			Name: "COUNT(*)",
			Raw:  true,
		}), 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.WhereEq(Callback = value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE (SELECT count(*) FROM `departments`) = ?").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").WhereEq(func(tx *gorm.DB) *gorm.DB {
			var count int64
			return tx.Table("departments").Count(&count)
		}, 1).Find(&dest)
		assert.Nil(t, err)
	})
}

func TestBuilder_OrWhereEq(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}
	t.Run("builder.OrWhereEq(string = value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? OR `id` = ?").WithArgs(1, 1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Where("status", 1).OrWhereEq("id", 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrWhereEq(*gorm.DB = value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? OR (SELECT COUNT(*) FROM `departments`) = ?").WithArgs(1, 1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Where("status", 1).OrWhereEq(mockDB.Table("departments").Select("COUNT(*)"), 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrWhereEq(*Builder = value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? OR (SELECT COUNT(*) FROM `departments`) = ?").WithArgs(1, 1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Where("status", 1).OrWhereEq(NewBuilder(mockDB).Table("departments").Select(clause.Column{
			Name: "COUNT(*)",
			Raw:  true,
		}), 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrWhereEq(Callback = value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? OR (SELECT count(*) FROM `departments`) = ?").WithArgs(1, 1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Where("status", 1).OrWhereEq(func(tx *gorm.DB) *gorm.DB {
			var count int64
			return tx.Table("departments").Count(&count)
		}, 1).Find(&dest)
		assert.Nil(t, err)
	})
}

func TestBuilder_WhereNeq(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}
	t.Run("builder.WhereNeq(string <> value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `id` <> ?").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").WhereNeq("id", 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.WhereNeq(*gorm.DB <> value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE (SELECT COUNT(*) FROM `departments`) <> ?").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").WhereNeq(mockDB.Table("departments").Select("COUNT(*)"), 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.WhereNeq(*Builder <> value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE (SELECT COUNT(*) FROM `departments`) <> ?").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").WhereNeq(NewBuilder(mockDB).Table("departments").Select(clause.Column{
			Name: "COUNT(*)",
			Raw:  true,
		}), 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.WhereNeq(Callback <> value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE (SELECT count(*) FROM `departments`) <> ?").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").WhereNeq(func(tx *gorm.DB) *gorm.DB {
			var count int64
			return tx.Table("departments").Count(&count)
		}, 1).Find(&dest)
		assert.Nil(t, err)
	})
}

func TestBuilder_OrWhereNeq(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}
	t.Run("builder.OrWhereNeq(string <> value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? OR `id` <> ?").WithArgs(1, 1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Where("status", 1).OrWhereNeq("id", 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrWhereNeq(*gorm.DB <> value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? OR (SELECT COUNT(*) FROM `departments`) <> ?").WithArgs(1, 1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Where("status", 1).OrWhereNeq(mockDB.Table("departments").Select("COUNT(*)"), 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrWhereNeq(*Builder <> value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? OR (SELECT COUNT(*) FROM `departments`) <> ?").WithArgs(1, 1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Where("status", 1).OrWhereNeq(NewBuilder(mockDB).Table("departments").Select(clause.Column{
			Name: "COUNT(*)",
			Raw:  true,
		}), 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrWhereNeq(Callback <> value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? OR (SELECT count(*) FROM `departments`) <> ?").WithArgs(1, 1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Where("status", 1).OrWhereNeq(func(tx *gorm.DB) *gorm.DB {
			var count int64
			return tx.Table("departments").Count(&count)
		}, 1).Find(&dest)
		assert.Nil(t, err)
	})
}

func TestBuilder_WhereGt(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}
	t.Run("builder.WhereGt(string > value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `id` > ?").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").WhereGt("id", 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.WhereGt(*gorm.DB > value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE (SELECT COUNT(*) FROM `departments`) > ?").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").WhereGt(mockDB.Table("departments").Select("COUNT(*)"), 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.WhereGt(*Builder > value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE (SELECT COUNT(*) FROM `departments`) > ?").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").WhereGt(NewBuilder(mockDB).Table("departments").Select(clause.Column{
			Name: "COUNT(*)",
			Raw:  true,
		}), 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.WhereGt(Callback > value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE (SELECT count(*) FROM `departments`) > ?").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").WhereGt(func(tx *gorm.DB) *gorm.DB {
			var count int64
			return tx.Table("departments").Count(&count)
		}, 1).Find(&dest)
		assert.Nil(t, err)
	})
}

func TestBuilder_OrWhereGt(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}
	t.Run("builder.OrWhereGt(string > value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? OR `id` > ?").WithArgs(1, 1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Where("status", 1).OrWhereGt("id", 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrWhereGt(*gorm.DB > value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? OR (SELECT COUNT(*) FROM `departments`) > ?").WithArgs(1, 1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Where("status", 1).OrWhereGt(mockDB.Table("departments").Select("COUNT(*)"), 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrWhereGt(*Builder > value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? OR (SELECT COUNT(*) FROM `departments`) > ?").WithArgs(1, 1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Where("status", 1).OrWhereGt(NewBuilder(mockDB).Table("departments").Select(clause.Column{
			Name: "COUNT(*)",
			Raw:  true,
		}), 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrWhereGt(Callback > value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? OR (SELECT count(*) FROM `departments`) > ?").WithArgs(1, 1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Where("status", 1).OrWhereGt(func(tx *gorm.DB) *gorm.DB {
			var count int64
			return tx.Table("departments").Count(&count)
		}, 1).Find(&dest)
		assert.Nil(t, err)
	})
}

func TestBuilder_WhereGte(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}
	t.Run("builder.WhereGte(string >= value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `id` >= ?").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").WhereGte("id", 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.WhereGte(*gorm.DB >= value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE (SELECT COUNT(*) FROM `departments`) >= ?").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").WhereGte(mockDB.Table("departments").Select("COUNT(*)"), 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.WhereGte(*Builder >= value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE (SELECT COUNT(*) FROM `departments`) >= ?").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").WhereGte(NewBuilder(mockDB).Table("departments").Select(clause.Column{
			Name: "COUNT(*)",
			Raw:  true,
		}), 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.WhereGte(Callback >= value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE (SELECT count(*) FROM `departments`) >= ?").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").WhereGte(func(tx *gorm.DB) *gorm.DB {
			var count int64
			return tx.Table("departments").Count(&count)
		}, 1).Find(&dest)
		assert.Nil(t, err)
	})
}

func TestBuilder_OrWhereGte(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}
	t.Run("builder.OrWhereGte(string >= value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? OR `id` >= ?").WithArgs(1, 1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Where("status", 1).OrWhereGte("id", 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrWhereGte(*gorm.DB >= value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? OR (SELECT COUNT(*) FROM `departments`) >= ?").WithArgs(1, 1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Where("status", 1).OrWhereGte(mockDB.Table("departments").Select("COUNT(*)"), 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrWhereGte(*Builder >= value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? OR (SELECT COUNT(*) FROM `departments`) >= ?").WithArgs(1, 1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Where("status", 1).OrWhereGte(NewBuilder(mockDB).Table("departments").Select(clause.Column{
			Name: "COUNT(*)",
			Raw:  true,
		}), 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrWhereGte(Callback >= value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? OR (SELECT count(*) FROM `departments`) >= ?").WithArgs(1, 1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Where("status", 1).OrWhereGte(func(tx *gorm.DB) *gorm.DB {
			var count int64
			return tx.Table("departments").Count(&count)
		}, 1).Find(&dest)
		assert.Nil(t, err)
	})
}

func TestBuilder_WhereLt(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}
	t.Run("builder.WhereLt(string < value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `id` < ?").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").WhereLt("id", 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.WhereLt(*gorm.DB < value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE (SELECT COUNT(*) FROM `departments`) < ?").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").WhereLt(mockDB.Table("departments").Select("COUNT(*)"), 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.WhereLt(*Builder < value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE (SELECT COUNT(*) FROM `departments`) < ?").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").WhereLt(NewBuilder(mockDB).Table("departments").Select(clause.Column{
			Name: "COUNT(*)",
			Raw:  true,
		}), 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.WhereLt(Callback < value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE (SELECT count(*) FROM `departments`) < ?").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").WhereLt(func(tx *gorm.DB) *gorm.DB {
			var count int64
			return tx.Table("departments").Count(&count)
		}, 1).Find(&dest)
		assert.Nil(t, err)
	})
}

func TestBuilder_OrWhereLt(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}
	t.Run("builder.OrWhereLt(string < value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? OR `id` < ?").WithArgs(1, 1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Where("status", 1).OrWhereLt("id", 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrWhereLt(*gorm.DB < value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? OR (SELECT COUNT(*) FROM `departments`) < ?").WithArgs(1, 1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Where("status", 1).OrWhereLt(mockDB.Table("departments").Select("COUNT(*)"), 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrWhereLt(*Builder < value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? OR (SELECT COUNT(*) FROM `departments`) < ?").WithArgs(1, 1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Where("status", 1).OrWhereLt(NewBuilder(mockDB).Table("departments").Select(clause.Column{
			Name: "COUNT(*)",
			Raw:  true,
		}), 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrWhereLt(Callback < value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? OR (SELECT count(*) FROM `departments`) < ?").WithArgs(1, 1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Where("status", 1).OrWhereLt(func(tx *gorm.DB) *gorm.DB {
			var count int64
			return tx.Table("departments").Count(&count)
		}, 1).Find(&dest)
		assert.Nil(t, err)
	})
}

func TestBuilder_WhereLte(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}
	t.Run("builder.WhereLte(string <= value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `id` <= ?").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").WhereLte("id", 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.WhereLte(*gorm.DB <= value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE (SELECT COUNT(*) FROM `departments`) <= ?").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").WhereLte(mockDB.Table("departments").Select("COUNT(*)"), 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.WhereLte(*Builder <= value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE (SELECT COUNT(*) FROM `departments`) <= ?").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").WhereLte(NewBuilder(mockDB).Table("departments").Select(clause.Column{
			Name: "COUNT(*)",
			Raw:  true,
		}), 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.WhereLte(Callback <= value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE (SELECT count(*) FROM `departments`) <= ?").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").WhereLte(func(tx *gorm.DB) *gorm.DB {
			var count int64
			return tx.Table("departments").Count(&count)
		}, 1).Find(&dest)
		assert.Nil(t, err)
	})
}

func TestBuilder_OrWhereLte(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}
	t.Run("builder.OrWhereLte(string <= value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? OR `id` <= ?").WithArgs(1, 1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Where("status", 1).OrWhereLte("id", 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrWhereLte(*gorm.DB <= value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? OR (SELECT COUNT(*) FROM `departments`) <= ?").WithArgs(1, 1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Where("status", 1).OrWhereLte(mockDB.Table("departments").Select("COUNT(*)"), 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrWhereLte(*Builder <= value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? OR (SELECT COUNT(*) FROM `departments`) <= ?").WithArgs(1, 1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Where("status", 1).OrWhereLte(NewBuilder(mockDB).Table("departments").Select(clause.Column{
			Name: "COUNT(*)",
			Raw:  true,
		}), 1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrWhereLte(Callback <= value)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? OR (SELECT count(*) FROM `departments`) <= ?").WithArgs(1, 1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Where("status", 1).OrWhereLte(func(tx *gorm.DB) *gorm.DB {
			var count int64
			return tx.Table("departments").Count(&count)
		}, 1).Find(&dest)
		assert.Nil(t, err)
	})
}

func TestBuilder_WhereIn(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}
	t.Run("builder.WhereIn(string IN values)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `id` IN (?,?,?)").WithArgs(1, 2, 3).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").WhereIn("id", 1, 2, 3).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.WhereIn(string IN *gorm.DB)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `id` IN (SELECT user_id FROM `vip_users`)").WithoutArgs().WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").WhereIn(
			"id",
			mockDB.Table("vip_users").Select("user_id"),
		).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.WhereIn(string IN *Builder)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `id` IN (SELECT `user_id` FROM `vip_users` )").WithoutArgs().WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").WhereIn(
			"id",
			NewBuilder(mockDB).Table("vip_users").Select("user_id"),
		).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.WhereIn(string IN Callback)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `id` IN (SELECT user_id FROM `vip_users`)").WithoutArgs().WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").WhereIn(
			"id",
			func(tx *gorm.DB) *gorm.DB {
				dest := make([]map[string]any, 0)
				return tx.Table("vip_users").Select("user_id").Find(&dest)
			},
		).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.WhereIn(*gorm.DB IN values)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE (SELECT COUNT(*) FROM `departments`) IN (?,?,?)").WithArgs(1, 2, 3).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").WhereIn(mockDB.Table("departments").Select("COUNT(*)"), 1, 2, 3).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.WhereIn(*Builder IN values)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE (SELECT COUNT(*) FROM `departments`) IN (?,?,?)").WithArgs(1, 2, 3).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").WhereIn(NewBuilder(mockDB).Table("departments").Select(clause.Column{
			Name: "COUNT(*)",
			Raw:  true,
		}), 1, 2, 3).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.WhereIn(Callback IN values)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE (SELECT count(*) FROM `departments`) IN (?,?,?)").WithArgs(1, 2, 3).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").WhereIn(func(tx *gorm.DB) *gorm.DB {
			var count int64
			return tx.Table("departments").Count(&count)
		}, 1, 2, 3).Find(&dest)
		assert.Nil(t, err)
	})
}

func TestBuilder_OrWhereIn(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}
	t.Run("builder.OrWhereIn(string IN values)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? OR `id` IN (?,?,?)").WithArgs(1, 1, 2, 3).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Where("status", 1).OrWhereIn("id", 1, 2, 3).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrWhereIn(*gorm.DB IN values)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? OR (SELECT COUNT(*) FROM `departments`) IN (?,?,?)").WithArgs(1, 1, 2, 3).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Where("status", 1).OrWhereIn(mockDB.Table("departments").Select("COUNT(*)"), 1, 2, 3).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrWhereIn(*Builder IN values)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? OR (SELECT COUNT(*) FROM `departments`) IN (?,?,?)").WithArgs(1, 1, 2, 3).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Where("status", 1).OrWhereIn(NewBuilder(mockDB).Table("departments").Select(clause.Column{
			Name: "COUNT(*)",
			Raw:  true,
		}), 1, 2, 3).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrWhereIn(Callback IN values)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? OR (SELECT count(*) FROM `departments`) IN (?,?,?)").WithArgs(1, 1, 2, 3).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Where("status", 1).OrWhereIn(func(tx *gorm.DB) *gorm.DB {
			var count int64
			return tx.Table("departments").Count(&count)
		}, 1, 2, 3).Find(&dest)
		assert.Nil(t, err)
	})
}

func TestBuilder_WhereNotIn(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}
	t.Run("builder.WhereNotIn(string NOT IN values)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `id` NOT IN (?,?,?)").WithArgs(1, 2, 3).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").WhereNotIn("id", 1, 2, 3).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.WhereNotIn(*gorm.DB NOT IN values)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE (SELECT COUNT(*) FROM `departments`) NOT IN (?,?,?)").WithArgs(1, 2, 3).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").WhereNotIn(mockDB.Table("departments").Select("COUNT(*)"), 1, 2, 3).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.WhereNotIn(*Builder NOT IN values)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE (SELECT COUNT(*) FROM `departments`) NOT IN (?,?,?)").WithArgs(1, 2, 3).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").WhereNotIn(NewBuilder(mockDB).Table("departments").Select(clause.Column{
			Name: "COUNT(*)",
			Raw:  true,
		}), 1, 2, 3).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.WhereNotIn(Callback NOT IN values)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE (SELECT count(*) FROM `departments`) NOT IN (?,?,?)").WithArgs(1, 2, 3).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").WhereNotIn(func(tx *gorm.DB) *gorm.DB {
			var count int64
			return tx.Table("departments").Count(&count)
		}, 1, 2, 3).Find(&dest)
		assert.Nil(t, err)
	})
}

func TestBuilder_OrWhereNotIn(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}
	t.Run("builder.OrWhereNotIn(string NOT IN values)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? OR `id` NOT IN (?,?,?)").WithArgs(1, 1, 2, 3).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Where("status", 1).OrWhereNotIn("id", 1, 2, 3).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrWhereNotIn(*gorm.DB NOT IN values)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? OR (SELECT COUNT(*) FROM `departments`) NOT IN (?,?,?)").WithArgs(1, 1, 2, 3).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Where("status", 1).OrWhereNotIn(mockDB.Table("departments").Select("COUNT(*)"), 1, 2, 3).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrWhereNotIn(*Builder NOT IN values)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? OR (SELECT COUNT(*) FROM `departments`) NOT IN (?,?,?)").WithArgs(1, 1, 2, 3).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Where("status", 1).OrWhereNotIn(NewBuilder(mockDB).Table("departments").Select(clause.Column{
			Name: "COUNT(*)",
			Raw:  true,
		}), 1, 2, 3).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrWhereNotIn(Callback NOT IN values)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? OR (SELECT count(*) FROM `departments`) NOT IN (?,?,?)").WithArgs(1, 1, 2, 3).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Where("status", 1).OrWhereNotIn(func(tx *gorm.DB) *gorm.DB {
			var count int64
			return tx.Table("departments").Count(&count)
		}, 1, 2, 3).Find(&dest)
		assert.Nil(t, err)
	})
}

func TestBuilder_WhereBetween(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}
	t.Run("builder.WhereBetween(string BETWEEN values)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `id` BETWEEN ? AND ?").WithArgs(1, 10).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").WhereBetween("id", 1, 10).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.WhereBetween(string BETWEEN *gorm.DB AND *gorm.DB)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `id` BETWEEN (SELECT MIN(`id`) FROM `vip_users`) AND (SELECT MAX(`id`) FROM `vip_users`)").WithoutArgs().WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").WhereBetween(
			"id",
			mockDB.Table("vip_users").Select("MIN(`id`)"),
			mockDB.Table("vip_users").Select("MAX(`id`)"),
		).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.WhereBetween(string BETWEEN *Builder AND *Builder)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `id` BETWEEN (SELECT MIN(`id`) FROM `vip_users` ) AND (SELECT MAX(`id`) FROM `vip_users` )").WithoutArgs().WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").WhereBetween(
			"id",
			NewBuilder(mockDB).Table("vip_users").Select(clause.Column{Name: "MIN(`id`)", Raw: true}),
			NewBuilder(mockDB).Table("vip_users").Select(clause.Column{Name: "MAX(`id`)", Raw: true}),
		).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.WhereBetween(string BETWEEN Callback AND Callback)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `id` BETWEEN (SELECT MIN(`id`) FROM `vip_users`) AND (SELECT MAX(`id`) FROM `vip_users`)").WithoutArgs().WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").WhereBetween(
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

func TestBuilder_WhereNotBetween(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}
	t.Run("builder.WhereNotBetween(string NOT BETWEEN values)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `id` NOT BETWEEN ? AND ?").WithArgs(1, 10).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").WhereNotBetween("id", 1, 10).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.WhereNotBetween(string NOT BETWEEN *gorm.DB AND *gorm.DB)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `id` NOT BETWEEN (SELECT MIN(`id`) FROM `vip_users`) AND (SELECT MAX(`id`) FROM `vip_users`)").WithoutArgs().WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").WhereNotBetween(
			"id",
			mockDB.Table("vip_users").Select("MIN(`id`)"),
			mockDB.Table("vip_users").Select("MAX(`id`)"),
		).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.WhereNotBetween(string NOT BETWEEN *Builder AND *Builder)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `id` NOT BETWEEN (SELECT MIN(`id`) FROM `vip_users` ) AND (SELECT MAX(`id`) FROM `vip_users` )").WithoutArgs().WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").WhereNotBetween(
			"id",
			NewBuilder(mockDB).Table("vip_users").Select(clause.Column{Name: "MIN(`id`)", Raw: true}),
			NewBuilder(mockDB).Table("vip_users").Select(clause.Column{Name: "MAX(`id`)", Raw: true}),
		).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.WhereNotBetween(string NOT BETWEEN Callback AND Callback)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `id` NOT BETWEEN (SELECT MIN(`id`) FROM `vip_users`) AND (SELECT MAX(`id`) FROM `vip_users`)").WithoutArgs().WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").WhereNotBetween(
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

func TestBuilder_OrWhereBetween(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}
	t.Run("builder.OrWhereBetween(string BETWEEN values)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? OR `id` BETWEEN ? AND ?").WithArgs(1, 1, 10).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Where("status", 1).OrWhereBetween("id", 1, 10).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrWhereBetween(string BETWEEN *gorm.DB AND *gorm.DB)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? OR `id` BETWEEN (SELECT MIN(`id`) FROM `vip_users`) AND (SELECT MAX(`id`) FROM `vip_users`)").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Where("status", 1).OrWhereBetween(
			"id",
			mockDB.Table("vip_users").Select("MIN(`id`)"),
			mockDB.Table("vip_users").Select("MAX(`id`)"),
		).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrWhereBetween(string BETWEEN *Builder AND *Builder)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? OR `id` BETWEEN (SELECT MIN(`id`) FROM `vip_users` ) AND (SELECT MAX(`id`) FROM `vip_users` )").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Where("status", 1).OrWhereBetween(
			"id",
			NewBuilder(mockDB).Table("vip_users").Select(clause.Column{Name: "MIN(`id`)", Raw: true}),
			NewBuilder(mockDB).Table("vip_users").Select(clause.Column{Name: "MAX(`id`)", Raw: true}),
		).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrWhereBetween(string BETWEEN Callback AND Callback)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? OR `id` BETWEEN (SELECT MIN(`id`) FROM `vip_users`) AND (SELECT MAX(`id`) FROM `vip_users`)").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Where("status", 1).OrWhereBetween(
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

func TestBuilder_OrWhereNotBetween(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}
	t.Run("builder.OrWhereNotBetween(string BETWEEN values)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? OR `id` NOT BETWEEN ? AND ?").WithArgs(1, 1, 10).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Where("status", 1).OrWhereNotBetween("id", 1, 10).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrWhereNotBetween(string BETWEEN *gorm.DB AND *gorm.DB)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? OR `id` NOT BETWEEN (SELECT MIN(`id`) FROM `vip_users`) AND (SELECT MAX(`id`) FROM `vip_users`)").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Where("status", 1).OrWhereNotBetween(
			"id",
			mockDB.Table("vip_users").Select("MIN(`id`)"),
			mockDB.Table("vip_users").Select("MAX(`id`)"),
		).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrWhereNotBetween(string BETWEEN *Builder AND *Builder)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? OR `id` NOT BETWEEN (SELECT MIN(`id`) FROM `vip_users` ) AND (SELECT MAX(`id`) FROM `vip_users` )").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Where("status", 1).OrWhereNotBetween(
			"id",
			NewBuilder(mockDB).Table("vip_users").Select(clause.Column{Name: "MIN(`id`)", Raw: true}),
			NewBuilder(mockDB).Table("vip_users").Select(clause.Column{Name: "MAX(`id`)", Raw: true}),
		).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("builder.OrWhereNotBetween(string BETWEEN Callback AND Callback)", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? OR `id` NOT BETWEEN (SELECT MIN(`id`) FROM `vip_users`) AND (SELECT MAX(`id`) FROM `vip_users`)").WithArgs(1).WillReturnRows(result)
		dest := make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Where("status", 1).OrWhereNotBetween(
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

func TestBuilder_WhereLike(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}

	mock.ExpectQuery("SELECT * FROM `users` WHERE `username` LIKE ?").WithArgs("%wardonne%").WillReturnRows(result)
	dest := make([]map[string]any, 0)
	err := NewBuilder(mockDB).Table("users").WhereLike("username", "%wardonne%").Find(&dest)
	assert.Nil(t, err)
}

func TestBuilder_WhereNotLike(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}

	mock.ExpectQuery("SELECT * FROM `users` WHERE `username` NOT LIKE ?").WithArgs("%wardonne%").WillReturnRows(result)
	dest := make([]map[string]any, 0)
	err := NewBuilder(mockDB).Table("users").WhereNotLike("username", "%wardonne%").Find(&dest)
	assert.Nil(t, err)
}

func TestBuilder_OrWhereLike(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}

	mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? OR `username` LIKE ?").WithArgs(1, "%wardonne%").WillReturnRows(result)
	dest := make([]map[string]any, 0)
	err := NewBuilder(mockDB).Table("users").Where("status", 1).OrWhereLike("username", "%wardonne%").Find(&dest)
	assert.Nil(t, err)
}

func TestBuilder_OrWhereNotLike(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}

	mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? OR `username` NOT LIKE ?").WithArgs(1, "%wardonne%").WillReturnRows(result)
	dest := make([]map[string]any, 0)
	err := NewBuilder(mockDB).Table("users").Where("status", 1).OrWhereNotLike("username", "%wardonne%").Find(&dest)
	assert.Nil(t, err)
}

func TestBuilder_WhereExists(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}
	mock.ExpectQuery("SELECT * FROM `users` WHERE EXISTS (SELECT * FROM `departments` WHERE `departments`.`id` = `users`.`department_id` )").WithoutArgs().WillReturnRows(result)
	var dest = make([]map[string]any, 0)
	err := NewBuilder(mockDB).Table("users").WhereExists(NewBuilder(mockDB).Table("departments").WhereRaw("`departments`.`id` = `users`.`department_id`")).Find(&dest)
	assert.Nil(t, err)
}

func TestBuilder_WhereNotExists(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}
	mock.ExpectQuery("SELECT * FROM `users` WHERE NOT EXISTS (SELECT * FROM `departments` WHERE `departments`.`id` = `users`.`department_id` )").WithoutArgs().WillReturnRows(result)
	var dest = make([]map[string]any, 0)
	err := NewBuilder(mockDB).Table("users").WhereNotExists(NewBuilder(mockDB).Table("departments").WhereRaw("`departments`.`id` = `users`.`department_id`")).Find(&dest)
	assert.Nil(t, err)
}

func TestBuilder_OrWhereExists(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}
	mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? OR EXISTS (SELECT * FROM `departments` WHERE `departments`.`id` = `users`.`department_id` )").WithArgs(1).WillReturnRows(result)
	var dest = make([]map[string]any, 0)
	err := NewBuilder(mockDB).Table("users").Where("status", 1).OrWhereExists(NewBuilder(mockDB).Table("departments").WhereRaw("`departments`.`id` = `users`.`department_id`")).Find(&dest)
	assert.Nil(t, err)
}

func TestBuilder_OrWhereNotExists(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}
	mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? OR NOT EXISTS (SELECT * FROM `departments` WHERE `departments`.`id` = `users`.`department_id` )").WithArgs(1).WillReturnRows(result)
	var dest = make([]map[string]any, 0)
	err := NewBuilder(mockDB).Table("users").Where("status", 1).OrWhereNotExists(NewBuilder(mockDB).Table("departments").WhereRaw("`departments`.`id` = `users`.`department_id`")).Find(&dest)
	assert.Nil(t, err)
}

func TestBuilder_WhereBuilder(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}
	mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? AND (`id` = 1 AND `department_id` = 1)").WithArgs(1).WillReturnRows(result)
	var dest = make([]map[string]any, 0)
	err := NewBuilder(mockDB).
		Table("users").
		Where("status", 1).
		WhereBuilder(NewBuilder(mockDB).Where("id", 1).Where("department_id", 1)).
		Find(&dest)
	assert.Nil(t, err)
}

func TestBuilder_WhereNotBuilder(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}
	mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? AND NOT (`id` = 1 AND `department_id` = 1)").WithArgs(1).WillReturnRows(result)
	var dest = make([]map[string]any, 0)
	err := NewBuilder(mockDB).
		Table("users").
		Where("status", 1).
		WhereNotBuilder(NewBuilder(mockDB).Where("id", 1).Where("department_id", 1)).
		Find(&dest)
	assert.Nil(t, err)
}

func TestBuilder_OrWhereBuilder(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}
	mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? OR (`id` = 1 AND `department_id` = 1)").WithArgs(1).WillReturnRows(result)
	var dest = make([]map[string]any, 0)
	err := NewBuilder(mockDB).
		Table("users").
		Where("status", 1).
		OrWhereBuilder(NewBuilder(mockDB).Where("id", 1).Where("department_id", 1)).
		Find(&dest)
	assert.Nil(t, err)
}
func TestBuilder_OrWhereNotBuilder(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}
	mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? OR NOT (`id` = 1 AND `department_id` = 1)").WithArgs(1).WillReturnRows(result)
	var dest = make([]map[string]any, 0)
	err := NewBuilder(mockDB).
		Table("users").
		Where("status", 1).
		OrWhereNotBuilder(NewBuilder(mockDB).Where("id", 1).Where("department_id", 1)).
		Find(&dest)
	assert.Nil(t, err)
}

func TestBuilder_WhereCallback(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}
	mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? AND (`id` = 1 AND `department_id` = 1)").WithArgs(1).WillReturnRows(result)
	var dest = make([]map[string]any, 0)
	err := NewBuilder(mockDB).
		Table("users").
		Where("status", 1).
		WhereCallback(func(builder *Builder) *Builder {
			return builder.Where("id", 1).Where("department_id", 1)
		}).
		Find(&dest)
	assert.Nil(t, err)
}

func TestBuilder_WhereNotCallback(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}
	mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? AND NOT (`id` = 1 AND `department_id` = 1)").WithArgs(1).WillReturnRows(result)
	var dest = make([]map[string]any, 0)
	err := NewBuilder(mockDB).
		Table("users").
		Where("status", 1).
		WhereNotCallback(func(builder *Builder) *Builder {
			return builder.Where("id", 1).Where("department_id", 1)
		}).
		Find(&dest)
	assert.Nil(t, err)
}

func TestBuilder_OrWhereCallback(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}
	mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? OR (`id` = 1 AND `department_id` = 1)").WithArgs(1).WillReturnRows(result)
	var dest = make([]map[string]any, 0)
	err := NewBuilder(mockDB).
		Table("users").
		Where("status", 1).
		OrWhereCallback(func(builder *Builder) *Builder {
			return builder.Where("id", 1).Where("department_id", 1)
		}).
		Find(&dest)
	assert.Nil(t, err)
}

func TestBuilder_OrWhereNotCallback(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}
	mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? OR NOT (`id` = 1 AND `department_id` = 1)").WithArgs(1).WillReturnRows(result)
	var dest = make([]map[string]any, 0)
	err := NewBuilder(mockDB).
		Table("users").
		Where("status", 1).
		OrWhereNotCallback(func(builder *Builder) *Builder {
			return builder.Where("id", 1).Where("department_id", 1)
		}).
		Find(&dest)
	assert.Nil(t, err)
}
