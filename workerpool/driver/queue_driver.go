package driver

import (
	"github.com/wardonne/gopi/support/queue"
	"github.com/wardonne/gopi/workerpool"
)

var _ workerpool.IWorkerPoolDriver = (*QueueDriver)(nil)

type QueueDriver struct {
	queue *queue.ArrayQueue[workerpool.Job]
}

func NewQueueDriver() *QueueDriver {
	return &QueueDriver{
		queue: queue.NewArrayQueue[workerpool.Job](),
	}
}

func (q *QueueDriver) Clear() {
	q.queue.Clear()
}

func (q *QueueDriver) Enqueue(job workerpool.Job) bool {
	return q.queue.Enqueue(job)
}

func (q *QueueDriver) Dequeue() (workerpool.Job, bool) {
	return q.queue.Dequeue()
}
