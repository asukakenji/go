package avl

import (
	"bytes"
)

// IntTreeSet implements a tree set based on AVL tree.
type IntTreeSet struct {
	root *IntTreeSetNode
	len  int
}

// Len returns the number of elements of tree t.
// The time complexity is O(1).
func (t *IntTreeSet) Len() int {
	return t.len
}

// Cap returns the capacity of tree t.
// It is always the same as t.Len().
func (t *IntTreeSet) Cap() int {
	return t.len
}

// IsEmpty returns whether tree t is empty.
func (t *IntTreeSet) IsEmpty() bool {
	return t.len == 0
}

// IsFull returns whether tree t is full.
// It always returns false.
func (t *IntTreeSet) IsFull() bool {
	return false
}

// TraversePreOrder TODO: Write this comment!
func (t *IntTreeSet) TraversePreOrder(consumer func(*IntTreeSetNode)) {
	t.root.TraversePreOrder(consumer)
}

// TraverseInOrder TODO: Write this comment!
func (t *IntTreeSet) TraverseInOrder(consumer func(*IntTreeSetNode)) {
	t.root.TraverseInOrder(consumer)
}

// TraversePostOrder TODO: Write this comment!
func (t *IntTreeSet) TraversePostOrder(consumer func(*IntTreeSetNode)) {
	t.root.TraversePostOrder(consumer)
}

// Fold TODO: Write this comment!
func Fold(
	traverseFunc func(func(*IntTreeSetNode)),
	f func(*IntTreeSetNode, interface{}) interface{},
	acc0 interface{},
) interface{} {
	acc := acc0
	traverseFunc(func(n *IntTreeSetNode) {
		acc = f(n, acc)
	})
	return acc
}

// ConditionalTraversePreOrder TODO: Write this comment!
// NOTE: This should move to ConditionalFold
// NOTE: Traverse should have no bool and no interface{}
// NOTE: ConditionalTraverse should have bool but no interface{}
// NOTE: Fold should have no bool but have interface{}
// NOTE: ConditionalFold should have both bool and interface{}
// NOTE: ConditionalFold should be a HOF that accepts either traverse order
func (t *IntTreeSet) ConditionalTraversePreOrder(predicate func(*IntTreeSetNode) bool) bool {
	return t.root.ConditionalTraversePreOrder(predicate)
}

// ConditionalTraverseInOrder TODO: Write this comment!
func (t *IntTreeSet) ConditionalTraverseInOrder(predicate func(*IntTreeSetNode) bool) bool {
	return t.root.ConditionalTraverseInOrder(predicate)
}

// ConditionalTraversePostOrder TODO: Write this comment!
func (t *IntTreeSet) ConditionalTraversePostOrder(predicate func(*IntTreeSetNode) bool) bool {
	return t.root.ConditionalTraversePostOrder(predicate)
}

// ConditionalFold TODO: Write this comment!
func ConditionalFold(
	conditionalTraverseFunc func(func(*IntTreeSetNode) bool) bool,
	f func(*IntTreeSetNode, interface{}) (bool, interface{}),
	acc0 interface{},
) (bool, interface{}) {
	var t bool
	acc := acc0
	conditionalTraverseFunc(func(n *IntTreeSetNode) bool {
		t, acc = f(n, acc)
		return t
	})
	return t, acc
}

// String TODO: Write this comment!
func (t *IntTreeSet) String() string {
	return t.root.String()
}

// Print prints the set.
// TODO: Delete this in production
func (t *IntTreeSet) Print(indentString string) string {
	var buffer bytes.Buffer
	buffer.WriteRune('\n')
	t.root.Print(&buffer, indentString, 0)
	return buffer.String()
}

// Contains returns whether the set contains v.
func (t *IntTreeSet) Contains(v int) bool {
	return t.root.Contains(v)
}

// Add adds v to the set.
func (t *IntTreeSet) Add(v int) {
	if isAdded, _ := t.root.Add(v, &t.root); isAdded {
		t.len++
	}
}

// Remove removes v from the set.
func (t *IntTreeSet) Remove(v int) {
	if t.root.Remove(v, &t.root) {
		t.len--
	}
}
