package event

import (
	"github.com/wardonne/gopi/workerpool/job"
)

type AfterHandle struct {
	Job job.JobInterface
}

func NewAfterHandle(job job.JobInterface) *AfterHandle {
	return &AfterHandle{job}
}

func (event *AfterHandle) Topic() string {
	return "after-handle"
}
