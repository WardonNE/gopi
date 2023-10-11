package queue

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDelayQueue_Count(t *testing.T) {
	q := NewDelayQueue[int]()
	q.Enqueue(NewDelayEntry[int](1, time.Second))
	q.Enqueue(NewDelayEntry[int](2, 2*time.Second))
	q.Enqueue(NewDelayEntry[int](3, 3*time.Second))
	assert.Equal(t, 3, q.Count())
}

func TestDelayQueue_IsEmpty(t *testing.T) {
	q := NewDelayQueue[int]()
	assert.True(t, q.IsEmpty())
	q.Enqueue(NewDelayEntry[int](1, time.Second))
	q.Enqueue(NewDelayEntry[int](2, 2*time.Second))
	q.Enqueue(NewDelayEntry[int](3, 3*time.Second))
	assert.False(t, q.IsEmpty())
}

func TestDelayQueue_IsNotEmpty(t *testing.T) {
	q := NewDelayQueue[int]()
	assert.False(t, q.IsNotEmpty())
	q.Enqueue(NewDelayEntry[int](1, time.Second))
	q.Enqueue(NewDelayEntry[int](2, 2*time.Second))
	q.Enqueue(NewDelayEntry[int](3, 3*time.Second))
	assert.True(t, q.IsNotEmpty())
}

func TestDelayQueue_Clear(t *testing.T) {
	q := NewDelayQueue[int]()
	assert.False(t, q.IsNotEmpty())
	q.Enqueue(NewDelayEntry[int](1, time.Second))
	q.Enqueue(NewDelayEntry[int](2, 2*time.Second))
	q.Enqueue(NewDelayEntry[int](3, 3*time.Second))
	assert.True(t, q.IsNotEmpty())
	q.Clear()
	assert.False(t, q.IsNotEmpty())
}

func TestDelayQueue_Peek(t *testing.T) {
	q := NewDelayQueue[int]()
	assert.Zero(t, q.Peek())
	q.Enqueue(NewDelayEntry[int](1, time.Second))
	q.Enqueue(NewDelayEntry[int](2, 2*time.Second))
	q.Enqueue(NewDelayEntry[int](3, 3*time.Second))
	assert.Equal(t, 1, q.Peek().Value())
}

func TestDelayQueue_Enqueue(t *testing.T) {
	q := NewDelayQueue[int]()
	q.Enqueue(NewDelayEntry[int](1, time.Second))
	q.Enqueue(NewDelayEntry[int](2, 2*time.Second))
	q.Enqueue(NewDelayEntry[int](3, 3*time.Second))
	values := []int{}
	entries := q.ToArray()
	for _, entry := range entries {
		values = append(values, entry.Value())
	}
	assert.ElementsMatch(t, []int{1, 2, 3}, values)
}

func TestDelayQueue_Dequeue(t *testing.T) {
	q := NewDelayQueue[int]()
	value, ok := q.Dequeue()
	assert.Zero(t, value)
	assert.False(t, ok)
	q.Enqueue(NewDelayEntry[int](1, time.Second))
	value, ok = q.Dequeue()
	assert.Zero(t, value)
	assert.False(t, ok)
	time.Sleep(time.Second)
	value, ok = q.Dequeue()
	assert.Equal(t, 1, value.Value())
	assert.True(t, ok)
}

func TestDelayQueue_EnqueueWithBlock(t *testing.T) {
	q := NewDelayQueue[int]()
	q.EnqueueWithBlock(NewDelayEntry[int](1, time.Second))
	q.EnqueueWithBlock(NewDelayEntry[int](2, 2*time.Second))
	q.EnqueueWithBlock(NewDelayEntry[int](3, 3*time.Second))
	values := []int{}
	entries := q.ToArray()
	for _, entry := range entries {
		values = append(values, entry.Value())
	}
	assert.ElementsMatch(t, []int{1, 2, 3}, values)
}

func TestDelayQueue_DequeueWithBlock(t *testing.T) {
	q := NewDelayQueue[int]()
	q.Enqueue(NewDelayEntry[int](1, time.Second))
	start := time.Now()
	value, ok := q.DequeueWithBlock()
	assert.Equal(t, time.Second, time.Duration(time.Since(start).Seconds())*time.Second)
	assert.Equal(t, 1, value.Value())
	assert.True(t, ok)
}

func TestDelayQueue_EnqueueWithTimeout(t *testing.T) {
	q := NewDelayQueue[int]()
	q.EnqueueWithTimeout(NewDelayEntry[int](1, time.Second), time.Second)
	q.EnqueueWithTimeout(NewDelayEntry[int](2, 2*time.Second), time.Second)
	q.EnqueueWithTimeout(NewDelayEntry[int](3, 3*time.Second), time.Second)
	values := []int{}
	entries := q.ToArray()
	for _, entry := range entries {
		values = append(values, entry.Value())
	}
	assert.ElementsMatch(t, []int{1, 2, 3}, values)
}

func TestDelayQueue_DequeueWithTimeout(t *testing.T) {
	q := NewDelayQueue[int]()
	q.Enqueue(NewDelayEntry[int](1, 3*time.Second))
	start := time.Now()
	value, ok := q.DequeueWithTimeout(1 * time.Second)
	assert.Equal(t, time.Second, time.Duration(time.Since(start).Seconds())*time.Second)
	assert.Zero(t, value)
	assert.False(t, ok)

	value, ok = q.DequeueWithTimeout(3 * time.Second)
	assert.Equal(t, 3*time.Second, time.Duration(time.Since(start).Seconds())*time.Second)
	assert.Equal(t, 1, value.Value())
	assert.True(t, ok)
}
