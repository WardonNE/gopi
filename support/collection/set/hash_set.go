package set

import (
	"encoding/json"
	"fmt"

	"github.com/wardonne/gopi/support/collection"
)

type HashSet[E comparable] struct {
	items map[E]struct{}
}

func NewHashSet[E comparable](values ...E) *HashSet[E] {
	hashSet := new(HashSet[E])
	hashSet.items = make(map[E]struct{})
	hashSet.FromArray(values)
	return hashSet
}

func (s *HashSet[E]) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.ToArray())
}

func (s *HashSet[E]) UnmarshalJSON(data []byte) error {
	values := make([]E, 0)
	if err := json.Unmarshal(data, &values); err != nil {
		return err
	}
	s.FromArray(values)
	return nil
}

func (s *HashSet[E]) ToArray() []E {
	values := make([]E, 0, len(s.items))
	for value := range s.items {
		values = append(values, value)
	}
	return values
}

func (s *HashSet[E]) FromArray(values []E) {
	for _, value := range values {
		if _, exists := s.items[value]; !exists {
			s.items[value] = struct{}{}
		}
	}
}

func (s *HashSet[E]) String() string {
	return fmt.Sprintf("%v", s.ToArray())
}

func (s *HashSet[E]) Clone() collection.Interface[E] {
	hashSet := NewHashSet[E]()
	hashSet.items = s.items
	return hashSet
}

func (s *HashSet[E]) Copy() *HashSet[E] {
	return s.Clone().(*HashSet[E])
}

func (s *HashSet[E]) Count() int {
	return len(s.items)
}

func (s *HashSet[E]) IsEmpty() bool {
	return s.Count() == 0
}

func (s *HashSet[E]) IsNotEmpty() bool {
	return s.Count() > 0
}

func (s *HashSet[E]) Contains(matcher func(value E) bool) bool {
	for value := range s.items {
		if matcher(value) {
			return true
		}
	}
	return false
}

func (s *HashSet[E]) Add(value E) {
	_, exists := s.items[value]
	if exists {
		return
	}
	s.items[value] = struct{}{}
}

func (s *HashSet[E]) AddAll(values ...E) {
	for _, value := range values {
		s.Add(value)
	}
}

func (s *HashSet[E]) Remove(matcher func(value E) bool) {
	for value := range s.items {
		if matcher(value) {
			delete(s.items, value)
		}
	}
}

func (s *HashSet[E]) Clear() {
	s.items = make(map[E]struct{})
}

func (s *HashSet[E]) Range(callback func(item E) bool) {
	for value := range s.items {
		if !callback(value) {
			break
		}
	}
}
