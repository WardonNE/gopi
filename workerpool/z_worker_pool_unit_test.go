package workerpool

import (
	"errors"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/wardonne/gopi/eventbus"
	"github.com/wardonne/gopi/workerpool/driver"
	"github.com/wardonne/gopi/workerpool/subscriber"
)

func TestWorkerPool_MaxWorkers(t *testing.T) {
	wp := NewWorkerPool(driver.NewMemoryDriver(), MaxWorkers(1), WorkerBatch(1))
	assert.Equal(t, 1, wp.maxWorkers)
	wp.Start()
	assert.Equal(t, 0, len(wp.Workers()))
	output := make(chan int)
	defer close(output)
	for i := 0; i < 10; i++ {
		go assert.True(t, wp.Dispatch(&testjob{callback: func() error {
			time.Sleep(time.Second)
			return nil
		}}))
	}
	assert.Equal(t, 1, len(wp.Workers()))
}

func TestWorkerPool_Start(t *testing.T) {
	wp := NewWorkerPool(driver.NewMemoryDriver(), MaxWorkers(10), WorkerBatch(1))
	assert.False(t, wp.IsRunning())
	wp.Start()
	assert.True(t, wp.IsRunning())
	var mu sync.Mutex
	outputs := []int{}
	for i := 0; i < 10; i++ {
		j := i
		assert.True(t, wp.Dispatch(&testjob{callback: func() error {
			time.Sleep(time.Second)
			mu.Lock()
			defer mu.Unlock()
			outputs = append(outputs, j)
			return nil
		}}))
	}
	workers := wp.Workers()
	assert.Equal(t, 10, len(workers))
	time.Sleep(100 * time.Millisecond)
	for _, worker := range workers {
		assert.True(t, worker.IsWorking())
	}
	time.Sleep(time.Second)
	assert.ElementsMatch(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, outputs)
}

func TestWorkerPool_Stop(t *testing.T) {
	wp := NewWorkerPool(driver.NewMemoryDriver(), MaxWorkers(1), WorkerBatch(1))
	wp.Start()
	var mu sync.Mutex
	outputs := []int{}
	for i := 0; i < 10; i++ {
		j := i
		assert.True(t, wp.Dispatch(&testjob{callback: func() error {
			time.Sleep(1 * time.Second)
			mu.Lock()
			defer mu.Unlock()
			outputs = append(outputs, j)
			return nil
		}}))
	}
	// sleep 500ms to make sure the worker is executing jobs
	time.Sleep(500 * time.Millisecond)
	wp.Stop()
	workers := wp.Workers()
	for _, worker := range workers {
		assert.True(t, worker.IsStopped())
	}
	assert.True(t, wp.IsStopped())
	// can't stop executing job immediately
	assert.Equal(t, 1, len(outputs))
}

func TestWorkerPool_Release(t *testing.T) {
	wp := NewWorkerPool(driver.NewMemoryDriver(), MaxWorkers(1), WorkerBatch(1))
	wp.Start()
	var mu sync.Mutex
	outputs := []int{}
	for i := 0; i < 10; i++ {
		j := i
		assert.True(t, wp.Dispatch(&testjob{callback: func() error {
			time.Sleep(time.Second)
			mu.Lock()
			defer mu.Unlock()
			outputs = append(outputs, j)
			return nil
		}}))
	}
	wp.Release()
	time.Sleep(time.Second)
	assert.Equal(t, 0, len(wp.Workers()))
	assert.Equal(t, 1, len(outputs))
}

func TestWorkerPool_Watcher(t *testing.T) {
	wp := NewWorkerPool(
		driver.NewMemoryDriver(),
		MaxWorkers(1),
		WorkerBatch(1),
		WorkerMaxIdleTime(time.Second),
		WorkerMaxStoppedTime(time.Second),
	)
	wp.Start()
	assert.True(t, wp.Dispatch(&testjob{callback: func() error {
		time.Sleep(time.Second)
		return nil
	}}))
	time.Sleep(500 * time.Millisecond)
	workers := wp.Workers()
	for _, worker := range workers {
		assert.True(t, worker.IsWorking())
	}
	time.Sleep(time.Second)
	// should be idled
	for _, worker := range workers {
		assert.True(t, worker.IsIdle())
	}
	time.Sleep(time.Second)
	// should be stopped
	for _, worker := range workers {
		assert.True(t, worker.IsStopped())
	}
	time.Sleep(time.Second)
	// should be released
	workers = wp.Workers()
	assert.Equal(t, 0, len(workers))
}

