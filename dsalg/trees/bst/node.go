package bst

import "github.com/asukakenji/go/dsalg/trees"

// Node represents a node in Tree.
type Node struct {
	// The value stored in this node.
	value interface{}
	// The left child of this node.
	childL *Node
	// The right child of this node.
	childR *Node
}

// Value returns the value stored in n.
func (n *Node) Value() interface{} {
	return n.value
}

// Height returns the height of n.
func (n *Node) Height() int {
	if n == nil {
		return -1
	}
	return max(n.childL.Height(), n.childR.Height()) + 1
}

func (n *Node) Depth(root trees.Node) int {
	// TODO: Write this!
	return 0
}

func (n *Node) IsNil() bool {
	return n == nil
}

func (n *Node) IsLeaf() bool {
	if n == nil {
		return false
	}
	return n.childL == nil && n.childR == nil
}

// Returns whether a new node is allocated
func (n *Node) insert(v interface{}, less func(interface{}, interface{}) bool) (*Node, bool) {
	if n == nil {
		return &Node{value: v}, true
	}
	if less(v, n.value) {
		needsAssignment := (n.childL == nil)
		targetNode, isCreated := n.childL.insert(v, less)
		if needsAssignment {
			n.childL = targetNode
		}
		return targetNode, isCreated
	} else if less(n.value, v) {
		needsAssignment := (n.childR == nil)
		targetNode, isCreated := n.childR.insert(v, less)
		if needsAssignment {
			n.childR = targetNode
		}
		return targetNode, isCreated
	} else {
		return n, false
	}
}

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
