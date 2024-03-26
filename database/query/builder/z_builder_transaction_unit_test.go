package builder

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestBuilder_Transaction(t *testing.T) {
	t.Run("Builder.Transaction failure", func(t *testing.T) {
		expectErr := errors.New("transaction error")
		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO `users` (`name`) VALUES (?)").WithArgs("wardonne").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectExec("UPDATE `users` SET `status`=? WHERE `id` = ?").WithArgs(1, 1).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectQuery("SELECT * FROM `departments` WHERE `status` = ?").WithArgs(1).WillReturnError(expectErr)
		mock.ExpectRollback()
		err := NewBuilder(mockDB).Transaction(func(builder *Builder) error {
			if err := builder.Table("users").Create(map[string]any{"name": "wardonne"}); err != nil {
				return err
			}
			if err := builder.Table("users").Where("id", 1).Update(map[string]any{"status": 1}); err != nil {
				return err
			}
			var dest = make([]map[string]any, 0)
			if err := builder.Table("departments").Where("status", 1).Find(&dest); err != nil {
				return err
			}
			return nil
		})
		assert.ErrorIs(t, err, expectErr)
	})

	t.Run("Builder.Transaction success", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO `users` (`name`) VALUES (?)").WithArgs("wardonne").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectExec("UPDATE `users` SET `status`=? WHERE `id` = ?").WithArgs(1, 1).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectQuery("SELECT * FROM `departments` WHERE `status` = ?").WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "wardonne"))
		mock.ExpectCommit()
		err := NewBuilder(mockDB).Transaction(func(builder *Builder) error {
			if err := builder.Table("users").Create(map[string]any{"name": "wardonne"}); err != nil {
				return err
			}
			if err := builder.Table("users").Where("id", 1).Update(map[string]any{"status": 1}); err != nil {
				return err
			}
			var dest = make([]map[string]any, 0)
			if err := builder.Table("departments").Where("status", 1).Find(&dest); err != nil {
				return err
			}
			return nil
		})
		assert.Nil(t, err)
	})

	t.Run("Builder.Transaction Nested", func(t *testing.T) {
		expectErr := errors.New("nest transaction error")
		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO `users` (`name`) VALUES (?)").WithArgs("wardonne").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectExec("SAVEPOINT sp1").WithoutArgs().WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectExec("UPDATE `users` SET `status`=? WHERE `id` = ?").WithArgs(1, 1).WillReturnError(expectErr)
		mock.ExpectExec("ROLLBACK TO SAVEPOINT sp1").WithoutArgs().WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectQuery("SELECT * FROM `departments` WHERE `status` = ?").WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "wardonne"))
		mock.ExpectCommit()

		err := NewBuilder(mockDB).Transaction(func(builder *Builder) error {
			err := builder.Table("users").Create(map[string]any{"name": "wardonne"})
			if err != nil {
				return err
			}
			builder.Transaction(func(builder *Builder) error {
				err = builder.Table("users").Where("id", 1).Update(map[string]any{"status": 1})
				if err != nil {
					assert.ErrorIs(t, err, expectErr)
					return err
				}
				return nil
			})
			var dest = make([]map[string]any, 0)
			err = builder.Table("departments").Where("status", 1).Find(&dest)
			if err != nil {
				return err
			}
			return nil
		})
		assert.Nil(t, err)
	})
}

func TestBuilder_Begin(t *testing.T) {
	t.Run("Builder.Begin failure", func(t *testing.T) {
		expectErr := errors.New("transaction error")
		builder := NewBuilder(mockDB)
		defer func() {
			if err := recover(); err != nil {
				builder.Rollback()
				assert.ErrorIs(t, err.(error), expectErr)
			}
		}()
		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO `users` (`name`) VALUES (?)").WithArgs("wardonne").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectExec("UPDATE `users` SET `status`=? WHERE `id` = ?").WithArgs(1, 1).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectQuery("SELECT * FROM `departments` WHERE `status` = ?").WithArgs(1).WillReturnError(expectErr)
		mock.ExpectRollback()
		builder.Begin()
		err := builder.Table("users").Create(map[string]any{"name": "wardonne"})
		if err != nil {
			panic(err)
		}
		err = builder.Table("users").Where("id", 1).Update(map[string]any{"status": 1})
		if err != nil {
			panic(err)
		}
		var dest = make([]map[string]any, 0)
		err = builder.Table("departments").Where("status", 1).Find(&dest)
		if err != nil {
			panic(err)
		}
		builder.Commit()
	})

	t.Run("Builder.Begin success", func(t *testing.T) {
		builder := NewBuilder(mockDB)
		defer func() {
			if err := recover(); err != nil {
				builder.Rollback()
			}
		}()
		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO `users` (`name`) VALUES (?)").WithArgs("wardonne").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectExec("UPDATE `users` SET `status`=? WHERE `id` = ?").WithArgs(1, 1).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectQuery("SELECT * FROM `departments` WHERE `status` = ?").WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "wardonne"))
		mock.ExpectCommit()
		builder.Begin()
		err := builder.Table("users").Create(map[string]any{"name": "wardonne"})
		if err != nil {
			panic(err)
		}
		err = builder.Table("users").Where("id", 1).Update(map[string]any{"status": 1})
		if err != nil {
			panic(err)
		}
		var dest = make([]map[string]any, 0)
		err = builder.Table("departments").Where("status", 1).Find(&dest)
		if err != nil {
			panic(err)
		}
		err = builder.Commit()
		assert.Nil(t, err)
	})

	t.Run("Builder.Begin Nest Transaction", func(t *testing.T) {
		expectErr := errors.New("nest transaction error")
		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO `users` (`name`) VALUES (?)").WithArgs("wardonne").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectExec("SAVEPOINT sp1").WithoutArgs().WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectExec("UPDATE `users` SET `status`=? WHERE `id` = ?").WithArgs(1, 1).WillReturnError(expectErr)
		mock.ExpectExec("ROLLBACK TO SAVEPOINT sp1").WithoutArgs().WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectQuery("SELECT * FROM `departments` WHERE `status` = ?").WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "wardonne"))
		mock.ExpectCommit()

		builder := NewBuilder(mockDB)
		builder.Begin()
		err := builder.Table("users").Create(map[string]any{"name": "wardonne"})
		if err != nil {
			builder.Rollback()
		}
		builder.Begin()
		err = builder.Table("users").Where("id", 1).Update(map[string]any{"status": 1})
		if err != nil {
			assert.ErrorIs(t, err, expectErr)
			builder.Rollback()
		}
		var dest = make([]map[string]any, 0)
		err = builder.Table("departments").Where("status", 1).Find(&dest)
		if err != nil {
			builder.Rollback()
		}
		builder.Commit()
	})
}
