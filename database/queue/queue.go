package queue

import (
	"encoding/json"
	"time"

	"github.com/wardonne/gopi/database/queue/model"
	"github.com/wardonne/gopi/eventbus"
	"github.com/wardonne/gopi/workerpool/driver"
	"github.com/wardonne/gopi/workerpool/event"
	"github.com/wardonne/gopi/workerpool/job"
	"github.com/wardonne/gopi/workerpool/subscriber"
	"gorm.io/gorm"
)

// Queue database workerpool driver
type Queue struct {
	driver.AbstractDriver
	*gorm.DB

	Queue     string
	TableName string
}

// NewQueue create a new database driver
func NewQueue(db *gorm.DB, tableName string, queueName string) *Queue {
	driver := new(Queue)
	driver.DB = db
	driver.TableName = tableName
	driver.AbstractDriver.EventBus = eventbus.NewEventBus()
	_ = driver.AbstractDriver.EventBus.AddEvent(new(event.BeforeHandle))
	_ = driver.AbstractDriver.EventBus.AddEvent(new(event.AfterHandle))
	_ = driver.AbstractDriver.EventBus.AddEvent(new(event.RetryHandle))
	_ = driver.AbstractDriver.EventBus.AddEvent(new(event.FailedHandle))
	return driver
}

// Count returns the count of pending jobs
func (d *Queue) Count() int64 {
	var total int64
	if err := d.Table(d.TableName).Where("queue = ?", d.Queue).
		Count(&total).Error; err != nil {
		panic(err)
	}
	return total
}

// IsEmpty returns if the count of pending jobs is zero
func (d *Queue) IsEmpty() bool {
	return d.Count() > 0
}

// Enqueue pushes a job to queue
func (d *Queue) Enqueue(job job.Interface) bool {
	payload, err := json.Marshal(job)
	if err != nil {
		panic(err)
	}
	var avaliableAt = time.Now()
	if job.Delay() != nil {
		avaliableAt = avaliableAt.Add(*job.Delay())
	}
	if err := d.Table(d.TableName).
		Create(&model.Job{
			Queue:       d.Queue,
			Payload:     payload,
			Attempts:    0,
			ExecutedAt:  nil,
			AvaliableAt: &avaliableAt,
		}).Error; err != nil {
		panic(err)
	}
	return true
}

// Dequeue pops a job from queue
func (d *Queue) Dequeue() job.Interface {
	panic("not implemented") // TODO: Implement
}

// Remove removes a job from queue
func (d *Queue) Remove(job job.Interface) {
	panic("not implemented") // TODO: Implement
}

// Ack acks a job
func (d *Queue) Ack(job job.Interface) {
	panic("not implemented") // TODO: Implement
}

// Fail handles a failed job
func (d *Queue) Fail(job job.Interface) {
	panic("not implemented") // TODO: Implement
}

// Flush removes all failed jobs
func (d *Queue) Flush() {
	panic("not implemented") // TODO: Implement
}

// Reload reloads all failed jobs into queue
func (d *Queue) Reload() {
	panic("not implemented") // TODO: Implement
}

// Subscribe add a subscriber to queue events
func (d *Queue) Subscribe(subscriber subscriber.SubscriberInterface) {
	panic("not implemented") // TODO: Implement
}

// DispatchEvent dispatches specifia event
func (d *Queue) DispatchEvent(event eventbus.EventInterface) {
	panic("not implemented") // TODO: Implement
}
