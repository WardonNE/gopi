package list

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/wardonne/gopi/support/collection"
	"github.com/wardonne/gopi/support/compare"
)

type ArrayList[E any] struct {
	items []E
}

func NewArrayList[E any](values ...E) *ArrayList[E] {
	arrayList := new(ArrayList[E])
	arrayList.FromArray(values)
	return arrayList
}

// MarshalJSON implements serializer.JSONSerializer.MarshalJSON
func (l *ArrayList[E]) MarshalJSON() ([]byte, error) {
	return json.Marshal(l.ToArray())
}

// UnmarshalJSON implements serializer.JSONSerializer.UnmarshalJSON
func (l *ArrayList[E]) UnmarshalJSON(data []byte) error {
	values := make([]E, 0)
	if err := json.Unmarshal(data, &values); err != nil {
		return err
	}
	l.items = values
	return nil
}

// ToArray implements serializer.ArraySerializer.ToArray
func (l *ArrayList[E]) ToArray() []E {
	return l.items
}

// FromArray implements serializer.ArraySerializer.FromArray
func (l *ArrayList[E]) FromArray(values []E) {
	l.items = values
}

// String implements support.Stringable
func (l *ArrayList[E]) String() string {
	return fmt.Sprintf("%v", l.items)
}

// Clone implements support.Clonable
func (l *ArrayList[E]) Clone() collection.Collection[E] {
	l2 := NewArrayList[E]()
	l2.items = append([]E{}, l.items...)
	return l2
}

func (l *ArrayList[E]) Copy() *ArrayList[E] {
	return l.Clone().(*ArrayList[E])
}

// Sort implements sort.Sortable
func (l *ArrayList[E]) Sort(comparator compare.Comparator[E]) {
	sort.Slice(l.items, func(i, j int) bool {
		return comparator.Compare(l.items[i], l.items[j]) < 0
	})
}

// Count implements support.Countable
func (l *ArrayList[E]) Count() int {
	return len(l.items)
}

func (l *ArrayList[E]) IsEmpty() bool {
	return l.Count() == 0
}

func (l *ArrayList[E]) IsNotEmpty() bool {
	return l.Count() > 0
}

func (l *ArrayList[E]) Get(index int) E {
	if index < 0 || index >= len(l.items) {
		panic(ErrIndexOutOfRange)
	} else {
		return l.items[index]
	}
}

func (l *ArrayList[E]) Pop() (value E) {
	if len(l.items) == 0 {
		return
	}
	size := len(l.items)
	el := l.items[size-1]
	l.items = l.items[0 : size-1]
	return el
}

func (l *ArrayList[E]) Shift() (value E) {
	if len(l.items) == 0 {
		return
	}
	el := l.items[0]
	l.items = l.items[1:]
	return el
}

func (l *ArrayList[E]) Contains(matcher func(value E) bool) bool {
	for _, item := range l.items {
		if matcher(item) {
			return true
		}
	}
	return false
}

func (l *ArrayList[E]) IndexOf(matcher func(value E) bool) int {
	for index, item := range l.items {
		if matcher(item) {
			return index
		}
	}
	return -1
}

func (l *ArrayList[E]) LastIndexOf(matcher func(value E) bool) int {
	size := len(l.items)
	for index := range l.items {
		if matcher(l.items[size-index-1]) {
			return size - index - 1
		}
	}
	return -1
}

func (l *ArrayList[E]) SubList(from, to int) List[E] {
	if from < 0 || from >= len(l.items) || to < 0 || to >= len(l.items) {
		panic(ErrIndexOutOfRange)
	}
	start, end := from, to
	if start > end {
		start, end = to, from
	}
	list := NewArrayList[E]()
	list.FromArray(l.items[start:end])
	return list
}

func (l *ArrayList[E]) SubArrayList(from, to int) *ArrayList[E] {
	return l.SubList(from, to).(*ArrayList[E])
}

func (l *ArrayList[E]) Add(value E) {
	l.Push(value)
}

func (l *ArrayList[E]) AddAll(values ...E) {
	l.PushAll(values...)
}

func (l *ArrayList[E]) Set(index int, value E) {
	if index < 0 || index >= len(l.items) {
		panic(ErrIndexOutOfRange)
	}
	l.items[index] = value
}

func (l *ArrayList[E]) Push(value E) {
	l.items = append(l.items, value)
}

func (l *ArrayList[E]) PushAll(values ...E) {
	l.items = append(l.items, values...)
}

func (l *ArrayList[E]) Unshift(value E) {
	l.items = append([]E{value}, l.items...)
}

func (l *ArrayList[E]) UnshiftAll(values ...E) {
	valueSize := len(values)
	for i, j := 0, valueSize-1; i < j; i, j = i+1, j-1 {
		values[i], values[j] = values[j], values[i]
	}
	l.items = append(values, l.items...)
}

func (l *ArrayList[E]) InsertBefore(index int, value E) {
	if index < 0 || index >= len(l.items) {
		panic(ErrIndexOutOfRange)
	}
	items := append([]E{}, l.items[0:index]...)
	items = append(items, value)
	items = append(items, l.items[index:]...)
	l.items = items
}

func (l *ArrayList[E]) InsertAfter(index int, value E) {
	if index < 0 || index >= len(l.items) {
		panic(ErrIndexOutOfRange)
	}
	items := append([]E{}, l.items[0:index+1]...)
	items = append(items, value)
	items = append(items, l.items[index+1:]...)
	l.items = items
}

func (l *ArrayList[E]) RemoveAt(index int) {
	if index < 0 || index >= len(l.items) {
		panic(ErrIndexOutOfRange)
	}
	items := append([]E{}, l.items[0:index]...)
	items = append(items, l.items[index+1:]...)
	l.items = items
}

func (l *ArrayList[E]) Remove(matcher func(value E) bool) {
	newItems := make([]E, 0)
	for _, item := range l.items {
		if !matcher(item) {
			newItems = append(newItems, item)
		}
	}
	l.items = newItems
}

func (l *ArrayList[E]) Clear() {
	l.items = []E{}
}

func (l *ArrayList[E]) Range(callback func(item E) bool) {
	for _, item := range l.items {
		if !callback(item) {
			break
		}
	}
}

func (l *ArrayList[E]) ReverseRange(callback func(item E) bool) {
	size := len(l.items)
	for index := range l.items {
		if !callback(l.items[size-index-1]) {
			break
		}
	}
}

func (l *ArrayList[E]) Map(callback func(value E) E) {
	for index, item := range l.items {
		l.items[index] = callback(item)
	}
}
