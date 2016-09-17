package dsalg

import (
	"testing"
)

func insertMultiple(tree *IntAVLTree, xs ...int) {
	for _, x := range xs {
		tree.Insert(x)
	}
}

func TestClone(t *testing.T) {
	tree := NewIntAVLTree()
	insertMultiple(tree, 14, 3, 16, 2, 6, 15, 17, 1, 5, 10, 18, 4, 8, 12)
	checkTraverse(t, tree, tree.TraversePreOrder, []int{14, 3, 2, 1, 6, 5, 4, 10, 8, 12, 16, 15, 17, 18})
	checkTraverse(t, tree, tree.TraverseInOrder, []int{1, 2, 3, 4, 5, 6, 8, 10, 12, 14, 15, 16, 17, 18})
	checkTraverse(t, tree, tree.TraversePostOrder, []int{1, 2, 4, 5, 8, 12, 10, 6, 3, 15, 18, 17, 16, 14})

	tree1 := tree.Clone()
	if tree1 == tree {
		t.Errorf("tree1 == tree; expected tree1 != tree")
	}
	if tree1.Len() != tree.Len() {
		t.Errorf("tree1.Len() == %d; expected %d", tree1.Len(), tree.Len())
	}
	if tree1.Cap() != tree.Cap() {
		t.Errorf("tree1.Cap() == %d; expected %d", tree1.Cap(), tree.Cap())
	}
	checkTraverse(t, tree1, tree1.TraversePreOrder, []int{14, 3, 2, 1, 6, 5, 4, 10, 8, 12, 16, 15, 17, 18})
	checkTraverse(t, tree1, tree1.TraverseInOrder, []int{1, 2, 3, 4, 5, 6, 8, 10, 12, 14, 15, 16, 17, 18})
	checkTraverse(t, tree1, tree1.TraversePostOrder, []int{1, 2, 4, 5, 8, 12, 10, 6, 3, 15, 18, 17, 16, 14})
}

