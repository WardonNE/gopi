package queue

import (
	"sync"
)

type SyncLinkedQueue[E any] struct {
	lock  *sync.Mutex
	queue *LinkedQueue[E]
}

func NewSyncLinkedQueue[E any]() *SyncLinkedQueue[E] {
	queue := new(SyncLinkedQueue[E])
	queue.lock = new(sync.Mutex)
	queue.queue = NewLinkedQueue[E]()
	return queue
}

func (q *SyncLinkedQueue[E]) MarshalJSON() ([]byte, error) {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.queue.MarshalJSON()
}

func (q *SyncLinkedQueue[E]) UnmarshalJSON(data []byte) error {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.queue.UnmarshalJSON(data)
}

func (q *SyncLinkedQueue[E]) ToArray() []E {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.queue.ToArray()
}

func (q *SyncLinkedQueue[E]) FromArray(values []E) {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.queue.FromArray(values)
}

func (q *SyncLinkedQueue[E]) Count() int {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.queue.Count()
}

func (q *SyncLinkedQueue[E]) IsEmpty() bool {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.queue.IsEmpty()
}

func (q *SyncLinkedQueue[E]) IsNotEmpty() bool {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.queue.IsNotEmpty()
}

func (q *SyncLinkedQueue[E]) Clear() {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.queue.Clear()
}

func (q *SyncLinkedQueue[E]) Peek() E {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.queue.Peek()
}

func (q *SyncLinkedQueue[E]) Enqueue(value E) bool {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.queue.Enqueue(value)
}

func (q *SyncLinkedQueue[E]) Dequeue() (value E, ok bool) {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.queue.Dequeue()
}
