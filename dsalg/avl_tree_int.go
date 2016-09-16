package dsalg

type intAVLTreeNode struct {
	leftChild, rightChild *intAVLTreeNode
	value                 int
	count                 int
	balanceFactor         int
}

// LL Single Rotation
func (node *intAVLTreeNode) rotateLL(ptrParent **intAVLTreeNode) {
	*ptrParent = node.leftChild
	node.balanceFactor, node.leftChild.balanceFactor = 0, 0
	node.leftChild, node.leftChild.rightChild = node.leftChild.rightChild, node
}

// RR Single Rotation
func (node *intAVLTreeNode) rotateRR(ptrParent **intAVLTreeNode) {
	*ptrParent = node.rightChild
	node.balanceFactor, node.rightChild.balanceFactor = 0, 0
	node.rightChild, node.rightChild.leftChild = node.rightChild.leftChild, node
}

// LR Double Rotation
func (node *intAVLTreeNode) rotateLR(ptrParent **intAVLTreeNode) {
	*ptrParent = node.leftChild.rightChild
	node.balanceFactor, node.leftChild.balanceFactor, node.rightChild.balanceFactor = 0, 0, 0
	node.leftChild, node.leftChild.rightChild, node.leftChild.rightChild.leftChild, node.leftChild.rightChild.rightChild = node.leftChild.rightChild.rightChild, node.leftChild.rightChild.leftChild, node.leftChild, node
}

// RL Double Rotation
func (node *intAVLTreeNode) rotateRL(ptrParent **intAVLTreeNode) {
	*ptrParent = node.rightChild.leftChild
	node.balanceFactor, node.leftChild.balanceFactor, node.rightChild.balanceFactor = 0, 0, 0
	node.rightChild, node.rightChild.leftChild, node.rightChild.leftChild.leftChild, node.rightChild.leftChild.rightChild = node.rightChild.leftChild.leftChild, node.rightChild.leftChild.rightChild, node, node.rightChild
}

// ptrParent: The pointer to the pointer from the parent node pointing to the "this" node
// x: The element to be inserted
// Returns:
// (1) whether a new node is allocated
// (2) whether the balance factor of the caller should be updated
// (3) whether the recursive call to insert is on the leftChild or the rightChild of "this" node
func (node *intAVLTreeNode) insert(ptrParent **intAVLTreeNode, x int) (bool, bool, direction) {
	if node == nil {
		*ptrParent = &intAVLTreeNode{value: x, count: 1}
		return true, true, nil
	}
	if x < node.value {
		isNewNodeCreated, shouldUpdateBalanceFactor, d := node.leftChild.insert(&node.leftChild, x)
		if !isNewNodeCreated {
			return false, false, left
		}
		if !shouldUpdateBalanceFactor {
			return true, false, left
		}
		node.balanceFactor++
		if node.balanceFactor != 2 {
			return true, true, left
		}
		if d == left {
			node.rotateLL(ptrParent)
		} else if d == right {
			node.rotateLR(ptrParent)
		} else {
			panic("")
		}
		return true, false, left
	} else if x > node.value {
		isNewNodeCreated, shouldUpdateBalanceFactor, d := node.rightChild.insert(&node.rightChild, x)
		if !isNewNodeCreated {
			return false, false, right
		}
		if !shouldUpdateBalanceFactor {
			return true, false, right
		}
		node.balanceFactor--
		if node.balanceFactor != -2 {
			return true, true, right
		}
		if d == left {
			node.rotateRL(ptrParent)
		} else if d == right {
			node.rotateRR(ptrParent)
		} else {
			panic("")
		}
		return true, false, right
	} else {
		node.count++
		return false, false, nil
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
	if isNewNodeCreated, _, _ := tree.root.insert(&tree.root, x); isNewNodeCreated {
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
