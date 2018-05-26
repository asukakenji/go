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

// Search TODO: Write this comment!
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

func (n *IntTreeSetNode) Remove(v int, ptrN **IntTreeSetNode) bool {
	// TODO: Write this!
	panic("Not implemented!")
}
