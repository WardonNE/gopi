package list

import "testing"

func BenchmarkArrayList_Add(b *testing.B) {
	list := NewArrayList[int]()
	for i := 0; i < b.N; i++ {
		list.Add(i)
	}
}

func BenchmarkArrayList_AddAll(b *testing.B) {
	list := NewArrayList[int]()
	for i := 0; i < b.N; i++ {
		list.AddAll(i)
	}
}

func BenchmarkArrayList_Push(b *testing.B) {
	list := NewArrayList[int]()
	for i := 0; i < b.N; i++ {
		list.Push(i)
	}
}

func BenchmarkArrayList_PushAll(b *testing.B) {
	list := NewArrayList[int]()
	for i := 0; i < b.N; i++ {
		list.PushAll(i)
	}
}

func BenchmarkArrayList_Unshift(b *testing.B) {
	list := NewArrayList[int]()
	for i := 0; i < b.N; i++ {
		list.Unshift(i)
	}
}

func BenchmarkArrayList_UnshiftAll(b *testing.B) {
	list := NewArrayList[int]()
	for i := 0; i < b.N; i++ {
		list.UnshiftAll(i)
	}
}

func BenchmarkArrayList_Get(b *testing.B) {
	list := NewArrayList[int]()
	for i := 0; i < b.N; i++ {
		list.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		list.Get(i)
	}
}

func BenchmarkArrayList_Pop(b *testing.B) {
	list := NewArrayList[int]()
	for i := 0; i < b.N; i++ {
		list.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		list.Pop()
	}
}

func BenchmarkArrayList_Shift(b *testing.B) {
	list := NewArrayList[int]()
	for i := 0; i < b.N; i++ {
		list.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		list.Shift()
	}
}

func BenchmarkArrayList_RemoveAt(b *testing.B) {
	list := NewArrayList[int]()
	for i := 0; i < b.N; i++ {
		list.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		list.RemoveAt(0)
	}
}

func BenchmarkArrayList_Remove(b *testing.B) {
	list := NewArrayList[int]()
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
