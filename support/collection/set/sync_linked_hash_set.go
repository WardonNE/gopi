package set

import (
	"sync"

	"github.com/wardonne/gopi/support/collection"
)

// SyncLinkedHashSet sync linked hash set
type SyncLinkedHashSet[E comparable] struct {
	mu  sync.RWMutex
	set *LinkedHashSet[E]
}

// NewSyncLinkedHashSet creates a new sync linked hash set
func NewSyncLinkedHashSet[E comparable](values ...E) *SyncLinkedHashSet[E] {
	hashSet := new(SyncLinkedHashSet[E])
	hashSet.set = NewLinkedHashSet[E](values...)
	return hashSet
}

func (s *SyncLinkedHashSet[E]) MarshalJSON() ([]byte, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.set.MarshalJSON()
}

func (s *SyncLinkedHashSet[E]) UnmarshalJSON(data []byte) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.set.UnmarshalJSON(data)
}

func (s *SyncLinkedHashSet[E]) ToArray() []E {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.set.ToArray()
}

func (s *SyncLinkedHashSet[E]) FromArray(values []E) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.set.FromArray(values)
}

func (s *SyncLinkedHashSet[E]) String() string {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.set.String()
}

func (s *SyncLinkedHashSet[E]) Clone() collection.Interface[E] {
	s.mu.Lock()
	defer s.mu.Unlock()
	hashSet := NewSyncLinkedHashSet[E]()
	hashSet.set = s.set.Copy()
	return hashSet
}

func (s *SyncLinkedHashSet[E]) Copy() *SyncLinkedHashSet[E] {
	return s.Clone().(*SyncLinkedHashSet[E])
}

func (s *SyncLinkedHashSet[E]) Count() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.set.Count()
}

func (s *SyncLinkedHashSet[E]) IsEmpty() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.set.IsEmpty()
}

func (s *SyncLinkedHashSet[E]) IsNotEmpty() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.set.IsNotEmpty()
}

func (s *SyncLinkedHashSet[E]) Get(index int) E {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.set.Get(index)
}

func (s *SyncLinkedHashSet[E]) Pop() E {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.set.Pop()
}

func (s *SyncLinkedHashSet[E]) Shift() E {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.set.Shift()
}

func (s *SyncLinkedHashSet[E]) IndexOf(matcher collection.Matcher[E]) int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.set.IndexOf(matcher)
}

func (s *SyncLinkedHashSet[E]) LastIndexOf(matcher collection.Matcher[E]) int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.set.LastIndexOf(matcher)
}

func (s *SyncLinkedHashSet[E]) Contains(matcher collection.Matcher[E]) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.set.Contains(matcher)
}

func (s *SyncLinkedHashSet[E]) Where(matcher collection.Matcher[E]) collection.Interface[E] {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.set.Where(matcher)
}

func (s *SyncLinkedHashSet[E]) Add(value E) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.set.Add(value)
}

func (s *SyncLinkedHashSet[E]) AddAll(values ...E) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.set.AddAll(values...)
}

func (s *SyncLinkedHashSet[E]) Push(value E) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.set.Push(value)
}

func (s *SyncLinkedHashSet[E]) PushAll(values ...E) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.set.PushAll(values...)
}

func (s *SyncLinkedHashSet[E]) Unshift(value E) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.set.Unshift(value)
}

func (s *SyncLinkedHashSet[E]) UnshiftAll(values ...E) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.set.UnshiftAll(values...)
}

func (s *SyncLinkedHashSet[E]) InsertBefore(index int, value E) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.set.InsertBefore(index, value)
}

func (s *SyncLinkedHashSet[E]) InsertAfter(index int, value E) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.set.InsertAfter(index, value)
}

func (s *SyncLinkedHashSet[E]) RemoveAt(index int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.set.RemoveAt(index)
}

func (s *SyncLinkedHashSet[E]) Remove(matcher collection.Matcher[E]) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.set.Remove(matcher)
}

func (s *SyncLinkedHashSet[E]) Clear() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.set.Clear()
}

func (s *SyncLinkedHashSet[E]) Range(callback func(item E) bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.set.Range(callback)
}

func (s *SyncLinkedHashSet[E]) ReverseRange(callback func(item E) bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.set.ReverseRange(callback)
}
