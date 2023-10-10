package maps

import (
	"testing"

	"github.com/wardonne/gopi/support/compare"
)

func BenchmarkTreeMap_Get(b *testing.B) {
	m := NewTreeMap[int, int](compare.NewNatureComparator[int](false))
	for i := 0; i < b.N; i++ {
		m.Set(1, 1)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Get(i)
	}
}

func BenchmarkTreeMap_GetOrDefault(b *testing.B) {
	m := NewTreeMap[int, int](compare.NewNatureComparator[int](false))
	for i := 0; i < b.N; i++ {
		m.Set(1, 1)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.GetOrDefault(i, i)
	}
}

func BenchmarkTreeMap_Set(b *testing.B) {
	m := NewTreeMap[int, int](compare.NewNatureComparator[int](false))
	for i := 0; i < b.N; i++ {
		m.Set(1, 1)
	}
}

func BenchmarkTreeMap_Remove(b *testing.B) {
	m := NewTreeMap[int, int](compare.NewNatureComparator[int](false))
	for i := 0; i < b.N; i++ {
		m.Set(1, 1)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Remove(i)
	}
}
