package job

import (
	"time"

	"github.com/wardonne/gopi/support/serializer"
)

// JobInterface is job's interface
type JobInterface interface {
	serializer.JSONSerializer
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
	// RetryDelay returns the max delay before retry
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

func (job *Job) Retryable() bool {
	return true
}

func (job *Job) ShouldRetry(err error) bool {
	return err != nil
}

func (job *Job) MaxAttempts() *int {
	return nil
}

func (job *Job) RetryDelay() *time.Duration {
	return nil
}

func (job *Job) RetryMaxDelay() *time.Duration {
	return nil
}

func (job *Job) RetryDelayStep() *time.Duration {
	return nil
}

func (job *Job) MaxExecuteTime() *time.Duration {
	return nil
}
