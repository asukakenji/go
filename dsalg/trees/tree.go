package trees

type Constructor func(less func(a interface{}, b interface{}) bool) Tree

type Tree interface {
	// IsEmpty returns whether this tree is empty.
	IsEmpty() bool
	Find(x interface{}) Node
	Insert(x interface{}) Node
	Delete(x interface{}) Node
	Min() interface{}
	Max() interface{}
	Glb(x interface{}) Node
	GlbEq(x interface{}) Node
	Lub(x interface{}) Node
	LubEq(x interface{}) Node
	VisitPreOrder(consumer func(interface{}))
	VisitInOrder(consumer func(interface{}))
	VisitPostOrder(consumer func(interface{}))
}

type Node interface {
	// Value returns the value stored in this node.
	Value() interface{}
	// Height returns the height of this node.
	Height() int
	// Depth returns the depth of this node from the root.
	Depth(root Node) int
	// IsNil returns whether this node is a nil node.
	IsNil() bool
	// IsLeaf returns whether this node is a leaf node.
	IsLeaf() bool
}
