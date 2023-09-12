package pagination

type IPaginator[T any] interface {
	Items() []T
	Total() int64
	FirstItem() int
	LastItem() int
	CurrentPage() int
	PageSize() int
	HasMore() bool
	LastPage() int
	ToMap() map[string]any
}

type Paginator[T any] struct {
	items       []T
	total       int64
	currentPage int
	pageSize    int
	lastPage    int
}

func NewPaginator[T any](total int64, items []T, pageSize, page int) *Paginator[T] {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	paginator := new(Paginator[T])
	paginator.items = items
	paginator.total = total
	paginator.currentPage = page
	paginator.pageSize = pageSize
	paginator.lastPage = int((total / int64(pageSize)) + 1)
	return paginator
}

func (p *Paginator[T]) Items() []T {
	return p.items
}

func (p *Paginator[T]) FirstItem() int {
	return (p.currentPage-1)*p.pageSize + 1
}

func (p *Paginator[T]) LastItem() int {
	return p.FirstItem() + len(p.items)
}

func (p *Paginator[T]) Total() int64 {
	return p.total
}

func (p *Paginator[T]) CurrentPage() int {
	return p.currentPage
}

func (p *Paginator[T]) PageSize() int {
	return p.pageSize
}

func (p *Paginator[T]) HasMore() bool {
	return p.lastPage > p.currentPage
}

func (p *Paginator[T]) LastPage() int {
	return p.lastPage
}

func (p *Paginator[T]) ToMap() map[string]any {
	return map[string]any{
		"items":        p.Items(),
		"total":        p.Total(),
		"current_page": p.CurrentPage(),
		"page_size":    p.PageSize(),
		"last_page":    p.LastPage(),
		"from":         p.FirstItem(),
		"to":           p.LastItem(),
	}
}
