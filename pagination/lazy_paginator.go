package pagination

import (
	"sync"

	"github.com/wardonne/gopi/utils"
)

type LazyPaginator[T any] struct {
	Paginator[T]

	loader func()
	once   *sync.Once
}

func NewLazyPaginator[T any](totalLoader func() int64, itemsLoader func() []T, pageSize, page int) *LazyPaginator[T] {
	paginator := new(LazyPaginator[T])
	paginator.currentPage = utils.If(page <= 0, 1, page)
	paginator.pageSize = utils.If(pageSize <= 0, 10, pageSize)
	paginator.loader = func() {
		paginator.total = totalLoader()
		paginator.items = itemsLoader()
		paginator.lastPage = int(paginator.total/int64(pageSize)) + 1
	}
	paginator.once = new(sync.Once)
	return paginator
}

func (p *LazyPaginator[T]) load() {
	p.once.Do(p.loader)
}

func (p *LazyPaginator[T]) Items() []T {
	p.load()
	return p.Paginator.Items()
}

func (p *LazyPaginator[T]) LastItemIndex() int {
	p.load()
	return p.Paginator.LastItem()
}

func (p *LazyPaginator[T]) Total() int64 {
	p.load()
	return p.Paginator.Total()
}

func (p *LazyPaginator[T]) CurrentPage() int {
	return p.Paginator.CurrentPage()
}

func (p *LazyPaginator[T]) PageSize() int {
	return p.Paginator.PageSize()
}

func (p *LazyPaginator[T]) HasMore() bool {
	return p.CurrentPage() < p.LastPage()
}

func (p *LazyPaginator[T]) LastPage() int {
	p.load()
	return p.Paginator.LastPage()
}

func (p *LazyPaginator[T]) ToMap() map[string]any {
	return map[string]any{
		"items":        p.Items(),
		"total":        p.Total(),
		"current_page": p.CurrentPage(),
		"page_size":    p.PageSize(),
		"last_page":    p.LastPage(),
		"from":         p.FirstItemIndex(),
		"to":           p.LastItemIndex(),
	}
}
