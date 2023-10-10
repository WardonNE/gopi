package set

import (
	"encoding/json"

	"github.com/wardonne/gopi/support/collection"
	"github.com/wardonne/gopi/support/collection/list"
)

type LinkedHashSet[E comparable] struct {
	set  *HashSet[E]
	list *list.LinkedList[E]
}

func NewLinkedHashSet[E comparable](values ...E) *LinkedHashSet[E] {
	hashSet := new(LinkedHashSet[E])
	hashSet.set = NewHashSet[E]()
	hashSet.list = list.NewLinkedList[E]()
	for _, value := range values {
		hashSet.Add(value)
	}
	return hashSet
}

func (s *LinkedHashSet[E]) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.ToArray())
}

func (s *LinkedHashSet[E]) UnmarshalJSON(data []byte) error {
	values := make([]E, 0)
	if err := json.Unmarshal(data, &values); err != nil {
		return err
	}
	for _, value := range values {
		if _, exists := s.set.items[value]; !exists {
			s.set.items[value] = struct{}{}
			s.list.Push(value)
		}
	}
	return nil
}

func (s *LinkedHashSet[E]) ToArray() []E {
	return s.list.ToArray()
}

func (s *LinkedHashSet[E]) FromArray(values []E) {
	for _, value := range values {
		if _, exists := s.set.items[value]; !exists {
			s.set.items[value] = struct{}{}
			s.list.Push(value)
		}
	}
}

func (s *LinkedHashSet[E]) String() string {
	return s.list.String()
}

func (s *LinkedHashSet[E]) Clone() collection.Collection[E] {
	hashSet := NewLinkedHashSet[E]()
	s.list.Range(func(value E) bool {
		hashSet.set.items[value] = struct{}{}
		hashSet.list.Add(value)
		return true
	})
	return hashSet
}

func (s *LinkedHashSet[E]) Copy() *LinkedHashSet[E] {
	return s.Clone().(*LinkedHashSet[E])
}

func (s *LinkedHashSet[E]) Count() int {
	return len(s.set.items)
}

func (s *LinkedHashSet[E]) IsEmpty() bool {
	return s.Count() == 0
}

func (s *LinkedHashSet[E]) IsNotEmpty() bool {
	return s.Count() > 0
}

func (s *LinkedHashSet[E]) Get(index int) E {
	return s.list.Get(index)
}

func (s *LinkedHashSet[E]) Pop() (value E) {
	if s.list.IsEmpty() {
		return
	}
	el := s.list.Pop()
	s.set.Remove(func(value E) bool {
		return value == el
	})
	return el
}

func (s *LinkedHashSet[E]) Shift() E {
	el := s.list.Shift()
	s.set.Remove(func(value E) bool {
		return value == el
	})
	return el
}

func (s *LinkedHashSet[E]) IndexOf(matcher func(value E) bool) int {
	return s.list.IndexOf(matcher)
}

func (s *LinkedHashSet[E]) LastIndexOf(matcher func(value E) bool) int {
	return s.list.LastIndexOf(matcher)
}

func (s *LinkedHashSet[E]) Contains(matcher func(value E) bool) bool {
	return s.set.Contains(matcher)
}

func (s *LinkedHashSet[E]) Add(value E) {
	if _, exists := s.set.items[value]; !exists {
		s.set.items[value] = struct{}{}
		s.list.Push(value)
	}
}

func (s *LinkedHashSet[E]) AddAll(values ...E) {
	for _, value := range values {
		s.Add(value)
	}
}

func (s *LinkedHashSet[E]) Push(value E) {
	s.Add(value)
}

func (s *LinkedHashSet[E]) PushAll(values ...E) {
	s.AddAll(values...)
}

func (s *LinkedHashSet[E]) Unshift(value E) {
	if _, exists := s.set.items[value]; !exists {
		s.set.items[value] = struct{}{}
		s.list.Unshift(value)
	}
}

func (s *LinkedHashSet[E]) UnshiftAll(values ...E) {
	size := len(values)
	for index := range values {
		s.Unshift(values[size-index-1])
	}
}

func (s *LinkedHashSet[E]) InsertBefore(index int, value E) {
	if _, exists := s.set.items[value]; !exists {
		s.set.items[value] = struct{}{}
		s.list.InsertBefore(index, value)
	}
}

func (s *LinkedHashSet[E]) InsertAfter(index int, value E) {
	if _, exists := s.set.items[value]; !exists {
		s.set.items[value] = struct{}{}
		s.list.InsertAfter(index, value)
	}
}

func (s *LinkedHashSet[E]) RemoveAt(index int) {
	el := s.list.Get(index)
	s.set.Remove(func(value E) bool {
		return value == el
	})
	s.list.RemoveAt(index)
}

func (s *LinkedHashSet[E]) Remove(matcher func(value E) bool) {
	s.set.Remove(matcher)
	s.list.Remove(matcher)
}

func (s *LinkedHashSet[E]) Clear() {
	s.set.Clear()
	s.list.Clear()
}

func (s *LinkedHashSet[E]) Range(callback func(item E) bool) {
	s.list.Range(callback)
}

func (s *LinkedHashSet[E]) ReverseRange(callback func(item E) bool) {
	s.list.ReverseRange(callback)
}
