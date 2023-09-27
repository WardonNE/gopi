package list

import (
	"testing"
)

func BenchmarkSyncArrayList_Add(b *testing.B) {
	list := NewSyncArrayList[int]()
	b.RunParallel(func(p *testing.PB) {
		i := 0
		for p.Next() {
			list.Add(i)
			i++
		}
	})

}

func BenchmarkSyncArrayList_AddAll(b *testing.B) {
	list := NewSyncArrayList[int]()
	b.RunParallel(func(p *testing.PB) {
		i := 0
		for p.Next() {
			list.AddAll(i)
			i++
		}
	})
}

func BenchmarkSyncArrayList_Push(b *testing.B) {
	list := NewSyncArrayList[int]()
	b.RunParallel(func(p *testing.PB) {
		i := 0
		for p.Next() {
			list.Push(i)
			i++
		}
	})
}

func BenchmarkSyncArrayList_PushAll(b *testing.B) {
	list := NewSyncArrayList[int]()
	b.RunParallel(func(p *testing.PB) {
		i := 0
		for p.Next() {
			list.PushAll(i)
			i++
		}
	})

}

func BenchmarkSyncArrayList_Unshift(b *testing.B) {
	list := NewSyncArrayList[int]()
	b.RunParallel(func(p *testing.PB) {
		i := 0
		for p.Next() {
			list.Unshift(i)
			i++
		}
	})

}

func BenchmarkSyncArrayList_UnshiftAll(b *testing.B) {
	list := NewSyncArrayList[int]()
	b.RunParallel(func(p *testing.PB) {
		i := 0
		for p.Next() {
			list.UnshiftAll(i)
			i++
		}
	})
}

func BenchmarkSyncArrayList_Get(b *testing.B) {
	list := NewSyncArrayList[int]()
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

func BenchmarkSyncArrayList_Pop(b *testing.B) {
	list := NewSyncArrayList[int]()
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

func BenchmarkSyncArrayList_Shift(b *testing.B) {
	list := NewSyncArrayList[int]()
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

func BenchmarkSyncArrayList_RemoveAt(b *testing.B) {
	list := NewSyncArrayList[int]()
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

func BenchmarkSyncArrayList_Remove(b *testing.B) {
	list := NewSyncArrayList[int]()
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
