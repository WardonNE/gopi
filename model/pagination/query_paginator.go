package pagination

import (
	"github.com/wardonne/gopi/pagination"
	"gorm.io/gorm"
)

type QueryPaginator[T any] struct {
	pagination.LazyPaginator[T]

	db *gorm.DB
}

func NewQueryPaginator[T any](db *gorm.DB, pageSize, page int) *QueryPaginator[T] {
	paginator := new(QueryPaginator[T])
	paginator.LazyPaginator = *pagination.NewLazyPaginator[T](paginator.fetchTotal, paginator.fetchItems, pageSize, page)
	paginator.db = db
	return paginator
}

func (q *QueryPaginator[T]) fetchTotal() int64 {
	var total int64
	query := q.db.Session(new(gorm.Session))
	if err := query.Count(&total).Error; err != nil {
		panic(err)
	}
	return total
}

func (q *QueryPaginator[T]) fetchItems() []T {
	var items []T
	query := q.db.Session(new(gorm.Session)).Limit(q.PageSize()).Offset(q.FirstItemIndex() - 1)
	if err := query.Find(&items).Error; err != nil {
		panic(err)
	}
	return items
}
