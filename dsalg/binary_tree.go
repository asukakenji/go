package dsalg

type intBinaryTreeNode struct {
	parent, leftChild, rightChild *intBinaryTreeNode
	value                         int
}

func (node *intBinaryTreeNode) traversePreOrder(consumer func(int)) {
	if node != nil {
		consumer(node.value)
		node.leftChild.traversePreOrder(consumer)
		node.rightChild.traversePreOrder(consumer)
	}
}

func (node *intBinaryTreeNode) traverseInOrder(consumer func(int)) {
	if node != nil {
        node.leftChild.traverseInOrder(consumer)
        consumer(node.value)
		node.rightChild.traverseInOrder(consumer)
	}
}

func (node *intBinaryTreeNode) traversePostOrder(consumer func(int)) {
	if node != nil {
        node.leftChild.traversePostOrder(consumer)
		node.rightChild.traversePostOrder(consumer)
        consumer(node.value)
	}
}

type IntBinaryTree struct {
	root   *intBinaryTreeNode
	length int
}

func NewIntBinaryTree() *IntBinaryTree {
	return &IntBinaryTree{}
}

func (t *IntBinaryTree) Insert(x int) {
	parent := (*intBinaryTreeNode)(nil)
	ptr := &t.root
	curr := t.root
	for {
		if curr == nil {
			*ptr = &intBinaryTreeNode{parent: parent, value: x}
			t.length++
			return
		}
		if x < curr.value {
			parent = curr
			ptr = &curr.leftChild
			curr = curr.leftChild
		} else if x > curr.value {
			parent = curr
			ptr = &curr.rightChild
			curr = curr.rightChild
		} else {
			return
		}
	}
}

func (t *IntBinaryTree) TraversePreOrder(consumer func(int)) {
	t.root.traversePreOrder(consumer)
}

func (t *IntBinaryTree) TraverseInOrder(consumer func(int)) {
	t.root.traverseInOrder(consumer)
}

func (t *IntBinaryTree) TraversePostOrder(consumer func(int)) {
	t.root.traversePostOrder(consumer)
}

func (t *IntBinaryTree) Len() int {
	return t.length
}

func (t *IntBinaryTree) IsEmpty() bool {
	return t.root == nil
}
