package list

import (
	"encoding/json"
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wardonne/gopi/support/compare"
)

func emptyArrayList(t *testing.T) *ArrayList[int] {
	t.Helper()
	return NewArrayList[int]()
}

func createArrayList(t *testing.T) *ArrayList[int] {
	t.Helper()
	return NewArrayList[int](-9, -56, 70, 77, -28, -27, -7, -32, -31, 98)
}

func TestArrayList_MarshalJSON(t *testing.T) {
	list := createArrayList(t)
	bytes, err := json.Marshal(list)
	assert.Nil(t, err)
	assert.JSONEq(t, `[-9, -56, 70, 77, -28, -27, -7, -32, -31, 98]`, string(bytes))
}

func TestArrayList_UnmarshalJSON(t *testing.T) {
	list := emptyArrayList(t)
	err := json.Unmarshal([]byte(`[-9, -56, 70, 77, -28, -27, -7, -32, -31, 98]`), &list)
	assert.Nil(t, err)
	assert.Equal(t, createArrayList(t), list)
}

func TestArrayList_ToArray(t *testing.T) {
	list := createArrayList(t)
	assert.Equal(t, []int{-9, -56, 70, 77, -28, -27, -7, -32, -31, 98}, list.ToArray())
}

func TestArrayList_FromArray(t *testing.T) {
	list := emptyArrayList(t)
	list.FromArray([]int{-9, -56, 70, 77, -28, -27, -7, -32, -31, 98})
	assert.Equal(t, createArrayList(t), list)
}

func TestArrayList_String(t *testing.T) {
	list := createArrayList(t)
	assert.Equal(t, fmt.Sprintf("%v", list), list.String())
}

func TestArrayList_Clone(t *testing.T) {
	list := createArrayList(t)
	clonedList := list.Clone()
	assert.Equal(t, list, clonedList)
}

func TestArrayList_Copy(t *testing.T) {
	list := createArrayList(t)
	copiedList := list.Copy()
	assert.Equal(t, list, copiedList)
}

func TestArrayList_Sort(t *testing.T) {
	list := createArrayList(t)
	list.Sort(compare.NewNatureComparator[int](false))
	s := []int{-9, -56, 70, 77, -28, -27, -7, -32, -31, 98}
	sort.Ints(s)
	assert.Equal(t, s, list.items)
}

func TestArrayList_Count(t *testing.T) {
	list := createArrayList(t)
	s := []int{-9, -56, 70, 77, -28, -27, -7, -32, -31, 98}
	assert.Equal(t, len(s), list.Count())
}

func TestArrayList_IsEmpty(t *testing.T) {
	list := emptyArrayList(t)
	assert.True(t, list.IsEmpty())
}

func TestArrayList_IsNotEmpty(t *testing.T) {
	list := createArrayList(t)
	assert.True(t, list.IsNotEmpty())
}

func TestArrayList_Get(t *testing.T) {
	list := createArrayList(t)
	t.Run("out-of-index", func(t *testing.T) {
		assert.PanicsWithValue(t, ErrIndexOutOfRange, func() { list.Get(-1) })
	})
	t.Run("non-out-of-index", func(t *testing.T) {
		value := list.Get(0)
		assert.Equal(t, -9, value)
	})
}

func TestArrayList_Pop(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		assert.Zero(t, emptyArrayList(t).Pop())
	})
	t.Run("non-empty", func(t *testing.T) {
		value := createArrayList(t).Pop()
		assert.Equal(t, 98, value)
	})
}

func TestArrayList_Shift(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		assert.Zero(t, emptyArrayList(t).Shift())
	})
	t.Run("non-empty", func(t *testing.T) {
		value := createArrayList(t).Shift()
		assert.Equal(t, -9, value)
	})
}

func TestArrayList_Contains(t *testing.T) {
	list := createArrayList(t)
	t.Run("contains", func(t *testing.T) {
		assert.True(t, list.Contains(func(value int) bool {
			return value == -9
		}))
	})
	t.Run("not-contains", func(t *testing.T) {
		assert.False(t, list.Contains(func(value int) bool {
			return value == 0
		}))
	})
}

func TestArrayList_IndexOf(t *testing.T) {
	list := createArrayList(t)
	t.Run("contains", func(t *testing.T) {
		assert.Equal(t, 0, list.IndexOf(func(value int) bool {
			return value == -9
		}))
	})
	t.Run("not-contains", func(t *testing.T) {
		assert.Equal(t, -1, list.IndexOf(func(value int) bool {
			return value == 0
		}))
	})
}

func TestArrayList_LastIndexOf(t *testing.T) {
	list := NewArrayList[int](1, 1, 1, 1, 1)
	t.Run("contains", func(t *testing.T) {
		assert.Equal(t, 4, list.LastIndexOf(func(value int) bool {
			return value == 1
		}))
	})
	t.Run("not-contains", func(t *testing.T) {
		assert.Equal(t, -1, list.LastIndexOf(func(value int) bool {
			return value == 0
		}))
	})
}

func TestArrayList_SubList(t *testing.T) {
	list := createArrayList(t)
	t.Run("not-out-of-index", func(t *testing.T) {
		assert.Equal(t, NewArrayList[int](-56, 70), list.SubList(1, 3))
		assert.Equal(t, NewArrayList[int](-56, 70), list.SubList(3, 1))
	})
	t.Run("out-of-index", func(t *testing.T) {
		assert.PanicsWithValue(t, ErrIndexOutOfRange, func() { list.SubList(-1, -2) })
	})
}

