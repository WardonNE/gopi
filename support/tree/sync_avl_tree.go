package tree

import (
	"sync"

	"github.com/wardonne/gopi/support/compare"
)

type SyncAVLTree[E any] struct {
	tree *AVLTree[E]
	mu   *sync.RWMutex
}

func NewSyncAVLTree[E any](comparator compare.Comparator[E], values ...E) *SyncAVLTree[E] {
	t := new(SyncAVLTree[E])
	t.tree = NewAVLTree[E](comparator, values...)
	t.mu = new(sync.RWMutex)
	return t
}

func (t *SyncAVLTree[E]) MarshalJSON() ([]byte, error) {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.tree.MarshalJSON()
}

func (t *SyncAVLTree[E]) UnmarshalJSON(data []byte) error {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.tree.UnmarshalJSON(data)
}

func (t *SyncAVLTree[E]) ToArray() []E {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.tree.ToArray()
}

func (t *SyncAVLTree[E]) FromArray(values []E) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.tree.FromArray(values)
}

func (t *SyncAVLTree[E]) String() string {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.tree.String()
}

func (t *SyncAVLTree[E]) Clone() Tree[E] {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return NewSyncAVLTree[E](t.tree.comparator, t.tree.ToArray()...)
}

func (t *SyncAVLTree[E]) Copy() *SyncAVLTree[E] {
	return t.Clone().(*SyncAVLTree[E])
}

func (t *SyncAVLTree[E]) Count() int {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.tree.Count()
}

func (t *SyncAVLTree[E]) IsEmpty() bool {
	return t.Count() == 0
}

func (t *SyncAVLTree[E]) IsNotEmpty() bool {
	return t.Count() > 0
}

func (t *SyncAVLTree[E]) Contains(values ...E) bool {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.tree.Contains(values...)
}

func (t *SyncAVLTree[E]) ContainsAny(values ...E) bool {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.tree.ContainsAny(values...)
}

func (t *SyncAVLTree[E]) Add(value E) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.tree.Add(value)
}

func (t *SyncAVLTree[E]) AddAll(values ...E) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.tree.AddAll(values...)
}

func (t *SyncAVLTree[E]) Remove(value E) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.tree.Remove(value)
}

func (t *SyncAVLTree[E]) Clear() {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.tree.Clear()
}

func (t *SyncAVLTree[E]) Comparator() compare.Comparator[E] {
	return t.tree.comparator
}

func (t *SyncAVLTree[E]) First() (value E, ok bool) {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.tree.First()
}

func (t *SyncAVLTree[E]) Last() (value E, ok bool) {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.tree.Last()
}

func (t *SyncAVLTree[E]) Range(callback func(value E) bool) {
	t.mu.RLock()
	defer t.mu.RUnlock()
	t.tree.Range(callback)
}
