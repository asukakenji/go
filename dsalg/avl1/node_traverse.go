package avl

// TraversePreOrder traverses the subtree rooted at n in pre-order (NLR).
func (n *IntTreeSetNode) TraversePreOrder(consumer func(int)) {
	if n == nil {
		return
	}
	consumer(n.value)
	n.childL.TraversePreOrder(consumer)
	n.childR.TraversePreOrder(consumer)
}

// TraverseInOrder traverses the subtree rooted at n in in-order (LNR).
func (n *IntTreeSetNode) TraverseInOrder(consumer func(int)) {
	if n == nil {
		return
	}
	n.childL.TraverseInOrder(consumer)
	consumer(n.value)
	n.childR.TraverseInOrder(consumer)
}

// TraversePostOrder traverses the subtree rooted at n in post-order (LRN).
func (n *IntTreeSetNode) TraversePostOrder(consumer func(int)) {
	if n == nil {
		return
	}
	n.childL.TraversePostOrder(consumer)
	n.childR.TraversePostOrder(consumer)
	consumer(n.value)
}

// ConditionalTraversePreOrder traverses the subtree rooted at n in pre-order (NLR).
// The traversal stops when predicate returns false.
// It returns true if predicate returns true throughout the travesal;
// it returns false otherwise.
func (n *IntTreeSetNode) ConditionalTraversePreOrder(predicate func(int) bool) bool {
	if n == nil {
		return true
	}
	if !predicate(n.value) {
		return false
	}
	if !n.childL.ConditionalTraversePreOrder(predicate) {
		return false
	}
	if !n.childR.ConditionalTraversePreOrder(predicate) {
		return false
	}
	return true
}

// ConditionalTraverseInOrder traverses the subtree rooted at n in in-order (LNR).
// The traversal stops when predicate returns false.
// It returns true if predicate returns true throughout the travesal;
// it returns false otherwise.
func (n *IntTreeSetNode) ConditionalTraverseInOrder(predicate func(int) bool) bool {
	if n == nil {
		return true
	}
	if !n.childL.ConditionalTraverseInOrder(predicate) {
		return false
	}
	if !predicate(n.value) {
		return false
	}
	if !n.childR.ConditionalTraverseInOrder(predicate) {
		return false
	}
	return true
}

// ConditionalTraversePostOrder traverses the subtree rooted at n in post-order (LRN).
// The traversal stops when predicate returns false.
// It returns true if predicate returns true throughout the travesal;
// it returns false otherwise.
func (n *IntTreeSetNode) ConditionalTraversePostOrder(predicate func(int) bool) bool {
	if n == nil {
		return true
	}
	if !n.childL.ConditionalTraversePostOrder(predicate) {
		return false
	}
	if !n.childR.ConditionalTraversePostOrder(predicate) {
		return false
	}
	if !predicate(n.value) {
		return false
	}
	return true
}

// TraverseNodePreOrder traverses the subtree rooted at n in pre-order (NLR).
// Any nil children of leaf nodes are also traversed.
func (n *IntTreeSetNode) TraverseNodePreOrder(consumer func(*IntTreeSetNode)) {
	if n == nil {
		consumer(nil)
		return
	}
	consumer(n)
	n.childL.TraverseNodePreOrder(consumer)
	n.childR.TraverseNodePreOrder(consumer)
}

// TraverseNodeInOrder traverses the subtree rooted at n in in-order (LNR).
// Any nil children of leaf nodes are also traversed.
func (n *IntTreeSetNode) TraverseNodeInOrder(consumer func(*IntTreeSetNode)) {
	if n == nil {
		consumer(nil)
		return
	}
	n.childL.TraverseNodeInOrder(consumer)
	consumer(n)
	n.childR.TraverseNodeInOrder(consumer)
}

// TraverseNodePostOrder traverses the subtree rooted at n in post-order (LRN).
// Any nil children of leaf nodes are also traversed.
func (n *IntTreeSetNode) TraverseNodePostOrder(consumer func(*IntTreeSetNode)) {
	if n == nil {
		consumer(nil)
		return
	}
	n.childL.TraverseNodePostOrder(consumer)
	n.childR.TraverseNodePostOrder(consumer)
	consumer(n)
}

// ConditionalTraverseNodePreOrder traverses the subtree rooted at n in pre-order (NLR).
// Any nil children of leaf nodes are also traversed.
// The traversal stops when predicate returns false.
// It returns true if predicate returns true throughout the travesal;
// it returns false otherwise.
func (n *IntTreeSetNode) ConditionalTraverseNodePreOrder(predicate func(*IntTreeSetNode) bool) bool {
	if n == nil {
		return predicate(nil)
	}
	if !predicate(n) {
		return false
	}
	if !n.childL.ConditionalTraverseNodePreOrder(predicate) {
		return false
	}
	if !n.childR.ConditionalTraverseNodePreOrder(predicate) {
		return false
	}
	return true
}

// ConditionalTraverseNodeInOrder traverses the subtree rooted at n in in-order (LNR).
// Any nil children of leaf nodes are also traversed.
// The traversal stops when predicate returns false.
// It returns true if predicate returns true throughout the travesal;
// it returns false otherwise.
func (n *IntTreeSetNode) ConditionalTraverseNodeInOrder(predicate func(*IntTreeSetNode) bool) bool {
	if n == nil {
		return predicate(nil)
	}
	if !n.childL.ConditionalTraverseNodeInOrder(predicate) {
		return false
	}
	if !predicate(n) {
		return false
	}
	if !n.childR.ConditionalTraverseNodeInOrder(predicate) {
		return false
	}
	return true
}

// ConditionalTraverseNodePostOrder traverses the subtree rooted at n in post-order (LRN).
// Any nil children of leaf nodes are also traversed.
// The traversal stops when predicate returns false.
// It returns true if predicate returns true throughout the travesal;
// it returns false otherwise.
func (n *IntTreeSetNode) ConditionalTraverseNodePostOrder(predicate func(*IntTreeSetNode) bool) bool {
	if n == nil {
		return predicate(nil)
	}
	if !n.childL.ConditionalTraverseNodePostOrder(predicate) {
		return false
	}
	if !n.childR.ConditionalTraverseNodePostOrder(predicate) {
		return false
	}
	if !predicate(n) {
		return false
	}
	return true
}