func TestArrayList_SubArrayList(t *testing.T) {
	list := createArrayList(t)
	t.Run("not-out-of-index", func(t *testing.T) {
		assert.Equal(t, NewArrayList[int](-56, 70), list.SubArrayList(1, 3))
	})
	t.Run("out-of-index", func(t *testing.T) {
		assert.PanicsWithValue(t, ErrIndexOutOfRange, func() { list.SubArrayList(-1, -2) })
	})
}

func TestArrayList_Add(t *testing.T) {
	list := emptyArrayList(t)
	list.Add(1)
	list.Add(2)
	list.Add(3)
	assert.Equal(t, NewArrayList[int](1, 2, 3), list)
}

func TestArrayList_AddAll(t *testing.T) {
	list := emptyArrayList(t)
	list.AddAll(1, 2, 3)
	assert.Equal(t, NewArrayList[int](1, 2, 3), list)
}

func TestArrayList_Set(t *testing.T) {
	list := emptyArrayList(t)
	list.AddAll(1, 2, 3)
	t.Run("out-of-index", func(t *testing.T) {
		assert.PanicsWithValue(t, ErrIndexOutOfRange, func() { list.Set(-1, 5) })
	})
	t.Run("not-out-of-index", func(t *testing.T) {
		list.Set(0, 5)
		assert.Equal(t, NewArrayList[int](5, 2, 3), list)
	})
}

func TestArrayList_Push(t *testing.T) {
	list := emptyArrayList(t)
	list.Push(1)
	list.Push(2)
	list.Push(3)
	assert.Equal(t, NewArrayList[int](1, 2, 3), list)
}

func TestArrayList_PushAll(t *testing.T) {
	list := emptyArrayList(t)
	list.AddAll(1, 2, 3)
	assert.Equal(t, NewArrayList[int](1, 2, 3), list)
}

func TestArrayList_Unshift(t *testing.T) {
	list := emptyArrayList(t)
	list.Unshift(1)
	list.Unshift(2)
	list.Unshift(3)
	assert.Equal(t, NewArrayList[int](3, 2, 1), list)
}

func TestArrayList_UnshiftAll(t *testing.T) {
	list := emptyArrayList(t)
	list.UnshiftAll(1, 2, 3)
	assert.Equal(t, NewArrayList[int](3, 2, 1), list)
}

func TestArrayList_InsertBefore(t *testing.T) {
	list := emptyArrayList(t)
	list.PushAll(1, 2, 3)
	list.InsertBefore(1, 5)
	assert.Equal(t, NewArrayList[int](1, 5, 2, 3), list)
	assert.PanicsWithValue(t, ErrIndexOutOfRange, func() { list.InsertBefore(-1, 2) })
}

func TestArrayList_InsertAfter(t *testing.T) {
	list := emptyArrayList(t)
	list.PushAll(1, 2, 3)
	list.InsertAfter(1, 5)
	assert.Equal(t, NewArrayList[int](1, 2, 5, 3), list)
	assert.PanicsWithValue(t, ErrIndexOutOfRange, func() { list.InsertAfter(-1, 2) })
}

func TestArrayList_RemoveAt(t *testing.T) {
	list := NewArrayList[int](1, 2, 3)
	list.RemoveAt(0)
	assert.Equal(t, NewArrayList[int](2, 3), list)
	assert.PanicsWithValue(t, ErrIndexOutOfRange, func() { list.RemoveAt(-1) })
}

func TestArrayList_Remove(t *testing.T) {
	list := NewArrayList[int](1, 2, 3)
	list.Remove(func(value int) bool {
		return value == 2 || value == 3
	})
	assert.Equal(t, NewArrayList[int](1), list)
}

func TestArrayList_Clear(t *testing.T) {
	list := createArrayList(t)
	list.Clear()
	assert.Equal(t, 0, list.Count())
}

func TestArrayList_Range(t *testing.T) {
	t.Run("range-all", func(t *testing.T) {
		s := []int{}
		list := NewArrayList[int](1, 2, 3, 4, 5)
		list.Range(func(item int) bool {
			s = append(s, item)
			return true
		})
		assert.Equal(t, []int{1, 2, 3, 4, 5}, s)
	})
	t.Run("range-break", func(t *testing.T) {
		s := []int{}
		list := NewArrayList[int](1, 2, 3, 4, 5)
		list.Range(func(item int) bool {
			s = append(s, item)
			return item < 3
		})
		assert.Equal(t, []int{1, 2, 3}, s)
	})
}

func TestArrayList_ReverseRange(t *testing.T) {
	t.Run("range-all", func(t *testing.T) {
		s := []int{}
		list := NewArrayList[int](1, 2, 3, 4, 5)
		list.ReverseRange(func(item int) bool {
			s = append(s, item)
			return true
		})
		assert.Equal(t, []int{5, 4, 3, 2, 1}, s)
	})
	t.Run("range-break", func(t *testing.T) {
		s := []int{}
		list := NewArrayList[int](1, 2, 3, 4, 5)
		list.ReverseRange(func(item int) bool {
			s = append(s, item)
			return item > 3
		})
		assert.Equal(t, []int{5, 4, 3}, s)
	})
}

func TestArrayList_Map(t *testing.T) {
	list := NewArrayList[int](1, 2, 3, 4, 5)
	list.Map(func(value int) int {
		return value * value
	})
	assert.Equal(t, NewArrayList[int](1, 4, 9, 16, 25), list)
}
