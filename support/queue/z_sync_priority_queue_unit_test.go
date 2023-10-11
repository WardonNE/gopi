package queue

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSyncPriorityQueue_MarshalJSON(t *testing.T) {
	q := NewSyncPriorityQueue[int](comparator())
	q.Enqueue(1)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(2)
	q.Enqueue(5)
	bytes, err := json.Marshal(q)
	assert.Nil(t, err)
	values := []int{}
	err = json.Unmarshal(bytes, &values)
	assert.Nil(t, err)
	assert.ElementsMatch(t, []int{1, 2, 3, 4, 5}, values)
}

func TestSyncPriorityQueue_UnmarshalJSON(t *testing.T) {
	q := NewSyncPriorityQueue[int](comparator())
	err := json.Unmarshal([]byte(`[1,3,4,2,5]`), q)
	assert.Nil(t, err)
	q2 := NewSyncPriorityQueue[int](comparator())
	q2.Enqueue(1)
	q2.Enqueue(3)
	q2.Enqueue(4)
	q2.Enqueue(2)
	q2.Enqueue(5)
	assert.Equal(t, q2, q)
}

func TestSyncPriorityQueue_ToArray(t *testing.T) {
	q := NewSyncPriorityQueue[int](comparator())
	q.Enqueue(1)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(2)
	q.Enqueue(5)
	assert.ElementsMatch(t, []int{1, 2, 3, 4, 5}, q.ToArray())
}

func TestSyncPriorityQueue_FromArray(t *testing.T) {
	q := NewSyncPriorityQueue[int](comparator())
	q.FromArray([]int{1, 3, 4, 2, 5})
	q2 := NewSyncPriorityQueue[int](comparator())
	q2.Enqueue(1)
	q2.Enqueue(3)
	q2.Enqueue(4)
	q2.Enqueue(2)
	q2.Enqueue(5)
	assert.Equal(t, q2, q)
}

func TestSyncPriorityQueue_Count(t *testing.T) {
	q := NewSyncPriorityQueue[int](comparator())
	q.Enqueue(1)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(2)
	q.Enqueue(5)
	assert.Equal(t, 5, q.Count())
}

func TestSyncPriorityQueue_IsEmpty(t *testing.T) {
	q := NewSyncPriorityQueue[int](comparator())
	assert.True(t, q.IsEmpty())
	q.Enqueue(1)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(2)
	q.Enqueue(5)
	assert.False(t, q.IsEmpty())
}

func TestSyncPriorityQueue_IsNotEmpty(t *testing.T) {
	q := NewSyncPriorityQueue[int](comparator())
	assert.False(t, q.IsNotEmpty())
	q.Enqueue(1)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(2)
	q.Enqueue(5)
	assert.True(t, q.IsNotEmpty())
}

func TestSyncPriorityQueue_Clear(t *testing.T) {
	q := NewSyncPriorityQueue[int](comparator())
	q.Enqueue(1)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(2)
	q.Enqueue(5)
	assert.True(t, q.IsNotEmpty())
	q.Clear()
	assert.False(t, q.IsNotEmpty())
}

func TestSyncPriorityQueue_Peek(t *testing.T) {
	q := NewSyncPriorityQueue[int](comparator())
	assert.Zero(t, q.Peek())
	q.Enqueue(1)
	assert.Equal(t, 1, q.Peek())
	q.Enqueue(3)
	assert.Equal(t, 3, q.Peek())
	q.Enqueue(4)
	assert.Equal(t, 4, q.Peek())
	q.Enqueue(2)
	assert.Equal(t, 4, q.Peek())
	q.Enqueue(5)
	assert.Equal(t, 5, q.Peek())
}

func TestSyncPriorityQueue_Enqueue(t *testing.T) {
	q := NewSyncPriorityQueue[int](comparator())
	q.Enqueue(1)
	assert.Contains(t, q.ToArray(), 1)
	q.Enqueue(3)
	assert.Contains(t, q.ToArray(), 3)
	q.Enqueue(4)
	assert.Contains(t, q.ToArray(), 4)
	q.Enqueue(2)
	assert.Contains(t, q.ToArray(), 2)
	q.Enqueue(5)
	assert.Contains(t, q.ToArray(), 5)
	assert.Equal(t, 5, q.Count())
}

func TestSyncPriorityQueue_Dequeue(t *testing.T) {
	q := NewSyncPriorityQueue[int](comparator())
	value, ok := q.Dequeue()
	assert.Zero(t, value)
	assert.False(t, ok)
	q.Enqueue(1)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(2)
	q.Enqueue(5)
	value, ok = q.Dequeue()
	assert.True(t, ok)
	assert.Equal(t, 5, value)
	assert.Equal(t, 4, q.Count())

	value, ok = q.Dequeue()
	assert.True(t, ok)
	assert.Equal(t, 4, value)
	assert.Equal(t, 3, q.Count())

	value, ok = q.Dequeue()
	assert.True(t, ok)
	assert.Equal(t, 3, value)
	assert.Equal(t, 2, q.Count())

	value, ok = q.Dequeue()
	assert.True(t, ok)
	assert.Equal(t, 2, value)
	assert.Equal(t, 1, q.Count())

	value, ok = q.Dequeue()
	assert.True(t, ok)
	assert.Equal(t, 1, value)
	assert.Equal(t, 0, q.Count())
}
