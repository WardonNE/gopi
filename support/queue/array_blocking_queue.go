package queue

import (
	"encoding/json"
	"sync"
	"time"

	"github.com/wardonne/gopi/support/collection/list"
)

type ArrayBlockingQueue[E any] struct {
	lock         *sync.Mutex
	items        *list.ArrayList[E]
	size         int
	cap          int
	enqueueIndex int
	dequeueIndex int
	takeLock     *sync.Cond
	putLock      *sync.Cond
}

func NewArrayBlockingQueue[E any](cap int) *ArrayBlockingQueue[E] {
	queue := new(ArrayBlockingQueue[E])
	queue.lock = new(sync.Mutex)
	queue.items = list.NewArrayList[E](make([]E, cap)...)
	queue.size = 0
	queue.cap = cap
	queue.enqueueIndex = 0
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

func (q *ArrayBlockingQueue[E]) MarshalJSON() ([]byte, error) {
	q.lock.Lock()
	defer q.lock.Unlock()
	values := make([]E, 0)
	if q.enqueueIndex > q.dequeueIndex {
		for index := q.dequeueIndex; index < q.enqueueIndex; index++ {
			values = append(values, q.items.Get(index))
		}
	} else {
		for index := q.dequeueIndex; index < q.cap; index++ {
			values = append(values, q.items.Get(index))
		}
		for index := 0; index < q.enqueueIndex; index++ {
			values = append(values, q.items.Get(index))
		}
	}
	return json.Marshal(values)
}

func (q *ArrayBlockingQueue[E]) UnmarshalJSON(data []byte) error {
	q.lock.Lock()
	defer q.lock.Unlock()
	values := make([]E, 0)
	if err := json.Unmarshal(data, &values); err != nil {
		return err
	}
	for _, value := range values {
		for q.size == q.cap {
			q.putLock.Wait()
		}
		q.items.Set(q.enqueueIndex, value)
		q.size++
		q.enqueueIndex = q.moveIndex(q.enqueueIndex)
		q.takeLock.Broadcast()
	}
	return nil
}

func (q *ArrayBlockingQueue[E]) ToArray() []E {
	q.lock.Lock()
	defer q.lock.Unlock()
	values := []E{}
	if q.enqueueIndex > q.dequeueIndex {
		for index := q.dequeueIndex; index < q.enqueueIndex; index++ {
			values = append(values, q.items.Get(index))
		}
	} else {
		for index := q.dequeueIndex; index < q.cap; index++ {
			values = append(values, q.items.Get(index))
		}
		for index := 0; index < q.enqueueIndex; index++ {
			values = append(values, q.items.Get(index))
		}
	}
	return values
}

func (q *ArrayBlockingQueue[E]) FromArray(values []E) {
	q.lock.Lock()
	defer q.lock.Unlock()
	for _, value := range values {
		for q.size == q.cap {
			q.putLock.Wait()
		}
		q.items.Set(q.enqueueIndex, value)
		q.size++
		q.enqueueIndex = q.moveIndex(q.enqueueIndex)
		q.takeLock.Broadcast()
	}
}

func (q *ArrayBlockingQueue[E]) Count() int {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.size
}

func (q *ArrayBlockingQueue[E]) IsEmpty() bool {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.size == 0
}

func (q *ArrayBlockingQueue[E]) IsNotEmpty() bool {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.size > 0
}

func (q *ArrayBlockingQueue[E]) Clear() {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.items.Clear()
	q.size = 0
	q.dequeueIndex = 0
	q.enqueueIndex = 0
}

func (q *ArrayBlockingQueue[E]) Peek() (value E) {
	q.lock.Lock()
	defer q.lock.Unlock()
	if q.size == 0 {
		return
	}
	return q.items.Get(0)
}

func (q *ArrayBlockingQueue[E]) Enqueue(value E) bool {
	q.lock.Lock()
	defer q.lock.Unlock()
	if q.cap == q.size {
		return false
	}
	q.items.Set(q.enqueueIndex, value)
	q.size++
	q.enqueueIndex = q.moveIndex(q.enqueueIndex)
	q.takeLock.Broadcast()
	return true
}

func (q *ArrayBlockingQueue[E]) Dequeue() (value E, ok bool) {
	q.lock.Lock()
	defer q.lock.Unlock()
	if q.size == 0 {
		return
	}
	value, ok = q.items.Get(q.dequeueIndex), true
	var zero E
	q.items.Set(q.dequeueIndex, zero)
	q.dequeueIndex = q.moveIndex(q.dequeueIndex)
	q.size--
	q.putLock.Broadcast()
	return
}

func (q *ArrayBlockingQueue[E]) EnqueueWithBlock(value E) bool {
	q.lock.Lock()
	defer q.lock.Unlock()
	for q.cap == q.size {
		q.putLock.Wait()
	}
	q.items.Set(q.enqueueIndex, value)
	q.size++
	q.enqueueIndex = q.moveIndex(q.enqueueIndex)
	q.takeLock.Broadcast()
	return true
}

func (q *ArrayBlockingQueue[E]) DequeueWithBlock() (value E, ok bool) {
	q.lock.Lock()
	defer q.lock.Unlock()
	for q.size == 0 {
		q.takeLock.Wait()
	}
	var zero E
	value = q.items.Get(q.dequeueIndex)
	q.items.Set(q.dequeueIndex, zero)
	q.size--
	q.dequeueIndex = q.moveIndex(q.dequeueIndex)
	q.putLock.Broadcast()
	return value, true
}

func (q *ArrayBlockingQueue[E]) EnqueueWithTimeout(value E, duration time.Duration) bool {
	timeout := time.After(duration)
	done := make(chan struct{})
	go func() {
		q.lock.Lock()
		defer q.lock.Unlock()
		for q.cap == q.size {
			q.putLock.Wait()
		}
		close(done)
	}()
	select {
	case <-done:
		q.items.Set(q.enqueueIndex, value)
		q.size++
		q.enqueueIndex = q.moveIndex(q.enqueueIndex)
		q.takeLock.Broadcast()
		return true
	case <-timeout:
		return false
	}
}

func (q *ArrayBlockingQueue[E]) DequeueWithTimeout(duration time.Duration) (value E, ok bool) {
	timeout := time.After(duration)
	done := make(chan struct{})
	go func() {
		q.lock.Lock()
		defer q.lock.Unlock()
		for q.size == 0 {
			q.takeLock.Wait()
		}
		close(done)
	}()
	select {
	case <-done:
		var zero E
		value = q.items.Get(q.dequeueIndex)
		q.items.Set(q.dequeueIndex, zero)
		q.size--
		q.dequeueIndex = q.moveIndex(q.dequeueIndex)
		q.putLock.Broadcast()
		return value, true
	case <-timeout:
		return
	}
}
