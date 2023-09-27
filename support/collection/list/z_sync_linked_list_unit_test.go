package list

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wardonne/gopi/support/compare"
)

func TestSyncLinkedList_MarshalJSON(t *testing.T) {
	bytes, err := json.Marshal(NewSyncLinkedList[int](0, 1, 2, 3, 4))
	assert.Nil(t, err)
	assert.JSONEq(t, `[0,1,2,3,4]`, string(bytes))
}

func TestSyncLinkedList_UnmarshalJSON(t *testing.T) {
	list := NewSyncLinkedList[int]()
	err := json.Unmarshal([]byte(`[0,1,2,3,4]`), list)
	assert.Nil(t, err)
	assert.Equal(t, NewSyncLinkedList[int](0, 1, 2, 3, 4), list)
}

func TestSyncLinkedList_ToArray(t *testing.T) {
	list := NewSyncLinkedList[int](0, 1, 2, 3, 4)
	assert.Equal(t, []int{0, 1, 2, 3, 4}, list.ToArray())
}

func TestSyncLinkedList_FromArray(t *testing.T) {
	list := NewSyncLinkedList[int]()
	list.FromArray([]int{0, 1, 2, 3, 4})
	assert.Equal(t, NewSyncLinkedList[int](0, 1, 2, 3, 4), list)
}

func TestSyncLinkedList_String(t *testing.T) {
	list := NewSyncLinkedList[int](0, 1, 2, 3, 4)
	assert.Equal(t, fmt.Sprintf("%v", []int{0, 1, 2, 3, 4}), list.String())
}

func TestSyncLinkedList_Sort(t *testing.T) {
	list := NewSyncLinkedList[int](0, 1, 2, 3, 4)
	list.Sort(compare.NewNatureComparator[int](true))
	assert.Equal(t, NewSyncLinkedList[int](4, 3, 2, 1, 0), list)
}

func TestSyncLinkedList_Count(t *testing.T) {
	list := NewSyncLinkedList[int](0, 1, 2, 3, 4)
	assert.Equal(t, 5, list.Count())
}

func TestSyncLinkedList_IsEmpty(t *testing.T) {
	list := NewSyncLinkedList[int](0, 1, 2, 3, 4)
	assert.False(t, list.IsEmpty())
}

func TestSyncLinkedList_IsNotEmpty(t *testing.T) {
	list := NewSyncLinkedList[int](0, 1, 2, 3, 4)
	assert.True(t, list.IsNotEmpty())
}

func TestSyncLinkedList_Clone(t *testing.T) {
	list := NewSyncLinkedList[int](0, 1, 2, 3, 4)
	clonedList := list.Clone()
	assert.Equal(t, list, clonedList)
}

func TestSyncLinkedList_Copy(t *testing.T) {
	list := NewSyncLinkedList[int](0, 1, 2, 3, 4)
	copiedList := list.Copy()
	assert.Equal(t, list, copiedList)
}

func TestSyncLinkedList_Get(t *testing.T) {
	list := NewSyncLinkedList[int](0, 1, 2, 3, 4)
	assert.Equal(t, 1, list.Get(1))
	assert.PanicsWithValue(t, ErrIndexOutOfRange, func() { list.Get(-1) })
}

func TestSyncLinkedList_Pop(t *testing.T) {
	list := NewSyncLinkedList[int](0, 1, 2, 3, 4)
	assert.Equal(t, 4, list.Pop())
	assert.Zero(t, NewSyncLinkedList[int]().Pop())
}

func TestSyncLinkedList_Shift(t *testing.T) {
	list := NewSyncLinkedList[int](0, 1, 2, 3, 4)
	assert.Equal(t, 0, list.Shift())
	assert.Zero(t, NewSyncLinkedList[int]().Shift())
}

func TestSyncLinkedList_Contains(t *testing.T) {
	list := NewSyncLinkedList[int](0, 1, 2, 3, 4)
	assert.True(t, list.Contains(func(value int) bool {
		return value == 1
	}))
	assert.False(t, list.Contains(func(value int) bool {
		return value == -1
	}))
}

func TestSyncLinkedList_IndexOf(t *testing.T) {
	list := NewSyncLinkedList[int](0, 1, 2, 3, 4)
	assert.Equal(t, -1, list.IndexOf(func(value int) bool {
		return value == -1
	}))
	assert.Equal(t, 1, list.IndexOf(func(value int) bool {
		return value == 1
	}))
}

func TestSyncLinkedList_LastIndexOf(t *testing.T) {
	list := NewSyncLinkedList[int](0, 1, 2, 1, 3, 4)
	assert.Equal(t, -1, list.LastIndexOf(func(value int) bool {
		return value == -1
	}))
	assert.Equal(t, 3, list.LastIndexOf(func(value int) bool {
		return value == 1
	}))
}

func TestSyncLinkedList_SubList(t *testing.T) {
	list := NewSyncLinkedList[int](0, 1, 2, 3, 4)
	assert.Equal(t, NewSyncLinkedList[int](1, 2), list.SubList(1, 3))
}

