package set

import (
	"encoding/json"
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wardonne/gopi/support/collection/list"
)

func TestSyncLinkedHashSet_MarshalJSON(t *testing.T) {
	set := NewSyncLinkedHashSet[int](1, 2, 3, 4, 1)
	bytes, err := json.Marshal(set)
	assert.Nil(t, err)
	assert.JSONEq(t, `[1,2,3,4]`, string(bytes))
}

func TestSyncLinkedHashSet_UnmarshalJSON(t *testing.T) {
	set := NewSyncLinkedHashSet[int]()
	err := json.Unmarshal([]byte(`[1,2,3,4,1]`), set)
	assert.Nil(t, err)
	assert.Equal(t, NewSyncLinkedHashSet[int](1, 2, 3, 4), set)
}

func TestSyncLinkedHashSet_ToArray(t *testing.T) {
	set := NewSyncLinkedHashSet[int](1, 2, 3, 4, 1)
	values := set.ToArray()
	sort.Ints(values)
	assert.Equal(t, []int{1, 2, 3, 4}, values)
}

func TestSyncLinkedHashSet_FromArray(t *testing.T) {
	set := NewSyncLinkedHashSet[int]()
	set.FromArray([]int{1, 2, 3, 4, 1})
	assert.Equal(t, NewSyncLinkedHashSet[int](1, 2, 3, 4), set)
}

func TestSyncLinkedHashSet_String(t *testing.T) {
	set := NewSyncLinkedHashSet[int](1, 2, 3, 4, 1)
	str := set.String()
	str = strings.TrimLeft(str, "[")
	str = strings.TrimRight(str, "]")
	values := strings.Split(str, " ")
	sort.Strings(values)
	assert.Equal(t, []string{"1", "2", "3", "4"}, values)
}

func TestSyncLinkedHashSet_Clone(t *testing.T) {
	set := NewSyncLinkedHashSet[int](1, 2, 3, 4, 5)
	clonedSet := set.Clone()
	assert.Equal(t, set, clonedSet)
}

func TestSyncLinkedHashSet_Copy(t *testing.T) {
	set := NewSyncLinkedHashSet[int](1, 2, 3, 4, 5)
	copiedSet := set.Copy()
	assert.Equal(t, set, copiedSet)
}

func TestSyncLinkedHashSet_Count(t *testing.T) {
	set := NewSyncLinkedHashSet[int](1, 2, 3, 4, 5)
	assert.Equal(t, 5, set.Count())
}

func TestSyncLinkedHashSet_IsEmpty(t *testing.T) {
	set := NewSyncLinkedHashSet[int]()
	assert.True(t, set.IsEmpty())
}

func TestSyncLinkedHashSet_IsNotEmpty(t *testing.T) {
	set := NewSyncLinkedHashSet[int](1, 2, 3)
	assert.True(t, set.IsNotEmpty())
}

func TestSyncLinkedHashSet_Contains(t *testing.T) {
	set := NewSyncLinkedHashSet[int](1, 2, 3, 4, 5)
	assert.True(t, set.Contains(func(value int) bool { return value == 1 }))
	assert.False(t, set.Contains(func(value int) bool { return value == 0 }))
}

func TestSyncLinkedHashSet_Add(t *testing.T) {
	set := NewSyncLinkedHashSet[int]()
	set.Add(1)
	set.Add(2)
	set.Add(3)
	assert.Equal(t, NewSyncLinkedHashSet[int](1, 2, 3), set)
}

func TestSyncLinkedHashSet_AddAll(t *testing.T) {
	set := NewSyncLinkedHashSet[int]()
	set.AddAll(1, 2, 3)
	assert.Equal(t, NewSyncLinkedHashSet[int](1, 2, 3), set)
}

func TestSyncLinkedHashSet_Remove(t *testing.T) {
	set := NewSyncLinkedHashSet[int](1, 2, 3)
	set.Remove(func(value int) bool {
		return value == 3
	})
	assert.Equal(t, NewSyncLinkedHashSet[int](1, 2), set)
	set.Remove(func(value int) bool { return value == 0 })
	assert.Equal(t, NewSyncLinkedHashSet[int](1, 2), set)
}

func TestSyncLinkedHashSet_Clear(t *testing.T) {
	set := NewSyncLinkedHashSet[int](1, 2, 3)
	set.Clear()
	assert.True(t, set.IsEmpty())
}

func TestSyncLinkedHashSet_Range(t *testing.T) {
	set := NewSyncLinkedHashSet[int](1, 2, 3, 4, 5)
	values1 := []int{}
	set.Range(func(item int) bool {
		values1 = append(values1, item)
		return true
	})
	assert.Equal(t, []int{1, 2, 3, 4, 5}, values1)
}

