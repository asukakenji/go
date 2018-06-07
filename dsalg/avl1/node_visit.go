package avl

// VisitPreOrder traverses the subtree rooted at n in pre-order (NLR).
func (n *IntTreeSetNode) VisitPreOrder(consumer func(int)) {
	if n == nil {
		return
	}
	consumer(n.value)
	n.childL.VisitPreOrder(consumer)
	n.childR.VisitPreOrder(consumer)
}

// VisitInOrder traverses the subtree rooted at n in in-order (LNR).
func (n *IntTreeSetNode) VisitInOrder(consumer func(int)) {
	if n == nil {
		return
	}
	n.childL.VisitInOrder(consumer)
	consumer(n.value)
	n.childR.VisitInOrder(consumer)
}

// VisitPostOrder traverses the subtree rooted at n in post-order (LRN).
func (n *IntTreeSetNode) VisitPostOrder(consumer func(int)) {
	if n == nil {
		return
	}
	n.childL.VisitPostOrder(consumer)
	n.childR.VisitPostOrder(consumer)
	consumer(n.value)
}

// ConditionalVisitPreOrder traverses the subtree rooted at n in pre-order (NLR).
// The traversal stops when predicate returns false.
// It returns true if predicate returns true throughout the travesal;
// it returns false otherwise.
func (n *IntTreeSetNode) ConditionalVisitPreOrder(predicate func(int) bool) bool {
	if n == nil {
		return true
	}
	if !predicate(n.value) {
		return false
	}
	if !n.childL.ConditionalVisitPreOrder(predicate) {
		return false
	}
	if !n.childR.ConditionalVisitPreOrder(predicate) {
		return false
	}
	return true
}

// ConditionalVisitInOrder traverses the subtree rooted at n in in-order (LNR).
// The traversal stops when predicate returns false.
// It returns true if predicate returns true throughout the travesal;
// it returns false otherwise.
func (n *IntTreeSetNode) ConditionalVisitInOrder(predicate func(int) bool) bool {
	if n == nil {
		return true
	}
	if !n.childL.ConditionalVisitInOrder(predicate) {
		return false
	}
	if !predicate(n.value) {
		return false
	}
	if !n.childR.ConditionalVisitInOrder(predicate) {
		return false
	}
	return true
}

// ConditionalVisitPostOrder traverses the subtree rooted at n in post-order (LRN).
// The traversal stops when predicate returns false.
// It returns true if predicate returns true throughout the travesal;
// it returns false otherwise.
func (n *IntTreeSetNode) ConditionalVisitPostOrder(predicate func(int) bool) bool {
	if n == nil {
		return true
	}
	if !n.childL.ConditionalVisitPostOrder(predicate) {
		return false
	}
	if !n.childR.ConditionalVisitPostOrder(predicate) {
		return false
	}
	if !predicate(n.value) {
		return false
	}
	return true
}

// VisitNodePreOrder traverses the subtree rooted at n in pre-order (NLR).
// Any nil children of leaf nodes are also traversed.
func (n *IntTreeSetNode) VisitNodePreOrder(consumer func(*IntTreeSetNode)) {
	if n == nil {
		consumer(nil)
		return
	}
	consumer(n)
	n.childL.VisitNodePreOrder(consumer)
	n.childR.VisitNodePreOrder(consumer)
}

// VisitNodeInOrder traverses the subtree rooted at n in in-order (LNR).
// Any nil children of leaf nodes are also traversed.
func (n *IntTreeSetNode) VisitNodeInOrder(consumer func(*IntTreeSetNode)) {
	if n == nil {
		consumer(nil)
		return
	}
	n.childL.VisitNodeInOrder(consumer)
	consumer(n)
	n.childR.VisitNodeInOrder(consumer)
}

// VisitNodePostOrder traverses the subtree rooted at n in post-order (LRN).
// Any nil children of leaf nodes are also traversed.
func (n *IntTreeSetNode) VisitNodePostOrder(consumer func(*IntTreeSetNode)) {
	if n == nil {
		consumer(nil)
		return
	}
	n.childL.VisitNodePostOrder(consumer)
	n.childR.VisitNodePostOrder(consumer)
	consumer(n)
}

// ConditionalVisitNodePreOrder traverses the subtree rooted at n in pre-order (NLR).
// Any nil children of leaf nodes are also traversed.
// The traversal stops when predicate returns false.
// It returns true if predicate returns true throughout the travesal;
// it returns false otherwise.
func (n *IntTreeSetNode) ConditionalVisitNodePreOrder(predicate func(*IntTreeSetNode) bool) bool {
	if n == nil {
		return predicate(nil)
	}
	if !predicate(n) {
		return false
	}
	if !n.childL.ConditionalVisitNodePreOrder(predicate) {
		return false
	}
	if !n.childR.ConditionalVisitNodePreOrder(predicate) {
		return false
	}
	return true
}

// ConditionalVisitNodeInOrder traverses the subtree rooted at n in in-order (LNR).
// Any nil children of leaf nodes are also traversed.
// The traversal stops when predicate returns false.
// It returns true if predicate returns true throughout the travesal;
// it returns false otherwise.
func (n *IntTreeSetNode) ConditionalVisitNodeInOrder(predicate func(*IntTreeSetNode) bool) bool {
	if n == nil {
		return predicate(nil)
	}
	if !n.childL.ConditionalVisitNodeInOrder(predicate) {
		return false
	}
	if !predicate(n) {
		return false
	}
	if !n.childR.ConditionalVisitNodeInOrder(predicate) {
		return false
	}
	return true
}

// ConditionalVisitNodePostOrder traverses the subtree rooted at n in post-order (LRN).
// Any nil children of leaf nodes are also traversed.
// The traversal stops when predicate returns false.
// It returns true if predicate returns true throughout the travesal;
// it returns false otherwise.
func (n *IntTreeSetNode) ConditionalVisitNodePostOrder(predicate func(*IntTreeSetNode) bool) bool {
	if n == nil {
		return predicate(nil)
	}
	if !n.childL.ConditionalVisitNodePostOrder(predicate) {
		return false
	}
	if !n.childR.ConditionalVisitNodePostOrder(predicate) {
		return false
	}
	if !predicate(n) {
		return false
	}
	return true
}
