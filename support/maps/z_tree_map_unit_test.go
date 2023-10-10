package maps

import (
	"encoding/json"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wardonne/gopi/support/compare"
)

func TestTreeMap_MarshalJSON(t *testing.T) {
	m := NewTreeMap[int, int](compare.NewNatureComparator[int](false))
	m.Set(1, 1)
	m.Set(2, 2)
	m.Set(3, 3)
	bytes, err := json.Marshal(m)
	assert.Nil(t, err)
	assert.JSONEq(t, `{"1":1,"2":2,"3":3}`, string(bytes))
}

func TestTreeMap_ToMap(t *testing.T) {
	m := NewTreeMap[int, int](compare.NewNatureComparator[int](false))
	m.Set(1, 1)
	m.Set(2, 2)
	m.Set(3, 3)
	assert.Equal(t, map[int]int{1: 1, 2: 2, 3: 3}, m.ToMap())
}

func TestTreeMap_Clone(t *testing.T) {
	m := NewTreeMap[int, int](compare.NewNatureComparator[int](false))
	m.Set(1, 1)
	m.Set(2, 2)
	m.Set(3, 3)
	assert.Equal(t, m, m.Clone())
}

func TestTreeMap_Get(t *testing.T) {
	m := NewTreeMap[int, int](compare.NewNatureComparator[int](false))
	m.Set(1, 1)
	m.Set(2, 2)
	m.Set(3, 3)
	assert.Zero(t, m.Get(4))
	assert.Equal(t, 1, m.Get(1))
	assert.Equal(t, 2, m.Get(2))
	assert.Equal(t, 3, m.Get(3))
}

func TestTreeMap_GetOrDefault(t *testing.T) {
	m := NewTreeMap[int, int](compare.NewNatureComparator[int](false))
	m.Set(1, 1)
	m.Set(2, 2)
	m.Set(3, 3)
	assert.Equal(t, 1, m.GetOrDefault(1, 0))
	assert.Equal(t, 2, m.GetOrDefault(2, 0))
	assert.Equal(t, 3, m.GetOrDefault(3, 0))
	assert.Equal(t, 1, m.GetOrDefault(4, 1))
}

func TestTreeMap_Set(t *testing.T) {
	m := NewTreeMap[int, int](compare.NewNatureComparator[int](false))
	m.Set(1, 1)
	m.Set(2, 2)
	m.Set(3, 3)
	assert.Zero(t, m.Get(4))
	assert.Equal(t, 1, m.Get(1))
	assert.Equal(t, 2, m.Get(2))
	assert.Equal(t, 3, m.Get(3))
}

func TestTreeMap_Remove(t *testing.T) {
	m := NewTreeMap[int, int](compare.NewNatureComparator[int](false))
	m.Set(1, 1)
	m.Set(2, 2)
	m.Set(3, 3)
	m.Remove(1)
	m.Remove(3)
	assert.Zero(t, m.Get(1))
	assert.Equal(t, 2, m.Get(2))
	assert.Zero(t, m.Get(3))
}

func TestTreeMap_Keys(t *testing.T) {
	m := NewTreeMap[int, int](compare.NewNatureComparator[int](false))
	m.Set(1, 1)
	m.Set(2, 2)
	m.Set(3, 3)
	keys := m.Keys()
	sort.Ints(keys)
	assert.Equal(t, []int{1, 2, 3}, keys)
}

func TestTreeMap_Values(t *testing.T) {
	m := NewTreeMap[int, int](compare.NewNatureComparator[int](false))
	m.Set(1, 1)
	m.Set(2, 2)
	m.Set(3, 3)
	values := m.Values()
	sort.Ints(values)
	assert.Equal(t, []int{1, 2, 3}, values)
}

func TestTreeMap_Clear(t *testing.T) {
	m := NewTreeMap[int, int](compare.NewNatureComparator[int](false))
	m.Set(1, 1)
	m.Set(2, 2)
	m.Set(3, 3)
	m.Clear()
	assert.True(t, m.IsEmpty())
}

func TestTreeMap_ContainsValue(t *testing.T) {
	m := NewTreeMap[int, int](compare.NewNatureComparator[int](false))
	m.Set(1, 1)
	m.Set(2, 2)
	m.Set(3, 3)
	assert.True(t, m.ContainsValue(func(value int) bool {
		return value == 1
	}))
	assert.True(t, m.ContainsValue(func(value int) bool {
		return value == 2
	}))
	assert.True(t, m.ContainsValue(func(value int) bool {
		return value == 3
	}))
}

func TestTreeMap_ContainsKey(t *testing.T) {
	m := NewTreeMap[int, int](compare.NewNatureComparator[int](false))
	m.Set(1, 1)
	m.Set(2, 2)
	m.Set(3, 3)
	assert.True(t, m.ContainsKey(1))
	assert.True(t, m.ContainsKey(2))
	assert.True(t, m.ContainsKey(3))
}

func TestTreeMap_Count(t *testing.T) {
	m := NewTreeMap[int, int](compare.NewNatureComparator[int](false))
	m.Set(1, 1)
	m.Set(2, 2)
	m.Set(3, 3)
	assert.Equal(t, 3, m.Count())
}

func TestTreeMap_IsEmpty(t *testing.T) {
	m := NewTreeMap[int, int](compare.NewNatureComparator[int](false))
	assert.True(t, m.IsEmpty())
	m.Set(1, 1)
	m.Set(2, 2)
	m.Set(3, 3)
	assert.False(t, m.IsEmpty())
}

func TestTreeMap_IsNotEmpty(t *testing.T) {
	m := NewTreeMap[int, int](compare.NewNatureComparator[int](false))
	assert.False(t, m.IsNotEmpty())
	m.Set(1, 1)
	m.Set(2, 2)
	m.Set(3, 3)
	assert.True(t, m.IsNotEmpty())
}

func TestTreeMap_FirstKey(t *testing.T) {
	m := NewTreeMap[int, int](compare.NewNatureComparator[int](false))
	m.Set(1, 1)
	m.Set(2, 2)
	m.Set(3, 3)
	assert.Equal(t, 1, m.FirstKey())
}

func TestTreeMap_LastKey(t *testing.T) {
	m := NewTreeMap[int, int](compare.NewNatureComparator[int](false))
	m.Set(1, 1)
	m.Set(2, 2)
	m.Set(3, 3)
	assert.Equal(t, 3, m.LastKey())
}

func TestTreeMap_Range(t *testing.T) {
	m := NewTreeMap[int, int](compare.NewNatureComparator[int](false))
	m.Set(1, 1)
	m.Set(2, 2)
	m.Set(3, 3)
	keys := []int{}
	values := []int{}
	m.Range(func(entry *Entry[int, int]) bool {
		values = append(values, entry.Value)
		keys = append(keys, entry.Key)
		return true
	})
	sort.Ints(keys)
	sort.Ints(values)
	assert.Equal(t, []int{1, 2, 3}, keys)
	assert.Equal(t, []int{1, 2, 3}, values)
}
