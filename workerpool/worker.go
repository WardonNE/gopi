package workerpool

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/wardonne/gopi/retry"
	"github.com/wardonne/gopi/support/utils"
	"github.com/wardonne/gopi/workerpool/driver"
	"github.com/wardonne/gopi/workerpool/event"
	"github.com/wardonne/gopi/workerpool/job"
)

// WorkerStatus worker status
type WorkerStatus int

// worker status enums
const (
	WorkerStatusIdle WorkerStatus = iota + 1
	WorkerStatusWorking
	WorkerStatusStopped
)

// Worker is a struct to handle jobs
type Worker struct {
	id        uuid.UUID    // unique id
	status    WorkerStatus // status
	createdAt time.Time    // created time
	startedAt time.Time    // last started time
	idledAt   time.Time    // last idled time
	stoppedAt time.Time    // last stopped time

	driver      driver.IDriver
	stopChannel chan struct{}
	// worker configs
	maxIdleTime    time.Duration
	maxStoppedTime time.Duration
	// job configs
	jobConfigs struct {
		maxAttempts    int
		retryDelay     time.Duration
		retryMaxDelay  time.Duration
		retryDelayStep time.Duration
		maxExecuteTime time.Duration
	}
}

func hire(wp *WorkerPool) *Worker {
	w := new(Worker)
	w.id = uuid.New()
	w.status = WorkerStatusIdle
	w.createdAt = time.Now()
	w.idledAt = time.Now()
	w.driver = wp.driver
	w.stopChannel = make(chan struct{})
	w.maxIdleTime = wp.workerConfigs.maxIdleTime
	w.maxStoppedTime = wp.workerConfigs.maxStoppedTime
	w.jobConfigs = wp.jobConfigs
	return w
}

func (w *Worker) working() {
	w.status = WorkerStatusWorking
	w.startedAt = time.Now()
}

func (w *Worker) idle() {
	w.status = WorkerStatusIdle
	w.idledAt = time.Now()
}

func (w *Worker) stop() {
	w.status = WorkerStatusStopped
	w.stoppedAt = time.Now()
}

// ID returns worker's unique id
func (w *Worker) ID() uuid.UUID {
	return w.id
}

// Status returns worker's active status
//   - [WorkerStatusIdle]
//   - [WorkerStatusWorking]
//   - [WorkerStatusStopped]
func (w *Worker) Status() WorkerStatus {
	return w.status
}

// CreatedAt returns the worker's created time
func (w *Worker) CreatedAt() time.Time {
	return w.createdAt
}

// IdledAt returns the worker's last idle time
func (w *Worker) IdledAt() time.Time {
	return w.idledAt
}

// StoppedAt returns the worker's last stopped time
func (w *Worker) StoppedAt() time.Time {
	return w.stoppedAt
}

// IsWorking returns whether the worker is working
func (w *Worker) IsWorking() bool {
	return w.status == WorkerStatusWorking
}

// IsIdle returns whether the worker is idle
func (w *Worker) IsIdle() bool {
	return w.status == WorkerStatusIdle
}

// IsStopped returns whether the worker is stopped
func (w *Worker) IsStopped() bool {
	return w.status == WorkerStatusStopped
}

// Stoppable returns whether the worker can be stopped
func (w *Worker) Stoppable() bool {
	return w.status == WorkerStatusIdle || w.status == WorkerStatusWorking
}

func (w *Worker) execute(job job.Interface) {
	w.working()
	var lifetime time.Duration
	if v := job.MaxExecuteTime(); v != nil {
		lifetime = *v
	} else if w.jobConfigs.maxExecuteTime > 0 {
		lifetime = w.jobConfigs.maxExecuteTime
	}
	fn := func() error {
		executor := func() (err error) {
			defer func() {
				if exp := recover(); exp != nil {
					switch e := exp.(type) {
					case error:
						err = e
					case string:
						err = errors.New(e)
					default:
						err = errors.New("should retry")
					}
				}
			}()
			return job.Handle()
		}
		if job.Retryable() {
			retryConfigs := &retry.Configs{}
			if job.MaxAttempts() != nil {
				retryConfigs.Attempts = *job.MaxAttempts()
			} else {
				retryConfigs.Attempts = w.jobConfigs.maxAttempts
			}
			if delay := job.RetryDelay(); delay != nil {
				retryConfigs.Delay = utils.Max(*delay, 0)
			} else {
				retryConfigs.Delay = w.jobConfigs.retryDelay
			}
			if maxDelay := job.RetryMaxDelay(); maxDelay != nil {
				retryConfigs.MaxDelay = utils.Max(*maxDelay, 0)
			} else {
				retryConfigs.MaxDelay = w.jobConfigs.retryMaxDelay
			}
			if delayStep := job.RetryDelayStep(); delayStep != nil {
				retryConfigs.DelayStep = utils.Max(*delayStep, 0)
			} else {
				retryConfigs.DelayStep = w.jobConfigs.retryDelayStep
			}
			retryConfigs.ShouldRetry = job.ShouldRetry
			retryConfigs.OnRetry = func(i int, err error) {
				w.driver.DispatchEvent(event.NewRetryHandle(job, i, err))
			}
			// if released w.driver will be nil
			if w.driver != nil {
				w.driver.DispatchEvent(event.NewBeforeHandle(job))
			}
			err := retry.DoWithConfigs(executor, retryConfigs)
			return err
		}
		// if released w.driver will be nil
		if w.driver != nil {
			w.driver.DispatchEvent(event.NewBeforeHandle(job))
		}
		err := executor()
		return err
	}
	var err error
	if lifetime > 0 {
		err = utils.RunOutWithErrorCause(fn, lifetime, ErrJobExecuteTimeout)
	} else {
		err = fn()
	}
	if err != nil && w.driver != nil {
		w.driver.DispatchEvent(event.NewFailedHandle(job, err))
		w.driver.Fail(job)
	} else if err == nil && w.driver != nil {
		w.driver.DispatchEvent(event.NewAfterHandle(job))
		w.driver.Ack(job)
	}
	w.idle()
}

// Start lets the worker starting worker
func (w *Worker) Start() {
	for {
		select {
		case <-w.stopChannel:
			w.stop()
			return
		default:
			if job, ok := w.driver.Dequeue(); ok {
				w.execute(job)
			}
		}
	}
}

// Stop stops the worker, if the worker's status is [WorkerStatusIdle] it will be stopped immediately
// if the worker's status is [WorkerStatusWorking], its status will change to [WorkerStatusStopping] and
// will be stopped after MaxExecuteTimePerAttempt
func (w *Worker) Stop() {
	w.stopChannel <- struct{}{}
}

// Release releases the worker
func (w *Worker) Release() {
	defer func() {
		w.driver = nil
	}()
	w.Stop()
	close(w.stopChannel)
}

// ShouldStop returns if the worker should be stopped
// it will return true when worker's status is [WorkerStatusIdle] and has been idled over max idle time
func (w *Worker) ShouldStop() bool {
	return w.IsIdle() && time.Since(w.idledAt) >= w.maxIdleTime
}

// ShouldRelease returns if the worker should be released,
// it will return true when worker's status is [WorkerStatusStopped] and has been stopped over max stopped time
func (w *Worker) ShouldRelease() bool {
	return w.IsStopped() && time.Since(w.stoppedAt) >= w.maxStoppedTime
}
