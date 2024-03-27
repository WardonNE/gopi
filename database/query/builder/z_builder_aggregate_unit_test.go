package builder

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestBuilder_Count(t *testing.T) {
	t.Run("Builder.Count failure", func(t *testing.T) {
		expectErr := errors.New("count error")
		mock.ExpectQuery("SELECT count(*) FROM `users`").WithoutArgs().WillReturnError(expectErr)
		count, err := NewBuilder(mockDB).Table("users").Count()
		assert.EqualValues(t, 0, count)
		assert.ErrorIs(t, err, expectErr)
	})

	t.Run("Builder.Count success", func(t *testing.T) {
		var result = sqlmock.NewRows([]string{"id"})
		for _, item := range mockUsers {
			result.AddRow(item["id"])
		}
		mock.ExpectQuery("SELECT count(*) FROM `users`").WithoutArgs().WillReturnRows(result)
		count, err := NewBuilder(mockDB).Table("users").Count()
		assert.EqualValues(t, len(mockUsers), count)
		assert.Nil(t, err)
	})
}
func TestBuilder_Sum(t *testing.T) {
	t.Run("Builder.Sum failure", func(t *testing.T) {
		expectErr := errors.New("sum error")
		mock.ExpectQuery("SELECT SUM(`amount`) AS `sum` FROM `orders`").WithoutArgs().WillReturnError(expectErr)
		sum, err := NewBuilder(mockDB).Table("orders").Sum("amount")
		assert.EqualValues(t, 0, sum)
		assert.ErrorIs(t, err, expectErr)
	})

	t.Run("Builder.Sum success", func(t *testing.T) {
		mock.ExpectQuery("SELECT SUM(`amount`) AS `sum` FROM `orders`").WithoutArgs().WillReturnRows(sqlmock.NewRows([]string{"sum"}).AddRow(100))
		sum, err := NewBuilder(mockDB).Table("orders").Sum("amount")
		assert.EqualValues(t, 100, sum)
		assert.Nil(t, err)
	})
}

func TestBuilder_Avg(t *testing.T) {
	t.Run("Builder.Avg failure", func(t *testing.T) {
		expectErr := errors.New("avg error")
		mock.ExpectQuery("SELECT AVG(`amount`) AS `avg` FROM `orders`").WithoutArgs().WillReturnError(expectErr)
		avg, err := NewBuilder(mockDB).Table("orders").Avg("amount")
		assert.EqualValues(t, 0, avg)
		assert.ErrorIs(t, err, expectErr)
	})

	t.Run("Builder.Avg success", func(t *testing.T) {
		mock.ExpectQuery("SELECT AVG(`amount`) AS `avg` FROM `orders`").WithoutArgs().WillReturnRows(sqlmock.NewRows([]string{"sum"}).AddRow(100))
		avg, err := NewBuilder(mockDB).Table("orders").Avg("amount")
		assert.EqualValues(t, 100, avg)
		assert.Nil(t, err)
	})
}

func TestBuilder_Max(t *testing.T) {
	t.Run("Builder.Max failure", func(t *testing.T) {
		expectErr := errors.New("avg error")
		mock.ExpectQuery("SELECT MAX(`amount`) AS `max` FROM `orders`").WithoutArgs().WillReturnError(expectErr)
		avg, err := NewBuilder(mockDB).Table("orders").Max("amount")
		assert.EqualValues(t, 0, avg)
		assert.ErrorIs(t, err, expectErr)
	})

	t.Run("Builder.Max success", func(t *testing.T) {
		mock.ExpectQuery("SELECT MAX(`amount`) AS `max` FROM `orders`").WithoutArgs().WillReturnRows(sqlmock.NewRows([]string{"sum"}).AddRow(100))
		avg, err := NewBuilder(mockDB).Table("orders").Max("amount")
		assert.EqualValues(t, 100, avg)
		assert.Nil(t, err)
	})
}

func TestBuilder_Min(t *testing.T) {
	t.Run("Builder.Min failure", func(t *testing.T) {
		expectErr := errors.New("avg error")
		mock.ExpectQuery("SELECT MIN(`amount`) AS `min` FROM `orders`").WithoutArgs().WillReturnError(expectErr)
		avg, err := NewBuilder(mockDB).Table("orders").Min("amount")
		assert.EqualValues(t, 0, avg)
		assert.ErrorIs(t, err, expectErr)
	})

	t.Run("Builder.Min success", func(t *testing.T) {
		mock.ExpectQuery("SELECT MIN(`amount`) AS `min` FROM `orders`").WithoutArgs().WillReturnRows(sqlmock.NewRows([]string{"sum"}).AddRow(100))
		avg, err := NewBuilder(mockDB).Table("orders").Min("amount")
		assert.EqualValues(t, 100, avg)
		assert.Nil(t, err)
	})
}
