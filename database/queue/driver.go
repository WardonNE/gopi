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

// Driver database workerpool driver
type Driver struct {
	driver.AbstractDriver
	*gorm.DB

	Queue     string
	TableName string
}

// NewDriver create a new database driver
func NewDriver(db *gorm.DB, tableName string, queueName string) *Driver {
	driver := new(Driver)
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
func (d *Driver) Count() int64 {
	var total int64
	if err := d.Table(d.TableName).Where("queue = ?", d.Queue).
		Count(&total).Error; err != nil {
		panic(err)
	}
	return total
}

// IsEmpty returns if the count of pending jobs is zero
func (d *Driver) IsEmpty() bool {
	return d.Count() > 0
}

// Enqueue pushes a job to queue
func (d *Driver) Enqueue(job job.Interface) bool {
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
func (d *Driver) Dequeue() job.Interface {
	// d.Table(d.TableName).
	// 	Where("queue = ?", d.Queue).
	// 	Where("executed_at IS NULL").
	// 	Where("avaliable_at <= ?", time.Now().UTC()).
	// 	First()
	return nil
}

// Remove removes a job from queue
func (d *Driver) Remove(job job.Interface) {
	panic("not implemented") // TODO: Implement
}

// Ack acks a job
func (d *Driver) Ack(job job.Interface) {
	panic("not implemented") // TODO: Implement
}

// Fail handles a failed job
func (d *Driver) Fail(job job.Interface) {
	panic("not implemented") // TODO: Implement
}

// Flush removes all failed jobs
func (d *Driver) Flush() {
	panic("not implemented") // TODO: Implement
}

// Reload reloads all failed jobs into queue
func (d *Driver) Reload() {
	panic("not implemented") // TODO: Implement
}

// Subscribe add a subscriber to queue events
func (d *Driver) Subscribe(subscriber subscriber.Interface) {
	panic("not implemented") // TODO: Implement
}

// DispatchEvent dispatches specifia event
func (d *Driver) DispatchEvent(event eventbus.EventInterface) {
	panic("not implemented") // TODO: Implement
}
