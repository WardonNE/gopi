package maps

import (
	"encoding/json"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashMap_MarshalJSON(t *testing.T) {
	m := NewHashMap[int, int]()
	m.Set(1, 1)
	m.Set(2, 2)
	m.Set(3, 3)
	bytes, err := json.Marshal(m)
	assert.Nil(t, err)
	assert.JSONEq(t, `{"1":1,"2":2,"3":3}`, string(bytes))
}

func TestHashMap_UnmarshalJSON(t *testing.T) {
	m := NewHashMap[int, int]()
	err := json.Unmarshal([]byte(`{"1":1,"2":2,"3":3}`), m)
	assert.Nil(t, err)
	mm := NewHashMap[int, int]()
	mm.Set(1, 1)
	mm.Set(2, 2)
	mm.Set(3, 3)
	assert.Equal(t, mm, m)
}

func TestHashMap_ToMap(t *testing.T) {
	m := NewHashMap[int, int]()
	m.Set(1, 1)
	m.Set(2, 2)
	m.Set(3, 3)
	assert.Equal(t, map[int]int{1: 1, 2: 2, 3: 3}, m.ToMap())
}

func TestHashMap_FromMap(t *testing.T) {
	m := NewHashMap[int, int]()
	m.FromMap(map[int]int{1: 1, 2: 2, 3: 3})
	mm := NewHashMap[int, int]()
	mm.Set(1, 1)
	mm.Set(2, 2)
	mm.Set(3, 3)
	assert.Equal(t, mm, m)
}

func TestHashMap_Clone(t *testing.T) {
	m := NewHashMap[int, int]()
	m.Set(1, 1)
	m.Set(2, 2)
	m.Set(3, 3)
	assert.Equal(t, m, m.Clone())
}

func TestHashMap_Copy(t *testing.T) {
	m := NewHashMap[int, int]()
	m.Set(1, 1)
	m.Set(2, 2)
	m.Set(3, 3)
	assert.Equal(t, m, m.Copy())
}

func TestHashMap_Count(t *testing.T) {
	m := NewHashMap[int, int]()
	m.Set(1, 1)
	m.Set(2, 2)
	m.Set(3, 3)
	assert.Equal(t, 3, m.Count())
}

func TestHashMap_Get(t *testing.T) {
	m := NewHashMap[int, int]()
	m.Set(1, 1)
	m.Set(2, 2)
	m.Set(3, 3)
	assert.Equal(t, 1, m.Get(1))
	assert.Zero(t, m.Get(4))
}

func TestHashMap_GetOrDefault(t *testing.T) {
	m := NewHashMap[int, int]()
	m.Set(1, 1)
	m.Set(2, 2)
	m.Set(3, 3)
	assert.Equal(t, 1, m.GetOrDefault(1, 2))
	assert.Equal(t, 4, m.GetOrDefault(4, 4))
}

func TestHashMap_Set(t *testing.T) {
	m := NewHashMap[int, int]()
	m.Set(1, 1)
	m.Set(2, 2)
	m.Set(3, 3)
	assert.Contains(t, m.ToMap(), 1)
	assert.Contains(t, m.ToMap(), 2)
	assert.Contains(t, m.ToMap(), 3)
}

func TestHashMap_Remove(t *testing.T) {
	m := NewHashMap[int, int]()
	m.Set(1, 1)
	m.Set(2, 2)
	m.Set(3, 3)
	m.Remove(1)
	m.Remove(3)
	assert.NotContains(t, m.ToMap(), 1)
	assert.Contains(t, m.ToMap(), 2)
	assert.NotContains(t, m.ToMap(), 3)
}

func TestHashMap_Keys(t *testing.T) {
	m := NewHashMap[int, int]()
	m.Set(1, 1)
	m.Set(2, 2)
	m.Set(3, 3)
	keys := m.Keys()
	sort.Ints(keys)
	assert.Equal(t, []int{1, 2, 3}, keys)
}

func TestHashMap_Values(t *testing.T) {
	m := NewHashMap[int, int]()
	m.Set(1, 1)
	m.Set(2, 2)
	m.Set(3, 3)
	values := m.Values()
	sort.Ints(values)
	assert.Equal(t, []int{1, 2, 3}, values)
}

func TestHashMap_Clear(t *testing.T) {
	m := NewHashMap[int, int]()
	m.Set(1, 1)
	m.Set(2, 2)
	m.Set(3, 3)
	m.Clear()
	assert.True(t, m.IsEmpty())
}

func TestHashMap_ContainsValue(t *testing.T) {
	m := NewHashMap[int, int]()
	m.Set(1, 1)
	m.Set(2, 2)
	m.Set(3, 3)
	assert.True(t, m.ContainsValue(func(value int) bool { return value == 1 }))
	assert.True(t, m.ContainsValue(func(value int) bool { return value == 2 }))
	assert.True(t, m.ContainsValue(func(value int) bool { return value == 3 }))
}

func TestHashMap_ContainsKey(t *testing.T) {
	m := NewHashMap[int, int]()
	m.Set(1, 1)
	m.Set(2, 2)
	m.Set(3, 3)
	assert.True(t, m.ContainsKey(1))
	assert.True(t, m.ContainsKey(2))
	assert.True(t, m.ContainsKey(3))
}

func TestHashMap_IsEmpty(t *testing.T) {
	m := NewHashMap[int, int]()
	assert.True(t, m.IsEmpty())
	m.Set(1, 1)
	m.Set(2, 2)
	m.Set(3, 3)
	assert.False(t, m.IsEmpty())
}

func TestHashMap_IsNotEmpty(t *testing.T) {
	m := NewHashMap[int, int]()
	assert.False(t, m.IsNotEmpty())
	m.Set(1, 1)
	m.Set(2, 2)
	m.Set(3, 3)
	assert.True(t, m.IsNotEmpty())

}

func TestHashMap_Range(t *testing.T) {
	m := NewHashMap[int, int]()
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
