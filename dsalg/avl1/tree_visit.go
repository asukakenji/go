package avl

// VisitPreOrder traverses the tree in pre-order (NLR).
func (t *IntTreeSet) VisitPreOrder(consumer func(int)) {
	t.root.VisitPreOrder(consumer)
}

// VisitInOrder traverses the tree in in-order (LNR).
func (t *IntTreeSet) VisitInOrder(consumer func(int)) {
	t.root.VisitInOrder(consumer)
}

// VisitPostOrder traverses the tree in post-order (LRN).
func (t *IntTreeSet) VisitPostOrder(consumer func(int)) {
	t.root.VisitPostOrder(consumer)
}

// Fold TODO: Write this comment!
func Fold(
	visitFunc func(func(int)),
	f func(int, interface{}) interface{},
	acc0 interface{},
) interface{} {
	acc := acc0
	visitFunc(func(v int) {
		acc = f(v, acc)
	})
	return acc
}

// ConditionalVisitPreOrder traverses the tree in pre-order (NLR).
// The traversal stops when predicate returns false.
// It returns true if predicate returns true throughout the travesal;
// it returns false otherwise.
func (t *IntTreeSet) ConditionalVisitPreOrder(predicate func(int) bool) bool {
	return t.root.ConditionalVisitPreOrder(predicate)
}

// ConditionalVisitInOrder traverses the tree in in-order (LNR).
// The traversal stops when predicate returns false.
// It returns true if predicate returns true throughout the travesal;
// it returns false otherwise.
func (t *IntTreeSet) ConditionalVisitInOrder(predicate func(int) bool) bool {
	return t.root.ConditionalVisitInOrder(predicate)
}

// ConditionalVisitPostOrder traverses the tree in post-order (LRN).
// The traversal stops when predicate returns false.
// It returns true if predicate returns true throughout the travesal;
// it returns false otherwise.
func (t *IntTreeSet) ConditionalVisitPostOrder(predicate func(int) bool) bool {
	return t.root.ConditionalVisitPostOrder(predicate)
}

// ConditionalFold TODO: Write this comment!
func ConditionalFold(
	conditionalVisitFunc func(func(int) bool) bool,
	f func(int, interface{}) (bool, interface{}),
	acc0 interface{},
) (bool, interface{}) {
	var t bool
	acc := acc0
	conditionalVisitFunc(func(v int) bool {
		t, acc = f(v, acc)
		return t
	})
	return t, acc
}

// VisitNodePreOrder traverses the tree in pre-order (NLR).
// Any nil children of leaf nodes are also traversed.
func (t *IntTreeSet) VisitNodePreOrder(consumer func(*IntTreeSetNode)) {
	t.root.VisitNodePreOrder(consumer)
}

// VisitNodeInOrder traverses the tree in in-order (LNR).
// Any nil children of leaf nodes are also traversed.
func (t *IntTreeSet) VisitNodeInOrder(consumer func(*IntTreeSetNode)) {
	t.root.VisitNodeInOrder(consumer)
}

// VisitNodePostOrder traverses the tree in post-order (LRN).
// Any nil children of leaf nodes are also traversed.
func (t *IntTreeSet) VisitNodePostOrder(consumer func(*IntTreeSetNode)) {
	t.root.VisitNodePostOrder(consumer)
}

// FoldNode TODO: Write this comment!
func FoldNode(
	visitNodeFunc func(func(*IntTreeSetNode)),
	f func(*IntTreeSetNode, interface{}) interface{},
	acc0 interface{},
) interface{} {
	acc := acc0
	visitNodeFunc(func(n *IntTreeSetNode) {
		acc = f(n, acc)
	})
	return acc
}

// ConditionalVisitNodePreOrder traverses the tree in pre-order (NLR).
// Any nil children of leaf nodes are also traversed.
// The traversal stops when predicate returns false.
// It returns true if predicate returns true throughout the travesal;
// it returns false otherwise.
func (t *IntTreeSet) ConditionalVisitNodePreOrder(predicate func(*IntTreeSetNode) bool) bool {
	return t.root.ConditionalVisitNodePreOrder(predicate)
}

// ConditionalVisitNodeInOrder traverses the tree in in-order (LNR).
// Any nil children of leaf nodes are also traversed.
// The traversal stops when predicate returns false.
// It returns true if predicate returns true throughout the travesal;
// it returns false otherwise.
func (t *IntTreeSet) ConditionalVisitNodeInOrder(predicate func(*IntTreeSetNode) bool) bool {
	return t.root.ConditionalVisitNodeInOrder(predicate)
}

// ConditionalVisitNodePostOrder traverses the tree in post-order (LRN).
// Any nil children of leaf nodes are also traversed.
// The traversal stops when predicate returns false.
// It returns true if predicate returns true throughout the travesal;
// it returns false otherwise.
func (t *IntTreeSet) ConditionalVisitNodePostOrder(predicate func(*IntTreeSetNode) bool) bool {
	return t.root.ConditionalVisitNodePostOrder(predicate)
}

// ConditionalFoldNode TODO: Write this comment!
func ConditionalFoldNode(
	conditionalVisitNodeFunc func(func(*IntTreeSetNode) bool) bool,
	f func(*IntTreeSetNode, interface{}) (bool, interface{}),
	acc0 interface{},
) (bool, interface{}) {
	var t bool
	acc := acc0
	conditionalVisitNodeFunc(func(n *IntTreeSetNode) bool {
		t, acc = f(n, acc)
		return t
	})
	return t, acc
}
