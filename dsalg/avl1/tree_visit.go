package avl

// VisitPreOrder traverses t in pre-order (NLR).
func (t *Tree) VisitPreOrder(consumer func(interface{})) {
	t.root.VisitPreOrder(consumer)
}

// VisitInOrder traverses t in in-order (LNR).
func (t *Tree) VisitInOrder(consumer func(interface{})) {
	t.root.VisitInOrder(consumer)
}

// VisitPostOrder traverses t in post-order (LRN).
func (t *Tree) VisitPostOrder(consumer func(interface{})) {
	t.root.VisitPostOrder(consumer)
}

// Fold TODO: Write this comment!
func Fold(
	visitFunc func(func(interface{})),
	f func(interface{}, interface{}) interface{},
	acc0 interface{},
) interface{} {
	acc := acc0
	visitFunc(func(v interface{}) {
		acc = f(v, acc)
	})
	return acc
}

// ConditionalVisitPreOrder traverses t in pre-order (NLR).
// The traversal stops when predicate returns false.
// It returns true if predicate returns true throughout the travesal;
// it returns false otherwise.
func (t *Tree) ConditionalVisitPreOrder(predicate func(interface{}) bool) bool {
	return t.root.ConditionalVisitPreOrder(predicate)
}

// ConditionalVisitInOrder traverses t in in-order (LNR).
// The traversal stops when predicate returns false.
// It returns true if predicate returns true throughout the travesal;
// it returns false otherwise.
func (t *Tree) ConditionalVisitInOrder(predicate func(interface{}) bool) bool {
	return t.root.ConditionalVisitInOrder(predicate)
}

// ConditionalVisitPostOrder traverses t in post-order (LRN).
// The traversal stops when predicate returns false.
// It returns true if predicate returns true throughout the travesal;
// it returns false otherwise.
func (t *Tree) ConditionalVisitPostOrder(predicate func(interface{}) bool) bool {
	return t.root.ConditionalVisitPostOrder(predicate)
}

// ConditionalFold TODO: Write this comment!
func ConditionalFold(
	conditionalVisitFunc func(func(interface{}) bool) bool,
	f func(interface{}, interface{}) (bool, interface{}),
	acc0 interface{},
) (bool, interface{}) {
	var t bool
	acc := acc0
	conditionalVisitFunc(func(v interface{}) bool {
		t, acc = f(v, acc)
		return t
	})
	return t, acc
}

// VisitNodePreOrder traverses t in pre-order (NLR).
// Any nil children of leaf nodes are also traversed.
func (t *Tree) VisitNodePreOrder(consumer func(*Node)) {
	t.root.VisitNodePreOrder(consumer)
}

// VisitNodeInOrder traverses t in in-order (LNR).
// Any nil children of leaf nodes are also traversed.
func (t *Tree) VisitNodeInOrder(consumer func(*Node)) {
	t.root.VisitNodeInOrder(consumer)
}

// VisitNodePostOrder traverses t in post-order (LRN).
// Any nil children of leaf nodes are also traversed.
func (t *Tree) VisitNodePostOrder(consumer func(*Node)) {
	t.root.VisitNodePostOrder(consumer)
}

// FoldNode TODO: Write this comment!
func FoldNode(
	visitNodeFunc func(func(*Node)),
	f func(*Node, interface{}) interface{},
	acc0 interface{},
) interface{} {
	acc := acc0
	visitNodeFunc(func(n *Node) {
		acc = f(n, acc)
	})
	return acc
}

// ConditionalVisitNodePreOrder traverses t in pre-order (NLR).
// Any nil children of leaf nodes are also traversed.
// The traversal stops when predicate returns false.
// It returns true if predicate returns true throughout the travesal;
// it returns false otherwise.
func (t *Tree) ConditionalVisitNodePreOrder(predicate func(*Node) bool) bool {
	return t.root.ConditionalVisitNodePreOrder(predicate)
}

// ConditionalVisitNodeInOrder traverses t in in-order (LNR).
// Any nil children of leaf nodes are also traversed.
// The traversal stops when predicate returns false.
// It returns true if predicate returns true throughout the travesal;
// it returns false otherwise.
func (t *Tree) ConditionalVisitNodeInOrder(predicate func(*Node) bool) bool {
	return t.root.ConditionalVisitNodeInOrder(predicate)
}

// ConditionalVisitNodePostOrder traverses t in post-order (LRN).
// Any nil children of leaf nodes are also traversed.
// The traversal stops when predicate returns false.
// It returns true if predicate returns true throughout the travesal;
// it returns false otherwise.
func (t *Tree) ConditionalVisitNodePostOrder(predicate func(*Node) bool) bool {
	return t.root.ConditionalVisitNodePostOrder(predicate)
}

// ConditionalFoldNode TODO: Write this comment!
func ConditionalFoldNode(
	conditionalVisitNodeFunc func(func(*Node) bool) bool,
	f func(*Node, interface{}) (bool, interface{}),
	acc0 interface{},
) (bool, interface{}) {
	var t bool
	acc := acc0
	conditionalVisitNodeFunc(func(n *Node) bool {
		t, acc = f(n, acc)
		return t
	})
	return t, acc
}
