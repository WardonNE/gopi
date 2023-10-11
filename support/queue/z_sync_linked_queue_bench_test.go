package queue

import "testing"

func BenchmarkSyncLinkedQueue_Enqueue(b *testing.B) {
	q := NewSyncLinkedQueue[int]()
	b.RunParallel(func(p *testing.PB) {
		i := 0
		for p.Next() {
			q.Enqueue(i)
			i++
		}
	})
}

func BenchmarkSyncLinkedQueue_Dequeue(b *testing.B) {
	q := NewSyncLinkedQueue[int]()
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
