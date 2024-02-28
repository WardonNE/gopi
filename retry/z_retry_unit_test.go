package retry

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRetry_SuccessAfterSomeFailure(t *testing.T) {
	retrier := New(Attempts(3), Delay(0))
	var result int
	var attempts int
	assert.Nil(t, retrier.Do(func() error {
		attempts++
		if attempts < 2 {
			return errors.New("err")
		}
		result = 100
		return nil
	}))
	assert.Equal(t, 100, result)
	assert.Equal(t, 2, attempts)
}

func TestRetry_AllFailed(t *testing.T) {
	retrier := New(Attempts(3), Delay(0))
	result := 0
	attempts := 0
	assert.Equal(t, "err", retrier.Do(func() error {
		attempts++
		return errors.New("err")
	}).Error())
	assert.Equal(t, 0, result)
	assert.Equal(t, 3, attempts)
}

func TestRetry_CustomShouldRetry(t *testing.T) {
	retrier := New(Attempts(3), Delay(0), ShouldRetry(func(err error) bool {
		return err.Error() == "shouldretry"
	}))
	result := 0
	attempts := 0
	assert.Equal(t, "shouldnotretry", retrier.Do(func() error {
		attempts++
		if attempts < 2 {
			return errors.New("shouldretry")
		}
		return errors.New("shouldnotretry")
	}).Error())
	assert.Equal(t, 0, result)
	assert.Equal(t, 2, attempts)
}

func TestRetry_CustomDelay(t *testing.T) {
	retrier := New(Attempts(3), Delay(time.Second))
	result := 0
	attempts := 0
	start := time.Now()
	assert.Equal(t, "err", retrier.Do(func() error {
		attempts++
		return errors.New("err")
	}).Error())
	assert.Equal(t, 2*time.Second, time.Duration(time.Since(start).Seconds())*time.Second)
	assert.Equal(t, 0, result)
	assert.Equal(t, 3, attempts)
}

func TestRetry_CustomDelayStep(t *testing.T) {
	retrier := New(Attempts(3), Delay(time.Second), DelayStep(time.Second))
	result := 0
	attempts := 0
	start := time.Now()
	assert.Equal(t, "err", retrier.Do(func() error {
		attempts++
		return errors.New("err")
	}).Error())
	assert.Equal(t, 3*time.Second, time.Duration(time.Since(start).Seconds())*time.Second)
	assert.Equal(t, 0, result)
	assert.Equal(t, 3, attempts)
}

func TestRetry_CustomMaxDelay(t *testing.T) {
	retrier := New(Attempts(3), Delay(time.Second), DelayStep(time.Second), MaxDelay(time.Second))
	result := 0
	attempts := 0
	start := time.Now()
	assert.Equal(t, "err", retrier.Do(func() error {
		attempts++
		return errors.New("err")
	}).Error())
	assert.Equal(t, 2*time.Second, time.Duration(time.Since(start).Seconds())*time.Second)
	assert.Equal(t, 0, result)
	assert.Equal(t, 3, attempts)
}

func TestRetry_CustomContextWithTimeout(t *testing.T) {
	start := time.Now()
	result := 0
	attempts := 0
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	retrier := New(Context(ctx), Attempts(3), Delay(2*time.Second), DelayStep(time.Second))
	err := retrier.Do(func() error {
		attempts++
		return errors.New("err")
	})
	assert.Equal(t, "context deadline exceeded", err.Error())
	assert.Equal(t, time.Second, time.Duration(time.Since(start).Seconds())*time.Second)
	assert.Equal(t, 0, result)
	assert.Equal(t, 1, attempts)
}

func TestRetry_CustomContextWithCancel(t *testing.T) {
	t.Run("cancel before do", func(t *testing.T) {
		ctx, cancel := context.WithCancelCause(context.Background())
		defer cancel(errors.New("defer cancel"))
		retrier := New(Context(ctx), Attempts(3), Delay(2*time.Second), DelayStep(time.Second))
		cancel(errors.New("cancel"))
		err := retrier.Do(func() error {
			return errors.New("err")
		})
		assert.Equal(t, "context canceled", err.Error())
		assert.Equal(t, "cancel", context.Cause(ctx).Error())
	})

	t.Run("cancel in func", func(t *testing.T) {
		ctx, cancel := context.WithCancelCause(context.Background())
		defer cancel(errors.New("defer cancel"))
		retrier := New(Context(ctx), Attempts(3), Delay(2*time.Second), DelayStep(time.Second))
		err := retrier.Do(func() error {
			cancel(errors.New("do cancel"))
			return errors.New("err")
		})
		assert.Equal(t, "context canceled", err.Error())
		assert.Equal(t, "do cancel", context.Cause(ctx).Error())
	})
}
