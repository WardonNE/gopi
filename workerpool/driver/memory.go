package driver

import (
	"sync"

	"github.com/wardonne/gopi/eventbus"
	"github.com/wardonne/gopi/support/collection/list"
	"github.com/wardonne/gopi/workerpool/event"
	"github.com/wardonne/gopi/workerpool/job"
	"github.com/wardonne/gopi/workerpool/subscriber"
)

var _ DriverInterface = (*MemoryDriver)(nil)

type MemoryDriver struct {
	sync.Mutex
	AbstractDriver
	list.ArrayList[job.JobInterface]

	total     int64
	pending   int64
	executing int64
	completed int64
}

func NewMemoryDriver() *MemoryDriver {
	driver := new(MemoryDriver)
	driver.AbstractDriver.EventBus = eventbus.NewEventBus()
	_ = driver.AbstractDriver.EventBus.AddEvent(new(event.BeforeHandle))
	_ = driver.AbstractDriver.EventBus.AddEvent(new(event.AfterHandle))
	_ = driver.AbstractDriver.EventBus.AddEvent(new(event.FailedHandle))
	_ = driver.AbstractDriver.EventBus.AddEvent(new(event.RetryHandle))
	_ = driver.AbstractDriver.EventBus.AddEvent(new(event.ProgressUpdated))
	return driver
}

func (driver *MemoryDriver) Count() int {
	driver.Lock()
	defer driver.Unlock()
	return driver.ArrayList.Count()
}

// IsEmpty returns if the count of pending jobs is zero
func (driver *MemoryDriver) IsEmpty() bool {
	driver.Lock()
	defer driver.Unlock()
	return driver.ArrayList.IsEmpty()
}

func (driver *MemoryDriver) onProgressUpdated() {
	_ = driver.EventBus.Dispatch(event.NewProgressUpdated(
		driver.total,
		driver.pending,
		driver.executing,
		driver.completed,
	), nil)
}

// Enqueue pushes a job to queue
func (driver *MemoryDriver) Enqueue(job job.JobInterface) bool {
	driver.Lock()
	defer driver.Unlock()
	driver.ArrayList.Push(job)
	driver.total++
	driver.pending++
	go driver.onProgressUpdated()
	return true
}

// Dequeue pops a job from queue
func (driver *MemoryDriver) Dequeue() (job.JobInterface, bool) {
	driver.Lock()
	defer driver.Unlock()
	if driver.ArrayList.IsEmpty() {
		return nil, false
	}
	job, ok := driver.ArrayList.Shift(), true
	driver.executing++
	driver.pending--
	go driver.onProgressUpdated()
	return job, ok
}

// Remove removes a job from queue
func (driver *MemoryDriver) Remove(value job.JobInterface) bool {
	driver.Lock()
	defer driver.Unlock()
	driver.ArrayList.Remove(func(item job.JobInterface) bool {
		return item == value
	})
	driver.total--
	driver.pending--
	go driver.onProgressUpdated()
	return true
}

// Ack acks a job
func (driver *MemoryDriver) Ack(job job.JobInterface) bool {
	driver.executing--
	driver.completed++
	go driver.onProgressUpdated()
	return true
}

// Fail handles a failed job
func (driver *MemoryDriver) Fail(job job.JobInterface) {
	driver.executing--
	driver.completed++
	go driver.onProgressUpdated()
}

// Flush removes all failed jobs
func (driver *MemoryDriver) Flush() {
}

// Reload reloads all failed jobs into queue
func (driver *MemoryDriver) Reload() {
}

// Subscribe add a subscriber to queue events
func (driver *MemoryDriver) Subscribe(subscriber subscriber.SubscriberInterface) {
	_ = driver.EventBus.Subscribe(subscriber)
}

// Lock locks queue
func (driver *MemoryDriver) Lock() {
	driver.Mutex.Lock()
}

// Unlock unlocks queue
func (driver *MemoryDriver) Unlock() {
	driver.Mutex.Unlock()
}
