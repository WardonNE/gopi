package list

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wardonne/gopi/support/compare"
)

func TestSyncArrayList_MarshalJSON(t *testing.T) {
	list := NewSyncArrayList[int](1, 2, 3, 4, 5)
	bytes, err := json.Marshal(list)
	assert.Nil(t, err)
	assert.JSONEq(t, `[1,2,3,4,5]`, string(bytes))
}

func TestSyncArrayList_UnmarshalJSON(t *testing.T) {
	list := NewSyncArrayList[int]()
	jsonBytes := []byte(`[1,2,3,4,5]`)
	err := json.Unmarshal(jsonBytes, list)
	assert.Nil(t, err)
	assert.Equal(t, NewSyncArrayList[int](1, 2, 3, 4, 5), list)
}

func TestSyncArrayList_ToArray(t *testing.T) {
	list := NewSyncArrayList[int](1, 2, 3, 4, 5)
	assert.Equal(t, []int{1, 2, 3, 4, 5}, list.ToArray())
}

func TestSyncArrayList_FromArray(t *testing.T) {
	list := NewSyncArrayList[int]()
	list.FromArray([]int{1, 2, 3, 4, 5})
	assert.Equal(t, NewSyncArrayList[int](1, 2, 3, 4, 5), list)
}

func TestSyncArrayList_String(t *testing.T) {
	list := NewSyncArrayList[int](1, 2, 3, 4, 5)
	assert.Equal(t, fmt.Sprintf("%v", []int{1, 2, 3, 4, 5}), list.String())
}

func TestSyncArrayList_Clone(t *testing.T) {
	list := NewSyncArrayList[int](1, 2, 3, 4, 5)
	clonedList := list.Clone()
	assert.Equal(t, list, clonedList)
}

func TestSyncArrayList_Copy(t *testing.T) {
	list := NewSyncArrayList[int](1, 2, 3, 4, 5)
	copiedList := list.Copy()
	assert.Equal(t, list, copiedList)
}

func TestSyncArrayList_Sort(t *testing.T) {
	list := NewSyncArrayList[int](1, 3, 5, 3, 21, 1)
	list.Sort(compare.NewNatureComparator[int](false))
	assert.Equal(t, []int{1, 1, 3, 3, 5, 21}, list.ToArray())
}

func TestSyncArrayList_Count(t *testing.T) {
	list := NewSyncArrayList[int](1, 2, 3)
	assert.Equal(t, 3, list.Count())
}

func TestSyncArrayList_IsEmpty(t *testing.T) {
	list := NewSyncArrayList[int](1, 2, 3)
	assert.False(t, list.IsEmpty())
}

func TestSyncArrayList_IsNotEmpty(t *testing.T) {
	list := NewSyncArrayList[int](1, 2, 3)
	assert.True(t, list.IsNotEmpty())
}

func TestSyncArrayList_Get(t *testing.T) {
	t.Run("out-of-index", func(t *testing.T) {
		assert.PanicsWithValue(t, ErrIndexOutOfRange, func() { NewSyncArrayList[int]().Get(-1) })
	})
	t.Run("not-out-of-index", func(t *testing.T) {
		list := NewSyncArrayList[int](1, 2, 3, 4, 5)
		assert.Equal(t, 2, list.Get(1))
	})
}

func TestSyncArrayList_Pop(t *testing.T) {
	t.Run("zero", func(t *testing.T) {
		assert.Zero(t, NewSyncArrayList[int]().Pop())
	})
	t.Run("non-zero", func(t *testing.T) {
		assert.Equal(t, 5, NewSyncArrayList[int](1, 2, 3, 4, 5).Pop())
	})
}

func TestSyncArrayList_Shift(t *testing.T) {
	t.Run("zero", func(t *testing.T) {
		assert.Zero(t, NewSyncArrayList[int]().Shift())
	})
	t.Run("non-zero", func(t *testing.T) {
		assert.Equal(t, 1, NewSyncArrayList[int](1, 2, 3).Shift())
	})
}

func TestSyncArrayList_Contains(t *testing.T) {
	t.Run("contains", func(t *testing.T) {
		assert.True(t, NewSyncArrayList[int](1, 2, 3, 4).Contains(func(v int) bool {
			return v == 2
		}))
	})
	t.Run("not-contains", func(t *testing.T) {
		assert.False(t, NewSyncArrayList[int](1, 2, 3, 4).Contains(func(v int) bool {
			return v == 0
		}))
	})
}

func TestSyncArrayList_IndexOf(t *testing.T) {
	t.Run("contains", func(t *testing.T) {
		assert.Equal(t, 1, NewSyncArrayList[int](0, 1, 2, 3).IndexOf(func(value int) bool { return value == 1 }))
	})
	t.Run("not-contains", func(t *testing.T) {
		assert.Equal(t, -1, NewSyncArrayList[int](0, 1, 2, 3).IndexOf(func(value int) bool { return value == -1 }))
	})
}

func TestSyncArrayList_LastIndexOf(t *testing.T) {
	t.Run("contains", func(t *testing.T) {
		assert.Equal(t, 3, NewSyncArrayList[int](0, 1, 2, 1).LastIndexOf(func(value int) bool { return value == 1 }))
	})
	t.Run("not-contains", func(t *testing.T) {
		assert.Equal(t, -1, NewSyncArrayList[int](0, 1, 2, 1).IndexOf(func(value int) bool { return value == -1 }))
	})
}

func TestSyncArrayList_SubList(t *testing.T) {
	t.Run("out-of-index", func(t *testing.T) {
		assert.PanicsWithValue(t, ErrIndexOutOfRange, func() { NewSyncArrayList[int]().SubList(1, 2) })
	})
	t.Run("not-out-of-index", func(t *testing.T) {
		assert.Equal(t, NewSyncArrayList[int](1, 2), NewSyncArrayList[int](0, 1, 2, 3, 4).SubList(1, 3))
		assert.Equal(t, NewSyncArrayList[int](1, 2), NewSyncArrayList[int](0, 1, 2, 3, 4).SubList(3, 1))
	})
}

