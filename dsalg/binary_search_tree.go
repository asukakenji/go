package dsalg

// Reference: sort.IntSlice.Less
func IntLess(a, b interface{}) bool {
	return a.(int) < b.(int)
}

// Reference: A copy of sort.Float64Slice.Less
func Float64Less(a, b interface{}) bool {
	return a.(float64) < b.(float64) || isNaN(a.(float64)) && !isNaN(b.(float64))
}

// Reference: A copy of sort.StringSlice.Less
func StringLess(a, b interface{}) bool {
	return a.(string) < b.(string)
}

type BinarySearchTreeNode struct {
	leftChild, rightChild *BinarySearchTreeNode
	value                 interface{}
	count                 int
}

// Returns whether a new node is allocated
func (node *BinarySearchTreeNode) insert(ptrFromParent **BinarySearchTreeNode, less func(interface{}, interface{}) bool, v interface{}) bool {
	if node == nil {
		*ptrFromParent = &BinarySearchTreeNode{value: v, count: 1}
		return true
	}
	if less(v, node.value) {
		return node.leftChild.insert(&node.leftChild, less, v)
	} else if less(node.value, v) {
		return node.rightChild.insert(&node.rightChild, less, v)
	} else {
		node.count++
		return false
	}
}

func (node *BinarySearchTreeNode) TraversePreOrder(consumer func(interface{})) {
	if node == nil {
		return
	}
	for i := 0; i < node.count; i++ {
		consumer(node.value)
	}
	node.leftChild.TraversePreOrder(consumer)
	node.rightChild.TraversePreOrder(consumer)
}

func (node *BinarySearchTreeNode) TraverseInOrder(consumer func(interface{})) {
	if node == nil {
		return
	}
	node.leftChild.TraverseInOrder(consumer)
	for i := 0; i < node.count; i++ {
		consumer(node.value)
	}
	node.rightChild.TraverseInOrder(consumer)
}

func (node *BinarySearchTreeNode) TraversePostOrder(consumer func(interface{})) {
	if node == nil {
		return
	}
	node.leftChild.TraversePostOrder(consumer)
	node.rightChild.TraversePostOrder(consumer)
	for i := 0; i < node.count; i++ {
		consumer(node.value)
	}
}

type BinarySearchTree struct {
	less func(interface{}, interface{}) bool
	root *BinarySearchTreeNode
	len  int
	cap  int
}

func NewBinarySearchTree(less func(interface{}, interface{}) bool) *BinarySearchTree {
	return &BinarySearchTree{less: less}
}

func NewIntBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{less: IntLess}
}

func NewFloat64BinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{less: Float64Less}
}

func NewStringBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{less: StringLess}
}

func (tree *BinarySearchTree) Insert(v interface{}) {
	if tree.root.insert(&tree.root, tree.less, v) {
		tree.cap++
	}
	tree.len++
}

func (tree *BinarySearchTree) TraversePreOrder(consumer func(interface{})) {
	tree.root.TraversePreOrder(consumer)
}

func (tree *BinarySearchTree) TraverseInOrder(consumer func(interface{})) {
	tree.root.TraverseInOrder(consumer)
}

func (tree *BinarySearchTree) TraversePostOrder(consumer func(interface{})) {
	tree.root.TraversePostOrder(consumer)
}

func (tree *BinarySearchTree) Len() int {
	return tree.len
}

func (tree *BinarySearchTree) Cap() int {
	return tree.cap
}

func (tree *BinarySearchTree) IsEmpty() bool {
	return tree.len == 0
}

func (tree *BinarySearchTree) IsFull() bool {
	return false
}
