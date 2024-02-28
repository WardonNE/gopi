package job

import (
	"time"

	"github.com/wardonne/gopi/support/serializer"
)

// Interface is job's interface
type Interface interface {
	serializer.JSONSerializer
	// Delay returns delay
	Delay() *time.Duration
	// Retryable returns whether a job is retryable
	Retryable() bool
	// ShouldRetry returns whether a job should be retried when err is occured
	ShouldRetry(err error) bool
	// MaxAttempts returns count of attempts.
	// if nil is returned, it will use the pool's config, see [WorkerPoolConfigs]
	// if negative is returned it will use 0
	MaxAttempts() *int
	// RetryDelay returns the min delay before retry
	// if nil is returned, it will use the pool's config, see [WorkerPoolConfigs]
	// if negative is returned it will use 0
	RetryDelay() *time.Duration
	// RetryMaxDelay returns the max delay before retry
	// if nil is returned, it will use the pool's config, see [WorkerPoolConfigs]
	// if negative is returned it will use 0
	RetryMaxDelay() *time.Duration
	// RetryDelayStep returns the delay step
	// if nil is retruned, it will use the pool's config, see [WorkerPoolConfigs]
	// if negative is returned it will use 0
	RetryDelayStep() *time.Duration
	// MaxExecuteTime returns the max total execution time
	// if nil is returned it will use the pool's config, see [WorkerPoolConfigs]
	// if negative is returned it will use 0
	MaxExecuteTime() *time.Duration
	// Handle is the main function to handle the job
	// when an error was returned, if necessary, it will be retried
	// if it still returns an error after max attempts, an error will be passed to the [Failed] function
	Handle() error
}

// Job is a basic job with retry enabled
//
// example:
//
//	type QueryJob struct {
//		Job
//	}
//
//	// write your own handle function
//	func (job *QueryJob) Handle() error {
//		// code...
//		return nil
//	}
type Job struct {
}

// Delay returns job delay
func (job *Job) Delay() *time.Duration {
	return nil
}

// Retryable returns whether a job is retryable
func (job *Job) Retryable() bool {
	return true
}

// ShouldRetry returns whether a job should be retried when err is occured
func (job *Job) ShouldRetry(err error) bool {
	return err != nil
}

// MaxAttempts returns count of attempts.
// if nil is returned, it will use the pool's config, see [WorkerPoolConfigs]
// if negative is returned it will use 0
func (job *Job) MaxAttempts() *int {
	return nil
}

// RetryDelay returns the min delay before retry
// if nil is returned, it will use the pool's config, see [WorkerPoolConfigs]
// if negative is returned it will use 0
func (job *Job) RetryDelay() *time.Duration {
	return nil
}

// RetryMaxDelay returns the max delay before retry
// if nil is returned, it will use the pool's config, see [WorkerPoolConfigs]
// if negative is returned it will use 0
func (job *Job) RetryMaxDelay() *time.Duration {
	return nil
}

// RetryDelayStep returns the delay step
// if nil is retruned, it will use the pool's config, see [WorkerPoolConfigs]
// if negative is returned it will use 0
func (job *Job) RetryDelayStep() *time.Duration {
	return nil
}

// MaxExecuteTime returns the max total execution time
// if nil is returned it will use the pool's config, see [WorkerPoolConfigs]
// if negative is returned it will use 0
func (job *Job) MaxExecuteTime() *time.Duration {
	return nil
}
