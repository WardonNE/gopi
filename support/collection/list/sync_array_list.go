package list

import (
	"sync"

	"github.com/wardonne/gopi/support/collection"
	"github.com/wardonne/gopi/support/compare"
)

type SyncArrayList[E any] struct {
	mu   *sync.RWMutex
	list ArrayList[E]
}

func NewSyncArrayList[E any](values ...E) *SyncArrayList[E] {
	syncArrayList := new(SyncArrayList[E])
	syncArrayList.FromArray(values)
	return syncArrayList
}

func (l *SyncArrayList[E]) MarshalJSON() ([]byte, error) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.list.MarshalJSON()
}

func (l *SyncArrayList[E]) UnmarshalJSON(data []byte) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.list.UnmarshalJSON(data)
}

func (l *SyncArrayList[E]) ToArray() []E {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.list.ToArray()
}

func (l *SyncArrayList[E]) FromArray(values []E) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.list.FromArray(values)
}

func (l *SyncArrayList[V]) String() string {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.list.String()
}

func (l *SyncArrayList[E]) Clone() collection.Collection[E] {
	l.mu.RLock()
	defer l.mu.RUnlock()
	l2 := NewSyncArrayList[E]()
	l2.list.items = append([]E{}, l.list.items...)
	return l2
}

func (l *SyncArrayList[E]) Copy() *SyncArrayList[E] {
	return l.Clone().(*SyncArrayList[E])
}

func (l *SyncArrayList[E]) Sort(comparator compare.Comparator[E]) {
	l.mu.Lock()
	defer l.mu.RUnlock()
	l.list.Sort(comparator)
}

func (l *SyncArrayList[E]) Count() int {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.list.Count()
}

func (l *SyncArrayList[E]) IsEmpty() bool {
	return l.Count() == 0
}

func (l *SyncArrayList[E]) IsNotEmpty() bool {
	return l.Count() > 0
}

func (l *SyncArrayList[E]) Get(index int) E {
	l.mu.RLock()
	defer l.mu.Unlock()
	return l.list.Get(index)
}

func (l *SyncArrayList[E]) Pop() E {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.list.Pop()
}

func (l *SyncArrayList[E]) Shift() E {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.list.Shift()
}

func (l *SyncArrayList[E]) Contains(matcher func(value E) bool) bool {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.list.Contains(matcher)
}

func (l *SyncArrayList[E]) ContainsAny(matcher func(value E) bool) bool {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.list.ContainsAny(matcher)
}

func (l *SyncArrayList[E]) IndexOf(matcher func(value E) bool) int {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.list.IndexOf(matcher)
}

func (l *SyncArrayList[E]) LastIndexOf(matcher func(value E) bool) int {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.list.LastIndexOf(matcher)
}

func (l *SyncArrayList[E]) SubList(from, to int) List[E] {
	l.mu.RLock()
	defer l.mu.Unlock()
	syncList := NewSyncArrayList[E]()
	list := l.list.SubArrayList(from, to)
	syncList.list = *list
	return syncList
}

func (l *SyncArrayList[E]) SubArrayList(from, to int) *SyncArrayList[E] {
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

func (l *SyncArrayList[E]) Remove(matcher func(value E) bool) {
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
	l.mu.RLock()
	defer l.mu.RUnlock()
	l.list.Range(callback)
}

func (l *SyncArrayList[V]) ReverseRange(callback func(item V) bool) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	l.list.ReverseRange(callback)
}

func (l *SyncArrayList[E]) Map(callback func(value E) E) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.list.Map(callback)
}
