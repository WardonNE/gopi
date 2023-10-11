package list

import (
	"encoding/json"
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wardonne/gopi/support/compare"
)

func TestLinkedList_MarshalJSON(t *testing.T) {
	list := NewLinkedList[int](1, 2, 3, 4, 5)
	bytes, err := json.Marshal(list)
	assert.Nil(t, err)
	assert.JSONEq(t, `[1,2,3,4,5]`, string(bytes))
}

func TestLinkedList_UnmarshalJSON(t *testing.T) {
	list := NewLinkedList[int]()
	err := json.Unmarshal([]byte(`[1,2,3,4,5]`), list)
	assert.Nil(t, err)
	assert.Equal(t, NewLinkedList[int](1, 2, 3, 4, 5), list)
}

func TestLinkedList_ToArray(t *testing.T) {
	list := NewLinkedList[int](1, 2, 3, 4, 5)
	assert.Equal(t, []int{1, 2, 3, 4, 5}, list.ToArray())
}

func TestLinkedList_FromArray(t *testing.T) {
	list := NewLinkedList[int]()
	list.FromArray([]int{1, 2, 3, 4, 5})
	assert.Equal(t, NewLinkedList[int](1, 2, 3, 4, 5), list)
}

func TestLinkedList_String(t *testing.T) {
	list := NewLinkedList[int](1, 2, 3, 4, 5)
	assert.Equal(t, fmt.Sprintf("%v", list.ToArray()), list.String())
}

func TestLinkedList_Sort(t *testing.T) {
	s := []int{2, 3, 1, 8, 3, 5, 3, 2, 91, 12}
	list := NewLinkedList[int](s...)
	list.Sort(compare.NewNatureComparator[int](false))
	sort.Ints(s)
	assert.Equal(t, s, list.ToArray())
}

func TestLinkedList_Clone(t *testing.T) {
	list := NewLinkedList[int](1, 2, 3, 4, 5)
	clonedList := list.Clone()
	assert.Equal(t, list, clonedList)
}

func TestLinkedList_Copy(t *testing.T) {
	list := NewLinkedList[int](1, 2, 3, 4, 5)
	copiedList := list.Copy()
	assert.Equal(t, list, copiedList)
}

func TestLinkedList_Count(t *testing.T) {
	list := NewLinkedList[int](1, 2, 3, 4, 5)
	assert.Equal(t, 5, list.Count())
}

func TestLinkedList_IsEmpty(t *testing.T) {
	list := NewLinkedList[int]()
	assert.True(t, list.IsEmpty())
}

func TestLinkedList_IsNotEmpty(t *testing.T) {
	list := NewLinkedList[int]()
	assert.False(t, list.IsNotEmpty())
}

func TestLinkedList_Get(t *testing.T) {
	list := NewLinkedList[int](1, 2, 3, 4, 5)
	assert.PanicsWithValue(t, ErrIndexOutOfRange, func() { list.Get(-1) })
	assert.Equal(t, 1, list.Get(0))
}

func TestLinkedList_Pop(t *testing.T) {
	list := NewLinkedList[int]()
	assert.Zero(t, list.Pop())
	list = NewLinkedList[int](1, 2, 3, 4, 5)
	assert.Equal(t, 5, list.Pop())
}

func TestLinkedList_Shift(t *testing.T) {
	list := NewLinkedList[int]()
	assert.Zero(t, list.Shift())
	list = NewLinkedList[int](1, 2, 3, 4, 5)
	assert.Equal(t, 1, list.Shift())
}

func TestLinkedList_Contains(t *testing.T) {
	list := NewLinkedList[int](1, 2, 3, 4, 5)
	assert.True(t, list.Contains(func(value int) bool {
		return value == 1
	}))
	assert.False(t, list.Contains(func(value int) bool {
		return value == 0
	}))
}

func TestLinkedList_IndexOf(t *testing.T) {
	assert.Equal(t, -1, NewLinkedList[int]().IndexOf(func(value int) bool { return value == 1 }))
	list := NewLinkedList[int](1, 2, 2, 3, 4, 0, 5)
	assert.Equal(t, -1, list.IndexOf(func(value int) bool {
		return value == -1
	}))
	assert.Equal(t, 5, list.IndexOf(func(value int) bool {
		return value == 0
	}))
}

func TestLinkedList_LastIndexOf(t *testing.T) {
	assert.Equal(t, -1, NewLinkedList[int]().LastIndexOf(func(value int) bool { return value == 1 }))
	list := NewLinkedList[int](1, 2, 0, 2, 3, 0, 4, 5)
	assert.Equal(t, -1, list.LastIndexOf(func(value int) bool { return value == -1 }))
	assert.Equal(t, 5, list.LastIndexOf(func(value int) bool {
		return value == 0
	}))
}

func TestLinkedList_SubList(t *testing.T) {
	list := NewLinkedList[int](1, 2, 3, 4, 5)
	assert.PanicsWithValue(t, ErrIndexOutOfRange, func() { list.SubList(-1, 2) })
	assert.Equal(t, NewLinkedList[int](2, 3), list.SubList(1, 3))
	assert.Equal(t, NewLinkedList[int](2, 3), list.SubList(3, 1))
}

