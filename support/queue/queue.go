package queue

type Queue[E comparable] interface {
	Peek() E
	Enqueue(value E)
	Dequeue() E
}
