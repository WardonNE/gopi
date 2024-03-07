package driver

import (
	"github.com/wardonne/gopi/eventbus"
	"github.com/wardonne/gopi/support/collection/list"
	"github.com/wardonne/gopi/workerpool/event"
	"github.com/wardonne/gopi/workerpool/job"
	"github.com/wardonne/gopi/workerpool/subscriber"
)

var _ IDriver = (*MemoryDriver)(nil)

type memoryJob struct {
	executing bool
	job       job.Interface
}

// MemoryDriver memory workerpool driver
type MemoryDriver struct {
	AbstractDriver
	jobs       *list.SyncLinkedList[*memoryJob]
	failedJobs *list.SyncLinkedList[job.Interface]
}

// NewMemoryDriver creates a new memory driver
func NewMemoryDriver() *MemoryDriver {
	driver := new(MemoryDriver)
	driver.jobs = list.NewSyncLinkedList[*memoryJob]()
	driver.failedJobs = list.NewSyncLinkedList[job.Interface]()
	driver.AbstractDriver.EventBus = eventbus.NewEventBus()
	_ = driver.AbstractDriver.EventBus.AddEvent(new(event.BeforeHandle))
	_ = driver.AbstractDriver.EventBus.AddEvent(new(event.AfterHandle))
	_ = driver.AbstractDriver.EventBus.AddEvent(new(event.FailedHandle))
	_ = driver.AbstractDriver.EventBus.AddEvent(new(event.RetryHandle))
	return driver
}

// Count returns the count of pending jobs
func (driver *MemoryDriver) Count() int64 {
	return int64(driver.jobs.Count())
}

// IsEmpty returns if the count of pending jobs is zero
func (driver *MemoryDriver) IsEmpty() bool {
	return driver.jobs.IsEmpty()
}

// Enqueue pushes a job to queue
func (driver *MemoryDriver) Enqueue(job job.Interface) bool {
	driver.jobs.Push(&memoryJob{
		job: job,
	})
	return true
}

// Dequeue pops a job from queue
func (driver *MemoryDriver) Dequeue() (job.Interface, bool) {
	if driver.jobs.IsEmpty() {
		return nil, false
	}
	if driver.jobs.IsEmpty() {
		return nil, false
	}
	job, err := driver.jobs.FirstWhere(func(value *memoryJob) bool {
		return !value.executing
	})
	return job.job, err == nil
}

// Remove removes a job from queue
func (driver *MemoryDriver) Remove(value job.Interface) bool {
	driver.jobs.Remove(func(item *memoryJob) bool {
		return item.job == value
	})
	return true
}

// Ack acks a job
func (driver *MemoryDriver) Ack(job job.Interface) bool {
	return driver.Remove(job)
}

// Fail handles a failed job
func (driver *MemoryDriver) Fail(job job.Interface) {
	driver.failedJobs.Push(job)
}

// Flush removes all failed jobs
func (driver *MemoryDriver) Flush() {
}

// Reload reloads all failed jobs into queue
func (driver *MemoryDriver) Reload() {
}

// Subscribe add a subscriber to queue events
func (driver *MemoryDriver) Subscribe(subscriber subscriber.Interface) {
	_ = driver.EventBus.Subscribe(subscriber)
}
