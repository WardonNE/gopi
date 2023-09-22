package workerpool

import "errors"

var (
	ErrJobExecuteTimeout  = errors.New("Job lifecycle has ended")
	ErrJobAttemptTimeout  = errors.New("Job execution timeouts")
	ErrJobUnknownErrPanic = errors.New("Job unknown error panic")
)
