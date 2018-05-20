package avl

// AVLTreeNode is a node of an AVL tree.
type AVLTreeNode struct {
	// Left child and right child pointers in the AVL tree.
	childL, childR *AVLTreeNode

	// The value stored with this node.
	value interface{}

	// The height of this node counting from the left / the right.
	// The height of a non-existing (nil) node is defined to be -1.
	// The height of a leaf node is 0.
	heightL, heightR int
}

func NewAVLTreeNode(childL, childR *AVLTreeNode, v interface{}) *AVLTreeNode {
	return &AVLTreeNode{
		childL,
		childR,
		v,
		childL.Height() + 1,
		childR.Height() + 1,
	}
}

// LeftChild returns the left child of this node or nil.
func (node *AVLTreeNode) LeftChild() *AVLTreeNode {
	return node.childL
}

// RightChild returns the right child of this node or nil.
func (node *AVLTreeNode) RightChild() *AVLTreeNode {
	return node.childR
}

// Value returns the value stored with this node.
func (node *AVLTreeNode) Value() interface{} {
	return node.value
}

// LeftHeight returns the left height of this node.
func (node *AVLTreeNode) LeftHeight() int {
	return node.heightL
}

// RightHeight returns the right height of this node.
func (node *AVLTreeNode) RightHeight() int {
	return node.heightR
}

// Height returns the height of this node,
// i.e. the maximum of its left height and right height.
// The height of a non-existing (nil) node is defined to be -1,
// while the height of a leaf node is 0.
func (node *AVLTreeNode) Height() int {
	if node == nil {
		return -1
	}
	return max(node.heightL, node.heightR)
}

// TODO: Benchmark to see whether this or the next implementation is faster
/*
// UpdateHeight updates the heights of this node.
func (node *AVLTreeNode) UpdateHeight(d direction) {
	if node == nil {
		return
	}
	if d == none || d == left {
		node.heightL = node.childL.Height() + 1
	}
	if d == none || d == right {
		node.heightR = node.childR.Height() + 1
	}
}
*/

// UpdateHeight updates the heights of this node.
func (node *AVLTreeNode) UpdateHeight() {
	node.heightL = node.childL.Height() + 1
	node.heightR = node.childR.Height() + 1
}

/*
rotateLeft performs a left rotation with this node being the root of the rotation.
It returns the new root of the subtree (the pivot).

Left Rotation
-------------
.     R       .       P     .
.    / \      .      / \    .
.   A   P    ==>    R   C   .
.      / \    .    / \      .
.     B   C   .   A   B     .
*/
func (root *AVLTreeNode) rotateLeft() *AVLTreeNode {
	pivot := root.childR
	root.childR, pivot.childL = pivot.childL, root
	pivot.childL.UpdateHeight()
	pivot.UpdateHeight()
	return pivot
}

/*
rotateRight performs a right rotation with this node being the root of the rotation.
It returns the new root of the subtree (the pivot).

Right Rotation
--------------
.       R     .     P       .
.      / \    .    / \      .
.     P   C  ==>  A   R     .
.    / \      .      / \    .
.   A   B     .     B   C   .
*/
func (root *AVLTreeNode) rotateRight() *AVLTreeNode {
	pivot := root.childL
	root.childL, pivot.childR = pivot.childR, root
	pivot.childR.UpdateHeight()
	pivot.UpdateHeight()
	return pivot
}

/*
rotateLL performs an LL-rotation with this node being the root of the rotation.
It returns the new root of the subtree (the pivot of the rotation).

LL Single Rotation
------------------
.         R     .                   .
.        / \    .         P         .
.       P   D   .      +-/ \-+      .
.      / \     ==>    Q       R     .
.     Q   C     .    / \     / \    .
.    / \        .   A   B   C   D   .
.   A   B       .                   .
*/
func (root *AVLTreeNode) rotateLL() *AVLTreeNode {
	return root.rotateRight()
}

/*
rotateRR performs an RR-rotation with this node being the root of the rotation.
It returns the new root of the subtree (the pivot of the rotation).

RR Single Rotation
------------------
.     R         .                   .
.    / \        .         P         .
.   A   P       .      +-/ \-+      .
.      / \     ==>    R       Q     .
.     B   Q     .    / \     / \    .
.        / \    .   A   B   C   D   .
.       C   D   .                   .
*/
func (root *AVLTreeNode) rotateRR() *AVLTreeNode {
	return root.rotateLeft()
}

