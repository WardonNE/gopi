package queue

import "time"

type SynchronousQueue[E any] struct {
	channel chan E
}

func NewSynchronousQueue[E any](ch ...chan E) *SynchronousQueue[E] {
	queue := new(SynchronousQueue[E])
	if len(ch) > 0 {
		queue.channel = ch[0]
	} else {
		queue.channel = make(chan E)
	}
	return queue
}

func (q *SynchronousQueue[E]) MarshalJSON() ([]byte, error) {
	return []byte{}, nil

}

func (q *SynchronousQueue[E]) UnmarshalJSON(data []byte) error {
	return nil
}

func (q *SynchronousQueue[E]) ToArray() []E {
	return make([]E, 0)
}

func (q *SynchronousQueue[E]) FromArray(values []E) {
	return
}

func (q *SynchronousQueue[E]) Count() int {
	return 0
}

func (q *SynchronousQueue[E]) IsEmpty() bool {
	return true
}

func (q *SynchronousQueue[E]) IsNotEmpty() bool {
	return false
}

func (q *SynchronousQueue[E]) Clear() {
	return
}

func (q *SynchronousQueue[E]) Peek() (value E) {
	return
}

func (q *SynchronousQueue[E]) Enqueue(value E) bool {
	q.channel <- value
	return true
}

func (q *SynchronousQueue[E]) Dequeue() (value E, ok bool) {
	value = <-q.channel
	return value, true
}

func (q *SynchronousQueue[E]) EnqueueWithBlock(value E) bool {
	return q.Enqueue(value)
}

func (q *SynchronousQueue[E]) DequeueWithBlock() (value E, ok bool) {
	return q.Dequeue()
}

func (q *SynchronousQueue[E]) EnqueueWithTimeout(value E, duration time.Duration) bool {
	timeout := time.After(duration)
	select {
	case <-timeout:
		return false
	case q.channel <- value:
		return true
	}
}

func (q *SynchronousQueue[E]) DequeueWithTimeout(duration time.Duration) (value E, ok bool) {
	timeout := time.After(duration)
	select {
	case <-timeout:
		return
	case <-q.channel:
		return
	}
}
