package queue

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedQueue_MarshalJSON(t *testing.T) {
	q := NewLinkedQueue[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)
	bytes, err := json.Marshal(q)
	assert.Nil(t, err)
	assert.JSONEq(t, `[1,2,3,4,5]`, string(bytes))
}

func TestLinkedQueue_UnmarshalJSON(t *testing.T) {
	q := NewLinkedQueue[int]()
	err := json.Unmarshal([]byte(`[1,2,3,4,5]`), q)
	assert.Nil(t, err)
	assert.Equal(t, []int{1, 2, 3, 4, 5}, q.ToArray())
}

func TestLinkedQueue_ToArray(t *testing.T) {
	q := NewLinkedQueue[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)
	assert.Equal(t, []int{1, 2, 3, 4, 5}, q.ToArray())
}

func TestLinkedQueue_FromArray(t *testing.T) {
	q := NewLinkedQueue[int]()
	q.FromArray([]int{1, 2, 3, 4, 5})
	assert.Equal(t, []int{1, 2, 3, 4, 5}, q.ToArray())
}

func TestLinkedQueue_Count(t *testing.T) {
	q := NewLinkedQueue[int]()
	q.FromArray([]int{1, 2, 3, 4, 5})
	assert.Equal(t, 5, q.Count())
}

func TestLinkedQueue_IsEmpty(t *testing.T) {
	q := NewLinkedQueue[int]()
	assert.True(t, q.IsEmpty())
	q.Enqueue(1)
	assert.False(t, q.IsEmpty())
}

func TestLinkedQueue_IsNotEmpty(t *testing.T) {
	q := NewLinkedQueue[int]()
	assert.False(t, q.IsNotEmpty())
	q.Enqueue(1)
	assert.True(t, q.IsNotEmpty())
}

func TestLinkedQueue_Clear(t *testing.T) {
	q := NewLinkedQueue[int]()
	q.FromArray([]int{1, 2, 3, 4, 5})
	assert.False(t, q.IsEmpty())
	q.Clear()
	assert.True(t, q.IsEmpty())
}

func TestLinkedQueue_Peek(t *testing.T) {
	q := NewLinkedQueue[int]()
	assert.Zero(t, q.Peek())
	q.FromArray([]int{1, 2, 3, 4, 5})
	assert.Equal(t, 1, q.Peek())
	assert.Equal(t, 5, q.Count())
}

func TestLinkedQueue_Enqueue(t *testing.T) {
	q := NewLinkedQueue[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)
	assert.Equal(t, []int{1, 2, 3, 4, 5}, q.ToArray())
}

func TestLinkedQueue_Dequeue(t *testing.T) {
	q := NewLinkedQueue[int]()
	value, ok := q.Dequeue()
	assert.Zero(t, value)
	assert.False(t, ok)
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)
	value, ok = q.Dequeue()
	assert.Equal(t, 1, value)
	assert.True(t, ok)
	value, ok = q.Dequeue()
	assert.Equal(t, 2, value)
	assert.True(t, ok)
	value, ok = q.Dequeue()
	assert.Equal(t, 3, value)
	assert.True(t, ok)
	value, ok = q.Dequeue()
	assert.Equal(t, 4, value)
	assert.True(t, ok)
	value, ok = q.Dequeue()
	assert.Equal(t, 5, value)
	assert.True(t, ok)
	value, ok = q.Dequeue()
	assert.Zero(t, value)
	assert.False(t, ok)
}
