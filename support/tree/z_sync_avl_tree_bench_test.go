package tree

import (
	"testing"

	"github.com/wardonne/gopi/support/compare"
)

func BenchmarkSyncAVLTree_Add(b *testing.B) {
	tree := NewSyncAVLTree[int](compare.NewNatureComparator[int](false))
	b.RunParallel(func(p *testing.PB) {
		i := 0
		for p.Next() {
			tree.Add(i)
			i++
		}
	})
}

func BenchmarkSyncAVLTree_AddAll(b *testing.B) {
	tree := NewSyncAVLTree[int](compare.NewNatureComparator[int](false))
	b.RunParallel(func(p *testing.PB) {
		i := 0
		for p.Next() {
			tree.AddAll(i)
			i++
		}
	})
}

func BenchmarkSyncAVLTree_Remove(b *testing.B) {
	tree := NewSyncAVLTree[int](compare.NewNatureComparator[int](false))
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

func BenchmarkSyncAVLTree_First(b *testing.B) {
	tree := NewSyncAVLTree[int](compare.NewNatureComparator[int](false))
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

func BenchmarkSyncAVLTree_Last(b *testing.B) {
	tree := NewSyncAVLTree[int](compare.NewNatureComparator[int](false))
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
