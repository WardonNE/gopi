package builder

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestBuilder_Delete(t *testing.T) {
	t.Run("Builder.Delete failure", func(t *testing.T) {
		expectErr := errors.New("delete error")
		mock.ExpectExec("DELETE FROM `users` WHERE `id` = ?").WithArgs(1).WillReturnError(expectErr)
		err := NewBuilder(mockDB).Table("users").Where("id", 1).Delete()
		assert.ErrorIs(t, err, expectErr)
	})

	t.Run("Builder.Delete success", func(t *testing.T) {
		mock.ExpectExec("DELETE FROM `users` WHERE `id` = ?").WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))
		err := NewBuilder(mockDB).Table("users").Where("id", 1).Delete()
		assert.Nil(t, err)
	})
}
