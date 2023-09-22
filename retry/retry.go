package retry

import (
	"context"
	"time"

	"github.com/wardonne/gopi/utils"
)

var (
	DefaultAttempts                    = 3
	DefaultDelay         time.Duration = 5 * time.Second
	DefaultMaxDelay      time.Duration = time.Minute
	DefaultDelayStep     time.Duration = 0
	DefaultShouldRetryFn               = func(err error) bool { return err != nil }
	DefaultOnRetry                     = func(int, error) {}
)

type RetryConfigs struct {
	Ctx         context.Context
	Attempts    int
	Delay       time.Duration
	MaxDelay    time.Duration
	DelayStep   time.Duration
	ShouldRetry func(err error) bool
	OnRetry     func(i int, err error)
}

func (r *RetryConfigs) ToOptions() []Option {
	return []Option{
		Context(r.Ctx),
		Attempts(r.Attempts),
		Delay(r.Delay),
		MaxDelay(r.MaxDelay),
		DelayStep(r.DelayStep),
		ShouldRetry(r.ShouldRetry),
		OnRetry(r.OnRetry),
	}
}

type retryConfigs struct {
	ctx         context.Context
	attempts    int
	delay       time.Duration
	maxDelay    time.Duration
	delayStep   time.Duration
	shouldRetry func(err error) bool
	onRetry     func(i int, err error)
}

func Default() *retryConfigs {
	return &retryConfigs{
		ctx:         context.Background(),
		attempts:    DefaultAttempts,
		delay:       DefaultDelay,
		maxDelay:    DefaultMaxDelay,
		delayStep:   DefaultDelayStep,
		shouldRetry: DefaultShouldRetryFn,
		onRetry:     DefaultOnRetry,
	}
}

func New(options ...Option) *retryConfigs {
	retryConfigs := Default()
	for _, option := range options {
		option(retryConfigs)
	}
	return retryConfigs
}

func NewWithConfigs(configs *RetryConfigs) *retryConfigs {
	return New(configs.ToOptions()...)
}

func Do(fn func() error, options ...Option) error {
	return New(options...).Do(fn)
}

func DoWithConfigs(fn func() error, configs *RetryConfigs) error {
	return New(configs.ToOptions()...).Do(fn)
}

func (r *retryConfigs) Do(fn func() error) (err error) {
	select {
	case <-r.ctx.Done():
		return r.ctx.Err()
	default:
		attempts := 0
		for {
			err = fn()
			if !r.shouldRetry(err) {
				return err
			}
			if r.attempts > 0 && attempts == int(r.attempts)-1 {
				return err
			}
			attempts++
			r.onRetry(attempts, err)
			delay := utils.Min(r.delay+r.delayStep*time.Duration(attempts), r.maxDelay)
			time.Sleep(delay)
		}
	}
}
