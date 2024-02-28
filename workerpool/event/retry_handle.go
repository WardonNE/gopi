package event

import "github.com/wardonne/gopi/workerpool/job"

type RetryHandle struct {
	Job      job.Interface
	Attempts int
	Error    error
}

func NewRetryHandle(job job.Interface, attempts int, err error) *RetryHandle {
	return &RetryHandle{job, attempts, err}
}

func (event *RetryHandle) Topic() string {
	return RetryHandleTopic
}
