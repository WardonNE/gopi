package queue

import "time"

type BlockingQueue[E comparable] interface {
	Queue[E]

	IsFull() bool
	EnqueueWithTimeout(value E, duration time.Duration) bool
	DequeueWithTimeout(duration time.Duration) (E, bool)
}