func TestInsertAVL1(t *testing.T) {
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

func TestInsertAVL2(t *testing.T) {
	tree := NewIntAVLTree()
	checkTraverse(t, tree, tree.TraversePreOrder, []int{})
	checkTraverse(t, tree, tree.TraverseInOrder, []int{})
	checkTraverse(t, tree, tree.TraversePostOrder, []int{})

	tree.Insert(7)
	checkTraverse(t, tree, tree.TraversePreOrder, []int{7})
	checkTraverse(t, tree, tree.TraverseInOrder, []int{7})
	checkTraverse(t, tree, tree.TraversePostOrder, []int{7})

	tree.Insert(4)
	checkTraverse(t, tree, tree.TraversePreOrder, []int{7, 4})
	checkTraverse(t, tree, tree.TraverseInOrder, []int{4, 7})
	checkTraverse(t, tree, tree.TraversePostOrder, []int{4, 7})

	tree.Insert(9)
	checkTraverse(t, tree, tree.TraversePreOrder, []int{7, 4, 9})
	checkTraverse(t, tree, tree.TraverseInOrder, []int{4, 7, 9})
	checkTraverse(t, tree, tree.TraversePostOrder, []int{4, 9, 7})

	tree.Insert(2)
	checkTraverse(t, tree, tree.TraversePreOrder, []int{7, 4, 2, 9})
	checkTraverse(t, tree, tree.TraverseInOrder, []int{2, 4, 7, 9})
	checkTraverse(t, tree, tree.TraversePostOrder, []int{2, 4, 9, 7})

	tree.Insert(6)
	checkTraverse(t, tree, tree.TraversePreOrder, []int{7, 4, 2, 6, 9})
	checkTraverse(t, tree, tree.TraverseInOrder, []int{2, 4, 6, 7, 9})
	checkTraverse(t, tree, tree.TraversePostOrder, []int{2, 6, 4, 9, 7})

	tree.Insert(1)
	checkTraverse(t, tree, tree.TraversePreOrder, []int{4, 2, 1, 7, 6, 9})
	checkTraverse(t, tree, tree.TraverseInOrder, []int{1, 2, 4, 6, 7, 9})
	checkTraverse(t, tree, tree.TraversePostOrder, []int{1, 2, 6, 9, 7, 4})

	tree.Insert(8)
	checkTraverse(t, tree, tree.TraversePreOrder, []int{4, 2, 1, 7, 6, 9, 8})
	checkTraverse(t, tree, tree.TraverseInOrder, []int{1, 2, 4, 6, 7, 8, 9})
	checkTraverse(t, tree, tree.TraversePostOrder, []int{1, 2, 6, 8, 9, 7, 4})
}

/*
Rotation root being left child
------------------------------

Insert: 14, 3, 16, 2, 6, 15,
        17, 1, 5, 10, 18, 4,
        8, 12

        14
     +-/ \-+
    3       16
   / \     / \
  2   6   15  17
 /   / \       \
1   5   10      18
   /   / \
  4   8   12

 8: (0,0)
10: (1,1)
 6: (2,2)
 3: (2,3)
14: (4,3)

Now insert: 7 (or 9 or 11 or 13)

          14
      +--/ \--+
     3         16
   +/ \+      / \
  2     6    15  17
 /    +/ \+       \
1    5     10      18
    /   +-/ \-+
   4   8       12
      / \     / \
    (7) [9] [11][13]

 7: (0,0)
 8: (0,0) -> (1,0)
10: (1,1) -> (2,1)
 6: (2,2) -> (2,3)
 3: (2,3) -> (2,4) X
14: (4,3)

Rotate with: root 3, pivot 6

               14
          +---/ \---+
         6           16
     +--/ \--+      / \
    3         10   15  17
   / \     +-/ \-+      \
  2   5   8       12     18
 /   /   / \     / \
1   4  (7) [9] [11][13]

 7: (0,0)
 8: (0,0) -> (1,0)
10: (1,1) -> (2,1)
 6: (2,2) -> (2,3) -> (3,3)
 3: (2,3) -> (2,4) -> (2,2)
14: (4,3)
*/

func TestInsertAVLRR(t *testing.T) {
	tree := NewIntAVLTree()
	insertMultiple(tree, 14, 3, 16, 2, 6, 15, 17, 1, 5, 10, 18, 4, 8, 12)
	checkTraverse(t, tree, tree.TraversePreOrder, []int{14, 3, 2, 1, 6, 5, 4, 10, 8, 12, 16, 15, 17, 18})
	checkTraverse(t, tree, tree.TraverseInOrder, []int{1, 2, 3, 4, 5, 6, 8, 10, 12, 14, 15, 16, 17, 18})
	checkTraverse(t, tree, tree.TraversePostOrder, []int{1, 2, 4, 5, 8, 12, 10, 6, 3, 15, 18, 17, 16, 14})

	tree1 := tree.Clone()
	tree1.Insert(7)
	checkTraverse(t, tree1, tree1.TraversePreOrder, []int{14, 6, 3, 2, 1, 5, 4, 10, 8, 7, 12, 16, 15, 17, 18})
	checkTraverse(t, tree1, tree1.TraverseInOrder, []int{1, 2, 3, 4, 5, 6, 7, 8, 10, 12, 14, 15, 16, 17, 18})
	checkTraverse(t, tree1, tree1.TraversePostOrder, []int{1, 2, 4, 5, 3, 7, 8, 12, 10, 6, 15, 18, 17, 16, 14})

	tree2 := tree.Clone()
	tree2.Insert(9)
	checkTraverse(t, tree2, tree2.TraversePreOrder, []int{14, 6, 3, 2, 1, 5, 4, 10, 8, 9, 12, 16, 15, 17, 18})
	checkTraverse(t, tree2, tree2.TraverseInOrder, []int{1, 2, 3, 4, 5, 6, 8, 9, 10, 12, 14, 15, 16, 17, 18})
	checkTraverse(t, tree2, tree2.TraversePostOrder, []int{1, 2, 4, 5, 3, 9, 8, 12, 10, 6, 15, 18, 17, 16, 14})

	tree3 := tree.Clone()
	tree3.Insert(11)
	checkTraverse(t, tree3, tree3.TraversePreOrder, []int{14, 6, 3, 2, 1, 5, 4, 10, 8, 12, 11, 16, 15, 17, 18})
	checkTraverse(t, tree3, tree3.TraverseInOrder, []int{1, 2, 3, 4, 5, 6, 8, 10, 11, 12, 14, 15, 16, 17, 18})
	checkTraverse(t, tree3, tree3.TraversePostOrder, []int{1, 2, 4, 5, 3, 8, 11, 12, 10, 6, 15, 18, 17, 16, 14})

	tree4 := tree.Clone()
	tree4.Insert(13)
	checkTraverse(t, tree4, tree4.TraversePreOrder, []int{14, 6, 3, 2, 1, 5, 4, 10, 8, 12, 13, 16, 15, 17, 18})
	checkTraverse(t, tree4, tree4.TraverseInOrder, []int{1, 2, 3, 4, 5, 6, 8, 10, 12, 13, 14, 15, 16, 17, 18})
	checkTraverse(t, tree4, tree4.TraversePostOrder, []int{1, 2, 4, 5, 3, 8, 13, 12, 10, 6, 15, 18, 17, 16, 14})

	tree5 := tree.Clone()
	insertMultiple(tree5, 7, 9, 11, 13)
	checkTraverse(t, tree5, tree5.TraversePreOrder, []int{14, 6, 3, 2, 1, 5, 4, 10, 8, 7, 9, 12, 11, 13, 16, 15, 17, 18})
	checkTraverse(t, tree5, tree5.TraverseInOrder, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18})
	checkTraverse(t, tree5, tree5.TraversePostOrder, []int{1, 2, 4, 5, 3, 7, 9, 8, 11, 13, 12, 10, 6, 15, 18, 17, 16, 14})
}
