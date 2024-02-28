package list

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/wardonne/gopi/support/collection"
	"github.com/wardonne/gopi/support/compare"
)

// LinkedList linked list
type LinkedList[E any] struct {
	first *element[E]
	last  *element[E]
	size  int
}

// NewLinkedList create a new LinkedList
func NewLinkedList[E any](values ...E) *LinkedList[E] {
	linkedList := new(LinkedList[E])
	linkedList.FromArray(values)
	return linkedList
}

func (l *LinkedList[E]) remove(el *element[E]) {
	if el.prev != nil {
		el.prev.next = el.next
	} else {
		l.first = el.next
	}
	if el.next != nil {
		el.next.prev = el.prev
	} else {
		l.last = el.prev
	}
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
		el := l.first
		for i := 0; i < index; i++ {
			el = el.next
		}
		return el
	}
	el := l.last
	for i := l.size - 1; i > index; i-- {
		el = el.prev
	}
	return el
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
	return fmt.Sprintf("%v", l.ToArray())
}

func (l *LinkedList[E]) Sort(comparator compare.Comparator[E]) {
	items := l.ToArray()
	sort.Slice(items, func(i, j int) bool {
		return comparator.Compare(items[i], items[j]) < 0
	})
	l.Clear()
	l.PushAll(items...)
}

func (l *LinkedList[E]) Clone() collection.Interface[E] {
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

func (l *LinkedList[E]) First() E {
	return l.node(0).Value
}

func (l *LinkedList[E]) Last() E {
	return l.node(l.size - 1).Value
}

func (l *LinkedList[E]) FirstWhere(matcher collection.Matcher[E]) (e E, err error) {
	for el := l.first; el != nil; el = el.next {
		if matcher(el.Value) {
			return el.Value, nil
		}
	}
	return e, ErrElementNotFound
}

func (l *LinkedList[E]) LastWhere(matcher collection.Matcher[E]) (e E, err error) {
	for el := l.last; el != nil; el = el.prev {
		if matcher(el.Value) {
			return el.Value, nil
		}
	}
	return e, ErrElementNotFound
}

func (l *LinkedList[E]) Pop() (value E) {
	if l.size == 0 {
		return
	}
	el := l.last
	l.remove(el)
	return el.Value
}

func (l *LinkedList[E]) Shift() (value E) {
	if l.size == 0 {
		return
	}
	el := l.first
	l.remove(el)
	return el.Value
}

func (l *LinkedList[E]) Contains(matcher collection.Matcher[E]) bool {
	for el := l.first; el != nil; el = el.next {
		if matcher(el.Value) {
			return true
		}
	}
	return false
}

func (l *LinkedList[E]) IndexOf(matcher collection.Matcher[E]) int {
	if l.size == 0 {
		return -1
	}
	for el, index := l.first, 0; el.next != nil; el, index = el.next, index+1 {
		if matcher(el.Value) {
			return index
		}
	}
	return -1
}

func (l *LinkedList[E]) LastIndexOf(matcher collection.Matcher[E]) int {
	if l.size == 0 {
		return -1
	}
	for el, index := l.last, l.size-1; el != nil; el, index = el.prev, index-1 {
		if matcher(el.Value) {
			return index
		}
	}
	return -1
}

func (l *LinkedList[E]) Where(matcher collection.Matcher[E]) collection.Interface[E] {
	l2 := NewLinkedList[E]()
	for el := l.first; el != nil; el = el.next {
		if matcher(el.Value) {
			l2.Add(el.Value)
		}
	}
	return l2
}

func (l *LinkedList[E]) SubList(from, to int) Interface[E] {
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
	el := &element[E]{Value: value}
	if l.first == nil {
		l.first = el
		l.last = el
	} else {
		el.prev = l.last
		l.last.next = el
		l.last = el
	}
	l.size++
}

func (l *LinkedList[E]) PushAll(values ...E) {
	for _, value := range values {
		l.Push(value)
	}
}

func (l *LinkedList[E]) Unshift(value E) {
	el := &element[E]{Value: value}
	if l.first == nil {
		l.first = el
		l.last = el
	} else {
		el.next = l.first
		l.first.prev = el
		l.first = el
	}
	l.size++
}

func (l *LinkedList[E]) UnshiftAll(values ...E) {
	for _, value := range values {
		l.Unshift(value)
	}
}

func (l *LinkedList[E]) InsertBefore(index int, value E) {
	if index < 0 || index >= l.size {
		panic(ErrIndexOutOfRange)
	}
	el := &element[E]{Value: value}
	at := l.node(index)
	el.next = at
	el.prev = at.prev
	if at.prev != nil {
		at.prev.next = el
	} else {
		l.first = el
	}
	at.prev = el
	l.size++
}

func (l *LinkedList[E]) InsertAfter(index int, value E) {
	if index < 0 || index >= l.size {
		panic(ErrIndexOutOfRange)
	}
	el := &element[E]{Value: value}
	at := l.node(index)
	el.prev = at
	el.next = at.next
	if at.next != nil {
		at.next.prev = el
	} else {
		l.last = el
	}
	at.next = el
	l.size++
}

func (l *LinkedList[E]) RemoveAt(index int) {
	if index < 0 || index >= l.size {
		panic(ErrIndexOutOfRange)
	}
	if el := l.node(index); el != nil {
		l.remove(el)
	}
}

func (l *LinkedList[E]) Remove(matcher collection.Matcher[E]) {
	if l.size == 0 {
		return
	}
	els := make([]*element[E], 0)
	for el := l.first; el != nil; el = el.next {
		if matcher(el.Value) {
			els = append(els, el)
		}
	}
	for _, el := range els {
		l.remove(el)
	}
}

func (l *LinkedList[E]) Clear() {
	l.first = nil
	l.last = nil
	l.size = 0
}

func (l *LinkedList[E]) Range(callback func(value E) bool) {
	for el := l.first; el != nil; el = el.next {
		if !callback(el.Value) {
			break
		}
	}
}

func (l *LinkedList[E]) ReverseRange(callback func(value E) bool) {
	for el := l.last; el != nil; el = el.prev {
		if !callback(el.Value) {
			break
		}
	}
}

func (l *LinkedList[E]) Map(callback func(value E) E) {
	for el := l.first; el != nil; el = el.next {
		el.Value = callback(el.Value)
	}
}
