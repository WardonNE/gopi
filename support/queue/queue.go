package queue

import (
	"time"

	"github.com/wardonne/gopi/support"
	"github.com/wardonne/gopi/support/serializer"
)

type Queue[E any] interface {
	serializer.JSONSerializer
	serializer.ArraySerializer[E]
	support.Countable

	IsEmpty() bool
	IsNotEmpty() bool
	Clear()
	Peek() E
	Enqueue(value E) bool
	Dequeue() (value E, ok bool)
}

type BlockingQueue[E any] interface {
	Queue[E]

	EnqueueWithBlock(value E) bool
	DequeueWithBlock() (E, bool)
	EnqueueWithTimeout(value E, duration time.Duration) bool
	DequeueWithTimeout(duration time.Duration) (E, bool)
}

func implements[E any]() {
	var _ Queue[E] = (*ArrayQueue[E])(nil)
	var _ Queue[E] = (*LinkedQueue[E])(nil)
	var _ Queue[E] = (*SyncArrayQueue[E])(nil)
	var _ Queue[E] = (*SyncLinkedQueue[E])(nil)
	var _ Queue[E] = (*PriorityQueue[E])(nil)
	var _ Queue[E] = (*SyncPriorityQueue[E])(nil)

	var _ BlockingQueue[E] = (*ArrayBlockingQueue[E])(nil)
	var _ BlockingQueue[E] = (*LinkedBlockingQueue[E])(nil)
	var _ BlockingQueue[E] = (*PriorityBlockingQueue[E])(nil)
}
