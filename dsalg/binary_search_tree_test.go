package dsalg

import (
	"testing"
)

func checkLRV(t *testing.T, n, expectedL, expectedR *BinarySearchTreeNode, expectedV int) {
	l := n.leftChild
	if l != expectedL {
		t.Errorf("n.leftChild == %v; expected %v", l, expectedL)
		panic("")
	}
	r := n.rightChild
	if r != expectedR {
		t.Errorf("n.rightChild == %v; expected %v", r, expectedR)
		panic("")
	}
	v := n.value
	if v != expectedV {
		t.Errorf("n.value == %v; expected %v", v, expectedV)
		panic("")
	}
}

func checkLen(t *testing.T, tree *BinarySearchTree, expectedLength int) {
	length := tree.Len()
	if length != expectedLength {
		t.Errorf("n.Len() == %d; expected %d", length, expectedLength)
		panic("")
	}
}

func TestNewBinaryTree(t *testing.T) {
	tree := NewIntBinarySearchTree()
	checkLen(t, tree, 0)
	isEmpty := tree.IsEmpty()
	if !isEmpty {
		t.Errorf("tree.IsEmpty() == false; expected true")
	}
}

func TestInsert(t *testing.T) {
	tree := NewIntBinarySearchTree()

	tree.Insert(4)
	checkLRV(t, tree.root, nil, nil, 4)
	checkLen(t, tree, 1)

	tree.Insert(2)
	checkLRV(t, tree.root, tree.root.leftChild, nil, 4)
	checkLRV(t, tree.root.leftChild, nil, nil, 2)
	checkLen(t, tree, 2)

	tree.Insert(6)
	checkLRV(t, tree.root, tree.root.leftChild, tree.root.rightChild, 4)
	checkLRV(t, tree.root.leftChild, nil, nil, 2)
	checkLRV(t, tree.root.rightChild, nil, nil, 6)
	checkLen(t, tree, 3)

	tree.Insert(1)
	checkLRV(t, tree.root, tree.root.leftChild, tree.root.rightChild, 4)
	checkLRV(t, tree.root.leftChild, tree.root.leftChild.leftChild, nil, 2)
	checkLRV(t, tree.root.rightChild, nil, nil, 6)
	checkLRV(t, tree.root.leftChild.leftChild, nil, nil, 1)
	checkLen(t, tree, 4)
}

func TestTraverseBST(t *testing.T) {
	tree := NewIntBinarySearchTree()
	checkTraverse(t, tree, tree.TraversePreOrder, []interface{}{})
	checkTraverse(t, tree, tree.TraverseInOrder, []interface{}{})
	checkTraverse(t, tree, tree.TraversePostOrder, []interface{}{})

	tree.Insert(4)
	checkTraverse(t, tree, tree.TraversePreOrder, []interface{}{4})
	checkTraverse(t, tree, tree.TraverseInOrder, []interface{}{4})
	checkTraverse(t, tree, tree.TraversePostOrder, []interface{}{4})

	tree.Insert(2)
	checkTraverse(t, tree, tree.TraversePreOrder, []interface{}{4, 2})
	checkTraverse(t, tree, tree.TraverseInOrder, []interface{}{2, 4})
	checkTraverse(t, tree, tree.TraversePostOrder, []interface{}{2, 4})

	tree.Insert(6)
	checkTraverse(t, tree, tree.TraversePreOrder, []interface{}{4, 2, 6})
	checkTraverse(t, tree, tree.TraverseInOrder, []interface{}{2, 4, 6})
	checkTraverse(t, tree, tree.TraversePostOrder, []interface{}{2, 6, 4})

	tree.Insert(1)
	checkTraverse(t, tree, tree.TraversePreOrder, []interface{}{4, 2, 1, 6})
	checkTraverse(t, tree, tree.TraverseInOrder, []interface{}{1, 2, 4, 6})
	checkTraverse(t, tree, tree.TraversePostOrder, []interface{}{1, 2, 6, 4})

	tree.Insert(3)
	checkTraverse(t, tree, tree.TraversePreOrder, []interface{}{4, 2, 1, 3, 6})
	checkTraverse(t, tree, tree.TraverseInOrder, []interface{}{1, 2, 3, 4, 6})
	checkTraverse(t, tree, tree.TraversePostOrder, []interface{}{1, 3, 2, 6, 4})

	tree.Insert(5)
	checkTraverse(t, tree, tree.TraversePreOrder, []interface{}{4, 2, 1, 3, 6, 5})
	checkTraverse(t, tree, tree.TraverseInOrder, []interface{}{1, 2, 3, 4, 5, 6})
	checkTraverse(t, tree, tree.TraversePostOrder, []interface{}{1, 3, 2, 5, 6, 4})

	tree.Insert(7)
	checkTraverse(t, tree, tree.TraversePreOrder, []interface{}{4, 2, 1, 3, 6, 5, 7})
	checkTraverse(t, tree, tree.TraverseInOrder, []interface{}{1, 2, 3, 4, 5, 6, 7})
	checkTraverse(t, tree, tree.TraversePostOrder, []interface{}{1, 3, 2, 5, 7, 6, 4})
}
