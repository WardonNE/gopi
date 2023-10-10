package set

import "testing"

func BenchmarkLinkedHashSet_Add(b *testing.B) {
	set := NewLinkedHashSet[int]()
	for i := 0; i < b.N; i++ {
		set.Add(i)
	}
}

func BenchmarkLinkedHashSet_AddAll(b *testing.B) {
	set := NewLinkedHashSet[int]()
	for i := 0; i < b.N; i++ {
		set.AddAll(i)
	}
}

func BenchmarkLinkedHashSet_Remove(b *testing.B) {
	set := NewLinkedHashSet[int]()
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

func BenchmarkLinkedHashSet_RemoveAt(b *testing.B) {
	set := NewLinkedHashSet[int]()
	for i := 0; i < b.N; i++ {
		set.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set.RemoveAt(0)
	}
}

func BenchmarkLinkedHashSet_Get(b *testing.B) {
	set := NewLinkedHashSet[int]()
	for i := 0; i < b.N; i++ {
		set.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set.Get(i)
	}
}

func BenchmarkLinkedHashSet_Pop(b *testing.B) {
	set := NewLinkedHashSet[int]()
	for i := 0; i < b.N; i++ {
		set.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set.Pop()
	}
}

func BenchmarkLinkedHashSet_Shift(b *testing.B) {
	set := NewLinkedHashSet[int]()
	for i := 0; i < b.N; i++ {
		set.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set.Shift()
	}
}

func BenchmarkLinkedHashSet_Unshift(b *testing.B) {
	set := NewLinkedHashSet[int]()
	for i := 0; i < b.N; i++ {
		set.Unshift(i)
	}
}

func BenchmarkLinkedHashSet_UnshiftAll(b *testing.B) {
	set := NewLinkedHashSet[int]()
	for i := 0; i < b.N; i++ {
		set.UnshiftAll(i)
	}
}
