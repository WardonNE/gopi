package retry

import (
	"context"
	"time"

	"github.com/wardonne/gopi/support/utils"
)

var (
	// DefaultAttempts default attempts
	DefaultAttempts = 3
	// DefaultRetryDelay default retry delay
	DefaultRetryDelay time.Duration = 5 * time.Second
	// DefaultMaxRetryDelay default max retry delay
	DefaultMaxRetryDelay time.Duration = time.Minute
	// DefaultRetryDelayStep default retry delay step
	DefaultRetryDelayStep time.Duration = 0
	// DefaultShouldRetryFn default function to check wheather a job should retry
	DefaultShouldRetryFn = func(err error) bool { return err != nil }
	// DefaultOnRetry default event on retry
	DefaultOnRetry = func(int, error) {}
)

// Configs retry configs
type Configs struct {
	Ctx         context.Context
	Attempts    int
	Delay       time.Duration
	MaxDelay    time.Duration
	DelayStep   time.Duration
	ShouldRetry func(err error) bool
	OnRetry     func(i int, err error)
}

// ToOptions convert [Configs] to options
func (r *Configs) ToOptions() []Option {
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

// Exector the retry exector
type Exector struct {
	Ctx         context.Context
	Attempts    int
	Delay       time.Duration
	MaxDelay    time.Duration
	DelayStep   time.Duration
	ShouldRetry func(err error) bool
	OnRetry     func(i int, err error)
}

// Default default retry configs
func Default() *Exector {
	return &Exector{
		Ctx:         context.Background(),
		Attempts:    DefaultAttempts,
		Delay:       DefaultRetryDelay,
		MaxDelay:    DefaultMaxRetryDelay,
		DelayStep:   DefaultRetryDelayStep,
		ShouldRetry: DefaultShouldRetryFn,
		OnRetry:     DefaultOnRetry,
	}
}

// New create a new exector
func New(options ...Option) *Exector {
	exector := Default()
	for _, option := range options {
		option(exector)
	}
	return exector
}

// NewWithConfigs create a new exector by configs
func NewWithConfigs(configs *Configs) *Exector {
	return New(configs.ToOptions()...)
}

// Do run a function with options
func Do(fn func() error, options ...Option) error {
	return New(options...).Do(fn)
}

// DoWithConfigs run a function with configs
func DoWithConfigs(fn func() error, configs *Configs) error {
	return New(configs.ToOptions()...).Do(fn)
}

// Do run job
func (r *Exector) Do(fn func() error) (err error) {
	attempts := 0
	for {
		err = fn()
		if !r.ShouldRetry(err) {
			return err
		}
		if r.Attempts > 0 && attempts == int(r.Attempts)-1 {
			return err
		}
		delay := utils.Min(r.Delay+r.DelayStep*time.Duration(attempts), r.MaxDelay)
		attempts++
		r.OnRetry(attempts, err)
		select {
		case <-r.Ctx.Done():
			return r.Ctx.Err()
		case <-time.After(delay):
		}
	}
}
