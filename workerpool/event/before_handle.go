package event

import "github.com/wardonne/gopi/workerpool/job"

type BeforeHandle struct {
	Job job.Interface
}

func NewBeforeHandle(job job.Interface) *BeforeHandle {
	return &BeforeHandle{job}
}

func (event *BeforeHandle) Topic() string {
	return BeforeHandleTopic
}
