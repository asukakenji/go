package avl

import (
	"bytes"
	"fmt"
)

// IntTreeSetNode is a node of IntTreeSet.
type IntTreeSetNode struct {
	value         int
	childL        *IntTreeSetNode
	childR        *IntTreeSetNode
	balanceFactor int
}

func (n *IntTreeSetNode) Value() interface{} {
	if n == nil {
		return nil
	}
	return n.value
}

func (n *IntTreeSetNode) IntValue() int {
	if n == nil {
		return 0
	}
	return n.value
}

func (n *IntTreeSetNode) BalanceFactor() int {
	if n == nil {
		return 0
	}
	return n.balanceFactor
}

// TraversePreOrder traverses the subtree rooted at n in pre-order (NLR).
func (n *IntTreeSetNode) TraversePreOrder(
	f func(*IntTreeSetNode, interface{}) (bool, interface{}),
	acc0 interface{},
) (bool, interface{}) {
	if n == nil {
		return f(nil, acc0)
	}
	b, acc1 := f(n, acc0)
	if !b {
		return false, acc1
	}
	b, acc2 := n.childL.TraversePreOrder(f, acc1)
	if !b {
		return false, acc2
	}
	return n.childR.TraversePreOrder(f, acc2)
}

// TraverseInOrder traverses the subtree rooted at n in in-order (LNR).
func (n *IntTreeSetNode) TraverseInOrder(
	f func(*IntTreeSetNode, interface{}) (bool, interface{}),
	acc0 interface{},
) (bool, interface{}) {
	if n == nil {
		return f(nil, acc0)
	}
	b, acc1 := n.childL.TraverseInOrder(f, acc0)
	if !b {
		return false, acc1
	}
	b, acc2 := f(n, acc1)
	if !b {
		return false, acc2
	}
	return n.childR.TraverseInOrder(f, acc2)
}

// TraversePostOrder traverses the subtree rooted at n in post-order (LRN).
func (n *IntTreeSetNode) TraversePostOrder(
	f func(*IntTreeSetNode, interface{}) (bool, interface{}),
	acc0 interface{},
) (bool, interface{}) {
	if n == nil {
		return f(nil, acc0)
	}
	b, acc1 := n.childL.TraversePostOrder(f, acc0)
	if !b {
		return false, acc1
	}
	b, acc2 := n.childR.TraversePostOrder(f, acc1)
	if !b {
		return false, acc2
	}
	return f(n, acc0)
}

func (n *IntTreeSetNode) Search(v int, consumer func(*IntTreeSetNode) (int, interface{})) interface{} {
	dir, result := consumer(n)
	if dir < 0 {
		return n.childL.Search(v, consumer)
	}
	if dir > 0 {
		return n.childR.Search(v, consumer)
	}
	return result
}

func (n *IntTreeSetNode) search(v int, ptrN **IntTreeSetNode, consumer func(*IntTreeSetNode, **IntTreeSetNode) (int, interface{})) interface{} {
	dir, result := consumer(n, ptrN)
	if dir < 0 {
		return n.childL.search(v, &n.childL, consumer)
	}
	if dir > 0 {
		return n.childR.search(v, &n.childR, consumer)
	}
	return result
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
func (n *IntTreeSetNode) Print(buffer *bytes.Buffer, indentString string, indentLevel int) {
	for i := 0; i < indentLevel; i++ {
		buffer.WriteString(indentString)
	}
	if n == nil {
		buffer.WriteString("nil\n")
	} else {
		buffer.WriteString(fmt.Sprintf("%d (height: %d, balance factor: %d)\n", n.value, n.height(), n.balanceFactor))
		n.childL.Print(buffer, indentString, indentLevel+1)
		n.childR.Print(buffer, indentString, indentLevel+1)
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
func (n *IntTreeSetNode) Add(v int, ptrN **IntTreeSetNode) (bool, bool) {
	if n == nil {
		*ptrN = &IntTreeSetNode{value: v}
		return true, true
	}
	if v < n.value {
		if isAdded, needsPropagation := n.childL.Add(v, &n.childL); needsPropagation {
			if n.balanceFactor > 0 {
				n.balanceFactor--
				return true, false
			} else if n.balanceFactor == 0 {
				n.balanceFactor--
				return true, true
			} else {
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
				return true, false
			}
		} else {
			return isAdded, false
		}
	}
	if v > n.value {
		if isAdded, needsPropagation := n.childR.Add(v, &n.childR); needsPropagation {
			if n.balanceFactor < 0 {
				n.balanceFactor++
				return true, false
			} else if n.balanceFactor == 0 {
				n.balanceFactor++
				return true, true
			} else {
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
						n.balanceFactor = 0
						p.balanceFactor = 1
					} else if q.balanceFactor > 0 {
						n.balanceFactor = -1
						p.balanceFactor = 0
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
				return true, false
			}
		} else {
			return isAdded, false
		}
	}
	// v already existed, nothing to be done
	return false, false
}

func (n *IntTreeSetNode) Remove(v int, ptrN **IntTreeSetNode) bool {
	// TODO: Write this!
	panic("Not implemented!")
}
