package builder

import (
	"database/sql/driver"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestBuilder_With(t *testing.T) {
	type Department struct {
		ID   uint64
		Name string
	}
	type User struct {
		ID           uint64
		Name         string
		DepartmentID uint64
		Department   *Department
	}
	mock.ExpectQuery("SELECT * FROM `users`").WithoutArgs().WillReturnRows(sqlmock.NewRows([]string{"id", "name", "department_id"}).AddRows(
		[]driver.Value{1, "wardonne1", 1},
		[]driver.Value{2, "wardonne2", 1},
		[]driver.Value{3, "wardonne3", 2},
	))
	mock.ExpectQuery("SELECT * FROM `departments` WHERE `departments`.`id` IN (?,?)").WithArgs(1, 2).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "department1").AddRow(2, "department2"))
	var dest = make([]User, 0)
	err := NewBuilder(mockDB).Model(new(User)).With("Department").Find(&dest)
	assert.Nil(t, err)
}
