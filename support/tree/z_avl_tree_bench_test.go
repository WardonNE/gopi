package tree

import (
	"testing"

	"github.com/wardonne/gopi/support/compare"
)

func BenchmarkAVLTree_Add(b *testing.B) {
	tree := NewAVLTree[int](compare.NewNatureComparator[int](false))
	for i := 0; i < b.N; i++ {
		tree.Add(i)
	}
}

func BenchmarkAVLTree_AddAll(b *testing.B) {
	tree := NewAVLTree[int](compare.NewNatureComparator[int](false))
	for i := 0; i < b.N; i++ {
		tree.AddAll(i)
	}
}

func BenchmarkAVLTree_Remove(b *testing.B) {
	tree := NewAVLTree[int](compare.NewNatureComparator[int](false))
	for i := 0; i < b.N; i++ {
		tree.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.Remove(i)
	}
}

func BenchmarkAVLTree_First(b *testing.B) {
	tree := NewAVLTree[int](compare.NewNatureComparator[int](false))
	for i := 0; i < b.N; i++ {
		tree.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.First()
	}
}

func BenchmarkAVLTree_Last(b *testing.B) {
	tree := NewAVLTree[int](compare.NewNatureComparator[int](false))
	for i := 0; i < b.N; i++ {
		tree.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.Last()
	}
}
