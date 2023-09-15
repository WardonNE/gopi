package tree

import (
	"sync"

	"github.com/wardonne/gopi/support/compare"
)

type SyncRBTree[E any] struct {
	tree *RBTree[E]
	mu   *sync.RWMutex
}

func NewSyncRBTree[E any](comparator compare.Comparator[E], values ...E) *SyncRBTree[E] {
	t := new(SyncRBTree[E])
	t.mu = new(sync.RWMutex)
	t.tree = NewRBTree[E](comparator, values...)
	return t
}

func (t *SyncRBTree[E]) MarshalJSON() ([]byte, error) {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.tree.MarshalJSON()
}

func (t *SyncRBTree[E]) UnmarshalJSON(data []byte) error {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.tree.UnmarshalJSON(data)
}

func (t *SyncRBTree[E]) ToArray() []E {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.tree.ToArray()
}

func (t *SyncRBTree[E]) FromArray(values []E) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.tree.FromArray(values)
}

func (t *SyncRBTree[E]) String() string {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.tree.String()
}

func (t *SyncRBTree[E]) Clone() Tree[E] {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return NewSyncRBTree[E](t.tree.comparator, t.tree.ToArray()...)
}

func (t *SyncRBTree[E]) Copy() *SyncRBTree[E] {
	return t.Clone().(*SyncRBTree[E])
}

func (t *SyncRBTree[E]) Count() int {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.tree.Count()
}

func (t *SyncRBTree[E]) IsEmpty() bool {
	return t.Count() == 0
}

func (t *SyncRBTree[E]) IsNotEmpty() bool {
	return t.Count() > 0
}

func (t *SyncRBTree[E]) Contains(values ...E) bool {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.tree.Contains(values...)
}

func (t *SyncRBTree[E]) ContainsAny(values ...E) bool {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.tree.ContainsAny(values...)
}

func (t *SyncRBTree[E]) Add(value E) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.tree.Add(value)
}

func (t *SyncRBTree[E]) AddAll(values ...E) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.tree.AddAll(values...)
}

func (t *SyncRBTree[E]) Remove(value E) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.tree.Remove(value)
}

func (t *SyncRBTree[E]) Clear() {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.tree.Clear()
}

func (t *SyncRBTree[E]) Comparator() compare.Comparator[E] {
	return t.tree.comparator
}

func (t *SyncRBTree[E]) First() (value E, ok bool) {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.tree.First()
}

func (t *SyncRBTree[E]) Last() (value E, ok bool) {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.tree.Last()
}

func (t *SyncRBTree[E]) Range(callback func(value E) bool) {
	t.mu.RLock()
	defer t.mu.RUnlock()
	t.tree.Range(callback)
}
