package event

import "github.com/wardonne/gopi/workerpool/job"

type FailedHandle struct {
	Job   job.Interface
	Error error
}

func NewFailedHandle(job job.Interface, err error) *FailedHandle {
	return &FailedHandle{job, err}
}

func (event *FailedHandle) Topic() string {
	return FailedHandleTopic
}
