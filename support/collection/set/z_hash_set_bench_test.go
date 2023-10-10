package set

import "testing"

func BenchmarkHashSet_Add(b *testing.B) {
	set := NewHashSet[int]()
	for i := 0; i < b.N; i++ {
		set.Add(i)
	}
}

func BenchmarkHashSet_AddAll(b *testing.B) {
	set := NewHashSet[int]()
	for i := 0; i < b.N; i++ {
		set.AddAll(i)
	}
}

func BenchmarkHashSet_Remove(b *testing.B) {
	set := NewHashSet[int]()
	for i := 0; i < b.N; i++ {
		set.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set.Remove(func(value int) bool {
			return value == i
		})
	}
}
