package dsalg

import (
    "reflect"
	"testing"
)

func checkPLRV(t *testing.T, n, expectedP, expectedL, expectedR *intBinaryTreeNode, expectedV int) {
	p := n.parent
	if p != expectedP {
		t.Errorf("n.parent == %v; expected %v", p, expectedP)
		panic("")
	}
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

func checkLen(t *testing.T, tree *IntBinaryTree, expectedLength int) {
    length := tree.Len()
    if length != expectedLength {
        t.Errorf("n.Len() == %d; expected %d", length, expectedLength)
        panic("")
    }
}

func checkTraverse(t *testing.T, tree *IntBinaryTree, traverse func(func(int)), expectedStack []int) {
    stack := make([]int, 0, tree.Len())
    consumer := func(x int) {
        stack = append(stack, x)
    }
    traverse(consumer)
    if !reflect.DeepEqual(stack, expectedStack) {
        t.Errorf("tree.TraverseXXX(consumer) == %v; expected %v", stack, expectedStack)
        panic("")
    }
}

func TestNewBinaryTree(t *testing.T) {
	tree := NewIntBinaryTree()
    checkLen(t, tree, 0)
	isEmpty := tree.IsEmpty()
	if !isEmpty {
		t.Errorf("tree.IsEmpty() == false; expected true")
	}
}

func TestInsert(t *testing.T) {
	tree := NewIntBinaryTree()

	tree.Insert(4)
	checkPLRV(t, tree.root, nil, nil, nil, 4)
    checkLen(t, tree, 1)

	tree.Insert(2)
	checkPLRV(t, tree.root, nil, tree.root.leftChild, nil, 4)
	checkPLRV(t, tree.root.leftChild, tree.root, nil, nil, 2)
    checkLen(t, tree, 2)

	tree.Insert(6)
	checkPLRV(t, tree.root, nil, tree.root.leftChild, tree.root.rightChild, 4)
	checkPLRV(t, tree.root.leftChild, tree.root, nil, nil, 2)
	checkPLRV(t, tree.root.rightChild, tree.root, nil, nil, 6)
    checkLen(t, tree, 3)

    tree.Insert(1)
	checkPLRV(t, tree.root, nil, tree.root.leftChild, tree.root.rightChild, 4)
	checkPLRV(t, tree.root.leftChild, tree.root, tree.root.leftChild.leftChild, nil, 2)
	checkPLRV(t, tree.root.rightChild, tree.root, nil, nil, 6)
    checkPLRV(t, tree.root.leftChild.leftChild, tree.root.leftChild, nil, nil, 1)
    checkLen(t, tree, 4)
}

func TestTraverse(t *testing.T) {
    tree := NewIntBinaryTree()
    checkTraverse(t, tree, tree.TraversePreOrder, []int{})
    checkTraverse(t, tree, tree.TraverseInOrder, []int{})
    checkTraverse(t, tree, tree.TraversePostOrder, []int{})

    tree.Insert(4)
    checkTraverse(t, tree, tree.TraversePreOrder, []int{4})
    checkTraverse(t, tree, tree.TraverseInOrder, []int{4})
    checkTraverse(t, tree, tree.TraversePostOrder, []int{4})

    tree.Insert(2)
    checkTraverse(t, tree, tree.TraversePreOrder, []int{4, 2})
    checkTraverse(t, tree, tree.TraverseInOrder, []int{2, 4})
    checkTraverse(t, tree, tree.TraversePostOrder, []int{2, 4})

    tree.Insert(6)
    checkTraverse(t, tree, tree.TraversePreOrder, []int{4, 2, 6})
    checkTraverse(t, tree, tree.TraverseInOrder, []int{2, 4, 6})
    checkTraverse(t, tree, tree.TraversePostOrder, []int{2, 6, 4})

    tree.Insert(1)
    checkTraverse(t, tree, tree.TraversePreOrder, []int{4, 2, 1, 6})
    checkTraverse(t, tree, tree.TraverseInOrder, []int{1, 2, 4, 6})
    checkTraverse(t, tree, tree.TraversePostOrder, []int{1, 2, 6, 4})

    tree.Insert(3)
    checkTraverse(t, tree, tree.TraversePreOrder, []int{4, 2, 1, 3, 6})
    checkTraverse(t, tree, tree.TraverseInOrder, []int{1, 2, 3, 4, 6})
    checkTraverse(t, tree, tree.TraversePostOrder, []int{1, 3, 2, 6, 4})

    tree.Insert(5)
    checkTraverse(t, tree, tree.TraversePreOrder, []int{4, 2, 1, 3, 6, 5})
    checkTraverse(t, tree, tree.TraverseInOrder, []int{1, 2, 3, 4, 5, 6})
    checkTraverse(t, tree, tree.TraversePostOrder, []int{1, 3, 2, 5, 6, 4})

    tree.Insert(7)
    checkTraverse(t, tree, tree.TraversePreOrder, []int{4, 2, 1, 3, 6, 5, 7})
    checkTraverse(t, tree, tree.TraverseInOrder, []int{1, 2, 3, 4, 5, 6, 7})
    checkTraverse(t, tree, tree.TraversePostOrder, []int{1, 3, 2, 5, 7, 6, 4})
}
