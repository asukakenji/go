package dsalg

type IntAVLTreeNode struct {
	LeftChild, RightChild   *IntAVLTreeNode
	Value                   int
	Count                   int
	LeftHeight, RightHeight int
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func (node *IntAVLTreeNode) Height() int {
	if node == nil {
		return -1
	}
	return max(node.LeftHeight, node.RightHeight)
}

// LL Single Rotation
func (node *IntAVLTreeNode) RotateLL(ptrParent **IntAVLTreeNode) {
	*ptrParent = node.LeftChild
	node.LeftChild, node.LeftChild.RightChild = node.LeftChild.RightChild, node
	node.LeftHeight = node.LeftChild.Height() + 1
	(*ptrParent).RightHeight = (*ptrParent).RightChild.Height() + 1
}

// RR Single Rotation
func (node *IntAVLTreeNode) RotateRR(ptrParent **IntAVLTreeNode) {
	*ptrParent = node.RightChild
	node.RightChild, node.RightChild.LeftChild = node.RightChild.LeftChild, node
	node.RightHeight = node.RightChild.Height() + 1
	(*ptrParent).LeftHeight = (*ptrParent).LeftChild.Height() + 1
}

// LR Double Rotation
func (node *IntAVLTreeNode) RotateLR(ptrParent **IntAVLTreeNode) {
	*ptrParent = node.LeftChild.RightChild
	node.LeftChild, node.LeftChild.RightChild, node.LeftChild.RightChild.LeftChild, node.LeftChild.RightChild.RightChild = node.LeftChild.RightChild.RightChild, node.LeftChild.RightChild.LeftChild, node.LeftChild, node
	node.LeftHeight = node.LeftChild.Height() + 1
	(*ptrParent).LeftChild.RightHeight = (*ptrParent).LeftChild.RightChild.Height() + 1
	(*ptrParent).LeftHeight = (*ptrParent).LeftChild.Height() + 1
	(*ptrParent).RightHeight = (*ptrParent).RightChild.Height() + 1
}

// RL Double Rotation
func (node *IntAVLTreeNode) RotateRL(ptrParent **IntAVLTreeNode) {
	*ptrParent = node.RightChild.LeftChild
	node.RightChild, node.RightChild.LeftChild, node.RightChild.LeftChild.LeftChild, node.RightChild.LeftChild.RightChild = node.RightChild.LeftChild.LeftChild, node.RightChild.LeftChild.RightChild, node, node.RightChild
	node.RightHeight = node.RightChild.Height() + 1
	(*ptrParent).RightChild.LeftHeight = (*ptrParent).RightChild.LeftChild.Height() + 1
	(*ptrParent).LeftHeight = (*ptrParent).LeftChild.Height() + 1
	(*ptrParent).RightHeight = (*ptrParent).RightChild.Height() + 1
}

// ptrParent: The pointer to the pointer from the parent node pointing to the "this" node
// x: The element to be inserted
// Returns:
// (1) whether a new node is allocated
// (2) whether the balance factor of the caller should be updated
// (3) whether the recursive call to insert is on the LeftChild or the RightChild of "this" node
func (node *IntAVLTreeNode) Insert(ptrParent **IntAVLTreeNode, x int) (bool, direction) {
	if node == nil {
		*ptrParent = &IntAVLTreeNode{Value: x, Count: 1}
		return true, nil
	}
	if x < node.Value {
		isNewNodeCreated, d := node.LeftChild.Insert(&node.LeftChild, x)
		if !isNewNodeCreated {
			return false, left
		}
		node.LeftHeight = node.LeftChild.Height() + 1
		if node.LeftHeight-node.RightHeight <= 1 {
			return true, left
		}
		if d == left {
			node.RotateLL(ptrParent)
		} else if d == right {
			node.RotateLR(ptrParent)
		} else {
			panic("")
		}
		return true, left
	} else if x > node.Value {
		isNewNodeCreated, d := node.RightChild.Insert(&node.RightChild, x)
		if !isNewNodeCreated {
			return false, right
		}
		node.RightHeight = node.RightChild.Height() + 1
		if node.RightHeight-node.LeftHeight <= 1 {
			return true, right
		}
		if d == left {
			node.RotateRL(ptrParent)
		} else if d == right {
			node.RotateRR(ptrParent)
		} else {
			panic("")
		}
		return true, right
	} else {
		node.Count++
		return false, nil
	}
}

func (node *IntAVLTreeNode) TraversePreOrder(consumer func(int)) {
	if node == nil {
		return
	}
	for i := 0; i < node.Count; i++ {
		consumer(node.Value)
	}
	node.LeftChild.TraversePreOrder(consumer)
	node.RightChild.TraversePreOrder(consumer)
}

func (node *IntAVLTreeNode) TraverseInOrder(consumer func(int)) {
	if node == nil {
		return
	}
	node.LeftChild.TraverseInOrder(consumer)
	for i := 0; i < node.Count; i++ {
		consumer(node.Value)
	}
	node.RightChild.TraverseInOrder(consumer)
}

func (node *IntAVLTreeNode) TraversePostOrder(consumer func(int)) {
	if node == nil {
		return
	}
	node.LeftChild.TraversePostOrder(consumer)
	node.RightChild.TraversePostOrder(consumer)
	for i := 0; i < node.Count; i++ {
		consumer(node.Value)
	}
}

func (node *IntAVLTreeNode) Clone() *IntAVLTreeNode {
	if node == nil {
		return nil
	}
	return &IntAVLTreeNode{
		node.LeftChild.Clone(),
		node.RightChild.Clone(),
		node.Value,
		node.Count,
		node.LeftHeight,
		node.RightHeight,
	}
}

type IntAVLTree struct {
	root     *IntAVLTreeNode
	length   int
	capacity int
}

func NewIntAVLTree() *IntAVLTree {
	return &IntAVLTree{}
}

func (tree *IntAVLTree) Insert(x int) {
	if isNewNodeCreated, _ := tree.root.Insert(&tree.root, x); isNewNodeCreated {
		tree.capacity++
	}
	tree.length++
}

func (tree *IntAVLTree) TraversePreOrder(consumer func(int)) {
	tree.root.TraversePreOrder(consumer)
}

func (tree *IntAVLTree) TraverseInOrder(consumer func(int)) {
	tree.root.TraverseInOrder(consumer)
}

func (tree *IntAVLTree) TraversePostOrder(consumer func(int)) {
	tree.root.TraversePostOrder(consumer)
}

func (tree *IntAVLTree) Clone() *IntAVLTree {
	return &IntAVLTree{
		tree.root.Clone(),
		tree.length,
		tree.capacity,
	}
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
