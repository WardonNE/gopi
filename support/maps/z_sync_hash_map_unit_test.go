package maps

import (
	"encoding/json"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSyncHashMap_MarshalJSON(t *testing.T) {
	m := NewSyncHashMap[int, int]()
	m.Set(1, 1)
	m.Set(2, 2)
	m.Set(3, 3)
	bytes, err := json.Marshal(m)
	assert.Nil(t, err)
	assert.JSONEq(t, `{"1":1,"2":2,"3":3}`, string(bytes))
}

func TestSyncHashMap_UnmarshalJSON(t *testing.T) {
	m := NewSyncHashMap[int, int]()
	err := json.Unmarshal([]byte(`{"1":1,"2":2,"3":3}`), m)
	assert.Nil(t, err)
	mm := NewSyncHashMap[int, int]()
	mm.Set(1, 1)
	mm.Set(2, 2)
	mm.Set(3, 3)
	assert.Equal(t, mm, m)
}

func TestSyncHashMap_ToMap(t *testing.T) {
	m := NewSyncHashMap[int, int]()
	m.Set(1, 1)
	m.Set(2, 2)
	m.Set(3, 3)
	assert.Equal(t, map[int]int{1: 1, 2: 2, 3: 3}, m.ToMap())
}

func TestSyncHashMap_FromMap(t *testing.T) {
	m := NewSyncHashMap[int, int]()
	m.FromMap(map[int]int{1: 1, 2: 2, 3: 3})
	mm := NewSyncHashMap[int, int]()
	mm.Set(1, 1)
	mm.Set(2, 2)
	mm.Set(3, 3)
	assert.Equal(t, mm, m)
}

func TestSyncHashMap_Clone(t *testing.T) {
	m := NewSyncHashMap[int, int]()
	m.Set(1, 1)
	m.Set(2, 2)
	m.Set(3, 3)
	assert.Equal(t, m, m.Clone())
}

func TestSyncHashMap_Copy(t *testing.T) {
	m := NewSyncHashMap[int, int]()
	m.Set(1, 1)
	m.Set(2, 2)
	m.Set(3, 3)
	assert.Equal(t, m, m.Copy())
}

func TestSyncHashMap_Count(t *testing.T) {
	m := NewSyncHashMap[int, int]()
	m.Set(1, 1)
	m.Set(2, 2)
	m.Set(3, 3)
	assert.Equal(t, 3, m.Count())
}

func TestSyncHashMap_Get(t *testing.T) {
	m := NewSyncHashMap[int, int]()
	m.Set(1, 1)
	m.Set(2, 2)
	m.Set(3, 3)
	assert.Equal(t, 1, m.Get(1))
	assert.Zero(t, m.Get(4))
}

func TestSyncHashMap_GetOrDefault(t *testing.T) {
	m := NewSyncHashMap[int, int]()
	m.Set(1, 1)
	m.Set(2, 2)
	m.Set(3, 3)
	assert.Equal(t, 1, m.GetOrDefault(1, 2))
	assert.Equal(t, 4, m.GetOrDefault(4, 4))
}

func TestSyncHashMap_Set(t *testing.T) {
	m := NewSyncHashMap[int, int]()
	m.Set(1, 1)
	m.Set(2, 2)
	m.Set(3, 3)
	assert.Contains(t, m.ToMap(), 1)
	assert.Contains(t, m.ToMap(), 2)
	assert.Contains(t, m.ToMap(), 3)
}

func TestSyncHashMap_Remove(t *testing.T) {
	m := NewSyncHashMap[int, int]()
	m.Set(1, 1)
	m.Set(2, 2)
	m.Set(3, 3)
	m.Remove(1)
	m.Remove(3)
	assert.NotContains(t, m.ToMap(), 1)
	assert.Contains(t, m.ToMap(), 2)
	assert.NotContains(t, m.ToMap(), 3)
}

func TestSyncHashMap_Keys(t *testing.T) {
	m := NewSyncHashMap[int, int]()
	m.Set(1, 1)
	m.Set(2, 2)
	m.Set(3, 3)
	keys := m.Keys()
	sort.Ints(keys)
	assert.Equal(t, []int{1, 2, 3}, keys)
}

func TestSyncHashMap_Values(t *testing.T) {
	m := NewSyncHashMap[int, int]()
	m.Set(1, 1)
	m.Set(2, 2)
	m.Set(3, 3)
	values := m.Values()
	sort.Ints(values)
	assert.Equal(t, []int{1, 2, 3}, values)
}

func TestSyncHashMap_Clear(t *testing.T) {
	m := NewSyncHashMap[int, int]()
	m.Set(1, 1)
	m.Set(2, 2)
	m.Set(3, 3)
	m.Clear()
	assert.True(t, m.IsEmpty())
}

func TestSyncHashMap_ContainsValue(t *testing.T) {
	m := NewSyncHashMap[int, int]()
	m.Set(1, 1)
	m.Set(2, 2)
	m.Set(3, 3)
	assert.True(t, m.ContainsValue(func(value int) bool { return value == 1 }))
	assert.True(t, m.ContainsValue(func(value int) bool { return value == 2 }))
	assert.True(t, m.ContainsValue(func(value int) bool { return value == 3 }))
}

func TestSyncHashMap_ContainsKey(t *testing.T) {
	m := NewSyncHashMap[int, int]()
	m.Set(1, 1)
	m.Set(2, 2)
	m.Set(3, 3)
	assert.True(t, m.ContainsKey(1))
	assert.True(t, m.ContainsKey(2))
	assert.True(t, m.ContainsKey(3))
}

func TestSyncHashMap_IsEmpty(t *testing.T) {
	m := NewSyncHashMap[int, int]()
	assert.True(t, m.IsEmpty())
	m.Set(1, 1)
	m.Set(2, 2)
	m.Set(3, 3)
	assert.False(t, m.IsEmpty())
}

func TestSyncHashMap_IsNotEmpty(t *testing.T) {
	m := NewSyncHashMap[int, int]()
	assert.False(t, m.IsNotEmpty())
	m.Set(1, 1)
	m.Set(2, 2)
	m.Set(3, 3)
	assert.True(t, m.IsNotEmpty())

}

func TestSyncHashMap_Range(t *testing.T) {
	m := NewSyncHashMap[int, int]()
	m.Set(1, 1)
	m.Set(2, 2)
	m.Set(3, 3)
	keys := []int{}
	values := []int{}
	m.Range(func(entry *Entry[int, int]) bool {
		keys = append(keys, entry.Key)
		values = append(values, entry.Value)
		return true
	})
	sort.Ints(keys)
	sort.Ints(values)
	assert.Equal(t, []int{1, 2, 3}, keys)
	assert.Equal(t, []int{1, 2, 3}, values)
}
