package workerpool

import (
	"time"

	"github.com/wardonne/gopi/workerpool/subscriber"
)

var (
	// DefaultMaxWorkers default max workers
	DefaultMaxWorkers = 3
	// DefaultJobMaxAttempts default job max attempts
	DefaultJobMaxAttempts = 3
	// DefaultJobRetryDelay default job retry delay
	DefaultJobRetryDelay = 5 * time.Second
	// DefaultJobRetryMaxDelay default job retry max delay
	DefaultJobRetryMaxDelay = time.Minute
	// DefaultJobRetryDelayStep default job retry delay step
	DefaultJobRetryDelayStep = 5 * time.Second
	// DefaultJobMaxExecuteTimeTotal default job max total execution time
	DefaultJobMaxExecuteTimeTotal = 10 * time.Minute
	// DefaultWorkerBatch default worker number per batch
	DefaultWorkerBatch = 10
	// DefaultWorkerMaxIdleTime default worker max idle time
	DefaultWorkerMaxIdleTime = 2 * time.Minute
	// DefaultWorkerStoppedTimeout default worker stopped timeout
	DefaultWorkerStoppedTimeout = 5 * time.Minute
)

// Configs is a struct contains all worker pool configurtions
type Configs struct {
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
	Subscriber subscriber.Interface
}

// ToOptions converts the configurations to [Option]s
func (configs *Configs) ToOptions() []Option {
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
