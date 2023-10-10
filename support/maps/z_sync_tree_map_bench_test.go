package maps

import (
	"testing"

	"github.com/wardonne/gopi/support/compare"
)

func BenchmarkSyncTreeMap_Get(b *testing.B) {
	m := NewSyncTreeMap[int, int](compare.NewNatureComparator[int](false))
	b.RunParallel(func(p *testing.PB) {
		i := 0
		for p.Next() {
			m.Set(i, i)
			i++
		}
	})
	b.ResetTimer()
	b.RunParallel(func(p *testing.PB) {
		i := 0
		for p.Next() {
			m.Get(i)
			i++
		}
	})
}

func BenchmarkSyncTreeMap_GetOrDefault(b *testing.B) {
	m := NewSyncTreeMap[int, int](compare.NewNatureComparator[int](false))
	b.RunParallel(func(p *testing.PB) {
		i := 0
		for p.Next() {
			m.Set(i, i)
			i++
		}
	})
	b.ResetTimer()
	b.RunParallel(func(p *testing.PB) {
		i := 0
		for p.Next() {
			m.GetOrDefault(i, i)
			i++
		}
	})
}

func BenchmarkSyncTreeMap_Set(b *testing.B) {
	m := NewSyncTreeMap[int, int](compare.NewNatureComparator[int](false))
	b.RunParallel(func(p *testing.PB) {
		i := 0
		for p.Next() {
			m.Set(i, i)
			i++
		}
	})
}

func BenchmarkSyncTreeMap_Remove(b *testing.B) {
	m := NewSyncTreeMap[int, int](compare.NewNatureComparator[int](false))
	b.RunParallel(func(p *testing.PB) {
		i := 0
		for p.Next() {
			m.Set(i, i)
			i++
		}
	})
	b.ResetTimer()
	b.RunParallel(func(p *testing.PB) {
		i := 0
		for p.Next() {
			m.Remove(i)
			i++
		}
	})
}
