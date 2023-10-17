package pagination

import "github.com/wardonne/gopi/support/utils"

// ArrayPaginator is a paginator based on an array, it extends from [Paginator]
//
// Example:
//   allItems := []int{}
//   for i := 0; i < 100; i++ {
//       allItems = append(allItems, i)
//   }
//
//   pageSize := 10
//   page := 1
//
//   paginator := Array[int](allItems, pageSize, page)
type ArrayPaginator[T any] struct {
	Paginator[T]

	allItems []T
}

// Array creates an [ArrayPaginator] instance
func Array[T any](allItems []T, pageSize, page int) *ArrayPaginator[T] {
	paginator := new(ArrayPaginator[T])
	paginator.allItems = allItems
	page = utils.If(page <= 0, 1, page)
	pageSize = utils.If(pageSize <= 0, 10, pageSize)
	paginator.total = int64(len(allItems))
	paginator.currentPage = page
	paginator.pageSize = pageSize
	var values []T
	startIndex := int64((page - 1) * pageSize)
	endIndex := startIndex + int64(pageSize)
	if startIndex > paginator.total {
		values = make([]T, 0)
	} else if endIndex > paginator.total {
		values = allItems[startIndex:]
	} else {
		values = allItems[startIndex:endIndex]
	}
	paginator.Paginator = *New[T](int64(len(allItems)), values, pageSize, page)
	return paginator
}

// All returns all items
func (p *ArrayPaginator[T]) All() []T {
	return p.allItems
}
