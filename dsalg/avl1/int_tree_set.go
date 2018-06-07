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

// Clone returns a copy of the tree.
func (t *IntTreeSet) Clone() *IntTreeSet {
	return &IntTreeSet{
		t.root.clone(),
		t.len,
	}
}