func TestSyncArrayList_SubSyncArrayList(t *testing.T) {
	t.Run("out-of-index", func(t *testing.T) {
		assert.PanicsWithValue(t, ErrIndexOutOfRange, func() { NewSyncArrayList[int]().SubSyncArrayList(1, 2) })
	})
	t.Run("not-out-of-index", func(t *testing.T) {
		assert.Equal(t, NewSyncArrayList[int](1, 2), NewSyncArrayList[int](0, 1, 2, 3, 4).SubSyncArrayList(1, 3))
	})
}

func TestSyncArrayList_Add(t *testing.T) {
	list := NewSyncArrayList[int]()
	list.Add(1)
	list.Add(2)
	list.Add(3)
	assert.Equal(t, NewSyncArrayList[int](1, 2, 3), list)
}

func TestSyncArrayList_AddAll(t *testing.T) {
	list := NewSyncArrayList[int]()
	list.AddAll(1, 2, 3)
	assert.Equal(t, NewSyncArrayList[int](1, 2, 3), list)
}

func TestSyncArrayList_Set(t *testing.T) {
	list := NewSyncArrayList[int]()
	list.AddAll(1, 2, 3)
	list.Set(1, 5)
	assert.Equal(t, NewSyncArrayList[int](1, 5, 3), list)
}

func TestSyncArrayList_Push(t *testing.T) {
	list := NewSyncArrayList[int]()
	list.Push(1)
	list.Push(2)
	list.Push(3)
	assert.Equal(t, NewSyncArrayList[int](1, 2, 3), list)
}

func TestSyncArrayList_PushAll(t *testing.T) {
	list := NewSyncArrayList[int]()
	list.PushAll(1, 2, 3)
	assert.Equal(t, NewSyncArrayList[int](1, 2, 3), list)
}

func TestSyncArrayList_Unshift(t *testing.T) {
	list := NewSyncArrayList[int]()
	list.Unshift(1)
	list.Unshift(2)
	list.Unshift(3)
	assert.Equal(t, NewSyncArrayList[int](3, 2, 1), list)
}

func TestSyncArrayList_UnshiftAll(t *testing.T) {
	list := NewSyncArrayList[int]()
	list.UnshiftAll(1, 2, 3)
	assert.Equal(t, NewSyncArrayList[int](3, 2, 1), list)
}

func TestSyncArrayList_InsertBefore(t *testing.T) {
	list := NewSyncArrayList[int](1, 2, 3)
	list.InsertBefore(1, 5)
	assert.Equal(t, NewSyncArrayList[int](1, 5, 2, 3), list)
	assert.PanicsWithValue(t, ErrIndexOutOfRange, func() { list.InsertBefore(-1, 2) })
}

func TestSyncArrayList_InsertAfter(t *testing.T) {
	list := NewSyncArrayList[int](1, 2, 3)
	list.InsertAfter(1, 5)
	assert.Equal(t, NewSyncArrayList[int](1, 2, 5, 3), list)
	assert.PanicsWithValue(t, ErrIndexOutOfRange, func() { list.InsertAfter(-1, 2) })
}

func TestSyncArrayList_RemoveAt(t *testing.T) {
	list := NewSyncArrayList[int](1, 2, 3)
	list.RemoveAt(0)
	assert.Equal(t, NewSyncArrayList[int](2, 3), list)
	assert.PanicsWithValue(t, ErrIndexOutOfRange, func() { list.RemoveAt(-1) })
}

func TestSyncArrayList_Remove(t *testing.T) {
	list := NewSyncArrayList[int](1, 2, 3, 4, 5)
	list.Remove(func(value int) bool { return value == 1 || value == 4 })
	assert.Equal(t, NewSyncArrayList[int](2, 3, 5), list)
}

func TestSyncArrayList_Clear(t *testing.T) {
	list := NewSyncArrayList[int](1, 2, 3, 4, 5, 6)
	list.Clear()
	assert.True(t, list.IsEmpty())
}

func TestSyncArrayList_Range(t *testing.T) {
	list := NewSyncArrayList[int](1, 2, 3, 4, 5, 6, 7)
	t.Run("range-all", func(t *testing.T) {
		s := []int{}
		list.Range(func(item int) bool {
			s = append(s, item)
			return true
		})
		assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7}, s)
	})
	t.Run("range-break", func(t *testing.T) {
		s := []int{}
		list.Range(func(item int) bool {
			s = append(s, item)
			return item < 4
		})
		assert.Equal(t, []int{1, 2, 3, 4}, s)
	})
}

func TestSyncArrayList_ReverseRange(t *testing.T) {
	list := NewSyncArrayList[int](7, 6, 5, 4, 3, 2, 1)
	t.Run("range-all", func(t *testing.T) {
		s := []int{}
		list.ReverseRange(func(item int) bool {
			s = append(s, item)
			return true
		})
		assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7}, s)
	})
	t.Run("range-break", func(t *testing.T) {
		s := []int{}
		list.ReverseRange(func(item int) bool {
			s = append(s, item)
			return item < 4
		})
		assert.Equal(t, []int{1, 2, 3, 4}, s)
	})
}

func TestSyncArrayList_Map(t *testing.T) {
	list := NewSyncArrayList[int](7, 6, 5, 4, 3, 2, 1)
	list.Map(func(value int) int {
		return value * value
	})
	assert.Equal(t, []int{49, 36, 25, 16, 9, 4, 1}, list.ToArray())
}
