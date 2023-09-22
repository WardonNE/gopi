package workerpool

import "errors"

var (
	ErrWorkerPoolNameExists     = errors.New("WorkerPool name is exists")
	ErrWorkerPoolInstanceExists = errors.New("WorkerPool instance exists")
)

var (
	ErrJobExecuteTimeout  = errors.New("Job lifecycle has ended")
	ErrJobAttemptTimeout  = errors.New("Job execution timeouts")
	ErrJobUnknownErrPanic = errors.New("Job unknown error panic")
)
