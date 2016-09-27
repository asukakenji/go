package dsalg

// Reference: sort.IntSlice.Less
func IntLess(a, b interface{}) bool {
	return a.(int) < b.(int)
}

// Reference: sort.Float64Slice.Less
func Float64Less(a, b interface{}) bool {
	return a.(float64) < b.(float64) || isNaN(a.(float64)) && !isNaN(b.(float64))
}

// Reference: sort.StringSlice.Less
func StringLess(a, b interface{}) bool {
	return a.(string) < b.(string)
}

type BinarySearchTreeNode struct {
	// Left child and right child pointers in the binary search tree.
	leftChild, rightChild *BinarySearchTreeNode

	// The value stored with this node.
	value interface{}
}

// LeftChild returns the left child of this node or nil.
func (node *BinarySearchTreeNode) LeftChild() *BinarySearchTreeNode {
	return node.leftChild
}

// RightChild returns the right child of this node or nil.
func (node *BinarySearchTreeNode) RightChild() *BinarySearchTreeNode {
	return node.rightChild
}

// Value returns the value stored with this node.
func (node *BinarySearchTreeNode) Value() interface{} {
	return node.value
}

// Returns whether a new node is allocated
func (node *BinarySearchTreeNode) insert(v interface{}, less func(interface{}, interface{}) bool) (*BinarySearchTreeNode, *BinarySearchTreeNode, bool) {
	if node == nil {
		newNode := &BinarySearchTreeNode{value: v}
		return newNode, newNode, true
	}
	if less(v, node.value) {
		newRoot, targetNode, isCreated := node.leftChild.insert(v, less)
		node.leftChild = newRoot
		return node, targetNode, isCreated
	} else if less(node.value, v) {
		newRoot, targetNode, isCreated := node.rightChild.insert(v, less)
		node.rightChild = newRoot
		return node, targetNode, isCreated
	} else {
		return node, node, false
	}
}

func (node *BinarySearchTreeNode) TraversePreOrder(consumer func(interface{})) {
	if node == nil {
		return
	}
	consumer(node.value)
	node.leftChild.TraversePreOrder(consumer)
	node.rightChild.TraversePreOrder(consumer)
}

func (node *BinarySearchTreeNode) TraverseInOrder(consumer func(interface{})) {
	if node == nil {
		return
	}
	node.leftChild.TraverseInOrder(consumer)
	consumer(node.value)
	node.rightChild.TraverseInOrder(consumer)
}

func (node *BinarySearchTreeNode) TraversePostOrder(consumer func(interface{})) {
	if node == nil {
		return
	}
	node.leftChild.TraversePostOrder(consumer)
	node.rightChild.TraversePostOrder(consumer)
	consumer(node.value)
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

func (tree *BinarySearchTree) Insert(v interface{}) *BinarySearchTreeNode {
	newRoot, targetNode, isCreated := tree.root.insert(v, tree.less)
	tree.root = newRoot
	tree.len++
	if isCreated {
		tree.cap++
	}
	return targetNode
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
