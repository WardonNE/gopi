package queue

import (
	"sync"

	"github.com/wardonne/gopi/support/compare"
)

type SyncPriorityQueue[E any] struct {
	mu    *sync.Mutex
	queue *PriorityQueue[E]
}

func NewSyncPriorityQueue[E any](comparator compare.Comparator[E]) *SyncPriorityQueue[E] {
	queue := new(SyncPriorityQueue[E])
	queue.mu = new(sync.Mutex)
	queue.queue = NewPriorityQueue[E](comparator)
	return queue
}

func (q *SyncPriorityQueue[E]) MarshalJSON() ([]byte, error) {
	q.mu.Lock()
	defer q.mu.Unlock()
	return q.queue.MarshalJSON()
}

func (q *SyncPriorityQueue[E]) UnmarshalJSON(data []byte) error {
	q.mu.Lock()
	defer q.mu.Unlock()
	return q.queue.UnmarshalJSON(data)
}

func (q *SyncPriorityQueue[E]) ToArray() []E {
	q.mu.Lock()
	defer q.mu.Unlock()
	return q.queue.ToArray()
}

func (q *SyncPriorityQueue[E]) FromArray(values []E) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.queue.FromArray(values)
}

func (q *SyncPriorityQueue[E]) Count() int {
	q.mu.Lock()
	defer q.mu.Unlock()
	return q.queue.Count()
}

func (q *SyncPriorityQueue[E]) IsEmpty() bool {
	return q.Count() == 0
}

func (q *SyncPriorityQueue[E]) IsNotEmpty() bool {
	return q.Count() > 0
}

func (q *SyncPriorityQueue[E]) Clear() {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.queue.Clear()
}

func (q *SyncPriorityQueue[E]) Peek() E {
	q.mu.Lock()
	defer q.mu.Unlock()
	return q.queue.Peek()
}

func (q *SyncPriorityQueue[E]) Enqueue(value E) bool {
	q.mu.Lock()
	defer q.mu.Unlock()
	return q.queue.Enqueue(value)
}

func (q *SyncPriorityQueue[E]) Dequeue() (E, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()
	return q.queue.Dequeue()
}