/*
rotateLR performs an LR-rotation with this node being the root of the rotation.
It returns the new root of the subtree.

LR Double Rotation
------------------
.       R     .         R     .                   .
.      / \    .        / \    .         Q         .
.     P   D   .       Q   D   .      +-/ \-+      .
.    / \     ==>     / \     ==>    P       R     .
.   A   Q     .     P   C     .    / \     / \    .
.      / \    .    / \        .   A   B   C   D   .
.     B   C   .   A   B       .                   .
*/
func (root *AVLTreeNode) rotateLR() *AVLTreeNode {
	root.childL = root.childL.rotateLeft()
	return root.rotateRight()
}

/*
rotateRL performs an RL-rotation with this node being the root of the rotation.
It returns the new root of the subtree.

RL Double Rotation
------------------
.     R       .         R     .                   .
.    / \      .        / \    .         Q         .
.   A   P     .       Q   D   .      +-/ \-+      .
.      / \   ==>     / \     ==>    P       R     .
.     Q   D   .     P   C     .    / \     / \    .
.    / \      .    / \        .   A   B   C   D   .
.   B   C     .   A   B       .                   .
*/
func (root *AVLTreeNode) rotateRL() *AVLTreeNode {
	root.childR = root.childR.rotateRight()
	return root.rotateLeft()
}

func (root *AVLTreeNode) balance() *AVLTreeNode {
	if root.heightL-root.heightR > 1 {
		pivot := root.childL
		if pivot.heightL < pivot.heightR {
			// LR
			root.childL = pivot.rotateLeft()
		}
		// LL
		return root.rotateRight()
	} else if root.heightR-root.heightL > 1 {
		pivot := root.childR
		if pivot.heightR < pivot.heightL {
			// RL
			root.childR = pivot.rotateRight()
		}
		// RR
		return root.rotateLeft()
	} else {
		return root
	}
}

// v: The value to be inserted
// less: The comparison function for the values
// Returns:
// (1) The new root of the subtree rooted at this node after insertion;
// (2) The node being inserted (containing the value "v");
// (3) The height of the subtree, or -1 if new node is not created;
// (4) Whether a new node is allocated;
// (5) The direction of the recursive call to "insert".
func (node *AVLTreeNode) insert(v interface{}, less func(interface{}, interface{}) bool) (*AVLTreeNode, *AVLTreeNode, int, flags) {
	if node == nil {
		newNode := &AVLTreeNode{value: v}
		return newNode, newNode, 0, NodeCreated | DirNone
	}
	if less(v, node.value) {
		// Insert
		newChild, targetNode, newHeight, flags := node.childL.insert(v, less)
		node.childL = newChild
		if flags&NodeCreated == 0 {
			return node, targetNode, -1, DirLeft
		}
		// Update Height
		node.heightL = newHeight + 1
		// Check Height
		if node.heightL-node.heightR <= 1 {
			return node, targetNode, node.Height(), NodeCreated | DirLeft
		}
		// Balance
		if flags&DirRight != 0 {
			// LR
			node.childL = newChild.rotateLeft()
		}
		// LL
		newRoot := node.rotateRight()
		return newRoot, targetNode, newRoot.Height(), NodeCreated | DirLeft
	} else if less(node.value, v) {
		// Insert
		newChild, targetNode, newHeight, flags := node.childR.insert(v, less)
		node.childR = newChild
		if flags&NodeCreated == 0 {
			return node, targetNode, -1, DirRight
		}
		// Update Height
		node.heightR = newHeight + 1
		// Check Height
		if node.heightR-node.heightL <= 1 {
			return node, targetNode, node.Height(), NodeCreated | DirRight
		}
		// Balance
		if flags&DirLeft != 0 {
			// RL
			node.childR = newChild.rotateRight()
		}
		// RR
		newRoot := node.rotateLeft()
		return newRoot, targetNode, newRoot.Height(), NodeCreated | DirRight
	} else {
		return node, node, -1, DirNone
	}
}

// Returns:
// (1) The new root of the subtree rooted at this node after deletion
// (2) The node being deleted (containing the minimum value of the subtree)
func (node *AVLTreeNode) deleteMinimum() (*AVLTreeNode, *AVLTreeNode) {
	if node.childL == nil {
		return node.childR, node
	}
	newChild, targetNode := node.childL.deleteMinimum()
	node.childL = newChild
	node.UpdateHeight()
	if node.heightR-node.heightL <= 1 {
		return node, targetNode
	}
	if node.childR.heightL > node.childR.heightR {
		return node.rotateRL(), targetNode
	} else {
		return node.rotateRR(), targetNode
	}
}

