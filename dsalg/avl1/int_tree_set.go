package avl

import (
	"bytes"
)

// IntTreeSet implements a tree set based on AVL tree.
type IntTreeSet struct {
	root *IntTreeSetNode
	len  int
}

// Len TODO: Write this comment!
func (t *IntTreeSet) Len() int {
	return t.len
}

// Cap TODO: Write this comment!
func (t *IntTreeSet) Cap() int {
	return t.len
}

// IsEmpty TODO: Write this comment!
func (t *IntTreeSet) IsEmpty() bool {
	return t.len == 0
}

// IsFull TODO: Write this comment!
func (t *IntTreeSet) IsFull() bool {
	return false
}

// TraversePreOrder TODO: Write this comment!
// NOTE: This should move to ConditionalFold
// NOTE: Traverse should have no bool and no interface{}
// NOTE: ConditionalTraverse should have bool but no interface{}
// NOTE: Fold should have no bool but have interface{}
// NOTE: ConditionalFold should have both bool and interface{}
// NOTE: ConditionalFold should be a HOF that accepts either traverse order
func (t *IntTreeSet) TraversePreOrder(
	f func(*IntTreeSetNode, interface{}) (bool, interface{}),
	acc0 interface{},
) (bool, interface{}) {
	return t.root.TraversePreOrder(f, acc0)
}

// TraverseInOrder TODO: Write this comment!
func (t *IntTreeSet) TraverseInOrder(
	f func(*IntTreeSetNode, interface{}) (bool, interface{}),
	acc0 interface{},
) (bool, interface{}) {
	return t.root.TraverseInOrder(f, acc0)
}

// TraversePostOrder TODO: Write this comment!
func (t *IntTreeSet) TraversePostOrder(
	f func(*IntTreeSetNode, interface{}) (bool, interface{}),
	acc0 interface{},
) (bool, interface{}) {
	return t.root.TraversePostOrder(f, acc0)
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
	if t.root.Add(v, &t.root) {
		t.len++
	}
}

// Remove removes v from the set.
func (t *IntTreeSet) Remove(v int) {
	if t.root.Remove(v, &t.root) {
		t.len--
	}
}
