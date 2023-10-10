package maps

import "testing"

func BenchmarkSyncHashMap_Get(b *testing.B) {
	m := NewSyncHashMap[int, int]()
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

func BenchmarkSyncHashMap_GetOrDefault(b *testing.B) {
	m := NewSyncHashMap[int, int]()
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

func BenchmarkSyncHashMap_Set(b *testing.B) {
	m := NewSyncHashMap[int, int]()
	b.RunParallel(func(p *testing.PB) {
		i := 0
		for p.Next() {
			m.Set(i, i)
			i++
		}
	})
}

func BenchmarkSyncHashMap_Remove(b *testing.B) {
	m := NewSyncHashMap[int, int]()
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
