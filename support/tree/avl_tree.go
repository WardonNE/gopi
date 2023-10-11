package tree

import (
	"encoding/json"
	"fmt"

	"github.com/wardonne/gopi/support/compare"
)

type AVLTree[E any] struct {
	root       *AVLTreeNode[E]
	comparator compare.Comparator[E]
}

func NewAVLTree[E any](comparator compare.Comparator[E], values ...E) *AVLTree[E] {
	avlTree := new(AVLTree[E])
	avlTree.comparator = comparator
	avlTree.FromArray(values)
	return avlTree
}

func (t *AVLTree[E]) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.ToArray())
}

func (t *AVLTree[E]) UnmarshalJSON(data []byte) error {
	values := make([]E, 0)
	if err := json.Unmarshal(data, &values); err != nil {
		return err
	}
	t.FromArray(values)
	return nil
}

func (t *AVLTree[E]) ToArray() []E {
	nodes := t.root.inOrderRange()
	values := make([]E, 0, len(nodes))
	fmt.Println(len(nodes))
	for _, node := range nodes {
		values = append(values, node.value)
	}
	return values
}

func (t *AVLTree[E]) FromArray(values []E) {
	t.Clear()
	t.AddAll(values...)
}

func (t *AVLTree[E]) String() string {
	return fmt.Sprintf("%v", t.ToArray())
}

func (t *AVLTree[E]) Clone() Tree[E] {
	avltree := NewAVLTree[E](t.comparator)
	avltree.FromArray(t.ToArray())
	return avltree
}

func (t *AVLTree[E]) Copy() *AVLTree[E] {
	return t.Clone().(*AVLTree[E])
}

func (t *AVLTree[E]) Count() int {
	return len(t.root.inOrderRange())
}

func (t *AVLTree[E]) IsEmpty() bool {
	return t.Count() == 0
}

func (t *AVLTree[E]) IsNotEmpty() bool {
	return t.Count() > 0
}

func (t *AVLTree[E]) Contains(values ...E) bool {
	if t.root == nil {
		return false
	}
	for _, value := range values {
		if t.root.find(value, t.comparator) == nil {
			return false
		}
	}
	return true
}

func (t *AVLTree[E]) ContainsAny(values ...E) bool {
	if t.root == nil {
		return false
	}
	for _, value := range values {
		if t.root.find(value, t.comparator) != nil {
			return true
		}
	}
	return false
}

func (t *AVLTree[E]) Add(value E) {
	t.root = t.root.insert(value, t.comparator)
}

func (t *AVLTree[E]) AddAll(values ...E) {
	for _, value := range values {
		t.Add(value)
	}
}

func (t *AVLTree[E]) Remove(value E) {
	if t.root == nil {
		return
	}
	t.root = t.root.remove(value, t.comparator)
}

func (t *AVLTree[E]) Clear() {
	t.root = nil
}

func (t *AVLTree[E]) Comparator() compare.Comparator[E] {
	return t.comparator
}

func (t *AVLTree[E]) First() (value E, ok bool) {
	if t.root == nil {
		return
	}
	return t.root.min().value, true
}

func (t *AVLTree[E]) Last() (value E, ok bool) {
	if t.root == nil {
		return
	}
	return t.root.max().value, true
}

func (t *AVLTree[E]) Range(callback func(value E) bool) {
	for _, node := range t.root.inOrderRange() {
		if !callback(node.value) {
			break
		}
	}
}