func TestSyncLinkedList_SubSyncLinkedList(t *testing.T) {
	list := NewSyncLinkedList[int](0, 1, 2, 3, 4)
	assert.Equal(t, NewSyncLinkedList[int](1, 2), list.SubSyncLinkedList(1, 3))
}

func TestSyncLinkedList_Add(t *testing.T) {
	list := NewSyncLinkedList[int]()
	list.Add(0)
	list.Add(1)
	list.Add(2)
	assert.Equal(t, NewSyncLinkedList[int](0, 1, 2), list)
}

func TestSyncLinkedList_AddAll(t *testing.T) {
	list := NewSyncLinkedList[int]()
	list.AddAll(0, 1, 2)
	assert.Equal(t, NewSyncLinkedList[int](0, 1, 2), list)
}

func TestSyncLinkedList_Set(t *testing.T) {
	list := NewSyncLinkedList[int](0, 1, 2, 3, 4)
	list.Set(1, 5)
	assert.Equal(t, NewSyncLinkedList[int](0, 5, 2, 3, 4), list)
	assert.PanicsWithValue(t, ErrIndexOutOfRange, func() { list.Set(-1, 0) })
}

func TestSyncLinkedList_Push(t *testing.T) {
	list := NewSyncLinkedList[int]()
	list.Push(0)
	list.Push(1)
	list.Push(2)
	assert.Equal(t, NewSyncLinkedList[int](0, 1, 2), list)
}

func TestSyncLinkedList_PushAll(t *testing.T) {
	list := NewSyncLinkedList[int]()
	list.PushAll(0, 1, 2)
	assert.Equal(t, NewSyncLinkedList[int](0, 1, 2), list)
}

func TestSyncLinkedList_Unshift(t *testing.T) {
	list := NewSyncLinkedList[int]()
	list.Unshift(0)
	list.Unshift(1)
	list.Unshift(2)
	assert.Equal(t, NewSyncLinkedList[int](2, 1, 0), list)
}

func TestSyncLinkedList_UnshiftAll(t *testing.T) {
	list := NewSyncLinkedList[int]()
	list.UnshiftAll(0, 1, 2)
	assert.Equal(t, NewSyncLinkedList[int](2, 1, 0), list)
}

func TestSyncLinkedList_InsertBefore(t *testing.T) {
	list := NewSyncLinkedList[int](0, 1, 2, 3, 4)
	list.InsertBefore(1, 5)
	assert.Equal(t, NewSyncLinkedList[int](0, 5, 1, 2, 3, 4), list)
}

func TestSyncLinkedList_InsertAfter(t *testing.T) {
	list := NewSyncLinkedList[int](0, 1, 2, 3, 4)
	list.InsertAfter(1, 5)
	assert.Equal(t, NewSyncLinkedList[int](0, 1, 5, 2, 3, 4), list)
}

func TestSyncLinkedList_RemoveAt(t *testing.T) {
	list := NewSyncLinkedList[int](0, 1, 2, 3, 4)
	list.RemoveAt(1)
	assert.Equal(t, NewSyncLinkedList[int](0, 2, 3, 4), list)
}

func TestSyncLinkedList_Remove(t *testing.T) {
	list := NewSyncLinkedList[int](0, 1, 2, 3, 4)
	list.Remove(func(value int) bool { return value == 1 })
	assert.Equal(t, NewSyncLinkedList[int](0, 2, 3, 4), list)
}

func TestSyncLinkedList_Clear(t *testing.T) {
	list := NewSyncLinkedList[int](0, 1, 2, 3, 4)
	list.Clear()
	assert.True(t, list.IsEmpty())
}

func TestSyncLinkedList_Range(t *testing.T) {
	list := NewSyncLinkedList[int](0, 1, 2, 3, 4)
	s1 := []int{}
	list.Range(func(value int) bool { s1 = append(s1, value); return true })
	assert.Equal(t, []int{0, 1, 2, 3, 4}, s1)

	s2 := []int{}
	list.Range(func(value int) bool { s2 = append(s2, value); return value < 3 })
	assert.Equal(t, []int{0, 1, 2, 3}, s2)
}

func TestSyncLinkedList_ReverseRange(t *testing.T) {
	list := NewSyncLinkedList[int](0, 1, 2, 3, 4)
	s1 := []int{}
	list.ReverseRange(func(value int) bool { s1 = append(s1, value); return true })
	assert.Equal(t, []int{4, 3, 2, 1, 0}, s1)

	s2 := []int{}
	list.ReverseRange(func(value int) bool { s2 = append(s2, value); return value > 3 })
	assert.Equal(t, []int{4, 3}, s2)
}

func TestSyncLinkedList_Map(t *testing.T) {
	list := NewSyncLinkedList[int](0, 1, 2, 3, 4)
	list.Map(func(value int) int { return value * value })
	assert.Equal(t, NewSyncLinkedList[int](0, 1, 4, 9, 16), list)
}
