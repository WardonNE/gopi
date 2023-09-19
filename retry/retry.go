package retry

import (
	"time"

	"github.com/wardonne/gopi/utils"
)

const (
	defaultAttempts                = 3
	defaultDelay     time.Duration = time.Second
	defaultMaxDelay  time.Duration = time.Minute
	defaultDelayStep time.Duration = 0
)

type Retry struct {
	fn        func() error
	attempts  int
	delay     time.Duration
	maxDelay  time.Duration
	delayStep time.Duration
	retryIf   func(err error) bool
	onRetry   func(i int, err error)
	stop      chan struct{}
	done      chan struct{}
}

func Do(fn func() error, options ...Option) error {
	retry := new(Retry)
	retry.fn = fn
	retry.attempts = defaultAttempts
	retry.delay = defaultDelay
	retry.maxDelay = defaultMaxDelay
	retry.delayStep = defaultDelayStep
	retry.retryIf = func(err error) bool { return err != nil }
	retry.onRetry = func(i int, err error) {}
	retry.stop = make(chan struct{})
	retry.done = make(chan struct{})
	for _, option := range options {
		option(retry)
	}
	return retry.Start()
}

func (r *Retry) Start() error {
	var err error
	go func() {
		attempts := 1
		for {
			err = r.fn()
			if !r.retryIf(err) {
				r.done <- struct{}{}
				return
			}
			if r.attempts > 0 && attempts > int(r.attempts)-1 {
				r.done <- struct{}{}
				return
			}
			attempts++
			r.onRetry(attempts, err)
			delay := utils.Min(r.delay+r.delayStep*time.Duration(attempts), r.maxDelay)
			time.Sleep(delay)
		}
	}()
	select {
	case <-r.done:
		return err
	case <-r.stop:
		return err
	}
}

func (r *Retry) Stop() {
	r.stop <- struct{}{}
}
