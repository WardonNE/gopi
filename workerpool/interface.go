package workerpool

type IWorkerPoolDriver interface {
	Clear()
	Enqueue(value Job) bool
	Dequeue() (Job, bool)
}

type IWorkerPool interface {
	Driver() IWorkerPoolDriver
	Dispatch(job Job) bool
	Running() bool
	Start()
	Stop()
	Release()
}

type Job interface {
	Handle()
}
