package list

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/wardonne/gopi/support/collection"
	"github.com/wardonne/gopi/support/compare"
)

// ArrayList array list
type ArrayList[E any] struct {
	items []E
}

// NewArrayList creates a new array list
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
func (l *ArrayList[E]) Clone() collection.Interface[E] {
	l2 := NewArrayList[E]()
	l2.items = append([]E{}, l.items...)
	return l2
}

// Copy copy
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

// IsEmpty returns wheather the list is empty
func (l *ArrayList[E]) IsEmpty() bool {
	return l.Count() == 0
}

// IsEmpty returns wheather the list is not empty
func (l *ArrayList[E]) IsNotEmpty() bool {
	return l.Count() > 0
}

// Get returns element by index
func (l *ArrayList[E]) Get(index int) E {
	if index < 0 || index >= len(l.items) {
		panic(ErrIndexOutOfRange)
	} else {
		return l.items[index]
	}
}

// First returns the first element
func (l *ArrayList[E]) First() E {
	return l.Get(0)
}

// Last returns the last element
func (l *ArrayList[E]) Last() E {
	return l.Get(l.Count() - 1)
}

// FirstWhere returns the first element which matches the matcher
func (l *ArrayList[E]) FirstWhere(matcher collection.Matcher[E]) (e E, err error) {
	for _, item := range l.items {
		if matcher(item) {
			return item, nil
		}
	}
	return e, ErrElementNotFound
}

// LastWhere returns the last element which matches the matcher
func (l *ArrayList[E]) LastWhere(matcher collection.Matcher[E]) (e E, err error) {
	size := len(l.items)
	for index := range l.items {
		if matcher(l.items[size-index-1]) {
			return l.items[size-index-1], nil
		}
	}
	return e, ErrElementNotFound
}

// Pop removes the last element from the list and returns it
func (l *ArrayList[E]) Pop() (value E) {
	if len(l.items) == 0 {
		return
	}
	size := len(l.items)
	el := l.items[size-1]
	l.items = l.items[0 : size-1]
	return el
}

// Shift removes the first element from the list and returns it
func (l *ArrayList[E]) Shift() (value E) {
	if len(l.items) == 0 {
		return
	}
	el := l.items[0]
	l.items = l.items[1:]
	return el
}

// Contains returns wheather the list contains element which matches the matcher
func (l *ArrayList[E]) Contains(matcher collection.Matcher[E]) bool {
	for _, item := range l.items {
		if matcher(item) {
			return true
		}
	}
	return false
}

// IndexOf returns index of the first element which matches the matcher
func (l *ArrayList[E]) IndexOf(matcher collection.Matcher[E]) int {
	for index, item := range l.items {
		if matcher(item) {
			return index
		}
	}
	return -1
}

// IndexOf returns index of the last element which matches the matcher
func (l *ArrayList[E]) LastIndexOf(matcher collection.Matcher[E]) int {
	size := len(l.items)
	for index := range l.items {
		if matcher(l.items[size-index-1]) {
			return size - index - 1
		}
	}
	return -1
}

// SubList creates a sub list
func (l *ArrayList[E]) SubList(from, to int) Interface[E] {
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

// Where returns all elements matches the
func (l *ArrayList[E]) Where(matcher collection.Matcher[E]) collection.Interface[E] {
	l2 := []E{}
	for _, item := range l.items {
		if matcher(item) {
			l2 = append(l2, item)
		}
	}
	return NewArrayList(l2...)
}

// SubArrayList creates a sub array list
func (l *ArrayList[E]) SubArrayList(from, to int) *ArrayList[E] {
	return l.SubList(from, to).(*ArrayList[E])
}

// Add alias of [ArrayList.Push]
func (l *ArrayList[E]) Add(value E) {
	l.Push(value)
}

// AddAll alias of [ArrayList.PushAll]
func (l *ArrayList[E]) AddAll(values ...E) {
	l.PushAll(values...)
}

// Set replace element with given value by specific index
func (l *ArrayList[E]) Set(index int, value E) {
	if index < 0 || index >= len(l.items) {
		panic(ErrIndexOutOfRange)
	}
	l.items[index] = value
}

// Push pushes a new element
func (l *ArrayList[E]) Push(value E) {
	l.items = append(l.items, value)
}

// PushAll pushes new elements
func (l *ArrayList[E]) PushAll(values ...E) {
	l.items = append(l.items, values...)
}

// Unshift unshifts a new element
func (l *ArrayList[E]) Unshift(value E) {
	l.items = append([]E{value}, l.items...)
}

// UnshiftAll unshifts new elements
func (l *ArrayList[E]) UnshiftAll(values ...E) {
	valueSize := len(values)
	for i, j := 0, valueSize-1; i < j; i, j = i+1, j-1 {
		values[i], values[j] = values[j], values[i]
	}
	l.items = append(values, l.items...)
}

// InsertBefore insert element before specific index
func (l *ArrayList[E]) InsertBefore(index int, value E) {
	if index < 0 || index >= len(l.items) {
		panic(ErrIndexOutOfRange)
	}
	items := append([]E{}, l.items[0:index]...)
	items = append(items, value)
	items = append(items, l.items[index:]...)
	l.items = items
}

// InsertAfter insert element after specific index
func (l *ArrayList[E]) InsertAfter(index int, value E) {
	if index < 0 || index >= len(l.items) {
		panic(ErrIndexOutOfRange)
	}
	items := append([]E{}, l.items[0:index+1]...)
	items = append(items, value)
	items = append(items, l.items[index+1:]...)
	l.items = items
}

// RemoveAt removes an element by specific index
func (l *ArrayList[E]) RemoveAt(index int) {
	if index < 0 || index >= len(l.items) {
		panic(ErrIndexOutOfRange)
	}
	items := append([]E{}, l.items[0:index]...)
	items = append(items, l.items[index+1:]...)
	l.items = items
}

// Remove removes all elements which matches the matcher
func (l *ArrayList[E]) Remove(matcher collection.Matcher[E]) {
	newItems := make([]E, 0)
	for _, item := range l.items {
		if !matcher(item) {
			newItems = append(newItems, item)
		}
	}
	l.items = newItems
}

// Clear clears the list
func (l *ArrayList[E]) Clear() {
	l.items = []E{}
}

// Range ranges the list
func (l *ArrayList[E]) Range(callback func(item E) bool) {
	for _, item := range l.items {
		if !callback(item) {
			break
		}
	}
}

// ReverseRange ranges the list from tail
func (l *ArrayList[E]) ReverseRange(callback func(item E) bool) {
	size := len(l.items)
	for index := range l.items {
		if !callback(l.items[size-index-1]) {
			break
		}
	}
}

// Map processes all elements by callback
func (l *ArrayList[E]) Map(callback func(value E) E) {
	for index, item := range l.items {
		l.items[index] = callback(item)
	}
}
