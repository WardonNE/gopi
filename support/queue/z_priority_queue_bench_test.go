package queue

import "testing"

func BenchmarkPriorityQueue_Enqueue(b *testing.B) {
	q := NewPriorityQueue[int](comparator())
	for i := 0; i < b.N; i++ {
		q.Enqueue(i)
	}
}

func BenchmarkPriorityQueue_Dequeue(b *testing.B) {
	q := NewPriorityQueue[int](comparator())
	for i := 0; i < b.N; i++ {
		q.Enqueue(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Dequeue()
	}
}
