package list

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/wardonne/gopi/support/builder"
	"github.com/wardonne/gopi/support/collection"
	"github.com/wardonne/gopi/support/compare"
)

type LinkedList[E any] struct {
	root element[E]
	size int
}

// NewLinkedList[E any] create an empty LinkedList
func NewLinkedList[E any](values ...E) *LinkedList[E] {
	linkedList := new(LinkedList[E])
	linkedList.FromArray(values)
	return linkedList
}

func (l *LinkedList[E]) insert(el, at *element[E]) *element[E] {
	el.prev = at
	el.next = at.next
	el.prev.next = el
	el.next.prev = el
	el.list = l
	l.size++
	return el
}

func (l *LinkedList[E]) insertValue(value E, at *element[E]) *element[E] {
	return l.insert(&element[E]{Value: value}, at)
}

func (l *LinkedList[E]) remove(el *element[E]) {
	el.prev.next = el.next
	el.next.prev = el.prev
	el.next = nil
	el.prev = nil
	el.list = nil
	l.size--
}

func (l *LinkedList[E]) node(index int) *element[E] {
	if index < 0 || index >= l.size {
		return nil
	}
	if index < (l.size >> 1) {
		el := l.root.next
		for i := 0; i < index; i++ {
			el = el.next
		}
		return el
	} else {
		el := l.root.prev
		for i := l.size - 1; i > index; i-- {
			el = el.prev
		}
		return el
	}
}

func (l *LinkedList[E]) MarshalJSON() ([]byte, error) {
	return json.Marshal(l.ToArray())
}

func (l *LinkedList[E]) UnmarshalJSON(data []byte) error {
	values := make([]E, 0)
	if err := json.Unmarshal(data, &values); err != nil {
		return err
	}
	l.PushAll(values...)
	return nil
}

func (l *LinkedList[E]) ToArray() []E {
	values := make([]E, 0, l.size)
	l.Range(func(value E) bool {
		values = append(values, value)
		return true
	})
	return values
}

func (l *LinkedList[E]) FromArray(values []E) {
	l.Clear()
	l.PushAll(values...)
}

func (l *LinkedList[E]) String() string {
	bytes, err := l.MarshalJSON()
	if err != nil {
		sb := builder.NewStringBuilder("[")
		l.Range(func(value E) bool {
			sb.WriteString(fmt.Sprintf("%v", value))
			sb.WriteRune(' ')
			return true
		})
		sb.TrimSpace()
		return sb.String()
	}
	return string(bytes)
}

func (l *LinkedList[E]) Sort(comparator compare.Comparator[E]) {
	items := l.ToArray()
	sort.Slice(items, func(i, j int) bool {
		return comparator.Compare(items[i], items[j]) < 0
	})
	l.Clear()
	l.PushAll(items...)
}

func (l *LinkedList[E]) Clone() collection.Collection[E] {
	l2 := NewLinkedList[E]()
	l.Range(func(value E) bool {
		l2.Push(value)
		return true
	})
	return l2
}

func (l *LinkedList[E]) Copy() *LinkedList[E] {
	return l.Clone().(*LinkedList[E])
}

func (l *LinkedList[E]) Count() int {
	return l.size
}

func (l *LinkedList[E]) IsEmpty() bool {
	return l.size == 0
}

func (l *LinkedList[E]) IsNotEmpty() bool {
	return l.size > 0
}

func (l *LinkedList[E]) Get(index int) E {
	if index < 0 || index >= l.size {
		panic(ErrIndexOutOfRange)
	}
	return l.node(index).Value
}

func (l *LinkedList[E]) Pop() (value E) {
	if l.size == 0 {
		return
	}
	el := l.root.next
	l.remove(el)
	return el.Value
}

func (l *LinkedList[E]) Shift() (value E) {
	if l.size == 0 {
		return
	}
	el := l.root.prev
	l.remove(el)
	return el.Value
}

func (l *LinkedList[E]) Contains(matcher func(value E) bool) bool {
	for el := l.root.next; el != nil; el = el.next {
		if !matcher(el.Value) {
			return false
		}
	}
	return true
}

