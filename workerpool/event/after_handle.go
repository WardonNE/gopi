package event

import (
	"github.com/wardonne/gopi/workerpool/job"
)

type AfterHandle struct {
	Job job.Interface
}

func NewAfterHandle(job job.Interface) *AfterHandle {
	return &AfterHandle{job}
}

func (event *AfterHandle) Topic() string {
	return AfterHandleTopic
}
