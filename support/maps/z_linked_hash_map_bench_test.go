package maps

import "testing"

func BenchmarkLinkedHashMap_Get(b *testing.B) {
	m := NewLinkedHashMap[int, int]()
	for i := 0; i < b.N; i++ {
		m.Set(1, 1)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Get(i)
	}
}

func BenchmarkLinkedHashMap_GetOrDefault(b *testing.B) {
	m := NewLinkedHashMap[int, int]()
	for i := 0; i < b.N; i++ {
		m.Set(1, 1)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.GetOrDefault(i, i)
	}
}

func BenchmarkLinkedHashMap_Set(b *testing.B) {
	m := NewLinkedHashMap[int, int]()
	for i := 0; i < b.N; i++ {
		m.Set(1, 1)
	}
}

func BenchmarkLinkedHashMap_Remove(b *testing.B) {
	m := NewLinkedHashMap[int, int]()
	for i := 0; i < b.N; i++ {
		m.Set(1, 1)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Remove(i)
	}
}
