package tree

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wardonne/gopi/support/compare"
)

func TestAVLTree_MarshalJSON(t *testing.T) {
	tree := NewAVLTree[int](compare.NewNatureComparator[int](false), 3, 2, 1, 4, 5)
	bytes, err := json.Marshal(tree)
	assert.Nil(t, err)
	assert.JSONEq(t, `[1,2,3,4,5]`, string(bytes))
}

func TestAVLTree_UnmarshalJSON(t *testing.T) {
	tree := NewAVLTree[int](compare.NewNatureComparator[int](false))
	err := json.Unmarshal([]byte(`[1,3,4,2,5]`), tree)
	assert.Nil(t, err)
	assert.Equal(t, NewAVLTree[int](compare.NewNatureComparator[int](false), 1, 3, 4, 2, 5), tree)
}

func TestAVLTree_ToArray(t *testing.T) {
	tree := NewAVLTree[int](compare.NewNatureComparator[int](false), 1, 3, 4, 2, 5)
	assert.Equal(t, []int{1, 2, 3, 4, 5}, tree.ToArray())
}

func TestAVLTree_FromArray(t *testing.T) {
	tree := NewAVLTree[int](compare.NewNatureComparator[int](false))
	tree.FromArray([]int{1, 3, 4, 2, 5})
	assert.Equal(t, NewAVLTree[int](compare.NewNatureComparator[int](false), 1, 3, 4, 2, 5), tree)
}

func TestAVLTree_String(t *testing.T) {
	tree := NewAVLTree[int](compare.NewNatureComparator[int](false), 1, 3, 4, 2, 5)
	assert.Equal(t, fmt.Sprintf("%v", []int{1, 2, 3, 4, 5}), tree.String())
}

func TestAVLTree_Clone(t *testing.T) {
	tree := NewAVLTree[int](compare.NewNatureComparator[int](false), 1, 3, 4, 2, 5)
	assert.Equal(t, tree.ToArray(), tree.Clone().ToArray())
}

func TestAVLTree_Copy(t *testing.T) {
	tree := NewAVLTree[int](compare.NewNatureComparator[int](false), 1, 3, 4, 2, 5)
	assert.Equal(t, tree.ToArray(), tree.Copy().ToArray())
}

func TestAVLTree_Count(t *testing.T) {
	tree := NewAVLTree[int](compare.NewNatureComparator[int](false), 1, 3, 4, 2, 5)
	assert.Equal(t, 5, tree.Count())
}

func TestAVLTree_IsEmpty(t *testing.T) {
	tree := NewAVLTree[int](compare.NewNatureComparator[int](false))
	assert.True(t, tree.IsEmpty())
	tree.AddAll(1, 3, 4, 2, 5)
	assert.False(t, tree.IsEmpty())
}

func TestAVLTree_IsNotEmpty(t *testing.T) {
	tree := NewAVLTree[int](compare.NewNatureComparator[int](false))
	assert.False(t, tree.IsNotEmpty())
	tree.AddAll(1, 3, 4, 2, 5)
	assert.True(t, tree.IsNotEmpty())
}

func TestAVLTree_Contains(t *testing.T) {
	tree := NewAVLTree[int](compare.NewNatureComparator[int](false), 1, 3, 4, 2, 5)
	assert.True(t, tree.Contains(1, 2, 3, 4, 5))
	assert.False(t, tree.Contains(1, 2, 3, 4, 5, 6))
}

func TestAVLTree_ContainsAny(t *testing.T) {
	tree := NewAVLTree[int](compare.NewNatureComparator[int](false), 1, 3, 4, 2, 5)
	assert.True(t, tree.ContainsAny(1, 6, 7, 8, 9))
	assert.False(t, tree.ContainsAny(6, 7, 8, 9, 10))
}

func TestAVLTree_Add(t *testing.T) {
	tree := NewAVLTree[int](compare.NewNatureComparator[int](false))
	tree.Add(1)
	tree.Add(3)
	tree.Add(4)
	tree.Add(5)
	tree.Add(2)
	assert.True(t, tree.Contains(1, 2, 3, 4, 5))
}

func TestAVLTree_AddAll(t *testing.T) {
	tree := NewAVLTree[int](compare.NewNatureComparator[int](false))
	tree.AddAll(1, 2, 3, 4, 5)
	assert.True(t, tree.Contains(1, 2, 3, 4, 5))
}

func TestAVLTree_Remove(t *testing.T) {
	tree := NewAVLTree[int](compare.NewNatureComparator[int](false), 1, 2, 3, 4, 5)
	tree.Remove(1)
	tree.Remove(4)
	assert.Equal(t, 3, tree.Count())
	assert.Equal(t, []int{2, 3, 5}, tree.ToArray())
}

func TestAVLTree_Clear(t *testing.T) {
	tree := NewAVLTree[int](compare.NewNatureComparator[int](false), 1, 2, 3, 4, 5)
	tree.Clear()
	assert.True(t, tree.IsEmpty())
}

func TestAVLTree_First(t *testing.T) {
	tree := NewAVLTree[int](compare.NewNatureComparator[int](false))
	value, ok := tree.First()
	assert.False(t, ok)
	assert.Zero(t, value)
	tree.AddAll(1, 2, 3, 4, 5)
	value2, ok := tree.First()
	assert.True(t, ok)
	assert.Equal(t, 1, value2)
}

func TestAVLTree_Last(t *testing.T) {
	tree := NewAVLTree[int](compare.NewNatureComparator[int](false))
	value, ok := tree.Last()
	assert.False(t, ok)
	assert.Zero(t, value)
	tree.AddAll(1, 2, 3, 4, 5)
	value2, ok := tree.Last()
	assert.True(t, ok)
	assert.Equal(t, 5, value2)
}

func TestAVLTree_Range(t *testing.T) {
	tree := NewAVLTree[int](compare.NewNatureComparator[int](false), 1, 2, 3, 4, 5)
	values := []int{}
	tree.Range(func(value int) bool {
		values = append(values, value)
		return true
	})
	assert.Equal(t, []int{1, 2, 3, 4, 5}, values)
}
