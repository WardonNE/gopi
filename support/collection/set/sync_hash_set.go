package set

import (
	"sync"

	"github.com/wardonne/gopi/support/collection"
)

// SyncHashSet sync hash set
type SyncHashSet[E comparable] struct {
	mu  *sync.RWMutex
	set HashSet[E]
}

// NewSyncHashSet creates a new sync hash set
func NewSyncHashSet[E comparable](values ...E) *SyncHashSet[E] {
	hashSet := new(SyncHashSet[E])
	hashSet.mu = new(sync.RWMutex)
	hashSet.set = *NewHashSet[E](values...)
	return hashSet
}

func (s *SyncHashSet[E]) MarshalJSON() ([]byte, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.set.MarshalJSON()
}

func (s *SyncHashSet[E]) UnmarshalJSON(data []byte) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.set.UnmarshalJSON(data)
}

func (s *SyncHashSet[E]) ToArray() []E {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.set.ToArray()
}

func (s *SyncHashSet[E]) FromArray(values []E) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.set.FromArray(values)
}

func (s *SyncHashSet[E]) String() string {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.set.String()
}

func (s *SyncHashSet[E]) Clone() collection.Interface[E] {
	s.mu.Lock()
	defer s.mu.Unlock()
	hashSet := NewSyncHashSet[E]()
	set := s.set.Copy()
	hashSet.set = *set
	return hashSet
}

func (s *SyncHashSet[E]) Copy() *SyncHashSet[E] {
	return s.Clone().(*SyncHashSet[E])
}

func (s *SyncHashSet[E]) Count() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.set.Count()
}

func (s *SyncHashSet[E]) IsEmpty() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.set.IsEmpty()
}

func (s *SyncHashSet[E]) IsNotEmpty() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.set.IsNotEmpty()
}

func (s *SyncHashSet[E]) Contains(matcher collection.Matcher[E]) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.set.Contains(matcher)
}

func (s *SyncHashSet[E]) Where(matcher collection.Matcher[E]) collection.Interface[E] {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.set.Where(matcher)
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

func (s *SyncHashSet[E]) Remove(matcher collection.Matcher[E]) {
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
	s.mu.Lock()
	defer s.mu.Unlock()
	s.set.Range(callback)
}
