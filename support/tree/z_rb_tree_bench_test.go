package tree

import (
	"testing"

	"github.com/wardonne/gopi/support/compare"
)

func BenchmarkRBTree_Add(b *testing.B) {
	tree := NewRBTree[int](compare.NewNatureComparator[int](false))
	for i := 0; i < b.N; i++ {
		tree.Add(i)
	}
}

func BenchmarkRBTree_AddAll(b *testing.B) {
	tree := NewRBTree[int](compare.NewNatureComparator[int](false))
	for i := 0; i < b.N; i++ {
		tree.AddAll(i)
	}
}

func BenchmarkRBTree_Remove(b *testing.B) {
	tree := NewRBTree[int](compare.NewNatureComparator[int](false))
	for i := 0; i < b.N; i++ {
		tree.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.Remove(i)
	}
}

func BenchmarkRBTree_First(b *testing.B) {
	tree := NewRBTree[int](compare.NewNatureComparator[int](false))
	for i := 0; i < b.N; i++ {
		tree.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.First()
	}
}

func BenchmarkRBTree_Last(b *testing.B) {
	tree := NewRBTree[int](compare.NewNatureComparator[int](false))
	for i := 0; i < b.N; i++ {
		tree.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.Last()
	}
}
