package queue

import (
	"testing"
)

func BenchmarkArrayBlockingQueue(b *testing.B) {
	q := NewArrayBlockingQueue[*int](1000)
	go func() {
		for i := 0; i < b.N; i++ {
			q.EnqueueWithBlock(intptr(i))
		}
	}()
	for i := 0; i < b.N; i++ {
		q.DequeueWithBlock()
	}
}
