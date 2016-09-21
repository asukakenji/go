package dsalg

type AVLTreeNode struct {
	leftChild, rightChild *AVLTreeNode
	value                 interface{}
	count                 int
	// The height of **this** node counting from the left / the right.
	// The height of a non-existing (nil) node is defined to be -1.
	// The height of a leaf node is 0.
	leftHeight, rightHeight int
}

func (node *AVLTreeNode) LeftChild() *AVLTreeNode {
	return node.leftChild
}

func (node *AVLTreeNode) RightChild() *AVLTreeNode {
	return node.rightChild
}

func (node *AVLTreeNode) Value() interface{} {
	return node.value
}

func (node *AVLTreeNode) Count() int {
	return node.count
}

// Height returns the height of this node, i.e. the maximum of its left height and right height.
// A nil node has a height of -1, while a leaf node has a height of 0.
func (node *AVLTreeNode) Height() int {
	if node == nil {
		return -1
	}
	return max(node.leftHeight, node.rightHeight)
}

// UpdateHeight updates the height of this node.
// TODO: Benchmark to see whether this or the next implementation is faster
/*
func (node *AVLTreeNode) UpdateHeight(d direction) {
	if node == nil {
		return
	}
	if d == nil || d == left {
		node.leftHeight = node.leftChild.Height() + 1
	}
	if d == nil || d == right {
		node.rightHeight = node.rightChild.Height() + 1
	}
}
*/

func (node *AVLTreeNode) UpdateHeight() {
	node.leftHeight = node.leftChild.Height() + 1
	node.rightHeight = node.rightChild.Height() + 1
}

/*
Left Rotation
-------------
.     P       .       Q     .
.    / \      .      / \    .
.   A   Q    ==>    P   C   .
.      / \    .    / \      .
.     B   C   .   A   B     .
*/
func (node *AVLTreeNode) rotateLeft() *AVLTreeNode {
	newRoot := node.rightChild
	node.rightChild, node.rightChild.leftChild = node.rightChild.leftChild, node
	newRoot.leftChild.UpdateHeight()
	newRoot.UpdateHeight()
	return newRoot
}

/*
Right Rotation
--------------
.       P     .     Q       .
.      / \    .    / \      .
.     Q   C  ==>  A   P     .
.    / \      .      / \    .
.   A   B     .     B   C   .
*/
func (node *AVLTreeNode) rotateRight() *AVLTreeNode {
	newRoot := node.leftChild
	node.leftChild, node.leftChild.rightChild = node.leftChild.rightChild, node
	newRoot.rightChild.UpdateHeight()
	newRoot.UpdateHeight()
	return newRoot
}

// LL Single Rotation
func (node *AVLTreeNode) rotateLL() *AVLTreeNode {
	return node.rotateRight()
}

// RR Single Rotation
func (node *AVLTreeNode) rotateRR() *AVLTreeNode {
	return node.rotateLeft()
}

// LR Double Rotation
func (node *AVLTreeNode) rotateLR() *AVLTreeNode {
	node.leftChild = node.leftChild.rotateLeft()
	return node.rotateRight()
}

// RL Double Rotation
func (node *AVLTreeNode) rotateRL() *AVLTreeNode {
	node.rightChild = node.rightChild.rotateRight()
	return node.rotateLeft()
}

// v: The value to be inserted
// less: The comparison function for the values
// Returns:
// (1) The pointer to the new root of the subtree rooted at "node" after insertion
// (2) The pointer to the node containing the value "v"
// (3) Whether a new node is allocated
// (4) The direction of the recursive call to "insert"
func (node *AVLTreeNode) insert(v interface{}, less func(interface{}, interface{}) bool) (*AVLTreeNode, *AVLTreeNode, bool, direction) {
	if node == nil {
		newNode := &AVLTreeNode{value: v, count: 1}
		return newNode, newNode, true, nil
	}
	if less(v, node.value) {
		newRoot, targetNode, isCreated, dir := node.leftChild.insert(v, less)
		node.leftChild = newRoot
		if !isCreated {
			return node, targetNode, false, left
		}
		node.UpdateHeight()
		if node.leftHeight-node.rightHeight <= 1 {
			return node, targetNode, true, left
		}
		if dir == left {
			return node.rotateLL(), targetNode, true, left
		} else if dir == right {
			return node.rotateLR(), targetNode, true, left
		} else {
			panic("Dirction is nil")
		}
	} else if less(node.value, v) {
		newRoot, targetNode, isCreated, dir := node.rightChild.insert(v, less)
		node.rightChild = newRoot
		if !isCreated {
			return node, targetNode, false, right
		}
		node.UpdateHeight()
		if node.rightHeight-node.leftHeight <= 1 {
			return node, targetNode, true, right
		}
		if dir == left {
			return node.rotateRL(), targetNode, true, right
		} else if dir == right {
			return node.rotateRR(), targetNode, true, right
		} else {
			panic("Dirction is nil")
		}
	} else {
		node.count++
		return node, node, false, nil
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

func (tree *AVLTree) Insert(v interface{}) *AVLTreeNode {
	newRoot, targetNode, isCreated, _ := tree.root.insert(v, tree.less)
	tree.root = newRoot
	tree.len++
	if isCreated {
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
