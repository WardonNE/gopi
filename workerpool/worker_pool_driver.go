package workerpool

type IWorkerPoolDriver interface {
	Clear()

	Total() int64
	PendingCount() int64
	ExecutingCount() int64
	CompletedCount() int64
	SuccessCount() int64
	FailedCount() int64

	Complete(success bool)

	Enqueue(value IJob) bool
	Dequeue() (IJob, bool)
}
