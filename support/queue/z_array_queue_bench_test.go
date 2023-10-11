package queue

import "testing"

func BenchmarkArrayQueue_Enqueue(b *testing.B) {
	q := NewArrayQueue[int]()
	for i := 0; i < b.N; i++ {
		q.Enqueue(i)
	}
}

func BenchmarkArrayQueue_Dequeue(b *testing.B) {
	q := NewArrayQueue[int]()
	for i := 0; i < b.N; i++ {
		q.Enqueue(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Dequeue()
	}
}
