package driver

import (
	"github.com/wardonne/gopi/eventbus"
	"github.com/wardonne/gopi/workerpool/job"
	"github.com/wardonne/gopi/workerpool/subscriber"
)

type DriverInterface interface {
	// Count returns the count of pending jobs
	Count() int
	// IsEmpty returns if the count of pending jobs is zero
	IsEmpty() bool
	// Enqueue pushes a job to queue
	Enqueue(job job.JobInterface) bool
	// Dequeue pops a job from queue
	Dequeue() (job.JobInterface, bool)
	// Remove removes a job from queue
	Remove(job job.JobInterface) bool
	// Ack acks a job
	Ack(job job.JobInterface) bool
	// Fail handles a failed job
	Fail(job job.JobInterface)
	// Flush removes all failed jobs
	Flush()
	// Reload reloads all failed jobs into queue
	Reload()
	// Subscribe add a subscriber to queue events
	Subscribe(subscriber subscriber.Subscriber)
	// DispatchEvent dispatches specifia event
	DispatchEvent(event eventbus.EventInterface)
}

type AbstractDriver struct {
	EventBus eventbus.EventBusInterface
}

// DispatchEvent dispatches specifia event
func (driver *AbstractDriver) DispatchEvent(event eventbus.EventInterface) {
	_ = driver.EventBus.Dispatch(event, nil)
}
