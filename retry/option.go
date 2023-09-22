package retry

import (
	"context"
	"time"
)

type Option func(configs *retryConfigs)

var noneOption = func(configs *retryConfigs) {
}

func Context(ctx context.Context) Option {
	if ctx == nil {
		return noneOption
	}
	return func(configs *retryConfigs) {
		configs.ctx = ctx
	}
}

// Attempts sets count of retry, setting to <= 0 means retry until success,
// default is 3
func Attempts(attempts int) Option {
	return func(configs *retryConfigs) {
		configs.attempts = attempts
	}
}

// Delay sets min delay before retry, setting to < 0 means 0s,
// default is 5s
func Delay(delay time.Duration) Option {
	if delay < 0 {
		delay = 0
	}
	return func(configs *retryConfigs) {
		configs.delay = delay
	}
}

// MaxDelay sets max delay before retry, settings to <0 means 0s,
// default is 1min
func MaxDelay(maxDelay time.Duration) Option {
	if maxDelay < 0 {
		maxDelay = 0
	}
	return func(configs *retryConfigs) {
		configs.maxDelay = maxDelay
	}
}

// DelayStep sets delay step, default is 0s
func DelayStep(delayStep time.Duration) Option {
	if delayStep < 0 {
		delayStep = 0
	}
	return func(configs *retryConfigs) {
		configs.delayStep = delayStep
	}
}

// ShouldRetry controls whether a retry should be attempted
func ShouldRetry(fn func(err error) bool) Option {
	if fn == nil {
		return noneOption
	}
	return func(configs *retryConfigs) {
		configs.shouldRetry = fn
	}
}

// OnRetry sets callback on every retry
func OnRetry(fn func(i int, err error)) Option {
	if fn == nil {
		return noneOption
	}
	return func(configs *retryConfigs) {
		configs.onRetry = fn
	}
}
