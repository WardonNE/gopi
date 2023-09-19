package driver

import (
	"github.com/wardonne/gopi/support/queue"
	"github.com/wardonne/gopi/workerpool"
)

var _ workerpool.IWorkerPoolDriver = (*BlockingQueueDriver)(nil)

type BlockingQueueDriver struct {
	queue *queue.ArrayBlockingQueue[workerpool.Job]
}

func NewBlockingQueueDriver(cap int) *BlockingQueueDriver {
	return &BlockingQueueDriver{
		queue: queue.NewArrayBlockingQueue[workerpool.Job](cap),
	}
}

func (q *BlockingQueueDriver) Clear() {
	q.queue.Clear()
}

func (q *BlockingQueueDriver) Enqueue(job workerpool.Job) bool {
	return q.queue.EnqueueWithBlock(job)
}

func (q *BlockingQueueDriver) Dequeue() (workerpool.Job, bool) {
	return q.queue.DequeueWithBlock()
}