func TestLinkedList_SubLinkedList(t *testing.T) {
	list := NewLinkedList[int](1, 2, 3, 4, 5)
	subList := list.SubLinkedList(1, 3)
	assert.Equal(t, NewLinkedList[int](2, 3), subList)
}

func TestLinkedList_Set(t *testing.T) {
	list := NewLinkedList[int](1, 2, 3, 4, 5)
	list.Set(1, 5)
	assert.Equal(t, NewLinkedList[int](1, 5, 3, 4, 5), list)
	assert.PanicsWithValue(t, ErrIndexOutOfRange, func() { list.Set(-1, 5) })
}

func TestLinkedList_Add(t *testing.T) {
	list := NewLinkedList[int]()
	list.Add(1)
	list.Add(2)
	list.Add(3)
	assert.Equal(t, NewLinkedList[int](1, 2, 3), list)
}

func TestLinkedList_AddAll(t *testing.T) {
	list := NewLinkedList[int]()
	list.AddAll(1, 2, 3)
	assert.Equal(t, NewLinkedList[int](1, 2, 3), list)
}

func TestLinkedList_Push(t *testing.T) {
	list := NewLinkedList[int]()
	list.Push(1)
	list.Push(2)
	list.Push(3)
	assert.Equal(t, NewLinkedList[int](1, 2, 3), list)
}

func TestLinkedList_PushAll(t *testing.T) {
	list := NewLinkedList[int]()
	list.PushAll(1, 2, 3)
	assert.Equal(t, NewLinkedList[int](1, 2, 3), list)
}

func TestLinkedList_Unshift(t *testing.T) {
	list := NewLinkedList[int]()
	list.Unshift(1)
	list.Unshift(2)
	list.Unshift(3)
	assert.Equal(t, NewLinkedList[int](3, 2, 1), list)
}

func TestLinkedList_UnshiftAll(t *testing.T) {
	list := NewLinkedList[int]()
	list.UnshiftAll(1, 2, 3)
	assert.Equal(t, NewLinkedList[int](3, 2, 1), list)
}

func TestLinkedList_InsertBefore(t *testing.T) {
	list := NewLinkedList[int](1, 2, 3, 4, 5)
	list.InsertBefore(1, 5)
	assert.Equal(t, NewLinkedList[int](1, 5, 2, 3, 4, 5), list)
	assert.PanicsWithValue(t, ErrIndexOutOfRange, func() { list.InsertBefore(-1, 1) })
}

func TestLinkedList_InsertAfter(t *testing.T) {
	list := NewLinkedList[int](1, 2, 3, 4, 5)
	list.InsertAfter(1, 5)
	assert.Equal(t, NewLinkedList[int](1, 2, 5, 3, 4, 5), list)
	assert.PanicsWithValue(t, ErrIndexOutOfRange, func() { list.InsertAfter(-1, 1) })
}

func TestLinkedList_RemoveAt(t *testing.T) {
	list := NewLinkedList[int](1, 2, 3, 4, 5)
	list.RemoveAt(1)
	assert.Equal(t, NewLinkedList[int](1, 3, 4, 5), list)
	assert.PanicsWithValue(t, ErrIndexOutOfRange, func() {
		NewLinkedList[int]().RemoveAt(-1)
	})
}

func TestLinkedList_Remove(t *testing.T) {
	list := NewLinkedList[int](1, 2, 3, 4, 5)
	list.Remove(func(value int) bool {
		return value == 2 || value == 3
	})
	assert.Equal(t, NewLinkedList[int](1, 4, 5).ToArray(), list.ToArray())
	assert.NotPanics(t, func() {
		NewLinkedList[int]().Remove(func(value int) bool {
			return value == 1
		})
	})
}

func TestLinkedList_Clear(t *testing.T) {
	list := NewLinkedList[int](1, 2, 3, 4, 5)
	list.Clear()
	assert.True(t, list.IsEmpty())
}

func TestLinkedList_Range(t *testing.T) {
	list := NewLinkedList[int](1, 2, 3, 4, 5)
	s1 := []int{}
	list.Range(func(value int) bool { s1 = append(s1, value); return true })
	assert.Equal(t, []int{1, 2, 3, 4, 5}, s1)

	s2 := []int{}
	list.Range(func(value int) bool { s2 = append(s2, value); return value < 3 })
	assert.Equal(t, []int{1, 2, 3}, s2)
}

func TestLinkedList_ReverseRange(t *testing.T) {
	list := NewLinkedList[int](1, 2, 3, 4, 5)
	s1 := []int{}
	list.ReverseRange(func(value int) bool { s1 = append(s1, value); return true })
	assert.Equal(t, []int{5, 4, 3, 2, 1}, s1)

	s2 := []int{}
	list.ReverseRange(func(value int) bool { s2 = append(s2, value); return value > 3 })
	assert.Equal(t, []int{5, 4, 3}, s2)
}

func TestLinkedList_Map(t *testing.T) {
	list := NewLinkedList[int](1, 2, 3, 4, 5)
	list.Map(func(value int) int {
		return value * value
	})
	assert.Equal(t, NewLinkedList[int](1, 4, 9, 16, 25), list)
}
