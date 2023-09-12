package list

import (
	"sync"

	"github.com/wardonne/gopi/support/compare"
	"github.com/wardonne/gopi/support/container"
)

type SyncLinkedList[E comparable] struct {
	mu   *sync.RWMutex
	list LinkedList[E]
}

func NewSyncLinkedList[E comparable](values ...E) *SyncLinkedList[E] {
	syncLinkedList := new(SyncLinkedList[E])
	syncLinkedList.mu = new(sync.RWMutex)
	syncLinkedList.list = *NewLinkedList[E](values...)
	return syncLinkedList
}

func (l *SyncLinkedList[E]) MarshalJSON() ([]byte, error) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.list.MarshalJSON()
}

func (l *SyncLinkedList[E]) UnmarshalJSON(data []byte) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.list.UnmarshalJSON(data)
}

func (l *SyncLinkedList[E]) ToArray() []E {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.list.ToArray()
}

func (l *SyncLinkedList[E]) FromArray(values []E) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.list.FromArray(values)
}

func (l *SyncLinkedList[V]) String() string {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.list.String()
}

func (l *SyncLinkedList[E]) Sort(comparer compare.Comparer[E]) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.list.Sort(comparer)
}

func (l *SyncLinkedList[E]) SortElements(comparer compare.Comparer[*element[E]]) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.list.SortElements(comparer)
}

func (l *SyncLinkedList[V]) Count() int {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.list.Count()
}

func (l *SyncLinkedList[E]) Clone() container.Collection[E] {
	l.mu.Lock()
	defer l.mu.Unlock()
	l2 := NewSyncLinkedList[E]()
	l.list.Range(func(value E) bool {
		l2.Push(value)
		return true
	})
	return l2
}

func (l *SyncLinkedList[E]) Copy() *SyncLinkedList[E] {
	return l.Clone().(*SyncLinkedList[E])
}

func (l *SyncLinkedList[E]) IsEmpty() bool {
	return l.Count() == 0
}

func (l *SyncLinkedList[E]) IsNotEmpty() bool {
	return l.Count() > 0
}

func (l *SyncLinkedList[E]) Get(index int) E {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.list.Get(index)
}

func (l *SyncLinkedList[E]) Pop() E {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.list.Pop()
}

func (l *SyncLinkedList[E]) Shift() E {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.list.Shift()
}

func (l *SyncLinkedList[E]) Contains(values ...E) bool {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.list.Contains(values...)
}

func (l *SyncLinkedList[E]) ContainsAny(values ...E) bool {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.list.ContainsAny(values...)
}

func (l *SyncLinkedList[E]) IndexOf(value E) int {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.list.IndexOf(value)
}

func (l *SyncLinkedList[E]) LastIndexOf(value E) int {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.list.LastIndexOf(value)
}

func (l *SyncLinkedList[E]) SubList(from, to int) List[E] {
	l.mu.RLock()
	defer l.mu.RUnlock()
	syncList := NewSyncLinkedList[E]()
	list := l.list.SubLinkedList(from, to)
	syncList.list = *list
	return syncList
}

func (l *SyncLinkedList[E]) SubLinkedList(from, to int) List[E] {
	return l.SubList(from, to).(*SyncLinkedList[E])
}

func (l *SyncLinkedList[E]) Add(value E) {
	l.Push(value)
}

func (l *SyncLinkedList[E]) AddAll(values ...E) {
	l.PushAll(values...)
}

func (l *SyncLinkedList[E]) Set(index int, value E) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.list.Set(index, value)
}

func (l *SyncLinkedList[V]) Push(value V) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.list.Push(value)
}

func (l *SyncLinkedList[V]) PushAll(values ...V) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.list.PushAll(values...)
}

func (l *SyncLinkedList[E]) Unshift(value E) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.list.Unshift(value)
}

func (l *SyncLinkedList[E]) UnshiftAll(values ...E) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.list.UnshiftAll(values...)
}

func (l *SyncLinkedList[E]) InsertBefore(index int, value E) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.list.InsertBefore(index, value)
}

func (l *SyncLinkedList[E]) InsertAfter(index int, value E) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.list.InsertAfter(index, value)
}

func (l *SyncLinkedList[V]) RemoveAt(index int) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.list.RemoveAt(index)
}

func (l *SyncLinkedList[E]) Remove(value E) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.list.Remove(value)
}

func (l *SyncLinkedList[E]) Clear() {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.list.Clear()
}

func (l *SyncLinkedList[E]) Range(callback func(value E) bool) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.list.Range(callback)
}

func (l *SyncLinkedList[E]) RangeElements(callback func(element *element[E]) bool) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.list.RangeElement(callback)
}

func (l *SyncLinkedList[E]) ReverseRange(ReverseRange func(value E) bool) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.list.ReverseRange(ReverseRange)
}

func (l *SyncLinkedList[E]) Map(callback func(value E) E) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.list.Map(callback)
}
