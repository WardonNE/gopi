package driver

import (
	"github.com/wardonne/gopi/eventbus"
	"github.com/wardonne/gopi/workerpool/job"
	"github.com/wardonne/gopi/workerpool/subscriber"
)

// IDriver workerpool driver interface
type IDriver interface {
	// Count returns the count of pending jobs
	Count() int64
	// IsEmpty returns if the count of pending jobs is zero
	IsEmpty() bool
	// Enqueue pushes a job to queue
	Enqueue(job job.Interface) bool
	// Dequeue pops a job from queue
	Dequeue() (job.Interface, bool)
	// Remove removes a job from queue
	Remove(job job.Interface) bool
	// Ack acks a job
	Ack(job job.Interface) bool
	// Fail handles a failed job
	Fail(job job.Interface)
	// Flush removes all failed jobs
	Flush()
	// Reload reloads all failed jobs into queue
	Reload()
	// Subscribe add a subscriber to queue events
	Subscribe(subscriber subscriber.Interface)
	// DispatchEvent dispatches specifia event
	DispatchEvent(event eventbus.EventInterface)
}

// AbstractDriver abstract driver
type AbstractDriver struct {
	EventBus eventbus.IEventBus
}

// DispatchEvent dispatches specifia event
func (driver *AbstractDriver) DispatchEvent(event eventbus.EventInterface) {
	_ = driver.EventBus.Dispatch(event, nil)
}
