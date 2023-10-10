package set

import "testing"

func BenchmarkSyncLinkedHashSet_Add(b *testing.B) {
	set := NewSyncLinkedHashSet[int]()
	for i := 0; i < b.N; i++ {
		set.Add(i)
	}
}

func BenchmarkSyncLinkedHashSet_AddAll(b *testing.B) {
	set := NewSyncLinkedHashSet[int]()
	for i := 0; i < b.N; i++ {
		set.AddAll(i)
	}
}

func BenchmarkSyncLinkedHashSet_Remove(b *testing.B) {
	set := NewSyncLinkedHashSet[int]()
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

func BenchmarkSyncLinkedHashSet_RemoveAt(b *testing.B) {
	set := NewSyncLinkedHashSet[int]()
	for i := 0; i < b.N; i++ {
		set.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set.RemoveAt(0)
	}
}

func BenchmarkSyncLinkedHashSet_Get(b *testing.B) {
	set := NewSyncLinkedHashSet[int]()
	for i := 0; i < b.N; i++ {
		set.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set.Get(i)
	}
}

func BenchmarkSyncLinkedHashSet_Pop(b *testing.B) {
	set := NewSyncLinkedHashSet[int]()
	for i := 0; i < b.N; i++ {
		set.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set.Pop()
	}
}

func BenchmarkSyncLinkedHashSet_Shift(b *testing.B) {
	set := NewSyncLinkedHashSet[int]()
	for i := 0; i < b.N; i++ {
		set.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set.Shift()
	}
}

func BenchmarkSyncLinkedHashSet_Unshift(b *testing.B) {
	set := NewSyncLinkedHashSet[int]()
	for i := 0; i < b.N; i++ {
		set.Unshift(i)
	}
}

func BenchmarkSyncLinkedHashSet_UnshiftAll(b *testing.B) {
	set := NewSyncLinkedHashSet[int]()
	for i := 0; i < b.N; i++ {
		set.UnshiftAll(i)
	}
}
