package pagination

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayPaginator(t *testing.T) {
	ints := []int{}
	for i := 0; i < 99; i++ {
		ints = append(ints, i)
	}
	firstPage := Array[int](ints, 10, 1)
	assert.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, firstPage.Items())
	assert.Equal(t, 0, firstPage.FirstItemIndex())
	assert.Equal(t, 9, firstPage.LastItemIndex())
	assert.EqualValues(t, 99, firstPage.Total())
	assert.Equal(t, 1, firstPage.CurrentPage())
	assert.Equal(t, 10, firstPage.PageSize())
	assert.True(t, firstPage.HasMore())
	assert.Equal(t, 10, firstPage.LastPage())
	assert.Equal(t, ints, firstPage.All())

	lastPage := Array[int](ints, 10, 10)
	assert.Equal(t, []int{90, 91, 92, 93, 94, 95, 96, 97, 98}, lastPage.Items())
	assert.Equal(t, 90, lastPage.FirstItemIndex())
	assert.Equal(t, 98, lastPage.LastItemIndex())
	assert.EqualValues(t, 99, lastPage.Total())
	assert.Equal(t, 10, lastPage.CurrentPage())
	assert.Equal(t, 10, lastPage.PageSize())
	assert.False(t, lastPage.HasMore())
	assert.Equal(t, 10, lastPage.LastPage())
	assert.Equal(t, ints, lastPage.All())
}
