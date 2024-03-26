package builder

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBuilder_MaxExecutionTime(t *testing.T) {
	var result = mock.NewRows([]string{"id", "name"}).AddRow(1, "wardonne1").AddRow(2, "wardonne2")
	mock.ExpectQuery("SELECT /*+ MAX_EXECUTION_TIME(10000) */ * FROM `users`").WithoutArgs().WillReturnRows(result)
	var dest = make([]map[string]any, 0)
	err := NewBuilder(mockDB).Table("users").MaxExecutionTime(10 * time.Second).Find(&dest)
	assert.Nil(t, err)
}

func TestBuilder_UseIndex(t *testing.T) {
	var result = mock.NewRows([]string{"id", "name"}).AddRow(1, "wardonne1")
	mock.ExpectQuery("SELECT * FROM `users` USE INDEX (`idx_department_id`)").WithoutArgs().WillReturnRows(result)
	var dest = make([]map[string]any, 0)
	err := NewBuilder(mockDB).Table("users").UseIndex("idx_department_id").Find(&dest)
	assert.Nil(t, err)
}

func TestBuilder_IgnoreIndex(t *testing.T) {
	var result = mock.NewRows([]string{"id", "name"}).AddRow(1, "wardonne1")
	mock.ExpectQuery("SELECT * FROM `users` IGNORE INDEX (`idx_department_id`)").WithoutArgs().WillReturnRows(result)
	var dest = make([]map[string]any, 0)
	err := NewBuilder(mockDB).Table("users").IgnoreIndex("idx_department_id").Find(&dest)
	assert.Nil(t, err)
}

func TestBuilder_ForceIndex(t *testing.T) {
	var result = mock.NewRows([]string{"id", "name"}).AddRow(1, "wardonne1")
	mock.ExpectQuery("SELECT * FROM `users` FORCE INDEX (`idx_department_id`)").WithoutArgs().WillReturnRows(result)
	var dest = make([]map[string]any, 0)
	err := NewBuilder(mockDB).Table("users").ForceIndex("idx_department_id").Find(&dest)
	assert.Nil(t, err)
}

func TestBuilder_ForceIndexForJoin(t *testing.T) {
	var result = mock.NewRows([]string{"id", "name"}).AddRow(1, "wardonne1")
	mock.ExpectQuery("SELECT * FROM `users` FORCE INDEX FOR JOIN (`idx_department_id`)").WithoutArgs().WillReturnRows(result)
	var dest = make([]map[string]any, 0)
	err := NewBuilder(mockDB).Table("users").ForceIndexForJoin("idx_department_id").Find(&dest)
	assert.Nil(t, err)
}
func TestBuilder_ForceIndexForOrderBy(t *testing.T) {
	var result = mock.NewRows([]string{"id", "name"}).AddRow(1, "wardonne1")
	mock.ExpectQuery("SELECT * FROM `users` FORCE INDEX FOR ORDER BY (`idx_department_id`)").WithoutArgs().WillReturnRows(result)
	var dest = make([]map[string]any, 0)
	err := NewBuilder(mockDB).Table("users").ForceIndexForOrderBy("idx_department_id").Find(&dest)
	assert.Nil(t, err)
}

func TestBuilder_ForceIndexForGroupBy(t *testing.T) {
	var result = mock.NewRows([]string{"id", "name"}).AddRow(1, "wardonne1")
	mock.ExpectQuery("SELECT * FROM `users` FORCE INDEX FOR GROUP BY (`idx_department_id`)").WithoutArgs().WillReturnRows(result)
	var dest = make([]map[string]any, 0)
	err := NewBuilder(mockDB).Table("users").ForceIndexForGroupBy("idx_department_id").Find(&dest)
	assert.Nil(t, err)
}
