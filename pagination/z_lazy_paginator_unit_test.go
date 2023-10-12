package pagination

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLazyPaginator(t *testing.T) {
	var totalLoader = func() int64 { return 99 }
	var itemsLoader = func() []int { return []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} }

	firstPage := NewLazyPaginator[int](totalLoader, itemsLoader, 10, 1)
	assert.EqualValues(t, 0, firstPage.total)
	assert.Equal(t, []int(nil), firstPage.items)
	assert.Equal(t, 0, firstPage.lastPage)
	assert.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, firstPage.Items())
	assert.Equal(t, 0, firstPage.FirstItemIndex())
	assert.Equal(t, 9, firstPage.LastItemIndex())
	assert.EqualValues(t, 99, firstPage.Total())
	assert.Equal(t, 1, firstPage.CurrentPage())
	assert.Equal(t, 10, firstPage.PageSize())
	assert.True(t, firstPage.HasMore())
	assert.Equal(t, 10, firstPage.LastPage())
}
