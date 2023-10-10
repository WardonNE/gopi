package maps

import "testing"

func BenchmarkHashMap_Get(b *testing.B) {
	m := NewHashMap[int, int]()
	for i := 0; i < b.N; i++ {
		m.Set(1, 1)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Get(i)
	}
}

func BenchmarkHashMap_GetOrDefault(b *testing.B) {
	m := NewHashMap[int, int]()
	for i := 0; i < b.N; i++ {
		m.Set(1, 1)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.GetOrDefault(i, i)
	}
}

func BenchmarkHashMap_Set(b *testing.B) {
	m := NewHashMap[int, int]()
	for i := 0; i < b.N; i++ {
		m.Set(1, 1)
	}
}

func BenchmarkHashMap_Remove(b *testing.B) {
	m := NewHashMap[int, int]()
	for i := 0; i < b.N; i++ {
		m.Set(1, 1)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Remove(i)
	}
}
