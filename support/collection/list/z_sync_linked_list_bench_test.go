package list

import "testing"

func BenchmarkSyncLinkedList_Add(b *testing.B) {
	list := NewSyncLinkedList[int]()
	b.RunParallel(func(p *testing.PB) {
		i := 0
		for p.Next() {
			list.Add(i)
			i++
		}
	})

}

func BenchmarkSyncLinkedList_AddAll(b *testing.B) {
	list := NewSyncLinkedList[int]()
	b.RunParallel(func(p *testing.PB) {
		i := 0
		for p.Next() {
			list.AddAll(i)
			i++
		}
	})
}

func BenchmarkSyncLinkedList_Push(b *testing.B) {
	list := NewSyncLinkedList[int]()
	b.RunParallel(func(p *testing.PB) {
		i := 0
		for p.Next() {
			list.Push(i)
			i++
		}
	})
}

func BenchmarkSyncLinkedList_PushAll(b *testing.B) {
	list := NewSyncLinkedList[int]()
	b.RunParallel(func(p *testing.PB) {
		i := 0
		for p.Next() {
			list.PushAll(i)
			i++
		}
	})

}

func BenchmarkSyncLinkedList_Unshift(b *testing.B) {
	list := NewSyncLinkedList[int]()
	b.RunParallel(func(p *testing.PB) {
		i := 0
		for p.Next() {
			list.Unshift(i)
			i++
		}
	})

}

func BenchmarkSyncLinkedList_UnshiftAll(b *testing.B) {
	list := NewSyncLinkedList[int]()
	b.RunParallel(func(p *testing.PB) {
		i := 0
		for p.Next() {
			list.UnshiftAll(i)
			i++
		}
	})
}

func BenchmarkSyncLinkedList_Get(b *testing.B) {
	list := NewSyncLinkedList[int]()
	for i := 0; i < b.N; i++ {
		list.Add(i)
	}
	b.ResetTimer()
	b.RunParallel(func(p *testing.PB) {
		i := 0
		for p.Next() {
			list.Get(i)
			i++
		}
	})
}

func BenchmarkSyncLinkedList_Pop(b *testing.B) {
	list := NewSyncLinkedList[int]()
	for i := 0; i < b.N; i++ {
		list.Add(i)
	}
	b.ResetTimer()
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			list.Pop()
		}
	})
}

func BenchmarkSyncLinkedList_Shift(b *testing.B) {
	list := NewSyncLinkedList[int]()
	for i := 0; i < b.N; i++ {
		list.Add(i)
	}
	b.ResetTimer()
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			list.Shift()
		}
	})
}

func BenchmarkSyncLinkedList_RemoveAt(b *testing.B) {
	list := NewSyncLinkedList[int]()
	for i := 0; i < b.N; i++ {
		list.Add(i)
	}
	b.ResetTimer()
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			list.RemoveAt(0)
		}
	})
}

func BenchmarkSyncLinkedList_Remove(b *testing.B) {
	list := NewSyncLinkedList[int]()
	for i := 0; i < b.N; i++ {
		list.Add(i)
	}
	b.ResetTimer()
	b.RunParallel(func(p *testing.PB) {
		i := 0
		for p.Next() {
			list.Remove(func(value int) bool {
				return value == i
			})
			i++
		}
	})
}
