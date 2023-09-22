package driver

import (
	"sync/atomic"

	"github.com/wardonne/gopi/support/queue"
	"github.com/wardonne/gopi/workerpool"
)

var _ workerpool.IWorkerPoolDriver = (*BlockingQueueDriver)(nil)

type BlockingQueueDriver struct {
	queue     *queue.ArrayBlockingQueue[workerpool.IJob]
	total     int64
	pending   int64
	executing int64
	completed int64
	success   int64
	failed    int64
}

func NewBlockingQueueDriver(cap int) *BlockingQueueDriver {
	return &BlockingQueueDriver{
		queue: queue.NewArrayBlockingQueue[workerpool.IJob](cap),
	}
}

func (q *BlockingQueueDriver) Clear() {
	q.queue.Clear()
}

func (q *BlockingQueueDriver) Total() int64 {
	return q.total
}

func (q *BlockingQueueDriver) PendingCount() int64 {
	return q.pending
}

func (q *BlockingQueueDriver) ExecutingCount() int64 {
	return q.executing
}

func (q *BlockingQueueDriver) CompletedCount() int64 {
	return q.completed
}

func (q *BlockingQueueDriver) SuccessCount() int64 {
	return q.success
}

func (q *BlockingQueueDriver) FailedCount() int64 {
	return q.failed
}

func (q *BlockingQueueDriver) Complete(success bool) {
	atomic.AddInt64(&q.completed, 1)
	atomic.AddInt64(&q.executing, -1)
	if success {
		atomic.AddInt64(&q.success, 1)
	} else {
		atomic.AddInt64(&q.failed, 1)
	}
}

func (q *BlockingQueueDriver) IsEmpty() bool {
	return q.queue.IsEmpty()
}

func (q *BlockingQueueDriver) IsNotEmpty() bool {
	return q.queue.IsNotEmpty()
}

func (q *BlockingQueueDriver) Enqueue(job workerpool.IJob) (ok bool) {
	if ok = q.queue.Enqueue(job); ok {
		atomic.AddInt64(&q.total, 1)
		atomic.AddInt64(&q.pending, 1)
	}
	return
}

func (q *BlockingQueueDriver) Dequeue() (job workerpool.IJob, ok bool) {
	if job, ok = q.queue.Dequeue(); ok {
		atomic.AddInt64(&q.pending, -1)
		atomic.AddInt64(&q.executing, 1)
	}
	return
}