func TestSyncLinkedHashSet_Get(t *testing.T) {
	set := NewSyncLinkedHashSet[int](1, 2, 3, 4, 5)
	assert.PanicsWithValue(t, list.ErrIndexOutOfRange, func() {
		set.Get(-1)
	})
	assert.Equal(t, 2, set.Get(1))
}

func TestSyncLinkedHashSet_Pop(t *testing.T) {
	set := NewSyncLinkedHashSet[int](1, 2, 3, 4, 5)
	assert.Equal(t, 5, set.Pop())
	assert.Equal(t, NewSyncLinkedHashSet[int](1, 2, 3, 4), set)
}

func TestSyncLinkedHashSet_Shift(t *testing.T) {
	set := NewSyncLinkedHashSet[int](1, 2, 3, 4, 5)
	assert.Equal(t, 1, set.Shift())
	assert.Equal(t, NewSyncLinkedHashSet[int](2, 3, 4, 5), set)
}

func TestSyncLinkedHashSet_IndexOf(t *testing.T) {
	set := NewSyncLinkedHashSet[int](1, 2, 3, 4, 5)
	assert.Equal(t, 0, set.IndexOf(func(value int) bool {
		return value == 1
	}))
	assert.Equal(t, -1, set.IndexOf(func(value int) bool {
		return value == 0
	}))
}

func TestSyncLinkedHashSet_LastIndexOf(t *testing.T) {
	set := NewSyncLinkedHashSet[int](1, 2, 3, 4, 5, 1)
	assert.Equal(t, 0, set.LastIndexOf(func(value int) bool {
		return value == 1
	}))
	assert.Equal(t, -1, set.LastIndexOf(func(value int) bool {
		return value == 9
	}))
}

func TestSyncLinkedHashSet_Push(t *testing.T) {
	set := NewSyncLinkedHashSet[int](1, 2, 3, 4, 5)
	set.Push(5)
	set.Push(6)
	assert.Equal(t, NewSyncLinkedHashSet[int](1, 2, 3, 4, 5, 6), set)
}

func TestSyncLinkedHashSet_PushAll(t *testing.T) {
	set := NewSyncLinkedHashSet[int](1, 2, 3, 4, 5)
	set.PushAll(5, 6)
	assert.Equal(t, NewSyncLinkedHashSet[int](1, 2, 3, 4, 5, 6), set)
}

func TestSyncLinkedHashSet_Unshift(t *testing.T) {
	set := NewSyncLinkedHashSet[int](1, 2, 3, 4, 5)
	set.Unshift(5)
	set.Unshift(6)
	assert.Equal(t, NewSyncLinkedHashSet[int](6, 1, 2, 3, 4, 5), set)
}

func TestSyncLinkedHashSet_UnshiftAll(t *testing.T) {
	set := NewSyncLinkedHashSet[int](1, 2, 3, 4, 5)
	set.UnshiftAll(5, 6)
	assert.Equal(t, NewSyncLinkedHashSet[int](6, 1, 2, 3, 4, 5), set)
}

func TestSyncLinkedHashSet_InsertBefore(t *testing.T) {
	set := NewSyncLinkedHashSet[int](1, 2, 3, 4, 5)
	set.InsertBefore(0, 1)
	set.InsertBefore(0, 6)
	assert.Equal(t, NewSyncLinkedHashSet[int](6, 1, 2, 3, 4, 5), set)
}

func TestSyncLinkedHashSet_InsertAfter(t *testing.T) {
	set := NewSyncLinkedHashSet[int](1, 2, 3, 4, 5)
	set.InsertAfter(0, 1)
	set.InsertAfter(0, 6)
	assert.Equal(t, NewSyncLinkedHashSet[int](1, 6, 2, 3, 4, 5), set)
}

func TestSyncLinkedHashSet_RemoveAt(t *testing.T) {
	set := NewSyncLinkedHashSet[int](1, 2, 3, 4, 5)
	set.RemoveAt(0)
	assert.Equal(t, NewSyncLinkedHashSet[int](2, 3, 4, 5), set)
}

func TestSyncLinkedHashSet_ReverseRange(t *testing.T) {
	set := NewSyncLinkedHashSet[int](1, 2, 3, 4, 5)
	values1 := []int{}
	set.ReverseRange(func(item int) bool {
		values1 = append(values1, item)
		return true
	})
	assert.Equal(t, []int{5, 4, 3, 2, 1}, values1)
	values2 := []int{}
	set.ReverseRange(func(item int) bool {
		values2 = append(values2, item)
		return item > 3
	})
	assert.Equal(t, []int{5, 4, 3}, values2)
}
