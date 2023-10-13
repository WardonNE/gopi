package utils

import (
	"fmt"
	"time"
)

func RunOut(callback func(), d time.Duration) error {
	done := make(chan struct{})
	go func() {
		callback()
		done <- struct{}{}
	}()
	select {
	case <-time.After(d):
		return fmt.Errorf("timeout")
	case <-done:
		return nil
	}
}

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

func RunOutWithError(callback func() error, d time.Duration) error {
	done := make(chan error)
	go func() {
		err := callback()
		done <- err
	}()
	select {
	case <-time.After(d):
		return fmt.Errorf("timeout")
	case err := <-done:
		return err
	}
}

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
