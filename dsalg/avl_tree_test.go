package dsalg

import (
	"testing"
)

func TestTraverseAVL(t *testing.T) {
	tree := NewIntAVLTree()
	checkTraverse(t, tree, tree.TraversePreOrder, []int{})
	checkTraverse(t, tree, tree.TraverseInOrder, []int{})
	checkTraverse(t, tree, tree.TraversePostOrder, []int{})

	tree.Insert(3)
	checkTraverse(t, tree, tree.TraversePreOrder, []int{3})
	checkTraverse(t, tree, tree.TraverseInOrder, []int{3})
	checkTraverse(t, tree, tree.TraversePostOrder, []int{3})

	tree.Insert(2)
	checkTraverse(t, tree, tree.TraversePreOrder, []int{3, 2})
	checkTraverse(t, tree, tree.TraverseInOrder, []int{2, 3})
	checkTraverse(t, tree, tree.TraversePostOrder, []int{2, 3})

	tree.Insert(1)
	checkTraverse(t, tree, tree.TraversePreOrder, []int{2, 1, 3})
	checkTraverse(t, tree, tree.TraverseInOrder, []int{1, 2, 3})
	checkTraverse(t, tree, tree.TraversePostOrder, []int{1, 3, 2})

	tree.Insert(4)
	checkTraverse(t, tree, tree.TraversePreOrder, []int{2, 1, 3, 4})
	checkTraverse(t, tree, tree.TraverseInOrder, []int{1, 2, 3, 4})
	checkTraverse(t, tree, tree.TraversePostOrder, []int{1, 4, 3, 2})

	tree.Insert(5)
	checkTraverse(t, tree, tree.TraversePreOrder, []int{2, 1, 4, 3, 5})
	checkTraverse(t, tree, tree.TraverseInOrder, []int{1, 2, 3, 4, 5})
	checkTraverse(t, tree, tree.TraversePostOrder, []int{1, 3, 5, 4, 2})

	tree.Insert(6)
	checkTraverse(t, tree, tree.TraversePreOrder, []int{4, 2, 1, 3, 5, 6})
	checkTraverse(t, tree, tree.TraverseInOrder, []int{1, 2, 3, 4, 5, 6})
	checkTraverse(t, tree, tree.TraversePostOrder, []int{1, 3, 2, 6, 5, 4})

	tree.Insert(7)
	checkTraverse(t, tree, tree.TraversePreOrder, []int{4, 2, 1, 3, 6, 5, 7})
	checkTraverse(t, tree, tree.TraverseInOrder, []int{1, 2, 3, 4, 5, 6, 7})
	checkTraverse(t, tree, tree.TraversePostOrder, []int{1, 3, 2, 5, 7, 6, 4})
}
