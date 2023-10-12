package pagination

import "github.com/wardonne/gopi/utils"

type ArrayPaginator[T any] struct {
	Paginator[T]

	allItems []T
}

func NewArrayPaginator[T any](items []T, pageSize, page int) *ArrayPaginator[T] {
	paginator := new(ArrayPaginator[T])
	paginator.allItems = items
	page = utils.If(page <= 0, 1, page)
	pageSize = utils.If(pageSize <= 0, 10, pageSize)
	paginator.total = int64(len(items))
	paginator.currentPage = page
	paginator.pageSize = pageSize
	var values []T
	startIndex := int64((page - 1) * pageSize)
	endIndex := startIndex + int64(pageSize)
	if startIndex > paginator.total {
		values = make([]T, 0)
	} else if endIndex > paginator.total {
		values = items[startIndex:]
	} else {
		values = items[startIndex:endIndex]
	}
	paginator.Paginator = *NewPaginator[T](int64(len(items)), values, pageSize, page)
	return paginator
}

func (p *ArrayPaginator[T]) All() []T {
	return p.allItems
}
