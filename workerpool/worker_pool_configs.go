package workerpool

import (
	"time"

	"github.com/wardonne/gopi/workerpool/subscriber"
)

var (
	DefaultMaxWorkers = 3
	// job configs
	DefaultJobMaxAttempts         = 3
	DefaultJobRetryDelay          = 5 * time.Second
	DefaultJobRetryMaxDelay       = time.Minute
	DefaultJobRetryDelayStep      = 5 * time.Second
	DefaultJobMaxExecuteTimeTotal = 10 * time.Minute
	// worker configs
	DefaultWorkerBatch          = 10
	DefaultWorkerMaxIdleTime    = 2 * time.Minute
	DefaultWorkerStoppedTimeout = 5 * time.Minute
)

// WorkerPoolConfigs is a struct contains all worker pool configurtions
type WorkerPoolConfigs struct {
	MaxWorkers int
	// Worker configs
	WorkerConfigs struct {
		Batch          int
		MaxIdleTime    time.Duration
		MaxStoppedTime time.Duration
	}
	// Job configs
	JobConfigs struct {
		MaxAttempts    int
		RetryDelay     time.Duration
		RetryMaxDelay  time.Duration
		RetryDelayStep time.Duration
		MaxExecuteTime time.Duration
	}
	// Subscriber
	Subscriber subscriber.SubscriberInterface
}

// ToOptions converts the configurations to [Option]s
func (configs *WorkerPoolConfigs) ToOptions() []Option {
	return []Option{
		MaxWorkers(configs.MaxWorkers),
		WorkerBatch(configs.WorkerConfigs.Batch),
		WorkerMaxIdleTime(configs.WorkerConfigs.MaxIdleTime),
		WorkerMaxStoppedTime(configs.WorkerConfigs.MaxStoppedTime),
		JobMaxAttempts(configs.JobConfigs.MaxAttempts),
		JobRetryDelay(configs.JobConfigs.RetryDelay),
		JobRetryMaxDelay(configs.JobConfigs.RetryMaxDelay),
		JobRetryDelayStep(configs.JobConfigs.RetryDelayStep),
		JobMaxExecuteTimeTotal(configs.JobConfigs.MaxExecuteTime),
		Subscriber(configs.Subscriber),
	}
}
