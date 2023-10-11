package queue

import (
	"encoding/json"
	"sync"
	"time"

	"github.com/wardonne/gopi/support/serializer"
)

type Delayed[T any] interface {
	serializer.JSONSerializer

	Expire() time.Time
	Value() T
}

type DelayQueueEntry[T any] struct {
	expire time.Time
	value  T
}

func NewDelayEntry[T any](value T, delay time.Duration) *DelayQueueEntry[T] {
	return &DelayQueueEntry[T]{
		expire: time.Now().Add(delay),
		value:  value,
	}
}

func (e *DelayQueueEntry[T]) Expire() time.Time {
	return e.expire
}

func (e *DelayQueueEntry[T]) Value() T {
	return e.value
}

func (e *DelayQueueEntry[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"expire": e.expire.String(),
		"value":  e.value,
	})
}

func (e *DelayQueueEntry[T]) UnmarshalJSON(data []byte) error {
	var item struct {
		Expire time.Time `json:"expire"`
		Value  T         `json:"value"`
	}
	if err := json.Unmarshal(data, &item); err != nil {
		return err
	}
	e.expire = item.Expire
	e.value = item.Value
	return nil
}

type DelayQueue[T any] struct {
	queue    *PriorityQueue[Delayed[T]]
	lock     *sync.Mutex
	takeLock *sync.Cond
}

func NewDelayQueue[T any]() *DelayQueue[T] {
	q := new(DelayQueue[T])
	q.queue = NewPriorityQueue[Delayed[T]](q)
	q.lock = new(sync.Mutex)
	q.takeLock = sync.NewCond(q.lock)
	return q
}

func (e *DelayQueue[T]) Compare(a, b Delayed[T]) int {
	expireA := a.Expire()
	expireB := b.Expire()
	if expireA.Before(expireB) {
		return -1
	} else if expireA.After(expireB) {
		return 1
	} else {
		return 0
	}
}

func (q *DelayQueue[T]) MarshalJSON() ([]byte, error) {
	q.lock.Lock()
	defer q.lock.Unlock()
	return json.Marshal(q.queue)
}

func (q *DelayQueue[T]) UnmarshalJSON(data []byte) error {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.queue.UnmarshalJSON(data)
}

func (q *DelayQueue[T]) ToArray() []Delayed[T] {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.queue.ToArray()
}

func (q *DelayQueue[T]) FromArray(values []Delayed[T]) {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.queue.FromArray(values)
}

func (q *DelayQueue[T]) Count() int {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.queue.Count()
}

func (q *DelayQueue[T]) IsEmpty() bool {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.queue.IsEmpty()
}

func (q *DelayQueue[T]) IsNotEmpty() bool {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.queue.IsNotEmpty()
}

func (q *DelayQueue[T]) Clear() {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.queue.Clear()
}

func (q *DelayQueue[T]) Peek() Delayed[T] {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.queue.Peek()
}

func (q *DelayQueue[T]) Enqueue(value Delayed[T]) (ok bool) {
	q.lock.Lock()
	defer q.lock.Unlock()
	ok = q.queue.Enqueue(value)
	q.takeLock.Broadcast()
	return
}

func (q *DelayQueue[T]) Dequeue() (value Delayed[T], ok bool) {
	q.lock.Lock()
	defer q.lock.Unlock()
	if q.queue.IsEmpty() {
		return
	}
	if q.queue.Peek().Expire().After(time.Now()) {
		return
	}
	return q.queue.Dequeue()
}

func (q *DelayQueue[T]) EnqueueWithBlock(value Delayed[T]) bool {
	return q.Enqueue(value)
}

func (q *DelayQueue[T]) DequeueWithBlock() (value Delayed[T], ok bool) {
	q.lock.Lock()
	defer q.lock.Unlock()
	for q.queue.IsEmpty() {
		q.takeLock.Wait()
	}
	first := q.queue.Peek()
	timer := time.NewTimer(time.Until(first.Expire()))
	defer timer.Stop()
	<-timer.C
	value, ok = q.queue.Dequeue()
	return
}

func (q *DelayQueue[T]) EnqueueWithTimeout(value Delayed[T], duration time.Duration) bool {
	return q.Enqueue(value)
}

func (q *DelayQueue[T]) DequeueWithTimeout(duration time.Duration) (value Delayed[T], ok bool) {
	timeout := time.After(duration)
	done := make(chan struct{})
	go func() {
		q.lock.Lock()
		defer q.lock.Unlock()
		for q.queue.IsEmpty() {
			q.takeLock.Wait()
		}
		first := q.queue.Peek()
		timer := time.NewTimer(time.Until(first.Expire()))
		defer timer.Stop()
		<-timer.C
		close(done)
	}()
	select {
	case <-timeout:
		return
	case <-done:
		return q.queue.Dequeue()
	}
}
