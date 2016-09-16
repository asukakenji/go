package dsalg

type intBinarySearchTreeNode struct {
	leftChild, rightChild *intBinarySearchTreeNode
	value                 int
	count                 int
}

// Returns whether a new node is allocated
func (node *intBinarySearchTreeNode) insert(ptrParent **intBinarySearchTreeNode, x int) bool {
	if node == nil {
		*ptrParent = &intBinarySearchTreeNode{value: x, count: 1}
		return true
	}
	if x < node.value {
		return node.leftChild.insert(&node.leftChild, x)
	} else if x > node.value {
		return node.rightChild.insert(&node.rightChild, x)
	} else {
		node.count++
		return false
	}
}

func (node *intBinarySearchTreeNode) traversePreOrder(consumer func(int)) {
	if node == nil {
		return
	}
	for i := 0; i < node.count; i++ {
		consumer(node.value)
	}
	node.leftChild.traversePreOrder(consumer)
	node.rightChild.traversePreOrder(consumer)
}

func (node *intBinarySearchTreeNode) traverseInOrder(consumer func(int)) {
	if node == nil {
		return
	}
	node.leftChild.traverseInOrder(consumer)
	for i := 0; i < node.count; i++ {
		consumer(node.value)
	}
	node.rightChild.traverseInOrder(consumer)
}

func (node *intBinarySearchTreeNode) traversePostOrder(consumer func(int)) {
	if node == nil {
		return
	}
	node.leftChild.traversePostOrder(consumer)
	node.rightChild.traversePostOrder(consumer)
	for i := 0; i < node.count; i++ {
		consumer(node.value)
	}
}

type IntBinarySearchTree struct {
	root     *intBinarySearchTreeNode
	length   int
	capacity int
}

func NewIntBinarySearchTree() *IntBinarySearchTree {
	return &IntBinarySearchTree{}
}

func (tree *IntBinarySearchTree) Insert(x int) {
	if tree.root.insert(&tree.root, x) {
		tree.capacity++
	}
	tree.length++
}

func (tree *IntBinarySearchTree) TraversePreOrder(consumer func(int)) {
	tree.root.traversePreOrder(consumer)
}

func (tree *IntBinarySearchTree) TraverseInOrder(consumer func(int)) {
	tree.root.traverseInOrder(consumer)
}

func (tree *IntBinarySearchTree) TraversePostOrder(consumer func(int)) {
	tree.root.traversePostOrder(consumer)
}

func (tree *IntBinarySearchTree) Len() int {
	return tree.length
}

func (tree *IntBinarySearchTree) Cap() int {
	return tree.capacity
}

func (tree *IntBinarySearchTree) IsEmpty() bool {
	return tree.length == 0
}

func (tree *IntBinarySearchTree) IsFull() bool {
	return false
}
