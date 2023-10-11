package tree

import (
	"testing"

	"github.com/wardonne/gopi/support/compare"
)

func BenchmarkSyncRBTree_Add(b *testing.B) {
	tree := NewSyncRBTree[int](compare.NewNatureComparator[int](false))
	b.RunParallel(func(p *testing.PB) {
		i := 0
		for p.Next() {
			tree.Add(i)
			i++
		}
	})
}

func BenchmarkSyncRBTree_AddAll(b *testing.B) {
	tree := NewSyncRBTree[int](compare.NewNatureComparator[int](false))
	b.RunParallel(func(p *testing.PB) {
		i := 0
		for p.Next() {
			tree.AddAll(i)
			i++
		}
	})
}

func BenchmarkSyncRBTree_Remove(b *testing.B) {
	tree := NewSyncRBTree[int](compare.NewNatureComparator[int](false))
	b.RunParallel(func(p *testing.PB) {
		i := 0
		for p.Next() {
			tree.Add(i)
			i++
		}
	})
	b.ResetTimer()
	b.RunParallel(func(p *testing.PB) {
		i := 0
		for p.Next() {
			tree.Remove(i)
			i++
		}
	})
}

func BenchmarkSyncRBTree_First(b *testing.B) {
	tree := NewSyncRBTree[int](compare.NewNatureComparator[int](false))
	b.RunParallel(func(p *testing.PB) {
		i := 0
		for p.Next() {
			tree.Add(i)
			i++
		}
	})
	b.ResetTimer()
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			tree.First()
		}
	})
}

func BenchmarkSyncRBTree_Last(b *testing.B) {
	tree := NewSyncRBTree[int](compare.NewNatureComparator[int](false))
	b.RunParallel(func(p *testing.PB) {
		i := 0
		for p.Next() {
			tree.Add(i)
			i++
		}
	})
	b.ResetTimer()
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			tree.Last()
		}
	})
}
