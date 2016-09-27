package dsalg

// BinaryTreeNode is a node of a binary tree.
type BinaryTreeNode struct {
	// Left child and right child pointers in the binary tree.
	leftChild, rightChild *BinaryTreeNode

	// The value stored with this node.
	value interface{}
}

// LeftChild returns the left child of this node or nil.
func (node *BinaryTreeNode) LeftChild() *BinaryTreeNode {
	return node.leftChild
}

// RightChild returns the right child of this node or nil.
func (node *BinaryTreeNode) RightChild() *BinaryTreeNode {
	return node.rightChild
}

// Value returns the value stored with this node.
func (node *BinaryTreeNode) Value() interface{} {
	return node.value
}

/*
rotateLeft performs a left rotation with this node being the root of the rotation.
It returns the new root of the subtree (the pivot of the rotation).

Left Rotation
-------------
.     R       .       P     .
.    / \      .      / \    .
.   A   P    ==>    R   C   .
.      / \    .    / \      .
.     B   C   .   A   B     .
*/
func (root *BinaryTreeNode) rotateLeft() *BinaryTreeNode {
	pivot := root.rightChild
	root.rightChild, pivot.leftChild = pivot.leftChild, root
	return pivot
}

/*
rotateRight performs a right rotation with this node being the root of the rotation.
It returns the new root of the subtree (the pivot of the rotation).

Right Rotation
--------------
.       R     .     P       .
.      / \    .    / \      .
.     P   C  ==>  A   R     .
.    / \      .      / \    .
.   A   B     .     B   C   .
*/
func (root *BinaryTreeNode) rotateRight() *BinaryTreeNode {
	pivot := root.leftChild
	root.leftChild, pivot.rightChild = pivot.rightChild, root
	return pivot
}
