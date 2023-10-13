package workerpool

import (
	"errors"
	"sync"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/wardonne/gopi/eventbus"
	"github.com/wardonne/gopi/workerpool/driver"
	"github.com/wardonne/gopi/workerpool/event"
	"github.com/wardonne/gopi/workerpool/job"
	"github.com/wardonne/gopi/workerpool/subscriber"
)

type testjob struct {
	job.Job
	retryable *bool
	callback  func() error
}

func (j *testjob) MarshalJSON() ([]byte, error) {
	return []byte{}, nil
}

func (j *testjob) UnmarshalJSON(data []byte) error {
	return nil
}

func (j *testjob) Retryable() bool {
	if j.retryable == nil {
		return true
	}
	return *j.retryable
}

func (j *testjob) Handle() (err error) {
	return j.callback()
}

func TestWorker_Start(t *testing.T) {
	driver := driver.NewMemoryDriver()
	w := &Worker{
		id:          uuid.New(),
		status:      WorkerStatusIdle,
		createdAt:   time.Now(),
		idledAt:     time.Now(),
		driver:      driver,
		stopChannel: make(chan struct{}),
	}
	mu := sync.Mutex{}
	outputs := []int{}
	for i := 0; i < 10; i++ {
		j := i
		assert.True(t, driver.Enqueue(&testjob{callback: func() error {
			time.Sleep(100 * time.Millisecond)
			mu.Lock()
			defer mu.Unlock()
			outputs = append(outputs, j)
			return nil
		}}))
	}
	go func() { w.Start() }()
	time.Sleep(50 * time.Millisecond)
	assert.True(t, w.IsWorking())
	time.Sleep(2 * time.Second)
	assert.ElementsMatch(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, outputs)
}

func TestWorker_Stop(t *testing.T) {
	driver := driver.NewMemoryDriver()
	w := &Worker{
		id:          uuid.New(),
		status:      WorkerStatusIdle,
		createdAt:   time.Now(),
		idledAt:     time.Now(),
		driver:      driver,
		stopChannel: make(chan struct{}),
	}
	go func() { w.Start() }()
	mu := sync.Mutex{}
	outputs := []int{}
	for i := 0; i < 10; i++ {
		j := i
		assert.True(t, driver.Enqueue(&testjob{callback: func() error {
			time.Sleep(3 * time.Second)
			mu.Lock()
			defer mu.Unlock()
			outputs = append(outputs, j)
			return nil
		}}))
	}
	time.Sleep(50 * time.Millisecond)
	assert.True(t, w.IsWorking())
	w.Stop()
	assert.Equal(t, 1, len(outputs))
	assert.True(t, w.IsStopped())
}

func TestWorker_MaxExecuteTimeTotal(t *testing.T) {
	t.Run("timeout-without-retry", func(t *testing.T) {
		var err error
		driver := driver.NewMemoryDriver()
		driver.Subscribe(&subscriber.Subscriber{
			FailedHandle: func(ei eventbus.EventInterface) bool {
				event := ei.(*event.FailedHandle)
				err = event.Error
				return true
			},
		})
		w := &Worker{
			id:          uuid.New(),
			status:      WorkerStatusIdle,
			createdAt:   time.Now(),
			idledAt:     time.Now(),
			driver:      driver,
			stopChannel: make(chan struct{}),
			jobConfigs: struct {
				maxAttempts    int
				retryDelay     time.Duration
				retryMaxDelay  time.Duration
				retryDelayStep time.Duration
				maxExecuteTime time.Duration
			}{
				maxAttempts:    1,
				retryDelay:     0,
				retryMaxDelay:  0,
				retryDelayStep: 0,
				maxExecuteTime: time.Second,
			},
		}
		go func() { w.Start() }()
		retryable := false
		assert.True(t, driver.Enqueue(&testjob{
			retryable: &retryable,
			callback: func() error {
				time.Sleep(3 * time.Second)
				return nil
			},
		}))
		time.Sleep(500 * time.Millisecond)
		assert.Nil(t, err)
		time.Sleep(time.Second)
		assert.Equal(t, ErrJobExecuteTimeout.Error(), err.Error())
	})

	t.Run("timeout-in-first-attempt", func(t *testing.T) {
		var err error
		driver := driver.NewMemoryDriver()
		driver.Subscribe(&subscriber.Subscriber{
			FailedHandle: func(ei eventbus.EventInterface) bool {
				event := ei.(*event.FailedHandle)
				err = event.Error
				return true
			},
		})
		w := &Worker{
			id:          uuid.New(),
			status:      WorkerStatusIdle,
			createdAt:   time.Now(),
			idledAt:     time.Now(),
			driver:      driver,
			stopChannel: make(chan struct{}),
			jobConfigs: struct {
				maxAttempts    int
				retryDelay     time.Duration
				retryMaxDelay  time.Duration
				retryDelayStep time.Duration
				maxExecuteTime time.Duration
			}{
				maxAttempts:    1,
				retryDelay:     0,
				retryMaxDelay:  0,
				retryDelayStep: 0,
				maxExecuteTime: time.Second,
			},
		}
		go func() { w.Start() }()
		assert.True(t, driver.Enqueue(&testjob{
			callback: func() error {
				time.Sleep(3 * time.Second)
				return nil
			},
		}))
		time.Sleep(500 * time.Millisecond)
		assert.Nil(t, err)
		time.Sleep(time.Second)
		assert.Equal(t, ErrJobExecuteTimeout.Error(), err.Error())
	})

	t.Run("timeout-not-in-first-attempt", func(t *testing.T) {
		var err error
		var attempt int
		driver := driver.NewMemoryDriver()
		driver.Subscribe(&subscriber.Subscriber{
			FailedHandle: func(ei eventbus.EventInterface) bool {
				event := ei.(*event.FailedHandle)
				err = event.Error
				return true
			},
			RetryHandle: func(ei eventbus.EventInterface) bool {
				attempt++
				return true
			},
		})
		w := &Worker{
			id:          uuid.New(),
			status:      WorkerStatusIdle,
			createdAt:   time.Now(),
			idledAt:     time.Now(),
			driver:      driver,
			stopChannel: make(chan struct{}),
			jobConfigs: struct {
				maxAttempts    int
				retryDelay     time.Duration
				retryMaxDelay  time.Duration
				retryDelayStep time.Duration
				maxExecuteTime time.Duration
			}{
				maxAttempts:    3,
				retryDelay:     0,
				retryMaxDelay:  0,
				retryDelayStep: 0,
				maxExecuteTime: 2 * time.Second,
			},
		}
		go func() { w.Start() }()
		assert.True(t, driver.Enqueue(&testjob{
			callback: func() error {
				time.Sleep(2 * time.Second)
				return errors.New("err")
			},
		}))
		time.Sleep(500 * time.Millisecond)
		assert.Nil(t, err)
		time.Sleep(3 * time.Second)
		assert.Equal(t, ErrJobExecuteTimeout.Error(), err.Error())
		assert.Equal(t, 1, attempt)
	})
}
