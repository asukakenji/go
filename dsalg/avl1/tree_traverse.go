package avl

// TraversePreOrder traverses the tree in pre-order (NLR).
func (t *IntTreeSet) TraversePreOrder(consumer func(int)) {
	t.root.TraversePreOrder(consumer)
}

// TraverseInOrder traverses the tree in in-order (LNR).
func (t *IntTreeSet) TraverseInOrder(consumer func(int)) {
	t.root.TraverseInOrder(consumer)
}

// TraversePostOrder traverses the tree in post-order (LRN).
func (t *IntTreeSet) TraversePostOrder(consumer func(int)) {
	t.root.TraversePostOrder(consumer)
}

// Fold TODO: Write this comment!
func Fold(
	traverseFunc func(func(int)),
	f func(int, interface{}) interface{},
	acc0 interface{},
) interface{} {
	acc := acc0
	traverseFunc(func(v int) {
		acc = f(v, acc)
	})
	return acc
}

// ConditionalTraversePreOrder traverses the tree in pre-order (NLR).
// The traversal stops when predicate returns false.
// It returns true if predicate returns true throughout the travesal;
// it returns false otherwise.
func (t *IntTreeSet) ConditionalTraversePreOrder(predicate func(int) bool) bool {
	return t.root.ConditionalTraversePreOrder(predicate)
}

// ConditionalTraverseInOrder traverses the tree in in-order (LNR).
// The traversal stops when predicate returns false.
// It returns true if predicate returns true throughout the travesal;
// it returns false otherwise.
func (t *IntTreeSet) ConditionalTraverseInOrder(predicate func(int) bool) bool {
	return t.root.ConditionalTraverseInOrder(predicate)
}

// ConditionalTraversePostOrder traverses the tree in post-order (LRN).
// The traversal stops when predicate returns false.
// It returns true if predicate returns true throughout the travesal;
// it returns false otherwise.
func (t *IntTreeSet) ConditionalTraversePostOrder(predicate func(int) bool) bool {
	return t.root.ConditionalTraversePostOrder(predicate)
}

// ConditionalFold TODO: Write this comment!
func ConditionalFold(
	conditionalTraverseFunc func(func(int) bool) bool,
	f func(int, interface{}) (bool, interface{}),
	acc0 interface{},
) (bool, interface{}) {
	var t bool
	acc := acc0
	conditionalTraverseFunc(func(v int) bool {
		t, acc = f(v, acc)
		return t
	})
	return t, acc
}

// TraverseNodePreOrder traverses the tree in pre-order (NLR).
// Any nil children of leaf nodes are also traversed.
func (t *IntTreeSet) TraverseNodePreOrder(consumer func(*IntTreeSetNode)) {
	t.root.TraverseNodePreOrder(consumer)
}

// TraverseNodeInOrder traverses the tree in in-order (LNR).
// Any nil children of leaf nodes are also traversed.
func (t *IntTreeSet) TraverseNodeInOrder(consumer func(*IntTreeSetNode)) {
	t.root.TraverseNodeInOrder(consumer)
}

// TraverseNodePostOrder traverses the tree in post-order (LRN).
// Any nil children of leaf nodes are also traversed.
func (t *IntTreeSet) TraverseNodePostOrder(consumer func(*IntTreeSetNode)) {
	t.root.TraverseNodePostOrder(consumer)
}

// FoldNode TODO: Write this comment!
func FoldNode(
	traverseNodeFunc func(func(*IntTreeSetNode)),
	f func(*IntTreeSetNode, interface{}) interface{},
	acc0 interface{},
) interface{} {
	acc := acc0
	traverseNodeFunc(func(n *IntTreeSetNode) {
		acc = f(n, acc)
	})
	return acc
}

// ConditionalTraverseNodePreOrder traverses the tree in pre-order (NLR).
// Any nil children of leaf nodes are also traversed.
// The traversal stops when predicate returns false.
// It returns true if predicate returns true throughout the travesal;
// it returns false otherwise.
func (t *IntTreeSet) ConditionalTraverseNodePreOrder(predicate func(*IntTreeSetNode) bool) bool {
	return t.root.ConditionalTraverseNodePreOrder(predicate)
}

// ConditionalTraverseNodeInOrder traverses the tree in in-order (LNR).
// Any nil children of leaf nodes are also traversed.
// The traversal stops when predicate returns false.
// It returns true if predicate returns true throughout the travesal;
// it returns false otherwise.
func (t *IntTreeSet) ConditionalTraverseNodeInOrder(predicate func(*IntTreeSetNode) bool) bool {
	return t.root.ConditionalTraverseNodeInOrder(predicate)
}

// ConditionalTraverseNodePostOrder traverses the tree in post-order (LRN).
// Any nil children of leaf nodes are also traversed.
// The traversal stops when predicate returns false.
// It returns true if predicate returns true throughout the travesal;
// it returns false otherwise.
func (t *IntTreeSet) ConditionalTraverseNodePostOrder(predicate func(*IntTreeSetNode) bool) bool {
	return t.root.ConditionalTraverseNodePostOrder(predicate)
}

// ConditionalFoldNode TODO: Write this comment!
func ConditionalFoldNode(
	conditionalTraverseNodeFunc func(func(*IntTreeSetNode) bool) bool,
	f func(*IntTreeSetNode, interface{}) (bool, interface{}),
	acc0 interface{},
) (bool, interface{}) {
	var t bool
	acc := acc0
	conditionalTraverseNodeFunc(func(n *IntTreeSetNode) bool {
		t, acc = f(n, acc)
		return t
	})
	return t, acc
}
