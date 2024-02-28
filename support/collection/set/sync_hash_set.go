package set

import (
	"sync"

	"github.com/wardonne/gopi/support/collection"
)

type SyncHashSet[E comparable] struct {
	mu  *sync.RWMutex
	set HashSet[E]
}

func NewSyncHashSet[E comparable](values ...E) *SyncHashSet[E] {
	hashSet := new(SyncHashSet[E])
	hashSet.mu = new(sync.RWMutex)
	hashSet.set = *NewHashSet[E](values...)
	return hashSet
}

func (s *SyncHashSet[E]) MarshalJSON() ([]byte, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.set.MarshalJSON()
}

func (s *SyncHashSet[E]) UnmarshalJSON(data []byte) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.set.UnmarshalJSON(data)
}

func (s *SyncHashSet[E]) ToArray() []E {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.set.ToArray()
}

func (s *SyncHashSet[E]) FromArray(values []E) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.set.FromArray(values)
}

func (s *SyncHashSet[E]) String() string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.set.String()
}

func (s *SyncHashSet[E]) Clone() collection.Interface[E] {
	s.mu.RLock()
	defer s.mu.RUnlock()
	hashSet := NewSyncHashSet[E]()
	set := s.set.Copy()
	hashSet.set = *set
	return hashSet
}

func (s *SyncHashSet[E]) Copy() *SyncHashSet[E] {
	return s.Clone().(*SyncHashSet[E])
}

func (s *SyncHashSet[E]) Count() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.set.Count()
}

func (s *SyncHashSet[E]) IsEmpty() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.set.IsEmpty()
}

func (s *SyncHashSet[E]) IsNotEmpty() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.set.IsNotEmpty()
}

func (s *SyncHashSet[E]) Contains(matcher func(value E) bool) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.set.Contains(matcher)
}

func (s *SyncHashSet[E]) Add(value E) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.set.Add(value)
}

func (s *SyncHashSet[E]) AddAll(values ...E) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.set.AddAll(values...)
}

func (s *SyncHashSet[E]) Remove(matcher func(value E) bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.set.Remove(matcher)
}

func (s *SyncHashSet[E]) Clear() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.set.Clear()
}

func (s *SyncHashSet[E]) Range(callback func(item E) bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	s.set.Range(callback)
}
