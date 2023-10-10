package set

import (
	"encoding/json"
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSyncHashSet_MarshalJSON(t *testing.T) {
	set := NewSyncHashSet[int](1, 2, 3, 4, 1)
	bytes, err := json.Marshal(set)
	assert.Nil(t, err)
	values := []int{}
	json.Unmarshal(bytes, &values)
	sort.Ints(values)
	assert.Equal(t, []int{1, 2, 3, 4}, values)
}

func TestSyncHashSet_UnmarshalJSON(t *testing.T) {
	set := NewSyncHashSet[int]()
	err := json.Unmarshal([]byte(`[1,2,3,4,1]`), set)
	assert.Nil(t, err)
	assert.Equal(t, NewSyncHashSet[int](1, 2, 3, 4), set)
}

func TestSyncHashSet_ToArray(t *testing.T) {
	set := NewSyncHashSet[int](1, 2, 3, 4, 1)
	values := set.ToArray()
	sort.Ints(values)
	assert.Equal(t, []int{1, 2, 3, 4}, values)
}

func TestSyncHashSet_FromArray(t *testing.T) {
	set := NewSyncHashSet[int]()
	set.FromArray([]int{1, 2, 3, 4, 1})
	assert.Equal(t, NewSyncHashSet[int](1, 2, 3, 4), set)
}

func TestSyncHashSet_String(t *testing.T) {
	set := NewSyncHashSet[int](1, 2, 3, 4, 1)
	str := set.String()
	str = strings.TrimLeft(str, "[")
	str = strings.TrimRight(str, "]")
	values := strings.Split(str, " ")
	sort.Strings(values)
	assert.Equal(t, []string{"1", "2", "3", "4"}, values)
}

func TestSyncHashSet_Clone(t *testing.T) {
	set := NewSyncHashSet[int](1, 2, 3, 4, 5)
	clonedSet := set.Clone()
	assert.Equal(t, set, clonedSet)
}

func TestSyncHashSet_Copy(t *testing.T) {
	set := NewSyncHashSet[int](1, 2, 3, 4, 5)
	copiedSet := set.Copy()
	assert.Equal(t, set, copiedSet)
}

func TestSyncHashSet_Count(t *testing.T) {
	set := NewSyncHashSet[int](1, 2, 3, 4, 5)
	assert.Equal(t, 5, set.Count())
}

func TestSyncHashSet_IsEmpty(t *testing.T) {
	set := NewSyncHashSet[int]()
	assert.True(t, set.IsEmpty())
}

func TestSyncHashSet_IsNotEmpty(t *testing.T) {
	set := NewSyncHashSet[int](1, 2, 3)
	assert.True(t, set.IsNotEmpty())
}

func TestSyncHashSet_Contains(t *testing.T) {
	set := NewSyncHashSet[int](1, 2, 3, 4, 5)
	assert.True(t, set.Contains(func(value int) bool { return value == 1 }))
	assert.False(t, set.Contains(func(value int) bool { return value == 0 }))
}

func TestSyncHashSet_Add(t *testing.T) {
	set := NewSyncHashSet[int]()
	set.Add(1)
	set.Add(2)
	set.Add(3)
	assert.Equal(t, NewSyncHashSet[int](1, 2, 3), set)
}

func TestSyncHashSet_AddAll(t *testing.T) {
	set := NewSyncHashSet[int]()
	set.AddAll(1, 2, 3)
	assert.Equal(t, NewSyncHashSet[int](1, 2, 3), set)
}

func TestSyncHashSet_Remove(t *testing.T) {
	set := NewSyncHashSet[int](1, 2, 3)
	set.Remove(func(value int) bool { return value == 1 })
	assert.Equal(t, NewSyncHashSet[int](2, 3), set)
	set.Remove(func(value int) bool { return value == 0 })
	assert.Equal(t, NewSyncHashSet[int](2, 3), set)
}

func TestSyncHashSet_Clear(t *testing.T) {
	set := NewSyncHashSet[int](1, 2, 3)
	set.Clear()
	assert.True(t, set.IsEmpty())
}

func TestSyncHashSet_Range(t *testing.T) {
	set := NewSyncHashSet[int](1, 2, 3, 4, 5)
	values1 := []int{}
	set.Range(func(item int) bool {
		values1 = append(values1, item)
		return true
	})
	sort.Ints(values1)
	assert.Equal(t, []int{1, 2, 3, 4, 5}, values1)
}