func (l *LinkedList[E]) ContainsAny(matcher func(value E) bool) bool {
	for el := l.root.next; el != nil; el = el.next {
		if matcher(el.Value) {
			return true
		}
	}
	return false
}

func (l *LinkedList[E]) IndexOf(matcher func(value E) bool) int {
	for el, index := l.root.next, -1; el != nil; el, index = el.next, index+1 {
		if matcher(el.Value) {
			return index
		}
	}
	return -1
}

func (l *LinkedList[E]) LastIndexOf(matcher func(value E) bool) int {
	for el, index := l.root.prev, -1; el != nil; el, index = el.prev, index+1 {
		if matcher(el.Value) {
			return l.size - index - 1
		}
	}
	return -1
}

func (l *LinkedList[E]) SubList(from, to int) List[E] {
	if from < 0 || from >= l.size || to < 0 || to >= l.size {
		panic(ErrIndexOutOfRange)
	}
	start, end := from, to
	if start > end {
		start, end = to, from
	}
	list := NewLinkedList[E]()
	for index := start; index < end; index++ {
		node := l.node(index)
		list.Push(node.Value)
	}
	return list
}

func (l *LinkedList[E]) SubLinkedList(from, to int) *LinkedList[E] {
	return l.SubList(from, to).(*LinkedList[E])
}

func (l *LinkedList[E]) Set(index int, value E) {
	if index < 0 || index >= l.size {
		panic(ErrIndexOutOfRange)
	}
	l.node(index).Value = value
}

func (l *LinkedList[E]) Add(value E) {
	l.Push(value)
}

func (l *LinkedList[E]) AddAll(values ...E) {
	l.PushAll(values...)
}

func (l *LinkedList[E]) Push(value E) {
	l.insertValue(value, l.root.prev)
}

func (l *LinkedList[E]) PushAll(values ...E) {
	for _, value := range values {
		l.Push(value)
	}
}

func (l *LinkedList[E]) Unshift(value E) {
	l.insertValue(value, &l.root)
}

func (l *LinkedList[E]) UnshiftAll(values ...E) {
	l2 := NewLinkedList[E]()
	l2.PushAll(values...)
	l2.ReverseRange(func(value E) bool {
		l.Unshift(value)
		return true
	})
}

func (l *LinkedList[E]) InsertBefore(index int, value E) {
	if index < 0 || index >= l.size {
		panic(ErrIndexOutOfRange)
	}
	l.insertValue(value, l.node(index).prev)
}

func (l *LinkedList[E]) InsertAfter(index int, value E) {
	if index < 0 || index >= l.size {
		panic(ErrIndexOutOfRange)
	}
	l.insertValue(value, l.node(index))
}

func (l *LinkedList[E]) RemoveAt(index int) {
	if el := l.node(index); el == nil {
		return
	} else {
		l.remove(el)
	}
}

func (l *LinkedList[E]) Remove(matcher func(value E) bool) {
	for el := l.root.next; el != nil; el = el.next {
		if matcher(el.Value) {
			l.remove(el)
		}
	}
}

func (l *LinkedList[E]) Clear() {
	l.root = element[E]{}
	l.size = 0
}

func (l *LinkedList[E]) Range(callback func(value E) bool) {
	for el := l.root.next; el != nil; el = el.next {
		if !callback(el.Value) {
			break
		}
	}
}

func (l *LinkedList[E]) ReverseRange(callback func(value E) bool) {
	for el := l.root.prev; el != nil; el = el.prev {
		if !callback(el.Value) {
			break
		}
	}
}

func (l *LinkedList[E]) ReverseRangeElement(callback func(element *element[E]) bool) {
	for el := l.root.prev; el != nil; el = el.prev {
		if !callback(el) {
			break
		}
	}
}

func (l *LinkedList[E]) Map(callback func(value E) E) {
	for el := l.root.prev; el != nil; el = el.next {
		el.Value = callback(el.Value)
	}
}
