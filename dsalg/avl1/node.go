package avl

import (
	"bytes"
	"fmt"
)

// Node represents a node in Tree.
type Node struct {
	// The value stored in this node.
	value interface{}
	// The left child of this node.
	childL *Node
	// The right child of this node.
	childR *Node
	// The balance factor of this node.
	balanceFactor int
}

// Value returns the value stored in n.
func (n *Node) Value() interface{} {
	return n.value
}

// IntValue TODO: Write this comment!
func (n *Node) IntValue() int {
	type inter interface {
		Int() int
	}
	if v, ok := n.value.(int); ok {
		return v
	} else if v, ok := n.value.(inter); ok {
		return v.Int()
	}
	panic("") // TODO: Write this!
}

// Float64Value TODO: Write this comment!
func (n *Node) Float64Value() float64 {
	type floater64 interface {
		Float64() float64
	}
	if v, ok := n.value.(float64); ok {
		return v
	} else if v, ok := n.value.(floater64); ok {
		return v.Float64()
	}
	panic("") // TODO: Write this!
}

// StringValue TODO: Write this comment!
func (n *Node) StringValue() string {
	if v, ok := n.value.(string); ok {
		return v
	} else if v, ok := n.value.(fmt.Stringer); ok {
		return v.String()
	}
	panic("") // TODO: Write this!
}

// BalanceFactor TODO: Write this comment!
func (n *Node) BalanceFactor() int {
	if n == nil {
		return 0
	}
	return n.balanceFactor
}

// Search TODO: Write this comment!
func (n *Node) Search(v int, consumer func(*Node) (int, interface{})) interface{} {
	dir, result := consumer(n)
	if dir < 0 {
		return n.childL.Search(v, consumer)
	}
	if dir > 0 {
		return n.childR.Search(v, consumer)
	}
	return result
}

func (n *Node) search(v int, ptrN **Node, consumer func(*Node, **Node) (int, interface{})) interface{} {
	dir, result := consumer(n, ptrN)
	if dir < 0 {
		return n.childL.search(v, &n.childL, consumer)
	}
	if dir > 0 {
		return n.childR.search(v, &n.childR, consumer)
	}
	return result
}

func (n *Node) height() int {
	if n == nil {
		return -1
	}
	return max(n.childL.height(), n.childR.height()) + 1
}

func (n *Node) String() string {
	if n == nil {
		return "/"
	}
	return fmt.Sprintf("(%d %s %s)", n.value, n.childL.String(), n.childR.String())
}

// Print prints the subtree rooted at n.
func (n *Node) Print(buffer *bytes.Buffer, indentString string, indentLevel int) {
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
func (n *Node) Contains(v interface{}, lessFunc func(interface{}, interface{}) bool) bool {
	if n == nil {
		return false
	}
	if lessFunc(v, n.value) {
		return n.childL.Contains(v, lessFunc)
	}
	if lessFunc(n.value, v) {
		return n.childR.Contains(v, lessFunc)
	}
	return true
}

// Clone returns a copy of the subtree rooted at n.
func (n *Node) clone() *Node {
	if n == nil {
		return nil
	}
	return &Node{
		n.value,
		n.childL.clone(),
		n.childR.clone(),
		n.balanceFactor,
	}
}
