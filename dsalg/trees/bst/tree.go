package bst

import "github.com/asukakenji/go/dsalg/trees"

// Tree implements an unbalanced binary search tree.
type Tree struct {
	less func(interface{}, interface{}) bool
	root *Node
	len  int
}

func New(less func(interface{}, interface{}) bool) *Tree {
	return &Tree{less: less}
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

func (t *Tree) Find(x interface{}) trees.Node {
	return t.root.Find(x, t.less)
}

func (t *Tree) Insert(x interface{}) trees.Node {
	newRoot, targetNode, isNewNodeInserted := t.root.insert(x, t.less)
	if isNewNodeInserted {
		t.root = newRoot
		t.len++
	}
	return targetNode
}

func (t *Tree) Delete(x interface{}) trees.Node {
	newRoot, targetNode, isExistingNodeDeleted := t.root.delete(x, t.less)
	if isExistingNodeDeleted {
		t.root = newRoot
		t.len--
	}
	return targetNode
}

func (t *Tree) Min() interface{} {
	return t.root.Min()
}

func (t *Tree) Max() interface{} {
	return t.root.Max()
}

func (t *Tree) Glb(x interface{}) trees.Node {
	// TODO: Write this!
	return nil
}

func (t *Tree) GlbEq(x interface{}) trees.Node {
	// TODO: Write this!
	return nil
}

func (t *Tree) Lub(x interface{}) trees.Node {
	// TODO: Write this!
	return nil
}

func (t *Tree) LubEq(x interface{}) trees.Node {
	// TODO: Write this!
	return nil
}

func (t *Tree) VisitPreOrder(consumer func(interface{})) {
	t.root.VisitPreOrder(consumer)
}

func (t *Tree) VisitInOrder(consumer func(interface{})) {
	t.root.VisitInOrder(consumer)
}

func (t *Tree) VisitPostOrder(consumer func(interface{})) {
	t.root.VisitPostOrder(consumer)
}
