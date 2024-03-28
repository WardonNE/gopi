package builder

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestBuilder_Exec(t *testing.T) {
	t.Run("Builder.Exec failure", func(t *testing.T) {
		expectErr := errors.New("exec error")
		mock.ExpectExec("INSERT INTO `users` (`name`) VALUES (?)").WithArgs("wardonne").WillReturnError(expectErr)
		rows, err := NewBuilder(mockDB).Exec("INSERT INTO `users` (`name`) VALUES (?)", "wardonne")
		assert.EqualValues(t, 0, rows)
		assert.ErrorIs(t, err, expectErr)
	})

	t.Run("Builder.Exec success", func(t *testing.T) {
		mock.ExpectExec("INSERT INTO `users` (`name`) VALUES (?)").WithArgs("wardonne").WillReturnResult(sqlmock.NewResult(1, 1))
		rows, err := NewBuilder(mockDB).Exec("INSERT INTO `users` (`name`) VALUES (?)", "wardonne")
		assert.EqualValues(t, 1, rows)
		assert.Nil(t, err)
	})
}

func TestBuilder_Fetch(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}
	t.Run("Builder.Fetch failure", func(t *testing.T) {
		expectErr := errors.New("fetch error")
		mock.ExpectQuery("SELECT * FROM `users`").WithoutArgs().WillReturnError(expectErr)
		var dest = make([]map[string]any, 0)
		err := NewBuilder(mockDB).Fetch(&dest, "SELECT * FROM `users`")
		assert.ErrorIs(t, err, expectErr)
	})

	t.Run("Builder.Fetch success", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users`").WithoutArgs().WillReturnRows(result)
		var dest = make([]map[string]any, 0)
		err := NewBuilder(mockDB).Fetch(&dest, "SELECT * FROM `users`")
		assert.Nil(t, err)
	})
}
