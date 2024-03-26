package builder

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestBuilder_FirstOrCreate(t *testing.T) {
	type User struct {
		ID   uint `gorm:"primaryKey"`
		Name string
	}

	t.Run("Builder.FirstOrCreate first failure", func(t *testing.T) {
		expectErr := errors.New("firstorcreate error")
		mock.ExpectQuery("SELECT * FROM `users` ORDER BY `users`.`id` LIMIT ?").WithArgs(1).WillReturnError(expectErr)
		var dest = &User{Name: "wardonne"}
		err := NewBuilder(mockDB).Model(new(User)).FirstOrCreate(&dest)
		assert.ErrorIs(t, err, expectErr)
	})

	t.Run("Builder.FirstOrCreate first success", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` ORDER BY `users`.`id` LIMIT ?").WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "wardonne"))
		var dest = &User{Name: "wardonne1"}
		err := NewBuilder(mockDB).Model(new(User)).FirstOrCreate(&dest)
		assert.EqualValues(t, 1, dest.ID)
		assert.EqualValues(t, "wardonne", dest.Name)
		assert.Nil(t, err)
	})

	t.Run("Builder.FirstOrCreate create failure", func(t *testing.T) {
		expectErr := errors.New("firstorcreate error")
		mock.ExpectQuery("SELECT * FROM `users` ORDER BY `users`.`id` LIMIT ?").WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}))
		mock.ExpectExec("INSERT INTO `users` (`name`) VALUES (?)").WithArgs("wardonne").WillReturnError(expectErr)
		var dest = &User{Name: "wardonne"}
		err := NewBuilder(mockDB).Model(new(User)).FirstOrCreate(&dest)
		assert.ErrorIs(t, err, expectErr)
	})

	t.Run("Builder.FirstOrCreate create success", func(t *testing.T) {
		var result = sqlmock.NewResult(1, 1)
		mock.ExpectQuery("SELECT * FROM `users` ORDER BY `users`.`id` LIMIT ?").WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}))
		mock.ExpectExec("INSERT INTO `users` (`name`) VALUES (?)").WithArgs("wardonne").WillReturnResult(result)
		var dest = &User{Name: "wardonne"}
		err := NewBuilder(mockDB).Model(new(User)).FirstOrCreate(&dest)
		assert.EqualValues(t, 1, dest.ID)
		assert.EqualValues(t, "wardonne", dest.Name)
		assert.Nil(t, err)
	})
}

func TestBuilder_FirstOrInit(t *testing.T) {
	type User struct {
		ID   uint `gorm:"primaryKey"`
		Name string
	}

	t.Run("Builder.FirstOrInit first failure", func(t *testing.T) {
		expectErr := errors.New("firstorcreate error")
		mock.ExpectQuery("SELECT * FROM `users` ORDER BY `users`.`id` LIMIT ?").WithArgs(1).WillReturnError(expectErr)
		var dest = &User{Name: "wardonne"}
		err := NewBuilder(mockDB).Model(new(User)).FirstOrInit(&dest)
		assert.ErrorIs(t, err, expectErr)
	})

	t.Run("Builder.FirstOrInit first success", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` ORDER BY `users`.`id` LIMIT ?").WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "wardonne"))
		var dest = &User{Name: "wardonne1"}
		err := NewBuilder(mockDB).Model(new(User)).FirstOrInit(&dest)
		assert.EqualValues(t, 1, dest.ID)
		assert.EqualValues(t, "wardonne", dest.Name)
		assert.Nil(t, err)
	})

	t.Run("Builder.FirstOrInit init success", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` ORDER BY `users`.`id` LIMIT ?").WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}))
		var dest = &User{Name: "wardonne"}
		err := NewBuilder(mockDB).Model(new(User)).FirstOrInit(&dest)
		assert.EqualValues(t, "wardonne", dest.Name)
		assert.Nil(t, err)
	})

	t.Run("Builder.FirstOrInit init with Builder.Attrs success", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` ORDER BY `users`.`id` LIMIT ?").WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}))
		var dest = &User{Name: "wardonne"}
		err := NewBuilder(mockDB).Model(new(User)).Attrs(&User{ID: 2}).FirstOrInit(&dest)
		assert.EqualValues(t, 2, dest.ID)
		assert.EqualValues(t, "wardonne", dest.Name)
		assert.Nil(t, err)
	})

	t.Run("Builder.FirstOrInit init with Builder.Assign success", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` ORDER BY `users`.`id` LIMIT ?").WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}))
		var dest = &User{Name: "wardonne"}
		err := NewBuilder(mockDB).Model(new(User)).Assign(&User{ID: 2}).FirstOrInit(&dest)
		assert.EqualValues(t, 2, dest.ID)
		assert.EqualValues(t, "wardonne", dest.Name)
		assert.Nil(t, err)
	})
}

func TestBuilder_Create(t *testing.T) {
	t.Run("Builder.Create failure", func(t *testing.T) {
		expectErr := errors.New("create error")
		mock.ExpectExec("INSERT INTO `users` (`id`,`name`) VALUES (?,?)").WithArgs(1, "wardonne").WillReturnError(expectErr)
		var dest = map[string]any{
			"id":   1,
			"name": "wardonne",
		}
		err := NewBuilder(mockDB).Table("users").Create(&dest)
		assert.ErrorIs(t, err, expectErr)
	})

	t.Run("Builder.Create success", func(t *testing.T) {
		mock.ExpectExec("INSERT INTO `users` (`id`,`name`) VALUES (?,?)").WithArgs(1, "wardonne").WillReturnResult(sqlmock.NewResult(1, 1))
		var dest = map[string]any{
			"id":   1,
			"name": "wardonne",
		}
		err := NewBuilder(mockDB).Table("users").Create(&dest)
		assert.EqualValues(t, 1, dest["id"])
		assert.EqualValues(t, "wardonne", dest["name"])
		assert.Nil(t, err)
	})
}

func TestBuilder_Upsert(t *testing.T) {
	t.Run("Builder.Upsert failure", func(t *testing.T) {
		expectErr := errors.New("upsert error")
		mock.ExpectExec("INSERT INTO `users` (`department_id`,`name`) VALUES (?,?) ON DUPLICATE KEY UPDATE `name`=VALUES(`name`)").WithArgs(1, "wardonne2").WillReturnError(expectErr)
		err := NewBuilder(mockDB).Table("users").Upsert(map[string]any{
			"department_id": 1,
			"name":          "wardonne2",
		}, nil, "name")
		assert.ErrorIs(t, err, expectErr)
	})

	t.Run("Builder.Upsert success", func(t *testing.T) {
		mock.ExpectExec("INSERT INTO `users` (`department_id`,`name`) VALUES (?,?) ON DUPLICATE KEY UPDATE `name`=VALUES(`name`)").WithArgs(1, "wardonne2").WillReturnResult(sqlmock.NewResult(1, 1))
		var dest = map[string]any{
			"department_id": 1,
			"name":          "wardonne2",
		}
		err := NewBuilder(mockDB).Table("users").Upsert(dest, nil, "name")
		assert.Nil(t, err)
	})
}
