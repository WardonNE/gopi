package maps

import "testing"

func BenchmarkSyncLinkedHashMap_Get(b *testing.B) {
	m := NewSyncLinkedHashMap[int, int]()
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

func BenchmarkSyncLinkedHashMap_GetOrDefault(b *testing.B) {
	m := NewSyncLinkedHashMap[int, int]()
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

func BenchmarkSyncLinkedHashMap_Set(b *testing.B) {
	m := NewSyncLinkedHashMap[int, int]()
	b.RunParallel(func(p *testing.PB) {
		i := 0
		for p.Next() {
			m.Set(i, i)
			i++
		}
	})
}

func BenchmarkSyncLinkedHashMap_Remove(b *testing.B) {
	m := NewSyncLinkedHashMap[int, int]()
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
