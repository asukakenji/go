package trees

type Constructor func(less func(a interface{}, b interface{}) bool) Tree

type Tree interface {
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
	Value() interface{}
	Height() int
	Depth(root Node) int
	IsNil() bool
	IsLeaf() bool
}
