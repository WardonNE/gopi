package list

import (
	"sync"

	"github.com/wardonne/gopi/support/collection"
	"github.com/wardonne/gopi/support/compare"
)

// SyncArrayList sync array list
type SyncArrayList[E any] struct {
	mu   sync.RWMutex
	list ArrayList[E]
}

// NewSyncArrayList creates a new sync array list
func NewSyncArrayList[E any](values ...E) *SyncArrayList[E] {
	syncArrayList := new(SyncArrayList[E])
	syncArrayList.FromArray(values)
	return syncArrayList
}

// MarshalJSON implements serializer.JSONSerializer.MarshalJSON
func (l *SyncArrayList[E]) MarshalJSON() ([]byte, error) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.list.MarshalJSON()
}

// UnmarshalJSON implements serializer.JSONSerializer.UnmarshalJSON
func (l *SyncArrayList[E]) UnmarshalJSON(data []byte) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.list.UnmarshalJSON(data)
}

// ToArray implements serializer.ArraySerializer.ToArray
func (l *SyncArrayList[E]) ToArray() []E {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.list.ToArray()
}

// FromArray implements serializer.ArraySerializer.FromArray
func (l *SyncArrayList[E]) FromArray(values []E) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.list.FromArray(values)
}

// FromArray implements serializer.ArraySerializer.FromArray
func (l *SyncArrayList[V]) String() string {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.list.String()
}

// Clone implements support.Clonable
func (l *SyncArrayList[E]) Clone() collection.Interface[E] {
	l.mu.RLock()
	defer l.mu.RUnlock()
	l2 := NewSyncArrayList[E]()
	l2.list.items = append([]E{}, l.list.items...)
	return l2
}

// Copy copy
func (l *SyncArrayList[E]) Copy() *SyncArrayList[E] {
	return l.Clone().(*SyncArrayList[E])
}

// Sort implements sort.Sortable
func (l *SyncArrayList[E]) Sort(comparator compare.Comparator[E]) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.list.Sort(comparator)
}

// Count implements support.Countable
func (l *SyncArrayList[E]) Count() int {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.list.Count()
}

// IsEmpty returns wheather the list is empty
func (l *SyncArrayList[E]) IsEmpty() bool {
	return l.Count() == 0
}

// IsEmpty returns wheather the list is not empty
func (l *SyncArrayList[E]) IsNotEmpty() bool {
	return l.Count() > 0
}

// Get returns element by index
func (l *SyncArrayList[E]) Get(index int) E {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.list.Get(index)
}

// First returns the first element
func (l *SyncArrayList[E]) First() E {
	return l.Get(0)
}

// Last returns the last element
func (l *SyncArrayList[E]) Last() E {
	return l.Get(l.Count() - 1)
}

// FirstWhere returns the first element which matches the matcher
func (l *SyncArrayList[E]) FirstWhere(matcher collection.Matcher[E]) (E, error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.list.FirstWhere(matcher)
}

// LastWhere returns the last element which matches the matcher
func (l *SyncArrayList[E]) LastWhere(matcher collection.Matcher[E]) (e E, err error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.list.LastWhere(matcher)
}

// Pop removes the last element from the list and returns it
func (l *SyncArrayList[E]) Pop() E {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.list.Pop()
}

// Shift removes the first element from the list and returns it
func (l *SyncArrayList[E]) Shift() E {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.list.Shift()
}

func (l *SyncArrayList[E]) Contains(matcher collection.Matcher[E]) bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.list.Contains(matcher)
}

func (l *SyncArrayList[E]) IndexOf(matcher collection.Matcher[E]) int {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.list.IndexOf(matcher)
}

func (l *SyncArrayList[E]) LastIndexOf(matcher collection.Matcher[E]) int {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.list.LastIndexOf(matcher)
}

func (l *SyncArrayList[E]) SubList(from, to int) Interface[E] {
	l.mu.Lock()
	defer l.mu.Unlock()
	syncList := NewSyncArrayList[E]()
	list := l.list.SubArrayList(from, to)
	syncList.list = *list
	return syncList
}

func (l *SyncArrayList[E]) Where(matcher collection.Matcher[E]) collection.Interface[E] {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.Where(matcher)
}

func (l *SyncArrayList[E]) SubSyncArrayList(from, to int) *SyncArrayList[E] {
	return l.SubList(from, to).(*SyncArrayList[E])
}

func (l *SyncArrayList[E]) Add(value E) {
	l.Push(value)
}

func (l *SyncArrayList[E]) AddAll(values ...E) {
	l.PushAll(values...)
}

func (l *SyncArrayList[E]) Set(index int, value E) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.list.Set(index, value)
}

func (l *SyncArrayList[E]) Push(value E) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.list.Push(value)
}

func (l *SyncArrayList[E]) PushAll(values ...E) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.list.PushAll(values...)
}

func (l *SyncArrayList[E]) Unshift(value E) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.list.Unshift(value)
}

func (l *SyncArrayList[E]) UnshiftAll(values ...E) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.list.UnshiftAll(values...)
}

func (l *SyncArrayList[E]) InsertBefore(index int, value E) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.list.InsertBefore(index, value)
}

func (l *SyncArrayList[E]) InsertAfter(index int, value E) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.list.InsertAfter(index, value)
}

func (l *SyncArrayList[E]) RemoveAt(index int) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.list.RemoveAt(index)
}

func (l *SyncArrayList[E]) Remove(matcher collection.Matcher[E]) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.list.Remove(matcher)
}

func (l *SyncArrayList[E]) Clear() {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.list.Clear()
}

func (l *SyncArrayList[V]) Range(callback func(item V) bool) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.list.Range(callback)
}

func (l *SyncArrayList[V]) ReverseRange(callback func(item V) bool) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.list.ReverseRange(callback)
}

func (l *SyncArrayList[E]) Map(callback func(value E) E) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.list.Map(callback)
}
