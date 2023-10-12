package workerpool

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/wardonne/gopi/retry"
	"github.com/wardonne/gopi/utils"
	"github.com/wardonne/gopi/workerpool/driver"
	"github.com/wardonne/gopi/workerpool/event"
	"github.com/wardonne/gopi/workerpool/job"
)

type WorkerStatus int

const (
	WorkerStatusIdle WorkerStatus = iota + 1
	WorkerStatusWorking
	WorkerStatusStopping
	WorkerStatusStopped
)

type stopSignal uint

const (
	stop stopSignal = 0
	kill stopSignal = 9
)

// Worker is a struct to handle jobs
type Worker struct {
	id        uuid.UUID    // unique id
	status    WorkerStatus // status
	createdAt time.Time    // created time
	startedAt time.Time    // last started time
	idledAt   time.Time    // last idled time
	stoppedAt time.Time    // last stopped time

	driver      driver.DriverInterface
	stopChannel chan stopSignal
	activeJob   job.JobInterface
	// job configs
	jobConfigs struct {
		maxAttempts              int
		retryDelay               time.Duration
		retryMaxDelay            time.Duration
		retryDelayStep           time.Duration
		maxExecuteTimePerAttempt time.Duration
		maxExecuteTimeTotal      time.Duration
	}
}

func hire(wp *WorkerPool) *Worker {
	w := new(Worker)
	w.id = uuid.New()
	w.status = WorkerStatusIdle
	w.createdAt = time.Now()
	w.idledAt = time.Now()
	w.driver = wp.driver
	w.stopChannel = make(chan stopSignal)
	w.jobConfigs = wp.jobConfigs
	return w
}

func (w *Worker) working() {
	w.status = WorkerStatusWorking
	w.startedAt = time.Now()
}

func (w *Worker) idle() {
	w.activeJob = nil
	w.status = WorkerStatusIdle
	w.idledAt = time.Now()
}

func (w *Worker) stopping() {
	w.status = WorkerStatusStopping
}

func (w *Worker) stop() {
	w.activeJob = nil
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
//   - [WorkerStatusStopping]
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

// IsStopping returns whether the worker is stopping
func (w *Worker) IsStopping() bool {
	return w.status == WorkerStatusStopping
}

// IsStopped returns whether the worker is stopped
func (w *Worker) IsStopped() bool {
	return w.status == WorkerStatusStopped
}

// Stoppable returns whether the worker can be stopped
func (w *Worker) Stoppable() bool {
	return w.status == WorkerStatusIdle || w.status == WorkerStatusWorking
}

func (w *Worker) jobExecuteCtx() (context.Context, context.CancelFunc) {
	if maxLifeTime := w.activeJob.MaxExecuteTimeTotal(); maxLifeTime != nil {
		return context.WithTimeout(context.Background(), *maxLifeTime)
	} else if w.jobConfigs.maxExecuteTimeTotal > 0 {
		return context.WithTimeout(context.Background(), w.jobConfigs.maxExecuteTimeTotal)
	}
	return context.WithCancel(context.Background())
}

func (w *Worker) jobAttemptCtx() (context.Context, context.CancelFunc) {
	if executeTimeout := w.activeJob.MaxExecuteTimePerAttempt(); executeTimeout != nil {
		return context.WithTimeout(context.Background(), *executeTimeout)
	} else if w.jobConfigs.maxExecuteTimePerAttempt > 0 {
		return context.WithTimeout(context.Background(), w.jobConfigs.maxExecuteTimePerAttempt)
	}
	return context.WithCancel(context.Background())
}

func (w *Worker) jobExecutor(job job.JobInterface) func() error {
	return func() (err error) {
		defer func() {
			if exp := recover(); err != nil {
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
		attemptCtx, timeoutCancel := w.jobAttemptCtx()
		defer timeoutCancel()
		select {
		case <-attemptCtx.Done():
			err = attemptCtx.Err()
		default:
			err = job.Handle()
		}
		w.idle()
		return
	}
}

func (w *Worker) execute(job job.JobInterface) {
	w.activeJob = job
	w.working()
	executeCtx, ltCancel := w.jobExecuteCtx()
	defer ltCancel()
	select {
	case <-executeCtx.Done():
		w.driver.Fail(job)
		w.idle()
	default:
		executor := w.jobExecutor(job)
		retryConfigs := &retry.RetryConfigs{}
		if job.Retryable() {
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
			w.driver.DispatchEvent(event.NewBeforeHandle(job))
			err := retry.DoWithConfigs(executor, &retry.RetryConfigs{
				Attempts: w.jobConfigs.maxAttempts,
			})
			if err != nil {
				w.driver.DispatchEvent(event.NewFailedHandle(job, err))
				w.driver.Fail(job)
			} else {
				w.driver.DispatchEvent(event.NewAfterHandle(job))
				w.driver.Ack(job)
			}
		} else {
			if err := executor(); err != nil {
				w.driver.DispatchEvent(event.NewFailedHandle(job, err))
				w.driver.Fail(job)
			} else {
				w.driver.DispatchEvent(event.NewAfterHandle(job))
				w.driver.Ack(job)
			}
		}
	}
}

// Start lets the worker starting worker
func (w *Worker) Start() {
	for {
		select {
		case signal, ok := <-w.stopChannel:
			if ok {
				if signal == stop {
					if w.status == WorkerStatusWorking {
						w.stopping()
						ctx, cancel := w.jobAttemptCtx()
						defer cancel()
						<-ctx.Done()
					}
				}
				w.stop()
			}
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
	w.stopChannel <- stop
}

// ForceStop will stop the worker immediately
func (w *Worker) ForceStop() {
	w.stopChannel <- kill
}

func (w *Worker) release() {
	close(w.stopChannel)
	w.driver = nil
	w.activeJob = nil
}
