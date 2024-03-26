package builder

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func TestBuilder_Count(t *testing.T) {
	t.Run("Builder.Count failure", func(t *testing.T) {
		expectErr := errors.New("count error")
		mock.ExpectQuery("SELECT count(*) FROM `users`").WithoutArgs().WillReturnError(expectErr)
		var dest int64
		err := NewBuilder(mockDB).Table("users").Count(&dest)
		assert.ErrorIs(t, err, expectErr)
	})

	t.Run("Builder.Count success", func(t *testing.T) {
		// var result = sqlmock.NewRows([]string{"count(*)"}).AddRow(len(mockUsers))
		var result = sqlmock.NewRows([]string{"id"})
		for _, item := range mockUsers {
			result.AddRow(item["id"])
		}
		mock.ExpectQuery("SELECT count(*) FROM `users`").WithoutArgs().WillReturnRows(result)
		var dest int64
		err := NewBuilder(mockDB).Table("users").Count(&dest)
		assert.EqualValues(t, len(mockUsers), dest)
		assert.Nil(t, err)
	})
}

func TestBuilder_Take(t *testing.T) {
	t.Run("Builder.Take failure", func(t *testing.T) {
		expectErr := errors.New("take error")
		mock.ExpectQuery("SELECT * FROM `users` ORDER BY `username` DESC LIMIT ?").WithArgs(1).WillReturnError(expectErr)
		var dest = map[string]any{}
		err := NewBuilder(mockDB).
			Table("users").
			OrderDesc("username").
			Take(&dest)
		assert.ErrorIs(t, err, expectErr)
	})

	t.Run("Builder.Take success", func(t *testing.T) {
		var result = sqlmock.NewRows([]string{"id", "username"}).AddRow(1, "wardonne")
		mock.ExpectQuery("SELECT * FROM `users` ORDER BY `username` DESC LIMIT ?").WithArgs(1).WillReturnRows(result)
		var dest = map[string]any{}
		err := NewBuilder(mockDB).
			Table("users").
			OrderDesc("username").
			Take(&dest)
		assert.Nil(t, err)
		assert.EqualValues(t, 1, dest["id"])
		assert.EqualValues(t, "wardonne", dest["username"])
	})
}

func TestBuilder_First(t *testing.T) {
	type User struct {
		ID   uint `gorm:"primarykey"`
		Name string
	}
	t.Run("Builder.First failure", func(t *testing.T) {
		expectErr := errors.New("first error")
		mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? ORDER BY `users`.`id` LIMIT ?").WithArgs(1, 1).WillReturnError(expectErr)
		var dest = map[string]any{}
		err := NewBuilder(mockDB).Model(new(User)).Where("status", 1).First(&dest)
		assert.ErrorIs(t, err, expectErr)
	})

	t.Run("Builder.First success", func(t *testing.T) {
		var result = sqlmock.NewRows([]string{"id", "username"}).AddRow(1, "wardonne")
		mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? ORDER BY `users`.`id` LIMIT ?").WithArgs(1, 1).WillReturnRows(result)
		var dest = map[string]any{}
		err := NewBuilder(mockDB).Model(new(User)).Where("status", 1).First(&dest)
		assert.Nil(t, err)
		assert.EqualValues(t, 1, dest["id"])
		assert.EqualValues(t, "wardonne", dest["username"])
	})
}

func TestBuilder_Last(t *testing.T) {
	type User struct {
		ID   uint `gorm:"primarykey"`
		Name string
	}
	t.Run("Builder.Last failure", func(t *testing.T) {
		expectErr := errors.New("first error")
		mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? ORDER BY `users`.`id` DESC LIMIT ?").WithArgs(1, 1).WillReturnError(expectErr)
		var dest = map[string]any{}
		err := NewBuilder(mockDB).Model(new(User)).Where("status", 1).Last(&dest)
		assert.ErrorIs(t, err, expectErr)
	})

	t.Run("Builder.Last success", func(t *testing.T) {
		var result = sqlmock.NewRows([]string{"id", "username"}).AddRow(1, "wardonne")
		mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? ORDER BY `users`.`id` DESC LIMIT ?").WithArgs(1, 1).WillReturnRows(result)
		var dest = map[string]any{}
		err := NewBuilder(mockDB).Model(new(User)).Where("status", 1).Last(&dest)
		assert.Nil(t, err)
		assert.EqualValues(t, 1, dest["id"])
		assert.EqualValues(t, "wardonne", dest["username"])
	})
}

