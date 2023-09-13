package queue

import (
	"sync"
	"time"
)

type ArrayBlockingQueue[E comparable] struct {
	lock         *sync.Mutex
	items        []E
	size         int
	cap          int
	enqueueIndex int
	dequeueIndex int
	takeLock     *sync.Cond
	putLock      *sync.Cond
}

func NewArrayBlockingQueue[E comparable](cap int) *ArrayBlockingQueue[E] {
	queue := new(ArrayBlockingQueue[E])
	queue.lock = new(sync.Mutex)
	queue.items = make([]E, cap)
	queue.size = 0
	queue.cap = cap
	queue.enqueueIndex = queue.size
	queue.dequeueIndex = 0
	queue.takeLock = sync.NewCond(queue.lock)
	queue.putLock = sync.NewCond(queue.lock)
	return queue
}

func (q *ArrayBlockingQueue[E]) moveIndex(index int) int {
	index++
	if index >= q.cap {
		return 0
	} else {
		return index
	}
}

func (q *ArrayBlockingQueue[E]) Peek() (value E) {
	q.lock.Lock()
	defer q.lock.Unlock()
	if q.size == 0 {
		return
	}
	return q.items[0]
}

func (q *ArrayBlockingQueue[E]) Enqueue(value E) {
	q.lock.Lock()
	defer q.lock.Unlock()
	for q.cap == q.size {
		q.putLock.Wait()
	}
	q.items[q.enqueueIndex] = value
	q.size++
	q.enqueueIndex = q.moveIndex(q.enqueueIndex)
	q.takeLock.Signal()
}

func (q *ArrayBlockingQueue[E]) Dequeue() E {
	q.lock.Lock()
	defer q.lock.Unlock()
	for q.size == 0 {
		q.takeLock.Wait()
	}
	var zero E
	value := q.items[q.dequeueIndex]
	q.items[q.dequeueIndex] = zero
	q.size--
	q.dequeueIndex = q.moveIndex(q.dequeueIndex)
	q.putLock.Signal()
	return value
}

func (q *ArrayBlockingQueue[E]) EnqueueWithTimeout(value E, duration time.Duration) bool {
	timeout := time.After(duration)
	done := make(chan struct{})
	go func() {
		q.Enqueue(value)
		done <- struct{}{}
	}()
	select {
	case <-done:
		return true
	case <-timeout:
		return false
	}
}

func (q *ArrayBlockingQueue[E]) DequeueWithTimeout(duration time.Duration) (value E, ok bool) {
	timeout := time.After(duration)
	done := make(chan struct{})
	go func() {
		value = q.Dequeue()
		ok = true
		done <- struct{}{}
	}()
	select {
	case <-done:
		return
	case <-timeout:
		return
	}
}
