package queue

import "testing"

func BenchmarkSyncArrayQueue_Enqueue(b *testing.B) {
	q := NewSyncArrayQueue[int]()
	b.RunParallel(func(p *testing.PB) {
		i := 0
		for p.Next() {
			q.Enqueue(i)
			i++
		}
	})
}

func BenchmarkSyncArrayQueue_Dequeue(b *testing.B) {
	q := NewSyncArrayQueue[int]()
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
