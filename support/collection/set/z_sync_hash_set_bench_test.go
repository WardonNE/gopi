package set

import "testing"

func BenchmarkSyncHashSet_Add(b *testing.B) {
	set := NewSyncHashSet[int]()
	b.RunParallel(func(p *testing.PB) {
		i := 0
		for p.Next() {
			set.Add(i)
			i++
		}
	})
}

func BenchmarkSyncHashSet_AddAll(b *testing.B) {
	set := NewSyncHashSet[int]()
	b.RunParallel(func(p *testing.PB) {
		i := 0
		for p.Next() {
			set.AddAll(i)
			i++
		}
	})
}

func BenchmarkSyncHashSet_Remove(b *testing.B) {
	set := NewSyncHashSet[int]()
	b.RunParallel(func(p *testing.PB) {
		i := 0
		for p.Next() {
			set.Add(i)
			i++
		}
	})
	b.ResetTimer()
	b.RunParallel(func(p *testing.PB) {
		i := 0
		for p.Next() {
			set.Remove(func(value int) bool {
				return value == i
			})
			i++
		}
	})
}
