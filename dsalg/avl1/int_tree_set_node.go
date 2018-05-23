package avl

import "fmt"

// IntTreeSetNode is a node of IntTreeSet.
type IntTreeSetNode struct {
	value         int
	childL        *IntTreeSetNode
	childR        *IntTreeSetNode
	balanceFactor int
}

// TraversePreOrder traverses the subtree rooted at n in pre-order (NLR).
func (n *IntTreeSetNode) TraversePreOrder(consumer func(*IntTreeSetNode)) {
	if n == nil {
		return
	}
	consumer(n)
	n.childL.TraversePreOrder(consumer)
	n.childR.TraversePreOrder(consumer)
}

// TraversePreOrder traverses the subtree rooted at n in in-order (LNR).
func (n *IntTreeSetNode) TraverseInOrder(consumer func(*IntTreeSetNode)) {
	if n == nil {
		return
	}
	n.childL.TraversePreOrder(consumer)
	consumer(n)
	n.childR.TraversePreOrder(consumer)
}

// TraversePreOrder traverses the subtree rooted at n in post-order (LRN).
func (n *IntTreeSetNode) TraversePostOrder(consumer func(*IntTreeSetNode)) {
	if n == nil {
		return
	}
	n.childL.TraversePreOrder(consumer)
	n.childR.TraversePreOrder(consumer)
	consumer(n)
}

func (n *IntTreeSetNode) height() int {
	if n == nil {
		return -1
	}
	return max(n.childL.height(), n.childR.height()) + 1
}

func (n *IntTreeSetNode) String() string {
	if n == nil {
		return "/"
	}
	return fmt.Sprintf("(%d %s %s)", n.value, n.childL.String(), n.childR.String())
}

// Print prints the subtree rooted at n.
func (n *IntTreeSetNode) Print(indentString string, indentLevel int) {
	for i := 0; i < indentLevel; i++ {
		fmt.Print(indentString)
	}
	if n == nil {
		fmt.Println("nil")
	} else {
		fmt.Printf("%d (height: %d, balance factor: %d)\n", n.value, n.height(), n.balanceFactor)
		n.childL.Print(indentString, indentLevel+1)
		n.childR.Print(indentString, indentLevel+1)
	}
}

// Contains returns whether the subtree rooted at n contains v.
func (n *IntTreeSetNode) Contains(v int) bool {
	if n == nil {
		return false
	}
	if v < n.value {
		return n.childL.Contains(v)
	}
	if v > n.value {
		return n.childR.Contains(v)
	}
	return true
}

// Add adds v to the subtree rooted at n.
func (n *IntTreeSetNode) Add(v int, ptrN **IntTreeSetNode) bool {
	if n == nil {
		*ptrN = &IntTreeSetNode{value: v}
		return true
	}
	if v < n.value {
		if n.childL.Add(v, &n.childL) {
			if n.balanceFactor < 0 {
				// Rotate
				if v > n.childL.value {
					// LR Case
					p := n.childL
					q := p.childR
					p.childR = q.childL
					n.childL = q.childR
					q.childL = p
					q.childR = n
					*ptrN = q
					if q.balanceFactor < 0 {
						n.balanceFactor = 1
						p.balanceFactor = 0
					} else if q.balanceFactor > 0 {
						n.balanceFactor = 0
						p.balanceFactor = -1
					} else {
						n.balanceFactor = 0
						p.balanceFactor = 0
					}
					q.balanceFactor = 0
				} else {
					// LL Case
					p := n.childL
					n.childL = p.childR
					p.childR = n
					*ptrN = p
					n.balanceFactor = 0
					p.balanceFactor = 0
				}
				return false
			}
			n.balanceFactor--
			return true
		}
	}
	if v > n.value {
		if n.childR.Add(v, &n.childR) {
			if n.balanceFactor > 0 {
				// Rotate
				if v < n.childR.value {
					// RL Case
					p := n.childR
					q := p.childL
					p.childL = q.childR
					n.childR = q.childL
					q.childR = p
					q.childL = n
					*ptrN = q
					if q.balanceFactor < 0 {
						n.balanceFactor = 1
						p.balanceFactor = 0
					} else if q.balanceFactor > 0 {
						n.balanceFactor = 0
						p.balanceFactor = -1
					} else {
						n.balanceFactor = 0
						p.balanceFactor = 0
					}
					q.balanceFactor = 0
				} else {
					// RR Case
					p := n.childR
					n.childR = p.childL
					p.childL = n
					*ptrN = p
					n.balanceFactor = 0
					p.balanceFactor = 0
				}
				return false
			}
			n.balanceFactor++
			return true
		}
	}
	// v already existed, nothing to be done
	return false
}

func (n *IntTreeSetNode) Remove(v int, ptrN **IntTreeSetNode) bool {
	// TODO: Write this!
	panic("Not implemented!")
}
