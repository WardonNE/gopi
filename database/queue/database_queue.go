package queue

import "github.com/wardonne/gopi/support/queue"

type Queue[E any] interface {
	queue.Queue[E]
}
