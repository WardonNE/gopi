package pagination

type ArrayPaginator[T any] struct {
	Paginator[T]

	allItems []T
}

func NewArrayPaginator[T any](items []T, pageSize, page int) *ArrayPaginator[T] {
	paginator := new(ArrayPaginator[T])
	paginator.allItems = items

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	paginator.total = int64(len(items))
	paginator.currentPage = page
	paginator.pageSize = pageSize
	startIndex := int64((page - 1) * pageSize)
	endIndex := startIndex + int64(pageSize)
	if startIndex > paginator.total {
		paginator.items = make([]T, 0)
	} else if endIndex > paginator.total {
		paginator.items = items[startIndex:]
	} else {
		paginator.items = items[startIndex:endIndex]
	}
	return paginator
}

func (p *ArrayPaginator[T]) All() []T {
	return p.allItems
}
