package queue

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSynchronousQueue_EnqueueWithTimeout(t *testing.T) {
	q := NewSynchronousQueue[int]()
	start := time.Now()
	assert.False(t, q.EnqueueWithTimeout(1, time.Second))
	assert.Equal(t, time.Second, time.Duration(time.Since(start).Seconds())*time.Second)

	go func() {
		q.EnqueueWithTimeout(1, time.Second)
	}()
	value, ok := q.Dequeue()
	assert.Equal(t, 1, value)
	assert.True(t, ok)
}

func TestSynchronousQueue_DequeueWithTimeout(t *testing.T) {
	q := NewSynchronousQueue[int]()
	start := time.Now()
	value, ok := q.DequeueWithTimeout(time.Second)
	assert.Equal(t, time.Second, time.Duration(time.Since(start).Seconds())*time.Second)
	assert.Zero(t, value)
	assert.False(t, ok)

	go func() {
		q.Enqueue(1)
	}()
	time.Sleep(time.Microsecond)
	value, ok = q.DequeueWithTimeout(time.Second)
	assert.Equal(t, 1, value)
	assert.True(t, ok)
}
