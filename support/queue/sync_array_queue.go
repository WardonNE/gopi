package queue

import "sync"

type SyncArrayQueue[E any] struct {
	lock  *sync.Mutex
	queue *ArrayQueue[E]
}

func NewSyncArrayQueue[E any]() *SyncArrayQueue[E] {
	queue := new(SyncArrayQueue[E])
	queue.lock = new(sync.Mutex)
	queue.queue = NewArrayQueue[E]()
	return queue
}

func (q *SyncArrayQueue[E]) MarshalJSON() ([]byte, error) {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.queue.MarshalJSON()
}

func (q *SyncArrayQueue[E]) UnmarshalJSON(data []byte) error {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.queue.UnmarshalJSON(data)
}

func (q *SyncArrayQueue[E]) ToArray() []E {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.queue.ToArray()
}

func (q *SyncArrayQueue[E]) FromArray(values []E) {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.queue.FromArray(values)
}

func (q *SyncArrayQueue[E]) Count() int {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.queue.Count()
}

func (q *SyncArrayQueue[E]) IsEmpty() bool {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.queue.IsEmpty()
}

func (q *SyncArrayQueue[E]) IsNotEmpty() bool {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.queue.IsNotEmpty()
}

func (q *SyncArrayQueue[E]) Clear() {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.queue.Clear()
}

func (q *SyncArrayQueue[E]) Peek() (value E) {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.queue.Peek()
}

func (q *SyncArrayQueue[E]) Enqueue(value E) bool {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.queue.Enqueue(value)
}

func (q *SyncArrayQueue[E]) Dequeue() (value E, ok bool) {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.queue.Dequeue()
}
