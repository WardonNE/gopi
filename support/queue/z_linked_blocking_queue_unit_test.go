package queue

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLinkedBlockingQueue_MarshalJSON(t *testing.T) {
	q := NewLinkedBlockingQueue[*int](5)
	q.Enqueue(intptr(1))
	q.Enqueue(intptr(2))
	q.Enqueue(intptr(3))
	bytes, err := json.Marshal(q)
	assert.Nil(t, err)
	assert.JSONEq(t, `[1,2,3]`, string(bytes))
	q.Enqueue(intptr(4))
	q.Enqueue(intptr(5))
	q.Dequeue()
	q.Dequeue()
	q.Enqueue(intptr(6))
	q.Enqueue(intptr(7))
	q.Dequeue()
	bytes, err = json.Marshal(q)
	assert.Nil(t, err)
	assert.JSONEq(t, `[4,5,6,7]`, string(bytes))
}

func TestLinkedBlockingQueue_UnmarshalJSON(t *testing.T) {
	q := NewLinkedBlockingQueue[*int](5)
	err := json.Unmarshal([]byte(`[1,2,3]`), q)
	assert.Nil(t, err)
	q2 := NewLinkedBlockingQueue[*int](5)
	q2.Enqueue(intptr(1))
	q2.Enqueue(intptr(2))
	q2.Enqueue(intptr(3))
	assert.Equal(t, q2.items, q.items)
}

func TestLinkedBlockingQueue_ToArray(t *testing.T) {
	q := NewLinkedBlockingQueue[*int](5)
	q.Enqueue(intptr(1))
	q.Enqueue(intptr(2))
	q.Enqueue(intptr(3))
	assert.Equal(t, []*int{intptr(1), intptr(2), intptr(3)}, q.ToArray())
}

func TestLinkedBlockingQueue_FromArray(t *testing.T) {
	q := NewLinkedBlockingQueue[*int](5)
	q.FromArray([]*int{intptr(1), intptr(2), intptr(3)})
	q2 := NewLinkedBlockingQueue[*int](5)
	q2.Enqueue(intptr(1))
	q2.Enqueue(intptr(2))
	q2.Enqueue(intptr(3))
	assert.Equal(t, q2.items, q.items)
}

func TestLinkedBlockingQueue_Count(t *testing.T) {
	q := NewLinkedBlockingQueue[*int](5)
	q.Enqueue(intptr(1))
	q.Enqueue(intptr(2))
	q.Enqueue(intptr(3))
	assert.Equal(t, 3, q.Count())
}

func TestLinkedBlockingQueue_IsEmpty(t *testing.T) {
	q := NewLinkedBlockingQueue[*int](5)
	assert.True(t, q.IsEmpty())
	q.Enqueue(intptr(1))
	q.Enqueue(intptr(2))
	q.Enqueue(intptr(3))
	assert.False(t, q.IsEmpty())
}

func TestLinkedBlockingQueue_IsNotEmpty(t *testing.T) {
	q := NewLinkedBlockingQueue[*int](5)
	assert.False(t, q.IsNotEmpty())
	q.Enqueue(intptr(1))
	q.Enqueue(intptr(2))
	q.Enqueue(intptr(3))
	assert.True(t, q.IsNotEmpty())
}

func TestLinkedBlockingQueue_Clear(t *testing.T) {
	q := NewLinkedBlockingQueue[*int](5)
	assert.False(t, q.IsNotEmpty())
	q.Enqueue(intptr(1))
	q.Enqueue(intptr(2))
	q.Enqueue(intptr(3))
	assert.True(t, q.IsNotEmpty())
	q.Clear()
	assert.False(t, q.IsNotEmpty())
}

func TestLinkedBlockingQueue_Peek(t *testing.T) {
	q := NewLinkedBlockingQueue[*int](5)
	assert.Zero(t, q.Peek())
	q.Enqueue(intptr(1))
	q.Enqueue(intptr(2))
	q.Enqueue(intptr(3))
	assert.Equal(t, intptr(1), q.Peek())
}

func TestLinkedBlockingQueue_Enqueue(t *testing.T) {
	q := NewLinkedBlockingQueue[*int](5)
	assert.True(t, q.Enqueue(intptr(1)))
	assert.True(t, q.Enqueue(intptr(2)))
	assert.True(t, q.Enqueue(intptr(3)))
	assert.True(t, q.Enqueue(intptr(4)))
	assert.True(t, q.Enqueue(intptr(5)))
	assert.Equal(t, 5, q.Count())
	assert.False(t, q.Enqueue(intptr(6)))
	assert.False(t, q.Enqueue(intptr(7)))
	q.Dequeue()
	assert.True(t, q.Enqueue(intptr(8)))
	values := q.ToArray()
	assert.Equal(t, []*int{intptr(2), intptr(3), intptr(4), intptr(5), intptr(8)}, values)
}

