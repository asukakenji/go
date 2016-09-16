package dsalg

type intAVLTreeNode struct {
	leftChild, rightChild   *intAVLTreeNode
	value                   int
	count                   int
	leftHeight, rightHeight int
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func (node *intAVLTreeNode) height() int {
	if node == nil {
		return -1
	}
	return max(node.leftHeight, node.rightHeight)
}

// LL Single Rotation
func (node *intAVLTreeNode) rotateLL(ptrParent **intAVLTreeNode) {
	*ptrParent = node.leftChild
	node.leftChild, node.leftChild.rightChild = node.leftChild.rightChild, node
	node.leftHeight = node.leftChild.height() + 1
	(*ptrParent).rightHeight = (*ptrParent).rightChild.height() + 1
}

// RR Single Rotation
func (node *intAVLTreeNode) rotateRR(ptrParent **intAVLTreeNode) {
	*ptrParent = node.rightChild
	node.rightChild, node.rightChild.leftChild = node.rightChild.leftChild, node
	node.rightHeight = node.rightChild.height() + 1
	(*ptrParent).leftHeight = (*ptrParent).leftChild.height() + 1
}

// LR Double Rotation
func (node *intAVLTreeNode) rotateLR(ptrParent **intAVLTreeNode) {
	*ptrParent = node.leftChild.rightChild
	node.leftChild, node.leftChild.rightChild, node.leftChild.rightChild.leftChild, node.leftChild.rightChild.rightChild = node.leftChild.rightChild.rightChild, node.leftChild.rightChild.leftChild, node.leftChild, node
	node.leftHeight = node.leftChild.height() + 1
	(*ptrParent).leftChild.rightHeight = (*ptrParent).leftChild.rightChild.height() + 1
	(*ptrParent).leftHeight = (*ptrParent).leftChild.height() + 1
	(*ptrParent).rightHeight = (*ptrParent).rightChild.height() + 1
}

// RL Double Rotation
func (node *intAVLTreeNode) rotateRL(ptrParent **intAVLTreeNode) {
	*ptrParent = node.rightChild.leftChild
	node.rightChild, node.rightChild.leftChild, node.rightChild.leftChild.leftChild, node.rightChild.leftChild.rightChild = node.rightChild.leftChild.leftChild, node.rightChild.leftChild.rightChild, node, node.rightChild
	node.rightHeight = node.rightChild.height() + 1
	(*ptrParent).rightChild.leftHeight = (*ptrParent).rightChild.leftChild.height() + 1
	(*ptrParent).leftHeight = (*ptrParent).leftChild.height() + 1
	(*ptrParent).rightHeight = (*ptrParent).rightChild.height() + 1
}

// ptrParent: The pointer to the pointer from the parent node pointing to the "this" node
// x: The element to be inserted
// Returns:
// (1) whether a new node is allocated
// (2) whether the balance factor of the caller should be updated
// (3) whether the recursive call to insert is on the leftChild or the rightChild of "this" node
func (node *intAVLTreeNode) insert(ptrParent **intAVLTreeNode, x int) (bool, direction) {
	if node == nil {
		*ptrParent = &intAVLTreeNode{value: x, count: 1}
		return true, nil
	}
	if x < node.value {
		isNewNodeCreated, d := node.leftChild.insert(&node.leftChild, x)
		if !isNewNodeCreated {
			return false, left
		}
		node.leftHeight = node.leftChild.height() + 1
		if node.leftHeight-node.rightHeight <= 1 {
			return true, left
		}
		if d == left {
			node.rotateLL(ptrParent)
		} else if d == right {
			node.rotateLR(ptrParent)
		} else {
			panic("")
		}
		return true, left
	} else if x > node.value {
		isNewNodeCreated, d := node.rightChild.insert(&node.rightChild, x)
		if !isNewNodeCreated {
			return false, right
		}
		node.rightHeight = node.rightChild.height() + 1
		if node.rightHeight-node.leftHeight <= 1 {
			return true, right
		}
		if d == left {
			node.rotateRL(ptrParent)
		} else if d == right {
			node.rotateRR(ptrParent)
		} else {
			panic("")
		}
		return true, right
	} else {
		node.count++
		return false, nil
	}
}

func (node *intAVLTreeNode) traversePreOrder(consumer func(int)) {
	if node == nil {
		return
	}
	for i := 0; i < node.count; i++ {
		consumer(node.value)
	}
	node.leftChild.traversePreOrder(consumer)
	node.rightChild.traversePreOrder(consumer)
}

func (node *intAVLTreeNode) traverseInOrder(consumer func(int)) {
	if node == nil {
		return
	}
	node.leftChild.traverseInOrder(consumer)
	for i := 0; i < node.count; i++ {
		consumer(node.value)
	}
	node.rightChild.traverseInOrder(consumer)
}

func (node *intAVLTreeNode) traversePostOrder(consumer func(int)) {
	if node == nil {
		return
	}
	node.leftChild.traversePostOrder(consumer)
	node.rightChild.traversePostOrder(consumer)
	for i := 0; i < node.count; i++ {
		consumer(node.value)
	}
}

type IntAVLTree struct {
	root     *intAVLTreeNode
	length   int
	capacity int
}

func NewIntAVLTree() *IntAVLTree {
	return &IntAVLTree{}
}

func (tree *IntAVLTree) Insert(x int) {
	if isNewNodeCreated, _ := tree.root.insert(&tree.root, x); isNewNodeCreated {
		tree.capacity++
	}
	tree.length++
}

func (tree *IntAVLTree) TraversePreOrder(consumer func(int)) {
	tree.root.traversePreOrder(consumer)
}

func (tree *IntAVLTree) TraverseInOrder(consumer func(int)) {
	tree.root.traverseInOrder(consumer)
}

func (tree *IntAVLTree) TraversePostOrder(consumer func(int)) {
	tree.root.traversePostOrder(consumer)
}

func (tree *IntAVLTree) Len() int {
	return tree.length
}

func (tree *IntAVLTree) Cap() int {
	return tree.capacity
}

func (tree *IntAVLTree) IsEmpty() bool {
	return tree.length == 0
}

func (tree *IntAVLTree) IsFull() bool {
	return false
}
