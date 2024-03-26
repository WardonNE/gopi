package builder

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestBuilder_Scopes(t *testing.T) {
	var result = mock.NewRows([]string{"id", "name"}).AddRow(1, "wardonne")
	mock.ExpectQuery("SELECT * FROM `users` WHERE `department_id` = ? AND status = ?").WithArgs(1, 1).WillReturnRows(result)
	var dest = make([]map[string]any, 0)
	err := NewBuilder(mockDB).Table("users").Where("department_id", 1).Scopes(func(tx *gorm.DB) *gorm.DB {
		return tx.Where("status = ?", 1)
	}).Find(&dest)
	assert.Nil(t, err)
}
