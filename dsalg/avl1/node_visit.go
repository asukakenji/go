package avl

// VisitPreOrder traverses the subtree rooted at n in pre-order (NLR).
func (n *Node) VisitPreOrder(consumer func(interface{})) {
	if n == nil {
		return
	}
	consumer(n.value)
	n.childL.VisitPreOrder(consumer)
	n.childR.VisitPreOrder(consumer)
}

// VisitInOrder traverses the subtree rooted at n in in-order (LNR).
func (n *Node) VisitInOrder(consumer func(interface{})) {
	if n == nil {
		return
	}
	n.childL.VisitInOrder(consumer)
	consumer(n.value)
	n.childR.VisitInOrder(consumer)
}

// VisitPostOrder traverses the subtree rooted at n in post-order (LRN).
func (n *Node) VisitPostOrder(consumer func(interface{})) {
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
func (n *Node) ConditionalVisitPreOrder(predicate func(interface{}) bool) bool {
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
func (n *Node) ConditionalVisitInOrder(predicate func(interface{}) bool) bool {
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
func (n *Node) ConditionalVisitPostOrder(predicate func(interface{}) bool) bool {
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
func (n *Node) VisitNodePreOrder(consumer func(*Node)) {
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
func (n *Node) VisitNodeInOrder(consumer func(*Node)) {
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
func (n *Node) VisitNodePostOrder(consumer func(*Node)) {
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
func (n *Node) ConditionalVisitNodePreOrder(predicate func(*Node) bool) bool {
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
func (n *Node) ConditionalVisitNodeInOrder(predicate func(*Node) bool) bool {
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
func (n *Node) ConditionalVisitNodePostOrder(predicate func(*Node) bool) bool {
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
