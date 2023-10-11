package queue

import "testing"

func BenchmarkLinkedBlockingQueue(b *testing.B) {
	q := NewLinkedBlockingQueue[*int](1000)
	go func() {
		for i := 0; i < b.N; i++ {
			q.EnqueueWithBlock(intptr(i))
		}
	}()
	for i := 0; i < b.N; i++ {
		q.DequeueWithBlock()
	}
}
