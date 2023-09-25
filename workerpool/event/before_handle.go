package event

import "github.com/wardonne/gopi/workerpool/job"

type BeforeHandle struct {
	Job job.JobInterface
}

func NewBeforeHandle(job job.JobInterface) *BeforeHandle {
	return &BeforeHandle{job}
}

func (event *BeforeHandle) Topic() string {
	return "before-handle"
}
