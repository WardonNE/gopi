package queue

import "testing"

func BenchmarkLinkedQueue_Enqueue(b *testing.B) {
	q := NewLinkedQueue[int]()
	for i := 0; i < b.N; i++ {
		q.Enqueue(i)
	}
}

func BenchmarkLinkedQueue_Dequeue(b *testing.B) {
	q := NewLinkedQueue[int]()
	for i := 0; i < b.N; i++ {
		q.Enqueue(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Dequeue()
	}
}
