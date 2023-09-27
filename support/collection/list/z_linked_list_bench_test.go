package list

import "testing"

func BenchmarkLinkedList_Add(b *testing.B) {
	list := NewLinkedList[int]()
	for i := 0; i < b.N; i++ {
		list.Add(i)
	}
}

func BenchmarkLinkedList_AddAll(b *testing.B) {
	list := NewLinkedList[int]()
	for i := 0; i < b.N; i++ {
		list.AddAll(i)
	}
}

func BenchmarkLinkedList_Push(b *testing.B) {
	list := NewLinkedList[int]()
	for i := 0; i < b.N; i++ {
		list.Push(i)
	}
}

func BenchmarkLinkedList_PushAll(b *testing.B) {
	list := NewLinkedList[int]()
	for i := 0; i < b.N; i++ {
		list.PushAll(i)
	}
}

func BenchmarkLinkedList_Unshift(b *testing.B) {
	list := NewLinkedList[int]()
	for i := 0; i < b.N; i++ {
		list.Unshift(i)
	}
}

func BenchmarkLinkedList_UnshiftAll(b *testing.B) {
	list := NewLinkedList[int]()
	for i := 0; i < b.N; i++ {
		list.UnshiftAll(i)
	}
}

func BenchmarkLinkedList_Get(b *testing.B) {
	list := NewLinkedList[int]()
	for i := 0; i < b.N; i++ {
		list.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		list.Get(i)
	}
}

func BenchmarkLinkedList_Pop(b *testing.B) {
	list := NewLinkedList[int]()
	for i := 0; i < b.N; i++ {
		list.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		list.Pop()
	}
}

func BenchmarkLinkedList_Shift(b *testing.B) {
	list := NewLinkedList[int]()
	for i := 0; i < b.N; i++ {
		list.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		list.Shift()
	}
}

func BenchmarkLinkedList_RemoveAt(b *testing.B) {
	list := NewLinkedList[int]()
	for i := 0; i < b.N; i++ {
		list.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		list.RemoveAt(0)
	}
}

func BenchmarkLinkedList_Remove(b *testing.B) {
	list := NewLinkedList[int]()
	for i := 0; i < b.N; i++ {
		list.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		list.Remove(func(value int) bool {
			return value == i
		})
	}
}
