package queue

import (
	"github.com/wardonne/gopi/support/collection/list"
)

type ArrayQueue[E any] struct {
	items *list.ArrayList[E]
}

func NewArrayQueue[E any]() *ArrayQueue[E] {
	queue := new(ArrayQueue[E])
	queue.items = list.NewArrayList[E]()
	return queue
}

func (q *ArrayQueue[E]) MarshalJSON() ([]byte, error) {
	return q.items.MarshalJSON()
}

func (q *ArrayQueue[E]) UnmarshalJSON(data []byte) error {
	return q.items.UnmarshalJSON(data)
}

func (q *ArrayQueue[E]) ToArray() []E {
	return q.items.ToArray()
}

func (q *ArrayQueue[E]) FromArray(values []E) {
	q.items.FromArray(values)
}

func (q *ArrayQueue[E]) Count() int {
	return q.items.Count()
}

func (q *ArrayQueue[E]) IsEmpty() bool {
	return q.items.IsEmpty()
}

func (q *ArrayQueue[E]) IsNotEmpty() bool {
	return q.items.IsNotEmpty()
}

func (q *ArrayQueue[E]) Clear() {
	q.items.Clear()
}

func (q *ArrayQueue[E]) Peek() (value E) {
	if q.items.IsEmpty() {
		return
	}
	return q.items.Get(0)
}

func (q *ArrayQueue[E]) Enqueue(value E) bool {
	q.items.Push(value)
	return true
}

func (q *ArrayQueue[E]) Dequeue() (value E, ok bool) {
	if q.items.IsEmpty() {
		return
	}
	return q.items.Shift(), true
}
