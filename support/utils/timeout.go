package utils

import (
	"errors"
	"time"
)

// RunOut runs callback with time limited
//
// It returns an error when timeout
func RunOut(callback func(), d time.Duration) error {
	return RunOutCause(callback, d, errors.New("timeout"))
}

// RunOutCause runs callback with time limited
//
// It returns custom error when timeout
func RunOutCause(callback func(), d time.Duration, err error) error {
	done := make(chan struct{})
	go func() {
		callback()
		done <- struct{}{}
	}()
	select {
	case <-time.After(d):
		return err
	case <-done:
		return nil
	}
}

// RunOutWithError runs callback with time limited
//
// When timeout, it returns an error.
//
// When callback returns error, it returns the error which callback returned.
func RunOutWithError(callback func() error, d time.Duration) error {
	return RunOutWithErrorCause(callback, d, errors.New("timeout"))
}

// RunOutWithErrorCause runs callback with time limited
//
// When timeout, it returns a custom error.
//
// When callback returns error, it returns the error which callback returned.
func RunOutWithErrorCause(callback func() error, d time.Duration, err error) error {
	done := make(chan error)
	go func() {
		err := callback()
		done <- err
	}()
	select {
	case <-time.After(d):
		return err
	case err := <-done:
		return err
	}
}
