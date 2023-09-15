package tree

import (
	"encoding/json"
	"fmt"

	"github.com/wardonne/gopi/support/builder"
	"github.com/wardonne/gopi/support/compare"
)

type RBTree[E any] struct {
	root       *RBTreeNode[E]
	comparator compare.Comparator[E]
}

func NewRBTree[E any](comparator compare.Comparator[E], values ...E) *RBTree[E] {
	rbTree := new(RBTree[E])
	rbTree.comparator = comparator
	rbTree.FromArray(values)
	return rbTree
}

func (t *RBTree[E]) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.ToArray())
}

func (t *RBTree[E]) UnmarshalJSON(data []byte) error {
	values := make([]E, 0)
	if err := json.Unmarshal(data, &values); err != nil {
		return err
	}
	t.FromArray(values)
	return nil
}

func (t *RBTree[E]) ToArray() []E {
	nodes := t.root.inOrderRange()
	values := make([]E, 0, len(nodes))
	for _, node := range nodes {
		values = append(values, node.value)
	}
	return values
}

func (t *RBTree[E]) FromArray(values []E) {
	t.AddAll(values...)
}

func (t *RBTree[E]) String() string {
	if bytes, err := t.MarshalJSON(); err != nil {
		builder := builder.NewStringBuilder('{')
		values := t.ToArray()
		for _, value := range values {
			builder.WriteString(fmt.Sprintf("%v", value))
			builder.WriteRune(' ')
		}
		builder.TrimSpace()
		return builder.String()
	} else {
		return string(bytes)
	}
}

func (t *RBTree[E]) Clone() Tree[E] {
	rbTree := NewRBTree[E](t.comparator)
	rbTree.FromArray(t.ToArray())
	return rbTree
}

func (t *RBTree[E]) Copy() *RBTree[E] {
	return t.Clone().(*RBTree[E])
}

func (t *RBTree[E]) Count() int {
	return len(t.root.inOrderRange())
}

func (t *RBTree[E]) IsEmpty() bool {
	return t.Count() == 0
}

func (t *RBTree[E]) IsNotEmpty() bool {
	return t.Count() > 0
}

func (t *RBTree[E]) Contains(values ...E) bool {
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

func (t *RBTree[E]) ContainsAny(values ...E) bool {
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

func (t *RBTree[E]) Add(value E) {
	t.root = t.root.insert(value, t.comparator)
	t.root.color = black
}

func (t *RBTree[E]) AddAll(values ...E) {
	for _, value := range values {
		t.Add(value)
	}
}

func (t *RBTree[E]) Remove(value E) {
	if t.root == nil {
		return
	}
	if t.root.find(value, t.comparator) == nil {
		return
	}
	if t.root.left.isBlack() && t.root.right.isBlack() {
		t.root.color = red
	}
	t.root = t.root.remove(value, t.comparator)
	if t.root.isRed() {
		t.root.color = black
	}
}

func (t *RBTree[E]) Clear() {
	t.root = nil
}

func (t *RBTree[E]) Comparator() compare.Comparator[E] {
	return t.comparator
}

func (t *RBTree[E]) First() (value E, ok bool) {
	if t.root == nil {
		return
	}
	value = t.root.min().value
	return value, true
}

func (t *RBTree[E]) Last() (value E, ok bool) {
	if t.root == nil {
		return
	}
	value = t.root.max().value
	return value, true
}

func (t *RBTree[E]) Range(callback func(value E) bool) {
	for _, node := range t.root.inOrderRange() {
		if !callback(node.value) {
			break
		}
	}
}
