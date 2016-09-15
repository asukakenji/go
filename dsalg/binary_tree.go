package dsalg

type intBinarySearchTreeNode struct {
	parent, leftChild, rightChild *intBinarySearchTreeNode
	value                         int
}

func (node *intBinarySearchTreeNode) traversePreOrder(consumer func(int)) {
	if node == nil {
        return
    }
	consumer(node.value)
	node.leftChild.traversePreOrder(consumer)
	node.rightChild.traversePreOrder(consumer)
}

func (node *intBinarySearchTreeNode) traverseInOrder(consumer func(int)) {
    if node == nil {
        return
    }
    node.leftChild.traverseInOrder(consumer)
    consumer(node.value)
	node.rightChild.traverseInOrder(consumer)
}

func (node *intBinarySearchTreeNode) traversePostOrder(consumer func(int)) {
    if node == nil {
        return
    }
    node.leftChild.traversePostOrder(consumer)
	node.rightChild.traversePostOrder(consumer)
    consumer(node.value)
}

type IntBinarySearchTree struct {
	root   *intBinarySearchTreeNode
	length int
}

func NewIntBinarySearchTree() *IntBinarySearchTree {
	return &IntBinarySearchTree{}
}

func (t *IntBinarySearchTree) Insert(x int) {
	parent := (*intBinarySearchTreeNode)(nil)
	ptr := &t.root
	curr := t.root
	for {
		if curr == nil {
			*ptr = &intBinarySearchTreeNode{parent: parent, value: x}
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

func (t *IntBinarySearchTree) TraversePreOrder(consumer func(int)) {
	t.root.traversePreOrder(consumer)
}

func (t *IntBinarySearchTree) TraverseInOrder(consumer func(int)) {
	t.root.traverseInOrder(consumer)
}

func (t *IntBinarySearchTree) TraversePostOrder(consumer func(int)) {
	t.root.traversePostOrder(consumer)
}

func (t *IntBinarySearchTree) Len() int {
	return t.length
}

func (t *IntBinarySearchTree) IsEmpty() bool {
	return t.root == nil
}
