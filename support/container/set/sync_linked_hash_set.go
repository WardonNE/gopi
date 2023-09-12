package set

import (
	"sync"

	"github.com/wardonne/gopi/support/container"
)

type SyncLinkedHashSet[E comparable] struct {
	mu  *sync.RWMutex
	set *LinkedHashSet[E]
}

func NewSyncLinkedHashSet[E comparable]() *SyncLinkedHashSet[E] {
	hashSet := new(SyncLinkedHashSet[E])
	hashSet.mu = new(sync.RWMutex)
	hashSet.set = NewLinkedHashSet[E]()
	return hashSet
}

func (s *SyncLinkedHashSet[E]) MarshalJSON() ([]byte, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.set.MarshalJSON()
}

func (s *SyncLinkedHashSet[E]) UnmarshalJSON(data []byte) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.set.UnmarshalJSON(data)
}

func (s *SyncLinkedHashSet[E]) ToArray() []E {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.set.ToArray()
}

func (s *SyncLinkedHashSet[E]) FromArray(values []E) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.set.FromArray(values)
}

func (s *SyncLinkedHashSet[E]) String() string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.set.String()
}

func (s *SyncLinkedHashSet[E]) Clone() container.Collection[E] {
	s.mu.RLock()
	defer s.mu.RUnlock()
	hashSet := NewSyncLinkedHashSet[E]()
	hashSet.set.list.Range(func(value E) bool {
		hashSet.set.set.items[value] = struct{}{}
		hashSet.set.list.Add(value)
		return true
	})
	return hashSet
}

func (s *SyncLinkedHashSet[E]) Copy() *SyncLinkedHashSet[E] {
	return s.Clone().(*SyncLinkedHashSet[E])
}

func (s *SyncLinkedHashSet[E]) Count() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.set.Count()
}

func (s *SyncLinkedHashSet[E]) IsEmpty() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.set.IsEmpty()
}

func (s *SyncLinkedHashSet[E]) IsNotEmpty() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.set.IsNotEmpty()
}

func (s *SyncLinkedHashSet[E]) Get(index int) E {
	s.mu.RLock()
	defer s.mu.RUnlock()
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

func (s *SyncLinkedHashSet[E]) IndexOf(value E) int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.set.IndexOf(value)
}

func (s *SyncLinkedHashSet[E]) LastIndexOf(value E) int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.set.LastIndexOf(value)
}

func (s *SyncLinkedHashSet[E]) Contains(values ...E) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.set.Contains(values...)
}

func (s *SyncLinkedHashSet[E]) ContainsAny(values ...E) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.set.ContainsAny(values...)
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

func (s *SyncLinkedHashSet[E]) Remove(value E) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.set.Remove(value)
}

func (s *SyncLinkedHashSet[E]) Clear() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.set.Clear()
}

func (s *SyncLinkedHashSet[E]) Range(callback func(item E) bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	s.set.Range(callback)
}

func (s *SyncLinkedHashSet[E]) ReverseRange(callback func(item E) bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	s.set.ReverseRange(callback)
}
