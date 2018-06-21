package dsalg

// BinarySearchTreeNode is a node of a binary search tree.
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
func (node *BinarySearchTreeNode) insert(v interface{}, less func(interface{}, interface{}) bool) (*BinarySearchTreeNode, bool) {
	if node == nil {
		return &BinarySearchTreeNode{value: v}, true
	}
	if less(v, node.value) {
		needsAssignment := (node.leftChild == nil)
		targetNode, isCreated := node.leftChild.insert(v, less)
		if needsAssignment {
			node.leftChild = targetNode
		}
		return targetNode, isCreated
	} else if less(node.value, v) {
		needsAssignment := (node.rightChild == nil)
		targetNode, isCreated := node.rightChild.insert(v, less)
		if needsAssignment {
			node.rightChild = targetNode
		}
		return targetNode, isCreated
	} else {
		return node, false
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
	needsAssignment := (tree.root == nil)
	targetNode, isCreated := tree.root.insert(v, tree.less)
	if needsAssignment {
		tree.root = targetNode
	}
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
