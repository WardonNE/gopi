package queue

import (
	"github.com/wardonne/gopi/support/collection/list"
)

type LinkedQueue[E any] struct {
	items *list.LinkedList[E]
}

func NewLinkedQueue[E any]() *LinkedQueue[E] {
	queue := new(LinkedQueue[E])
	queue.items = list.NewLinkedList[E]()
	return queue
}

func (q *LinkedQueue[E]) MarshalJSON() ([]byte, error) {
	return q.items.MarshalJSON()
}

func (q *LinkedQueue[E]) UnmarshalJSON(data []byte) error {
	return q.items.UnmarshalJSON(data)
}

func (q *LinkedQueue[E]) ToArray() []E {
	return q.items.ToArray()
}

func (q *LinkedQueue[E]) FromArray(values []E) {
	q.items.FromArray(values)
}

func (q *LinkedQueue[E]) Count() int {
	return q.items.Count()
}

func (q *LinkedQueue[E]) IsEmpty() bool {
	return q.items.IsEmpty()
}

func (q *LinkedQueue[E]) IsNotEmpty() bool {
	return q.items.IsNotEmpty()
}

func (q *LinkedQueue[E]) Clear() {
	q.items.Clear()
}

func (q *LinkedQueue[E]) Peek() (value E) {
	if q.items.IsEmpty() {
		return
	}
	return q.items.Get(0)
}

func (q *LinkedQueue[E]) Enqueue(value E) bool {
	q.items.Push(value)
	return true
}

func (q *LinkedQueue[E]) Dequeue() (value E, ok bool) {
	if q.items.IsEmpty() {
		return
	}
	return q.items.Shift(), true
}