func TestLinkedBlockingQueue_Dequeue(t *testing.T) {
	q := NewLinkedBlockingQueue[*int](5)
	value, ok := q.Dequeue()
	assert.False(t, ok)
	assert.Zero(t, value)

	assert.True(t, q.Enqueue(intptr(1)))
	assert.True(t, q.Enqueue(intptr(2)))
	assert.True(t, q.Enqueue(intptr(3)))
	assert.True(t, q.Enqueue(intptr(4)))
	assert.True(t, q.Enqueue(intptr(5)))

	value, ok = q.Dequeue()
	assert.Equal(t, intptr(1), value)
	assert.True(t, ok)

	value, ok = q.Dequeue()
	assert.Equal(t, intptr(2), value)
	assert.True(t, ok)

	value, ok = q.Dequeue()
	assert.Equal(t, intptr(3), value)
	assert.True(t, ok)

	value, ok = q.Dequeue()
	assert.Equal(t, intptr(4), value)
	assert.True(t, ok)

	value, ok = q.Dequeue()
	assert.Equal(t, intptr(5), value)
	assert.True(t, ok)
}

func TestLinkedBlockingQueue_EnqueueWithBlock(t *testing.T) {
	q := NewLinkedBlockingQueue[*int](5)
	assert.True(t, q.EnqueueWithBlock(intptr(1)))
	assert.True(t, q.EnqueueWithBlock(intptr(2)))
	assert.True(t, q.EnqueueWithBlock(intptr(3)))
	assert.True(t, q.EnqueueWithBlock(intptr(4)))
	assert.True(t, q.EnqueueWithBlock(intptr(5)))
	assert.Equal(t, 5, q.Count())
	start := time.Now()
	go func() {
		time.Sleep(3 * time.Second)
		q.Dequeue()
	}()
	assert.True(t, q.EnqueueWithBlock(intptr(6)))
	assert.Equal(t, 3*time.Second, time.Duration(time.Since(start).Seconds())*time.Second)
	values := q.ToArray()
	assert.Equal(t, []*int{intptr(2), intptr(3), intptr(4), intptr(5), intptr(6)}, values)
}

func TestLinkedBlockingQueue_DequeueWithBlock(t *testing.T) {
	q := NewLinkedBlockingQueue[*int](5)
	assert.True(t, q.Enqueue(intptr(1)))
	value, ok := q.DequeueWithBlock()
	assert.Equal(t, intptr(1), value)
	assert.True(t, ok)
	start := time.Now()
	go func() {
		time.Sleep(3 * time.Second)
		assert.True(t, q.Enqueue(intptr(2)))
	}()
	value, ok = q.DequeueWithBlock()
	assert.Equal(t, 3*time.Second, time.Duration(time.Since(start).Seconds())*time.Second)
	assert.True(t, ok)
	assert.Equal(t, intptr(2), value)
}

func TestLinkedBlockingQueue_EnqueueWithTimeout(t *testing.T) {
	q := NewLinkedBlockingQueue[*int](5)
	assert.True(t, q.EnqueueWithTimeout(intptr(1), 3*time.Second))
	assert.True(t, q.EnqueueWithTimeout(intptr(2), 3*time.Second))
	assert.True(t, q.EnqueueWithTimeout(intptr(3), 3*time.Second))
	assert.True(t, q.EnqueueWithTimeout(intptr(4), 3*time.Second))
	assert.True(t, q.EnqueueWithTimeout(intptr(5), 3*time.Second))
	assert.Equal(t, 5, q.Count())
	start := time.Now()
	go func() {
		time.Sleep(2 * time.Second)
		q.Dequeue()
	}()
	assert.True(t, q.EnqueueWithTimeout(intptr(6), 3*time.Second))
	assert.Equal(t, 2*time.Second, time.Duration(time.Since(start).Seconds())*time.Second)
	assert.False(t, q.EnqueueWithTimeout(intptr(7), 1*time.Second))
	values := q.ToArray()
	assert.Equal(t, []*int{intptr(2), intptr(3), intptr(4), intptr(5), intptr(6)}, values)
}

func TestLinkedBlockingQueue_DequeueWithTimeout(t *testing.T) {
	q := NewLinkedBlockingQueue[*int](5)
	assert.True(t, q.Enqueue(intptr(1)))
	value, ok := q.DequeueWithTimeout(3 * time.Second)
	assert.Equal(t, intptr(1), value)
	assert.True(t, ok)
	start := time.Now()
	go func() {
		time.Sleep(2 * time.Second)
		assert.True(t, q.Enqueue(intptr(2)))
	}()
	value, ok = q.DequeueWithTimeout(3 * time.Second)
	assert.Equal(t, 2*time.Second, time.Duration(time.Since(start).Seconds())*time.Second)
	assert.True(t, ok)
	assert.Equal(t, intptr(2), value)
	value, ok = q.DequeueWithTimeout(1 * time.Second)
	assert.False(t, ok)
	assert.Zero(t, value)
}
