package event

import "github.com/wardonne/gopi/workerpool/job"

type FailedHandle struct {
	Job   job.JobInterface
	Error error
}

func NewFailedHandle(job job.JobInterface, err error) *FailedHandle {
	return &FailedHandle{job, err}
}

func (event *FailedHandle) Topic() string {
	return FailedHandleTopic
}