func TestBuilder_Find(t *testing.T) {
	t.Run("Builder.Find failure", func(t *testing.T) {
		expectErr := errors.New("find error")
		mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ?").WithArgs(1).WillReturnError(expectErr)
		var dest = []map[string]any{}
		err := NewBuilder(mockDB).Table("users").Where("status", 1).Find(&dest)
		assert.ErrorIs(t, err, expectErr)
	})

	t.Run("Builder.Find success", func(t *testing.T) {
		var result = sqlmock.NewRows([]string{"id", "username"}).AddRow(1, "wardonne")
		mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ?").WithArgs(1).WillReturnRows(result)
		var dest = []map[string]any{}
		err := NewBuilder(mockDB).Table("users").Where("status", 1).Find(&dest)
		assert.Nil(t, err)
		assert.EqualValues(t, 1, dest[0]["id"])
		assert.EqualValues(t, "wardonne", dest[0]["username"])
	})
}

func TestBuilder_Pluck(t *testing.T) {
	t.Run("Builder.Pluck failure", func(t *testing.T) {
		expectErr := errors.New("pluck error")
		mock.ExpectQuery("SELECT `id` FROM `users` WHERE `status` = ?").WithArgs(1).WillReturnError(expectErr)
		var dest = []int{}
		err := NewBuilder(mockDB).Table("users").Where("status", 1).Pluck("id", &dest)
		assert.ErrorIs(t, err, expectErr)
	})

	t.Run("Builder.Pluck success", func(t *testing.T) {
		var result = sqlmock.NewRows([]string{"id"}).AddRow(1).AddRow(2)
		mock.ExpectQuery("SELECT `users`.`id` FROM `users` WHERE `status` = ?").WithArgs(1).WillReturnRows(result)
		var dest = []int{}
		err := NewBuilder(mockDB).Table("users").Where("status", 1).Pluck(clause.Column{
			Name:  "id",
			Table: "users",
		}, &dest)
		assert.Nil(t, err)
		assert.EqualValues(t, []int{1, 2}, dest)
	})
}

func TestBuilder_Chunk(t *testing.T) {
	type User struct {
		ID uint64 `gorm:"primaryKey"`
	}
	t.Run("Builder.Chunk failure", func(t *testing.T) {
		expectErr := errors.New("chunk error")
		mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? ORDER BY `users`.`id` LIMIT ?").WithArgs(1, 1).WillReturnError(expectErr)
		var dest = []User{}
		err := NewBuilder(mockDB).Model(new(User)).Where("status", 1).Chunk(&dest, 1, func(tx *gorm.DB, batch int) error {
			return nil
		})
		assert.ErrorIs(t, err, expectErr)
	})
}

func TestBuilder_Cursor(t *testing.T) {
	t.Run("Builder.Cursor failure", func(t *testing.T) {
		expectErr := errors.New("cursor error")
		mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ?").WithArgs(1).WillReturnError(expectErr)
		var dest = map[string]any{}
		err := NewBuilder(mockDB).
			Table("users").
			Where("status", 1).
			Cursor(&dest, func() error {
				return nil
			})
		assert.ErrorIs(t, err, expectErr)
	})

	t.Run("Builder.Cursor success", func(t *testing.T) {
		var result = mock.NewRows([]string{"id", "name"}).
			AddRow(1, "wardonne1").
			AddRow(2, "wardonne2")
		mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ?").WithArgs(1).WillReturnRows(result)
		var dest = map[string]any{}
		var index = 1
		err := NewBuilder(mockDB).
			Table("users").
			Where("status", 1).
			Cursor(&dest, func() error {
				assert.EqualValues(t, index, dest["id"])
				index++
				return nil
			})
		assert.Nil(t, err)
	})
}
