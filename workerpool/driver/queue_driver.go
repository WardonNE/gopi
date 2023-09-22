package driver

import (
	"sync/atomic"

	"github.com/wardonne/gopi/support/queue"
	"github.com/wardonne/gopi/workerpool"
)

var _ workerpool.IWorkerPoolDriver = (*QueueDriver)(nil)

type QueueDriver struct {
	queue     *queue.SyncArrayQueue[workerpool.IJob]
	total     int64
	pending   int64
	executing int64
	completed int64
	success   int64
	failed    int64
}

func NewQueueDriver() *QueueDriver {
	return &QueueDriver{
		queue: queue.NewSyncArrayQueue[workerpool.IJob](),
	}
}

func (q *QueueDriver) Clear() {
	q.queue.Clear()
}

func (q *QueueDriver) Total() int64 {
	return q.total
}

func (q *QueueDriver) PendingCount() int64 {
	return q.pending
}

func (q *QueueDriver) ExecutingCount() int64 {
	return q.executing
}

func (q *QueueDriver) CompletedCount() int64 {
	return q.completed
}

func (q *QueueDriver) SuccessCount() int64 {
	return q.success
}

func (q *QueueDriver) FailedCount() int64 {
	return q.failed
}

func (q *QueueDriver) Complete(success bool) {
	atomic.AddInt64(&q.completed, 1)
	atomic.AddInt64(&q.executing, -1)
	if success {
		atomic.AddInt64(&q.success, 1)
	} else {
		atomic.AddInt64(&q.failed, 1)
	}
}

func (q *QueueDriver) IsEmpty() bool {
	return q.queue.IsEmpty()
}

func (q *QueueDriver) IsNotEmpty() bool {
	return q.queue.IsNotEmpty()
}

func (q *QueueDriver) Enqueue(job workerpool.IJob) (ok bool) {
	if ok = q.queue.Enqueue(job); ok {
		atomic.AddInt64(&q.total, 1)
		atomic.AddInt64(&q.pending, 1)
	}
	return
}

func (q *QueueDriver) Dequeue() (job workerpool.IJob, ok bool) {
	if job, ok = q.queue.Dequeue(); ok {
		atomic.AddInt64(&q.pending, -1)
		atomic.AddInt64(&q.executing, 1)
	}
	return
}
