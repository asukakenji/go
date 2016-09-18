package dsalg

type AVLTreeNode struct {
	leftChild, rightChild   *AVLTreeNode
	value                   interface{}
	count                   int
	leftHeight, rightHeight int
}

func (node *AVLTreeNode) Height() int {
	if node == nil {
		return -1
	}
	return max(node.leftHeight, node.rightHeight)
}

// LL Single Rotation
func (node *AVLTreeNode) rotateLL(ptrFromParent **AVLTreeNode) {
	*ptrFromParent = node.leftChild
	node.leftChild, node.leftChild.rightChild = node.leftChild.rightChild, node
	node.leftHeight = node.leftChild.Height() + 1
	(*ptrFromParent).rightHeight = (*ptrFromParent).rightChild.Height() + 1
}

// RR Single Rotation
func (node *AVLTreeNode) rotateRR(ptrFromParent **AVLTreeNode) {
	*ptrFromParent = node.rightChild
	node.rightChild, node.rightChild.leftChild = node.rightChild.leftChild, node
	node.rightHeight = node.rightChild.Height() + 1
	(*ptrFromParent).leftHeight = (*ptrFromParent).leftChild.Height() + 1
}

// LR Double Rotation
func (node *AVLTreeNode) rotateLR(ptrFromParent **AVLTreeNode) {
	*ptrFromParent = node.leftChild.rightChild
	node.leftChild, node.leftChild.rightChild, node.leftChild.rightChild.leftChild, node.leftChild.rightChild.rightChild = node.leftChild.rightChild.rightChild, node.leftChild.rightChild.leftChild, node.leftChild, node
	node.leftHeight = node.leftChild.Height() + 1
	(*ptrFromParent).leftChild.rightHeight = (*ptrFromParent).leftChild.rightChild.Height() + 1
	(*ptrFromParent).leftHeight = (*ptrFromParent).leftChild.Height() + 1
	(*ptrFromParent).rightHeight = (*ptrFromParent).rightChild.Height() + 1
}

// RL Double Rotation
func (node *AVLTreeNode) rotateRL(ptrFromParent **AVLTreeNode) {
	*ptrFromParent = node.rightChild.leftChild
	node.rightChild, node.rightChild.leftChild, node.rightChild.leftChild.leftChild, node.rightChild.leftChild.rightChild = node.rightChild.leftChild.leftChild, node.rightChild.leftChild.rightChild, node, node.rightChild
	node.rightHeight = node.rightChild.Height() + 1
	(*ptrFromParent).rightChild.leftHeight = (*ptrFromParent).rightChild.leftChild.Height() + 1
	(*ptrFromParent).leftHeight = (*ptrFromParent).leftChild.Height() + 1
	(*ptrFromParent).rightHeight = (*ptrFromParent).rightChild.Height() + 1
}

// ptrFromParent: The pointer to the pointer from the parent node pointing to the "this" node
// v: The value to be inserted
// Returns:
// (1) whether a new node is allocated
// (2) whether the recursive call to insert is on the leftChild or the rightChild of "this" node
func (node *AVLTreeNode) insert(ptrFromParent **AVLTreeNode, less func(interface{}, interface{}) bool, v interface{}) (bool, direction) {
	if node == nil {
		*ptrFromParent = &AVLTreeNode{value: v, count: 1}
		return true, nil
	}
	if less(v, node.value) {
		isNewNodeCreated, d := node.leftChild.insert(&node.leftChild, less, v)
		if !isNewNodeCreated {
			return false, left
		}
		node.leftHeight = node.leftChild.Height() + 1
		if node.leftHeight-node.rightHeight <= 1 {
			return true, left
		}
		if d == left {
			node.rotateLL(ptrFromParent)
		} else if d == right {
			node.rotateLR(ptrFromParent)
		} else {
			panic("")
		}
		return true, left
	} else if less(node.value, v) {
		isNewNodeCreated, d := node.rightChild.insert(&node.rightChild, less, v)
		if !isNewNodeCreated {
			return false, right
		}
		node.rightHeight = node.rightChild.Height() + 1
		if node.rightHeight-node.leftHeight <= 1 {
			return true, right
		}
		if d == left {
			node.rotateRL(ptrFromParent)
		} else if d == right {
			node.rotateRR(ptrFromParent)
		} else {
			panic("")
		}
		return true, right
	} else {
		node.count++
		return false, nil
	}
}

func (node *AVLTreeNode) TraversePreOrder(consumer func(interface{})) {
	if node == nil {
		return
	}
	for i := 0; i < node.count; i++ {
		consumer(node.value)
	}
	node.leftChild.TraversePreOrder(consumer)
	node.rightChild.TraversePreOrder(consumer)
}

func (node *AVLTreeNode) TraverseInOrder(consumer func(interface{})) {
	if node == nil {
		return
	}
	node.leftChild.TraverseInOrder(consumer)
	for i := 0; i < node.count; i++ {
		consumer(node.value)
	}
	node.rightChild.TraverseInOrder(consumer)
}

func (node *AVLTreeNode) TraversePostOrder(consumer func(interface{})) {
	if node == nil {
		return
	}
	node.leftChild.TraversePostOrder(consumer)
	node.rightChild.TraversePostOrder(consumer)
	for i := 0; i < node.count; i++ {
		consumer(node.value)
	}
}

func (node *AVLTreeNode) Clone() *AVLTreeNode {
	if node == nil {
		return nil
	}
	return &AVLTreeNode{
		node.leftChild.Clone(),
		node.rightChild.Clone(),
		node.value,
		node.count,
		node.leftHeight,
		node.rightHeight,
	}
}

type AVLTree struct {
	less func(interface{}, interface{}) bool
	root *AVLTreeNode
	len  int
	cap  int
}

func NewAVLTree(less func(interface{}, interface{}) bool) *AVLTree {
	return &AVLTree{less: less}
}

func NewIntAVLTree() *AVLTree {
	return &AVLTree{less: IntLess}
}

func NewFloat64AVLTree() *AVLTree {
	return &AVLTree{less: Float64Less}
}

func NewStringAVLTree() *AVLTree {
	return &AVLTree{less: StringLess}
}

func (tree *AVLTree) Insert(v interface{}) {
	if isNewNodeCreated, _ := tree.root.insert(&tree.root, tree.less, v); isNewNodeCreated {
		tree.cap++
	}
	tree.len++
}

func (tree *AVLTree) TraversePreOrder(consumer func(interface{})) {
	tree.root.TraversePreOrder(consumer)
}

func (tree *AVLTree) TraverseInOrder(consumer func(interface{})) {
	tree.root.TraverseInOrder(consumer)
}

func (tree *AVLTree) TraversePostOrder(consumer func(interface{})) {
	tree.root.TraversePostOrder(consumer)
}

func (tree *AVLTree) Clone() *AVLTree {
	return &AVLTree{
		tree.less,
		tree.root.Clone(),
		tree.len,
		tree.cap,
	}
}

func (tree *AVLTree) Len() int {
	return tree.len
}

func (tree *AVLTree) Cap() int {
	return tree.cap
}

func (tree *AVLTree) IsEmpty() bool {
	return tree.len == 0
}

func (tree *AVLTree) IsFull() bool {
	return false
}
