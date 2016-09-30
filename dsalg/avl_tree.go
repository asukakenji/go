package dsalg

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

func (tree *AVLTree) Insert(v interface{}) *AVLTreeNode {
	newRoot, targetNode, _, flags := tree.root.insert(v, tree.less)
	tree.root = newRoot
	tree.len++
	if flags&NodeCreated != 0 {
		tree.cap++
	}
	return targetNode
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
