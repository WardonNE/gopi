package queue

import (
	"encoding/json"
	"sync"
	"time"

	"github.com/wardonne/gopi/support/collection/list"
)

type LinkedBlockingQueue[E any] struct {
	lock     *sync.Mutex
	items    *list.LinkedList[E]
	cap      int
	takeLock *sync.Cond
	putLock  *sync.Cond
}

func NewLinkedBlockingQueue[E any](cap int) *LinkedBlockingQueue[E] {
	queue := new(LinkedBlockingQueue[E])
	queue.lock = new(sync.Mutex)
	queue.items = list.NewLinkedList[E]()
	queue.cap = cap
	queue.takeLock = sync.NewCond(queue.lock)
	queue.putLock = sync.NewCond(queue.lock)
	return queue
}

func (q *LinkedBlockingQueue[E]) MarshalJSON() ([]byte, error) {
	return q.items.MarshalJSON()
}

func (q *LinkedBlockingQueue[E]) UnmarshalJSON(data []byte) error {
	q.lock.Lock()
	defer q.lock.Unlock()
	values := make([]E, 0)
	if err := json.Unmarshal(data, &values); err != nil {
		return err
	}
	for _, value := range values {
		for q.items.Count() == q.cap {
			q.putLock.Wait()
		}
		q.items.Push(value)
		q.takeLock.Broadcast()
	}
	return nil
}

func (q *LinkedBlockingQueue[E]) ToArray() []E {
	return q.items.ToArray()
}

func (q *LinkedBlockingQueue[E]) FromArray(values []E) {
	q.lock.Lock()
	defer q.lock.Unlock()
	for _, value := range values {
		for q.items.Count() == q.cap {
			q.putLock.Wait()
		}
		q.items.Push(value)
		q.takeLock.Broadcast()
	}
}

func (q *LinkedBlockingQueue[E]) Count() int {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.items.Count()
}

func (q *LinkedBlockingQueue[E]) IsEmpty() bool {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.items.IsEmpty()
}

func (q *LinkedBlockingQueue[E]) IsNotEmpty() bool {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.items.IsNotEmpty()
}

func (q *LinkedBlockingQueue[E]) Clear() {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.items.Clear()
}

func (q *LinkedBlockingQueue[E]) Peek() (value E) {
	q.lock.Lock()
	defer q.lock.Unlock()
	if q.items.IsEmpty() {
		return
	}
	return q.items.Get(0)
}

func (q *LinkedBlockingQueue[E]) Enqueue(value E) bool {
	q.lock.Lock()
	defer q.lock.Unlock()
	if q.cap == q.items.Count() {
		return false
	}
	q.items.Push(value)
	return true
}

func (q *LinkedBlockingQueue[E]) Dequeue() (value E, ok bool) {
	q.lock.Lock()
	defer q.lock.Unlock()
	if q.items.IsEmpty() {
		return
	}
	return q.items.Shift(), true
}

func (q *LinkedBlockingQueue[E]) EnqueueWithBlock(value E) bool {
	q.lock.Lock()
	defer q.lock.Unlock()
	for q.cap == q.items.Count() {
		q.putLock.Wait()
	}
	q.items.Push(value)
	q.takeLock.Broadcast()
	return true
}

func (q *LinkedBlockingQueue[E]) DequeueWithBlock() (value E, ok bool) {
	q.lock.Lock()
	defer q.lock.Unlock()
	for q.items.IsEmpty() {
		q.takeLock.Wait()
	}
	value = q.items.Shift()
	q.putLock.Broadcast()
	return value, true
}

func (q *LinkedBlockingQueue[E]) EnqueueWithTimeout(value E, duration time.Duration) bool {
	timeout := time.After(duration)
	done := make(chan struct{})
	go func() {
		q.EnqueueWithBlock(value)
		done <- struct{}{}
	}()
	select {
	case <-done:
		return true
	case <-timeout:
		return false
	}
}

func (q *LinkedBlockingQueue[E]) DequeueWithTimeout(duration time.Duration) (value E, ok bool) {
	timeout := time.After(duration)
	done := make(chan struct{})
	go func() {
		value, ok = q.DequeueWithBlock()
		done <- struct{}{}
	}()
	select {
	case <-done:
		return
	case <-timeout:
		return
	}
}
