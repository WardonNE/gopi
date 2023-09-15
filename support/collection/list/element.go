package list

type element[V any] struct {
	prev  *element[V]
	next  *element[V]
	list  *LinkedList[V]
	Value V
}

func (e *element[V]) Prev() *element[V] {
	return e.prev
}

func (e *element[V]) Next() *element[V] {
	return e.next
}
