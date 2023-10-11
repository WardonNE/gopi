package queue

import "testing"

func BenchmarkPriorityBlockingQueue(b *testing.B) {
	q := NewPriorityBlockingQueue[*int](1000, intptrCmp{})
	go func() {
		for i := 0; i < b.N; i++ {
			q.EnqueueWithBlock(intptr(i))
		}
	}()
	for i := 0; i < b.N; i++ {
		q.DequeueWithBlock()
	}
}
