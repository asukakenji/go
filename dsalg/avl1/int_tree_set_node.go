package avl

import "fmt"

// IntTreeSetNode is a node of IntTreeSet.
type IntTreeSetNode struct {
	value         int
	leftChild     *IntTreeSetNode
	rightChild    *IntTreeSetNode
	balanceFactor int
}

// TraversePreOrder traverses the subtree rooted at n in pre-order (NLR).
func (n *IntTreeSetNode) TraversePreOrder(consumer func(*IntTreeSetNode)) {
	if n == nil {
		return
	}
	consumer(n)
	n.leftChild.TraversePreOrder(consumer)
	n.rightChild.TraversePreOrder(consumer)
}

// TraversePreOrder traverses the subtree rooted at n in in-order (LNR).
func (n *IntTreeSetNode) TraverseInOrder(consumer func(*IntTreeSetNode)) {
	if n == nil {
		return
	}
	n.leftChild.TraversePreOrder(consumer)
	consumer(n)
	n.rightChild.TraversePreOrder(consumer)
}

// TraversePreOrder traverses the subtree rooted at n in post-order (LRN).
func (n *IntTreeSetNode) TraversePostOrder(consumer func(*IntTreeSetNode)) {
	if n == nil {
		return
	}
	n.leftChild.TraversePreOrder(consumer)
	n.rightChild.TraversePreOrder(consumer)
	consumer(n)
}

func (n *IntTreeSetNode) height() int {
	if n == nil {
		return -1
	}
	return max(n.leftChild.height(), n.rightChild.height()) + 1
}

func (n *IntTreeSetNode) String() string {
	if n == nil {
		return "/"
	}
	return fmt.Sprintf("(%d %s %s)", n.value, n.leftChild.String(), n.rightChild.String())
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
		n.leftChild.Print(indentString, indentLevel+1)
		n.rightChild.Print(indentString, indentLevel+1)
	}
}

// Contains returns whether the subtree rooted at n contains v.
func (n *IntTreeSetNode) Contains(v int) bool {
	if n == nil {
		return false
	}
	if v < n.value {
		return n.leftChild.Contains(v)
	}
	if v > n.value {
		return n.rightChild.Contains(v)
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
		if n.leftChild.Add(v, &n.leftChild) {
			if n.balanceFactor < 0 {
				// Rotate
				if v > n.leftChild.value {
					// LR Case
					n.rotateLeftRight(ptrN)
				} else {
					// LL Case
					n.rotateRight(ptrN)
				}
				return false
			}
			n.balanceFactor--
			return true
		}
	}
	if v > n.value {
		if n.rightChild.Add(v, &n.rightChild) {
			if n.balanceFactor > 0 {
				// Rotate
				if v < n.rightChild.value {
					// RL Case
					n.rotateRightLeft(ptrN)
				} else {
					// RR Case
					n.rotateLeft(ptrN)
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
