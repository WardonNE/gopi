package queue

import (
	"encoding/json"
	"sync"
	"time"

	"github.com/wardonne/gopi/support/serializer"
)

type Delayer[E any] interface {
	serializer.JSONSerializer

	Expire() time.Time
	Value() E
}

type DelayQueueEntry[E any] struct {
	expire time.Time
	value  E
}

func NewDelayEntry[E any](value E, expire time.Time) *DelayQueueEntry[E] {
	return &DelayQueueEntry[E]{
		expire: expire,
		value:  value,
	}
}

func (e *DelayQueueEntry[E]) Expire() time.Time {
	return e.expire
}

func (e *DelayQueueEntry[E]) Value() E {
	return e.value
}

func (e *DelayQueueEntry[E]) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"expire": e.expire.Format("2006-01-02 15:04:05"),
		"value":  e.value,
	})
}

func (e *DelayQueueEntry[E]) UnmarshalJSON(data []byte) error {
	var item struct {
		Expire time.Time `json:"expire"`
		Value  E         `json:"value"`
	}
	if err := json.Unmarshal(data, &item); err != nil {
		return err
	}
	e.expire = item.Expire
	e.value = item.Value
	return nil
}

type DelayQueue[T any, E Delayer[T]] struct {
	queue    *PriorityQueue[E]
	lock     *sync.Mutex
	takeLock *sync.Cond
}

func NewDelayQueue[T any, E Delayer[T]]() *DelayQueue[T, E] {
	q := new(DelayQueue[T, E])
	q.queue = NewPriorityQueue[E](q)
	q.lock = new(sync.Mutex)
	q.takeLock = sync.NewCond(q.lock)
	return q
}

func (e *DelayQueue[T, E]) Compare(a, b E) int {
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

func (q *DelayQueue[T, E]) MarshalJSON() ([]byte, error) {
	q.lock.Lock()
	defer q.lock.Unlock()
	return json.Marshal(q.queue)
}

func (q *DelayQueue[T, E]) UnmarshalJSON(data []byte) error {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.queue.UnmarshalJSON(data)
}

func (q *DelayQueue[T, E]) ToArray() []E {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.queue.ToArray()
}

func (q *DelayQueue[T, E]) FromArray(values []E) {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.queue.FromArray(values)
}

func (q *DelayQueue[T, E]) Count() int {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.queue.Count()
}

func (q *DelayQueue[T, E]) IsEmpty() bool {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.queue.IsEmpty()
}

func (q *DelayQueue[T, E]) IsNotEmpty() bool {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.queue.IsNotEmpty()
}

func (q *DelayQueue[T, E]) Clear() {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.queue.Clear()
}

func (q *DelayQueue[T, E]) Peek() E {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.queue.Peek()
}

func (q *DelayQueue[T, E]) Enqueue(value E) (ok bool) {
	q.lock.Lock()
	defer q.lock.Unlock()
	ok = q.queue.Enqueue(value)
	q.takeLock.Broadcast()
	return
}

func (q *DelayQueue[T, E]) Dequeue() (value E, ok bool) {
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

func (q *DelayQueue[T, E]) EnqueueWithBlock(value E) bool {
	return q.Enqueue(value)
}

func (q *DelayQueue[T, E]) DequeueWithBlock() (value E, ok bool) {
	q.lock.Lock()
	defer q.lock.Unlock()
	for q.queue.IsEmpty() {
		q.takeLock.Wait()
	}
	first := q.queue.Peek()
	timer := time.NewTimer(time.Until(first.Expire()))
	<-timer.C
	value, ok = q.queue.Dequeue()
	return
}

func (q *DelayQueue[T, E]) EnqueueWithTimeout(value E, duration time.Duration) bool {
	return q.Enqueue(value)
}

func (q *DelayQueue[T, E]) DequeueWithTimeout(duration time.Duration) (value E, ok bool) {
	timeout := time.After(duration)
	done := make(chan struct{})
	go func() {
		value, ok = q.DequeueWithBlock()
		done <- struct{}{}
	}()
	select {
	case <-timeout:
		return
	case <-done:
		return
	}
}
