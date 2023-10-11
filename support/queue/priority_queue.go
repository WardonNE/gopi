package queue

import (
	"encoding/json"
	"sync"

	"github.com/wardonne/gopi/support/compare"
)

type PriorityQueue[E any] struct {
	lock       *sync.Mutex
	items      []E
	comparator compare.Comparator[E]
	size       int
}

func NewPriorityQueue[E any](comparator compare.Comparator[E]) *PriorityQueue[E] {
	queue := new(PriorityQueue[E])
	queue.lock = new(sync.Mutex)
	queue.comparator = comparator
	return queue
}

func (q *PriorityQueue[E]) MarshalJSON() ([]byte, error) {
	return json.Marshal(q.items)
}

func (q *PriorityQueue[E]) UnmarshalJSON(data []byte) error {
	items := []E{}
	err := json.Unmarshal(data, &items)
	if err != nil {
		return nil
	}
	for _, item := range items {
		q.Enqueue(item)
	}
	return nil
}

func (q *PriorityQueue[E]) ToArray() []E {
	return q.items
}

func (q *PriorityQueue[E]) FromArray(values []E) {
	for _, value := range values {
		q.Enqueue(value)
	}
}

func (q *PriorityQueue[E]) Count() int {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.size
}

func (q *PriorityQueue[E]) IsEmpty() bool {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.size == 0
}

func (q *PriorityQueue[E]) IsNotEmpty() bool {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.size > 0
}

func (q *PriorityQueue[E]) Clear() {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.items = make([]E, 0)
	q.size = 0
}

func (q *PriorityQueue[E]) Peek() (value E) {
	q.lock.Lock()
	defer q.lock.Unlock()
	if q.size == 0 {
		return
	}
	return q.items[0]
}

func (q *PriorityQueue[E]) Enqueue(value E) bool {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.items = append(q.items, value)
	q.size++
	for index := q.size - 1; q.less(index, (index-1)/2); index = (index - 1) / 2 {
		q.swap(index, (index-1)/2)
	}
	return true
}

func (q *PriorityQueue[E]) Dequeue() (value E, ok bool) {
	q.lock.Lock()
	defer q.lock.Unlock()
	if q.size == 0 {
		return
	}
	value = q.items[0]
	ok = true
	q.swap(0, q.size-1)
	q.items = q.items[:q.size-1]
	q.size--
	index := 0
	lastIndex := q.size - 1
	for {
		leftIndex := index*2 + 1
		if leftIndex > lastIndex || leftIndex < 0 {
			break
		}
		swapIndex := leftIndex
		if rightIndex := leftIndex + 1; rightIndex <= lastIndex && q.less(rightIndex, leftIndex) {
			swapIndex = rightIndex
		}
		if !q.less(swapIndex, index) {
			break
		}
		q.swap(swapIndex, index)
		index = swapIndex
	}
	return
}

func (q *PriorityQueue[E]) less(i, j int) bool {
	return q.comparator.Compare(q.items[i], q.items[j]) < 0
}

func (q *PriorityQueue[E]) swap(i, j int) {
	q.items[i], q.items[j] = q.items[j], q.items[i]
}
