package workerpool

import (
	"time"

	"github.com/wardonne/gopi/workerpool/subscriber"
)

// Option worker pool option fn
type Option func(*WorkerPool)

var noneOption = func(*WorkerPool) {}

// MaxWorkers sets count of workers, default is 3
func MaxWorkers(count int) Option {
	return func(wp *WorkerPool) {
		if count <= 0 {
			count = 1
		}
		wp.maxWorkers = count
	}
}

// JobMaxAttempts sets count of retry, default is 3
func JobMaxAttempts(maxAttempts int) Option {
	return func(wp *WorkerPool) {
		wp.jobConfigs.maxAttempts = maxAttempts
	}
}

// JobRetryDelay sets min delay before retry, default is 5s
func JobRetryDelay(delay time.Duration) Option {
	return func(wp *WorkerPool) {
		if delay < 0 {
			delay = 0
		}
		wp.jobConfigs.retryDelay = delay
	}
}

// JobRetryMaxDelay sets max delay before retry, default is 60s
func JobRetryMaxDelay(maxDelay time.Duration) Option {
	return func(wp *WorkerPool) {
		if maxDelay < 0 {
			maxDelay = 0
		}
		wp.jobConfigs.retryMaxDelay = maxDelay
	}
}

// JobRetryDelayStep sets retry delay steps, default is 5s
func JobRetryDelayStep(delayStep time.Duration) Option {
	return func(wp *WorkerPool) {
		if delayStep < 0 {
			delayStep = 0
		}
		wp.jobConfigs.retryDelayStep = delayStep
	}
}

// JobMaxExecuteTimeTotal sets max execution time, default is 300s
func JobMaxExecuteTimeTotal(d time.Duration) Option {
	return func(wp *WorkerPool) {
		if d < 0 {
			d = 0
		}
		wp.jobConfigs.maxExecuteTime = d
	}
}

// WorkerBatch sets count of worker creation batches, default is 10
func WorkerBatch(batch int) Option {
	return func(wp *WorkerPool) {
		if batch <= 0 {
			batch = 1
		}
		wp.workerConfigs.batch = batch
	}
}

// WorkerMaxIdleTime sets max idle time of a worker, default is 120s
func WorkerMaxIdleTime(d time.Duration) Option {
	return func(wp *WorkerPool) {
		if d < 0 {
			d = 0
		}
		wp.workerConfigs.maxIdleTime = d
	}
}

// WorkerMaxStoppedTime sets max stopped time of a worker, default is 300s
func WorkerMaxStoppedTime(d time.Duration) Option {
	return func(wp *WorkerPool) {
		if d < 0 {
			d = 0
		}
		wp.workerConfigs.maxStoppedTime = d
	}
}

// Subscriber adds a subscriber to queue events
func Subscriber(subscriber subscriber.Interface) Option {
	if subscriber == nil {
		return noneOption
	}
	return func(wp *WorkerPool) {
		wp.driver.Subscribe(subscriber)
	}
}
