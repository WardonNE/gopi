package builder

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestBuilder_Limit(t *testing.T) {
	result := sqlmock.NewRows([]string{"id", "name"})
	for _, user := range mockUsers {
		result.AddRow(user["id"], user["name"])
	}

	t.Run("Only Limit", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` LIMIT ?").WithArgs(1).WillReturnRows(result)
		var dest = make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Limit(1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("Only Offset", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` OFFSET ?").WithArgs(1).WillReturnRows(result)
		var dest = make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Offset(1).Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("Mix Limit and Offset", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` LIMIT ? OFFSET ?").WithArgs(1, 1).WillReturnRows(result)
		var dest = make([]map[string]any, 0)
		err := NewBuilder(mockDB).Table("users").Limit(1).Offset(1).Find(&dest)
		assert.Nil(t, err)
	})
}
