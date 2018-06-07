package avl

import (
	"bytes"
)

// Tree implements an AVL tree.
type Tree struct {
	lessFunc func(interface{}, interface{}) bool
	root     *Node
	len      int
}

func NewTree(lessFunc func(interface{}, interface{}) bool) *Tree {
	return &Tree{lessFunc: lessFunc}
}

// Len returns the number of elements in t.
// The time complexity is O(1).
func (t *Tree) Len() int {
	return t.len
}

// Cap returns the capacity of t.
// It is always the same as t.Len().
func (t *Tree) Cap() int {
	return t.len
}

// IsEmpty returns whether t is empty.
func (t *Tree) IsEmpty() bool {
	return t.len == 0
}

// IsFull returns whether t is full.
// It always returns false.
func (t *Tree) IsFull() bool {
	return false
}

// String TODO: Write this comment!
func (t *Tree) String() string {
	return t.root.String()
}

// Print prints t.
// TODO: Delete this in production
func (t *Tree) Print(indentString string) string {
	var buffer bytes.Buffer
	buffer.WriteRune('\n')
	t.root.Print(&buffer, indentString, 0)
	return buffer.String()
}

// Contains returns whether t contains v.
func (t *Tree) Contains(v interface{}) bool {
	return t.root.Contains(v, t.lessFunc)
}

// Clone returns a copy of t.
func (t *Tree) Clone() *Tree {
	return &Tree{
		t.lessFunc,
		t.root.clone(),
		t.len,
	}
}
