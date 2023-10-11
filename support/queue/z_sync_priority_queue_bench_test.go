package queue

import "testing"

func BenchmarkSyncPriorityQueue_Enqueue(b *testing.B) {
	q := NewSyncPriorityQueue[int](comparator())
	b.RunParallel(func(p *testing.PB) {
		i := 0
		for p.Next() {
			q.Enqueue(i)
			i++
		}
	})
}

func BenchmarkSyncPriorityQueue_Dequeue(b *testing.B) {
	q := NewSyncPriorityQueue[int](comparator())
	b.RunParallel(func(p *testing.PB) {
		i := 0
		for p.Next() {
			q.Enqueue(i)
			i++
		}
	})
	b.ResetTimer()
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			q.Dequeue()
		}
	})
}
