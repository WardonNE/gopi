package queue

import (
	"encoding/json"
	"sync"
	"time"

	"github.com/wardonne/gopi/support/compare"
)

type PriorityBlockingQueue[E any] struct {
	lock     *sync.Mutex
	queue    *PriorityQueue[E]
	cap      int
	takeLock *sync.Cond
	putLock  *sync.Cond
}

func NewPriorityBlockingQueue[E any](cap int, comparator compare.Comparator[E]) *PriorityBlockingQueue[E] {
	queue := new(PriorityBlockingQueue[E])
	queue.lock = new(sync.Mutex)
	queue.queue = NewPriorityQueue[E](comparator)
	queue.cap = cap
	queue.takeLock = sync.NewCond(queue.lock)
	queue.putLock = sync.NewCond(queue.lock)
	return queue
}

func (q *PriorityBlockingQueue[E]) MarshalJSON() ([]byte, error) {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.queue.MarshalJSON()
}

func (q *PriorityBlockingQueue[E]) UnmarshalJSON(data []byte) error {
	q.lock.Lock()
	defer q.lock.Unlock()
	values := make([]E, 0)
	if err := json.Unmarshal(data, &values); err != nil {
		return err
	}
	q.queue.Clear()
	for _, value := range values {
		for q.cap == q.queue.size {
			q.putLock.Wait()
		}
		q.queue.Enqueue(value)
		q.takeLock.Broadcast()
	}
	return nil
}

func (q *PriorityBlockingQueue[E]) ToArray() []E {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.queue.ToArray()
}

func (q *PriorityBlockingQueue[E]) FromArray(values []E) {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.queue.Clear()
	for _, value := range values {
		for q.cap == q.queue.size {
			q.putLock.Wait()
		}
		q.queue.Enqueue(value)
		q.takeLock.Broadcast()
	}
}

func (q *PriorityBlockingQueue[E]) Count() int {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.queue.Count()
}

func (q *PriorityBlockingQueue[E]) IsEmpty() bool {
	return q.Count() == 0
}

func (q *PriorityBlockingQueue[E]) IsNotEmpty() bool {
	return q.Count() > 0
}

func (q *PriorityBlockingQueue[E]) Clear() {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.queue.Clear()
}

func (q *PriorityBlockingQueue[E]) Peek() (value E) {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.queue.Peek()
}

func (q *PriorityBlockingQueue[E]) Enqueue(value E) bool {
	q.lock.Lock()
	defer q.lock.Unlock()
	if q.cap == q.queue.size {
		return false
	}
	return q.queue.Enqueue(value)
}

func (q *PriorityBlockingQueue[E]) Dequeue() (value E, ok bool) {
	q.lock.Lock()
	defer q.lock.Unlock()
	if q.queue.IsEmpty() {
		return
	}
	return q.queue.Dequeue()
}

func (q *PriorityBlockingQueue[E]) EnqueueWithBlock(value E) (ok bool) {
	q.lock.Lock()
	defer q.lock.Unlock()
	for q.cap == q.queue.size {
		q.putLock.Wait()
	}
	ok = q.queue.Enqueue(value)
	q.takeLock.Broadcast()
	return
}

func (q *PriorityBlockingQueue[E]) DequeueWithBlock() (value E, ok bool) {
	q.lock.Lock()
	defer q.lock.Unlock()
	for q.queue.IsEmpty() {
		q.takeLock.Wait()
	}
	value, ok = q.queue.Dequeue()
	q.putLock.Broadcast()
	return
}

func (q *PriorityBlockingQueue[E]) EnqueueWithTimeout(value E, duration time.Duration) bool {
	timeout := time.After(duration)
	done := make(chan struct{})
	go func() {
		for q.cap == q.queue.Count() {
			q.putLock.Wait()
		}
		close(done)
	}()
	select {
	case <-done:
		ok := q.queue.Enqueue(value)
		q.takeLock.Broadcast()
		return ok
	case <-timeout:
		return false
	}
}

func (q *PriorityBlockingQueue[E]) DequeueWithTimeout(duration time.Duration) (value E, ok bool) {
	timeout := time.After(duration)
	done := make(chan struct{})
	go func() {
		for q.queue.IsEmpty() {
			q.takeLock.Wait()
		}
		close(done)
	}()
	select {
	case <-done:
		value, ok = q.queue.Dequeue()
		q.putLock.Broadcast()
		return
	case <-timeout:
		return
	}
}
