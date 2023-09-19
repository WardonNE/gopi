package retry

import "time"

type Option func(retry *Retry)

func Attempts(attempts int) Option {
	return func(retry *Retry) {
		retry.attempts = attempts
	}
}

func Delay(delay time.Duration) Option {
	return func(retry *Retry) {
		retry.delay = delay
	}
}

func MaxDelay(maxDelay time.Duration) Option {
	return func(retry *Retry) {
		retry.maxDelay = maxDelay
	}
}

func DelayStep(delayStep time.Duration) Option {
	return func(retry *Retry) {
		retry.delayStep = delayStep
	}
}

func RetryIf(fn func(err error) bool) Option {
	return func(retry *Retry) {
		retry.retryIf = fn
	}
}

func OnRetry(fn func(i int, err error)) Option {
	return func(retry *Retry) {
		retry.onRetry = fn
	}
}