func TestWorkerPool_Events(t *testing.T) {
	t.Run("job success", func(t *testing.T) {
		var beforeHandle bool
		var afterHandle bool
		var failedHandle bool
		var retryHandle bool
		var progressUpdated bool
		wp := NewWorkerPool(
			driver.NewMemoryDriver(),
			MaxWorkers(1),
			WorkerBatch(1),
			Subscriber(&subscriber.Subscriber{
				BeforeHandle: func(eb eventbus.EventInterface) bool {
					beforeHandle = true
					return true
				},
				AfterHandle: func(eb eventbus.EventInterface) bool {
					afterHandle = true
					return true
				},
				FailedHandle: func(eb eventbus.EventInterface) bool {
					failedHandle = true
					return true
				},
				RetryHandle: func(eb eventbus.EventInterface) bool {
					retryHandle = true
					return true
				},
				ProgressUpdated: func(eb eventbus.EventInterface) bool {
					progressUpdated = true
					return true
				},
			}),
		)
		wp.Start()
		assert.True(t, wp.Dispatch(&testjob{
			callback: func() error {
				time.Sleep(100 * time.Millisecond)
				return nil
			},
		}))
		time.Sleep(500 * time.Millisecond)
		assert.True(t, beforeHandle)
		assert.True(t, afterHandle)
		assert.True(t, progressUpdated)
		assert.False(t, failedHandle)
		assert.False(t, retryHandle)
	})

	t.Run("failed-without-retry", func(t *testing.T) {
		var beforeHandle bool
		var afterHandle bool
		var failedHandle bool
		var retryHandle bool
		var progressUpdated bool
		wp := NewWorkerPool(
			driver.NewMemoryDriver(),
			MaxWorkers(1),
			WorkerBatch(1),
			Subscriber(&subscriber.Subscriber{
				BeforeHandle: func(eb eventbus.EventInterface) bool {
					beforeHandle = true
					return true
				},
				AfterHandle: func(eb eventbus.EventInterface) bool {
					afterHandle = true
					return true
				},
				FailedHandle: func(eb eventbus.EventInterface) bool {
					failedHandle = true
					return true
				},
				RetryHandle: func(eb eventbus.EventInterface) bool {
					retryHandle = true
					return true
				},
				ProgressUpdated: func(eb eventbus.EventInterface) bool {
					progressUpdated = true
					return true
				},
			}),
		)
		wp.Start()
		retryable := false
		wp.Dispatch(&testjob{
			retryable: &retryable,
			callback: func() error {
				return errors.New("err")
			},
		})
		time.Sleep(500 * time.Millisecond)
		assert.True(t, beforeHandle)
		assert.False(t, afterHandle)
		assert.True(t, progressUpdated)
		assert.True(t, failedHandle)
		assert.False(t, retryHandle)
	})

	t.Run("all-failed-with-retry", func(t *testing.T) {
		var beforeHandle bool
		var afterHandle bool
		var failedHandle bool
		var retryHandle bool
		var progressUpdated bool
		wp := NewWorkerPool(
			driver.NewMemoryDriver(),
			MaxWorkers(1),
			WorkerBatch(1),
			JobMaxAttempts(3),
			JobRetryDelay(0),
			JobRetryDelayStep(0),
			Subscriber(&subscriber.Subscriber{
				BeforeHandle: func(eb eventbus.EventInterface) bool {
					beforeHandle = true
					return true
				},
				AfterHandle: func(eb eventbus.EventInterface) bool {
					afterHandle = true
					return true
				},
				FailedHandle: func(eb eventbus.EventInterface) bool {
					failedHandle = true
					return true
				},
				RetryHandle: func(eb eventbus.EventInterface) bool {
					retryHandle = true
					return true
				},
				ProgressUpdated: func(eb eventbus.EventInterface) bool {
					progressUpdated = true
					return true
				},
			}),
		)
		wp.Start()
		retryable := true
		wp.Dispatch(&testjob{
			retryable: &retryable,
			callback: func() error {
				panic(errors.New("err"))
			},
		})
		time.Sleep(500 * time.Millisecond)
		assert.True(t, beforeHandle)
		assert.False(t, afterHandle)
		assert.True(t, progressUpdated)
		assert.True(t, failedHandle)
		assert.True(t, retryHandle)
	})

	t.Run("success-after-retry", func(t *testing.T) {
		var beforeHandle bool
		var afterHandle bool
		var failedHandle bool
		var retryHandle bool
		var progressUpdated bool
		wp := NewWorkerPool(
			driver.NewMemoryDriver(),
			MaxWorkers(1),
			WorkerBatch(1),
			JobMaxAttempts(3),
			JobRetryDelay(0),
			JobRetryDelayStep(0),
			Subscriber(&subscriber.Subscriber{
				BeforeHandle: func(eb eventbus.EventInterface) bool {
					beforeHandle = true
					return true
				},
				AfterHandle: func(eb eventbus.EventInterface) bool {
					afterHandle = true
					return true
				},
				FailedHandle: func(eb eventbus.EventInterface) bool {
					failedHandle = true
					return true
				},
				RetryHandle: func(eb eventbus.EventInterface) bool {
					retryHandle = true
					return true
				},
				ProgressUpdated: func(eb eventbus.EventInterface) bool {
					progressUpdated = true
					return true
				},
			}),
		)
		wp.Start()
		retryable := true
		attempt := 0
		wp.Dispatch(&testjob{
			retryable: &retryable,
			callback: func() error {
				attempt++
				if attempt <= 2 {
					return fmt.Errorf("err")
				}
				return nil
			},
		})
		time.Sleep(500 * time.Millisecond)
		assert.True(t, beforeHandle)
		assert.True(t, afterHandle)
		assert.True(t, progressUpdated)
		assert.False(t, failedHandle)
		assert.True(t, retryHandle)
	})
}
