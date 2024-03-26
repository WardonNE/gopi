package builder

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestBuilder_Update(t *testing.T) {
	t.Run("Builder.Update failure", func(t *testing.T) {
		expectErr := errors.New("update error")
		mock.ExpectExec("UPDATE `users` SET `status`=? WHERE `id` = ?").WithArgs(1, 1).WillReturnError(expectErr)
		var dest = map[string]any{
			"status": 1,
		}
		err := NewBuilder(mockDB).Table("users").Where("id", 1).Update(dest)
		assert.ErrorIs(t, err, expectErr)
	})

	t.Run("Builder.Update success", func(t *testing.T) {
		mock.ExpectExec("UPDATE `users` SET `status`=? WHERE `id` = ?").WithArgs(1, 1).WillReturnResult(sqlmock.NewResult(1, 1))
		var dest = map[string]any{
			"status": 1,
		}
		err := NewBuilder(mockDB).Table("users").Where("id", 1).Update(dest)
		assert.Nil(t, err)
	})
}