// Returns:
// (1) The new root of the subtree rooted at this node after deletion
// (2) The node being deleted (containing the maximum value of the subtree)
func (node *AVLTreeNode) deleteMaximum() (*AVLTreeNode, *AVLTreeNode) {
	if node.childR == nil {
		return node.childL, node
	}
	newChild, targetNode := node.childR.deleteMaximum()
	node.childR = newChild
	node.UpdateHeight()
	if node.heightL-node.heightR <= 1 {
		return node, targetNode
	}
	if node.childL.heightL > node.childL.heightR {
		return node.rotateLL(), targetNode
	} else {
		return node.rotateLR(), targetNode
	}
}

func (node *AVLTreeNode) delete(v interface{}, less func(interface{}, interface{}) bool) (*AVLTreeNode, *AVLTreeNode, bool, bool) {
	if node == nil {
		return nil, nil, false, false
	}
	if less(v, node.value) {
		newRoot, targetNode, isFound, isDeleted := node.childL.delete(v, less)
		if !isFound {
			return node, nil, false, false
		}
		if !isDeleted {
			return node, targetNode, true, false
		}
		node.childL = newRoot
		node.UpdateHeight()
		if node.heightR-node.heightL <= 1 {
			return node, targetNode, true, true
		}
		if node.childR.heightL > node.childR.heightR {
			return node.rotateRL(), targetNode, true, true
		} else {
			return node.rotateRR(), targetNode, true, true
		}
	} else if less(node.value, v) {
		newRoot, targetNode, isFound, isDeleted := node.childR.delete(v, less)
		if !isFound {
			return node, nil, false, false
		}
		if !isDeleted {
			return node, targetNode, true, false
		}
		node.childR = newRoot
		if node.heightL-node.heightR <= 1 {
			return node, targetNode, true, true
		}
		if node.childL.heightL > node.childL.heightR {
			return node.rotateLL(), targetNode, true, true
		} else {
			return node.rotateLR(), targetNode, true, true
		}
	} else {
		if node.heightL == 0 && node.heightR == 0 {
			return nil, node, true, true
		} else if node.heightL > node.heightR {
			newRoot, maximumNode := node.childL.deleteMaximum()
			maximumNode.childL = newRoot
			maximumNode.childR = node.childR
			maximumNode.UpdateHeight()
			if maximumNode.heightR-maximumNode.heightL <= 1 {
				return maximumNode, node, true, true
			}
			if maximumNode.childR.heightL > maximumNode.childR.heightR {
				return maximumNode.rotateRL(), node, true, true
			} else {
				return maximumNode.rotateRR(), node, true, true
			}
		} else {
			newRoot, minimumNode := node.childR.deleteMinimum()
			minimumNode.childR = newRoot
			minimumNode.childL = node.childL
			minimumNode.UpdateHeight()
			if minimumNode.heightL-minimumNode.heightR <= 1 {
				return minimumNode, node, true, true
			}
			if minimumNode.childL.heightL > minimumNode.childL.heightR {
				return minimumNode.rotateLL(), node, true, true
			} else {
				return minimumNode.rotateLR(), node, true, true
			}
		}
	}
}

func (node *AVLTreeNode) TraversePreOrder(consumer func(interface{})) {
	if node == nil {
		return
	}
	consumer(node.value)
	node.childL.TraversePreOrder(consumer)
	node.childR.TraversePreOrder(consumer)
}

func (node *AVLTreeNode) TraverseInOrder(consumer func(interface{})) {
	if node == nil {
		return
	}
	node.childL.TraverseInOrder(consumer)
	consumer(node.value)
	node.childR.TraverseInOrder(consumer)
}

func (node *AVLTreeNode) TraversePostOrder(consumer func(interface{})) {
	if node == nil {
		return
	}
	node.childL.TraversePostOrder(consumer)
	node.childR.TraversePostOrder(consumer)
	consumer(node.value)
}

func (node *AVLTreeNode) Clone() *AVLTreeNode {
	if node == nil {
		return nil
	}
	return &AVLTreeNode{
		node.childL.Clone(),
		node.childR.Clone(),
		node.value,
		node.heightL,
		node.heightR,
	}
}
