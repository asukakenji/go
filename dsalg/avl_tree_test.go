package dsalg

import (
	"testing"
)

func insertMultiple(tree *AVLTree, vs ...interface{}) {
	for _, v := range vs {
		tree.Insert(v)
	}
}

func TestClone(t *testing.T) {
	tree := NewIntAVLTree()
	insertMultiple(tree, 14, 3, 16, 2, 6, 15, 17, 1, 5, 10, 18, 4, 8, 12)
	checkTraverseInt(t, tree, tree.TraversePreOrder, []int{14, 3, 2, 1, 6, 5, 4, 10, 8, 12, 16, 15, 17, 18})
	checkTraverseInt(t, tree, tree.TraverseInOrder, []int{1, 2, 3, 4, 5, 6, 8, 10, 12, 14, 15, 16, 17, 18})
	checkTraverseInt(t, tree, tree.TraversePostOrder, []int{1, 2, 4, 5, 8, 12, 10, 6, 3, 15, 18, 17, 16, 14})

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
	checkTraverseInt(t, tree1, tree1.TraversePreOrder, []int{14, 3, 2, 1, 6, 5, 4, 10, 8, 12, 16, 15, 17, 18})
	checkTraverseInt(t, tree1, tree1.TraverseInOrder, []int{1, 2, 3, 4, 5, 6, 8, 10, 12, 14, 15, 16, 17, 18})
	checkTraverseInt(t, tree1, tree1.TraversePostOrder, []int{1, 2, 4, 5, 8, 12, 10, 6, 3, 15, 18, 17, 16, 14})
}

func TestInsertAVL1(t *testing.T) {
	tree := NewIntAVLTree()
	checkTraverseInt(t, tree, tree.TraversePreOrder, []int{})
	checkTraverseInt(t, tree, tree.TraverseInOrder, []int{})
	checkTraverseInt(t, tree, tree.TraversePostOrder, []int{})

	tree.Insert(3)
	checkTraverseInt(t, tree, tree.TraversePreOrder, []int{3})
	checkTraverseInt(t, tree, tree.TraverseInOrder, []int{3})
	checkTraverseInt(t, tree, tree.TraversePostOrder, []int{3})

	tree.Insert(2)
	checkTraverseInt(t, tree, tree.TraversePreOrder, []int{3, 2})
	checkTraverseInt(t, tree, tree.TraverseInOrder, []int{2, 3})
	checkTraverseInt(t, tree, tree.TraversePostOrder, []int{2, 3})

	tree.Insert(1)
	checkTraverseInt(t, tree, tree.TraversePreOrder, []int{2, 1, 3})
	checkTraverseInt(t, tree, tree.TraverseInOrder, []int{1, 2, 3})
	checkTraverseInt(t, tree, tree.TraversePostOrder, []int{1, 3, 2})

	tree.Insert(4)
	checkTraverseInt(t, tree, tree.TraversePreOrder, []int{2, 1, 3, 4})
	checkTraverseInt(t, tree, tree.TraverseInOrder, []int{1, 2, 3, 4})
	checkTraverseInt(t, tree, tree.TraversePostOrder, []int{1, 4, 3, 2})

	tree.Insert(5)
	checkTraverseInt(t, tree, tree.TraversePreOrder, []int{2, 1, 4, 3, 5})
	checkTraverseInt(t, tree, tree.TraverseInOrder, []int{1, 2, 3, 4, 5})
	checkTraverseInt(t, tree, tree.TraversePostOrder, []int{1, 3, 5, 4, 2})

	tree.Insert(6)
	checkTraverseInt(t, tree, tree.TraversePreOrder, []int{4, 2, 1, 3, 5, 6})
	checkTraverseInt(t, tree, tree.TraverseInOrder, []int{1, 2, 3, 4, 5, 6})
	checkTraverseInt(t, tree, tree.TraversePostOrder, []int{1, 3, 2, 6, 5, 4})

	tree.Insert(7)
	checkTraverseInt(t, tree, tree.TraversePreOrder, []int{4, 2, 1, 3, 6, 5, 7})
	checkTraverseInt(t, tree, tree.TraverseInOrder, []int{1, 2, 3, 4, 5, 6, 7})
	checkTraverseInt(t, tree, tree.TraversePostOrder, []int{1, 3, 2, 5, 7, 6, 4})
}

func TestInsertAVL2(t *testing.T) {
	tree := NewIntAVLTree()
	checkTraverseInt(t, tree, tree.TraversePreOrder, []int{})
	checkTraverseInt(t, tree, tree.TraverseInOrder, []int{})
	checkTraverseInt(t, tree, tree.TraversePostOrder, []int{})

	tree.Insert(7)
	checkTraverseInt(t, tree, tree.TraversePreOrder, []int{7})
	checkTraverseInt(t, tree, tree.TraverseInOrder, []int{7})
	checkTraverseInt(t, tree, tree.TraversePostOrder, []int{7})

	tree.Insert(4)
	checkTraverseInt(t, tree, tree.TraversePreOrder, []int{7, 4})
	checkTraverseInt(t, tree, tree.TraverseInOrder, []int{4, 7})
	checkTraverseInt(t, tree, tree.TraversePostOrder, []int{4, 7})

	tree.Insert(9)
	checkTraverseInt(t, tree, tree.TraversePreOrder, []int{7, 4, 9})
	checkTraverseInt(t, tree, tree.TraverseInOrder, []int{4, 7, 9})
	checkTraverseInt(t, tree, tree.TraversePostOrder, []int{4, 9, 7})

	tree.Insert(2)
	checkTraverseInt(t, tree, tree.TraversePreOrder, []int{7, 4, 2, 9})
	checkTraverseInt(t, tree, tree.TraverseInOrder, []int{2, 4, 7, 9})
	checkTraverseInt(t, tree, tree.TraversePostOrder, []int{2, 4, 9, 7})

	tree.Insert(6)
	checkTraverseInt(t, tree, tree.TraversePreOrder, []int{7, 4, 2, 6, 9})
	checkTraverseInt(t, tree, tree.TraverseInOrder, []int{2, 4, 6, 7, 9})
	checkTraverseInt(t, tree, tree.TraversePostOrder, []int{2, 6, 4, 9, 7})

	tree.Insert(1)
	checkTraverseInt(t, tree, tree.TraversePreOrder, []int{4, 2, 1, 7, 6, 9})
	checkTraverseInt(t, tree, tree.TraverseInOrder, []int{1, 2, 4, 6, 7, 9})
	checkTraverseInt(t, tree, tree.TraversePostOrder, []int{1, 2, 6, 9, 7, 4})

	tree.Insert(8)
	checkTraverseInt(t, tree, tree.TraversePreOrder, []int{4, 2, 1, 7, 6, 9, 8})
	checkTraverseInt(t, tree, tree.TraverseInOrder, []int{1, 2, 4, 6, 7, 8, 9})
	checkTraverseInt(t, tree, tree.TraversePostOrder, []int{1, 2, 6, 8, 9, 7, 4})
}

/*
RR-Rotation: Rotation root being left child
-------------------------------------------

. Insert: 14, 3, 16, 2, 6, 15,  . Now insert: 7 (9 / 11 / 13)   . Rotate with: root 3, pivot 6  .
.         17, 1, 5, 10, 18, 4,  .                               .                               .
.         8, 12                 .                               .                               .
.                               .                               .                               .
.         14                    .           14                  .                14             .
.      +-/ \-+                  .       +--/ \--+               .           +---/ \---+         .
.     3       16                .      3         16             .          6           16       .
.    / \     / \                .    +/ \+      / \             .      +--/ \--+      / \       .
.   2   6   15  17              .   2     6    15  17           .     3         10   15  17     .
.  /   / \       \              .  /    +/ \+       \           .    / \     +-/ \-+      \     .
. 1   5   10      18            . 1    5     10      18         .   2   5   8       12     18   .
.    /   / \                    .     /   +-/ \-+               .  /   /   / \     / \          .
.   4   8   12                  .    4   8       12             . 1   4  (7) [9] [11][13]       .
.                               .   / \     / \                 .                               .
.                               . (7) [9] [11][13]              .                               .
.                               .                               .                               .
.                               .  7: (0,0)                     .  7: (0,0)                     .
.  8: (0,0)                     .  8: (0,0) -> (1,0)            .  8: (0,0) -> (1,0)            .
. 10: (1,1)                     . 10: (1,1) -> (2,1)            . 10: (1,1) -> (2,1)            .
.  6: (2,2)                     .  6: (2,2) -> (2,3)            .  6: (2,2) -> (2,3) -> (3,3)   .
.  3: (2,3)                     .  3: (2,3) -> (2,4) X          .  3: (2,3) -> (2,4) -> (2,2)   .
. 14: (4,3)                     . 14: (4,3)                     . 14: (4,3)                     .
*/

func TestInsertAVLRRa(t *testing.T) {
	tree := NewIntAVLTree()
	insertMultiple(tree, 14, 3, 16, 2, 6, 15, 17, 1, 5, 10, 18, 4, 8, 12)
	checkTraverseInt(t, tree, tree.TraversePreOrder, []int{14, 3, 2, 1, 6, 5, 4, 10, 8, 12, 16, 15, 17, 18})
	checkTraverseInt(t, tree, tree.TraverseInOrder, []int{1, 2, 3, 4, 5, 6, 8, 10, 12, 14, 15, 16, 17, 18})
	checkTraverseInt(t, tree, tree.TraversePostOrder, []int{1, 2, 4, 5, 8, 12, 10, 6, 3, 15, 18, 17, 16, 14})

	tree1 := tree.Clone()
	tree1.Insert(7)
	checkTraverseInt(t, tree1, tree1.TraversePreOrder, []int{14, 6, 3, 2, 1, 5, 4, 10, 8, 7, 12, 16, 15, 17, 18})
	checkTraverseInt(t, tree1, tree1.TraverseInOrder, []int{1, 2, 3, 4, 5, 6, 7, 8, 10, 12, 14, 15, 16, 17, 18})
	checkTraverseInt(t, tree1, tree1.TraversePostOrder, []int{1, 2, 4, 5, 3, 7, 8, 12, 10, 6, 15, 18, 17, 16, 14})

	tree2 := tree.Clone()
	tree2.Insert(9)
	checkTraverseInt(t, tree2, tree2.TraversePreOrder, []int{14, 6, 3, 2, 1, 5, 4, 10, 8, 9, 12, 16, 15, 17, 18})
	checkTraverseInt(t, tree2, tree2.TraverseInOrder, []int{1, 2, 3, 4, 5, 6, 8, 9, 10, 12, 14, 15, 16, 17, 18})
	checkTraverseInt(t, tree2, tree2.TraversePostOrder, []int{1, 2, 4, 5, 3, 9, 8, 12, 10, 6, 15, 18, 17, 16, 14})

	tree3 := tree.Clone()
	tree3.Insert(11)
	checkTraverseInt(t, tree3, tree3.TraversePreOrder, []int{14, 6, 3, 2, 1, 5, 4, 10, 8, 12, 11, 16, 15, 17, 18})
	checkTraverseInt(t, tree3, tree3.TraverseInOrder, []int{1, 2, 3, 4, 5, 6, 8, 10, 11, 12, 14, 15, 16, 17, 18})
	checkTraverseInt(t, tree3, tree3.TraversePostOrder, []int{1, 2, 4, 5, 3, 8, 11, 12, 10, 6, 15, 18, 17, 16, 14})

	tree4 := tree.Clone()
	tree4.Insert(13)
	checkTraverseInt(t, tree4, tree4.TraversePreOrder, []int{14, 6, 3, 2, 1, 5, 4, 10, 8, 12, 13, 16, 15, 17, 18})
	checkTraverseInt(t, tree4, tree4.TraverseInOrder, []int{1, 2, 3, 4, 5, 6, 8, 10, 12, 13, 14, 15, 16, 17, 18})
	checkTraverseInt(t, tree4, tree4.TraversePostOrder, []int{1, 2, 4, 5, 3, 8, 13, 12, 10, 6, 15, 18, 17, 16, 14})

	tree5 := tree.Clone()
	insertMultiple(tree5, 7, 9, 11, 13)
	checkTraverseInt(t, tree5, tree5.TraversePreOrder, []int{14, 6, 3, 2, 1, 5, 4, 10, 8, 7, 9, 12, 11, 13, 16, 15, 17, 18})
	checkTraverseInt(t, tree5, tree5.TraverseInOrder, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18})
	checkTraverseInt(t, tree5, tree5.TraversePostOrder, []int{1, 2, 4, 5, 3, 7, 9, 8, 11, 13, 12, 10, 6, 15, 18, 17, 16, 14})
}

/*
RR-Rotation: Rotation root being right child
--------------------------------------------

. Insert: 5, 3, 8, 2, 4, 7, 11, . Now insert: 12 (14 / 16 / 18) . Rotate with: root 8, pivot 11 .
.         1, 6, 10, 15, 9, 13,  .                               .                               .
.         17                    .                               .                               .
.                               .                               .                               .
.         5                     .          5                    .           5                   .
.      +-/ \-+                  .      +--/ \--+                .      +---/ \---+              .
.     3       8                 .     3         8               .     3           11            .
.    / \     / \                .    / \      +/ \+             .    / \      +--/ \--+         .
.   2   4   7   11              .   2   4    7     11           .   2   4    8         15       .
.  /       /   / \              .  /        /    +/ \+          .  /        / \     +-/ \-+     .
. 1       6   10  15            . 1        6    10    15        . 1        7   10  13      17   .
.            /   / \            .              /   +-/ \-+      .         /   /   / \     / \   .
.           9   13  17          .             9   13      17    .        6   9  (12)[14][16][18].
.                               .                / \     / \    .                               .
.                               .              (12)[14][16][18] .                               .
.                               .                               .                               .
.                               . 12: (0,0)                     . 12: (0,0)                     .
. 13: (0,0)                     . 13: (0,0) -> (1,0)            . 13: (0,0) -> (1,0)            .
. 15: (1,1)                     . 15: (1,1) -> (2,1)            . 15: (1,1) -> (2,1)            .
. 11: (2,2)                     . 11: (2,2) -> (2,3)            . 11: (2,2) -> (2,3) -> (3,3)   .
.  8: (2,3)                     .  8: (2,3) -> (2,4) X          .  8: (2,3) -> (2,4) -> (2,2)   .
.  5: (3,4)                     .  5: (3,4)                     .  5: (3,4)                     .
*/

func TestInsertAVLRRb(t *testing.T) {
	tree := NewIntAVLTree()
	insertMultiple(tree, 5, 3, 8, 2, 4, 7, 11, 1, 6, 10, 15, 9, 13, 17)
	checkTraverseInt(t, tree, tree.TraversePreOrder, []int{5, 3, 2, 1, 4, 8, 7, 6, 11, 10, 9, 15, 13, 17})
	checkTraverseInt(t, tree, tree.TraverseInOrder, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 13, 15, 17})
	checkTraverseInt(t, tree, tree.TraversePostOrder, []int{1, 2, 4, 3, 6, 7, 9, 10, 13, 17, 15, 11, 8, 5})

	tree1 := tree.Clone()
	tree1.Insert(12)
	checkTraverseInt(t, tree1, tree1.TraversePreOrder, []int{5, 3, 2, 1, 4, 11, 8, 7, 6, 10, 9, 15, 13, 12, 17})
	checkTraverseInt(t, tree1, tree1.TraverseInOrder, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 15, 17})
	checkTraverseInt(t, tree1, tree1.TraversePostOrder, []int{1, 2, 4, 3, 6, 7, 9, 10, 8, 12, 13, 17, 15, 11, 5})

	tree2 := tree.Clone()
	tree2.Insert(14)
	checkTraverseInt(t, tree2, tree2.TraversePreOrder, []int{5, 3, 2, 1, 4, 11, 8, 7, 6, 10, 9, 15, 13, 14, 17})
	checkTraverseInt(t, tree2, tree2.TraverseInOrder, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 13, 14, 15, 17})
	checkTraverseInt(t, tree2, tree2.TraversePostOrder, []int{1, 2, 4, 3, 6, 7, 9, 10, 8, 14, 13, 17, 15, 11, 5})

	tree3 := tree.Clone()
	tree3.Insert(16)
	checkTraverseInt(t, tree3, tree3.TraversePreOrder, []int{5, 3, 2, 1, 4, 11, 8, 7, 6, 10, 9, 15, 13, 17, 16})
	checkTraverseInt(t, tree3, tree3.TraverseInOrder, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 13, 15, 16, 17})
	checkTraverseInt(t, tree3, tree3.TraversePostOrder, []int{1, 2, 4, 3, 6, 7, 9, 10, 8, 13, 16, 17, 15, 11, 5})

	tree4 := tree.Clone()
	tree4.Insert(18)
	checkTraverseInt(t, tree4, tree4.TraversePreOrder, []int{5, 3, 2, 1, 4, 11, 8, 7, 6, 10, 9, 15, 13, 17, 18})
	checkTraverseInt(t, tree4, tree4.TraverseInOrder, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 13, 15, 17, 18})
	checkTraverseInt(t, tree4, tree4.TraversePostOrder, []int{1, 2, 4, 3, 6, 7, 9, 10, 8, 13, 18, 17, 15, 11, 5})

	tree5 := tree.Clone()
	insertMultiple(tree5, 12, 14, 16, 18)
	checkTraverseInt(t, tree5, tree5.TraversePreOrder, []int{5, 3, 2, 1, 4, 11, 8, 7, 6, 10, 9, 15, 13, 12, 14, 17, 16, 18})
	checkTraverseInt(t, tree5, tree5.TraverseInOrder, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18})
	checkTraverseInt(t, tree5, tree5.TraversePostOrder, []int{1, 2, 4, 3, 6, 7, 9, 10, 8, 12, 14, 13, 16, 18, 17, 15, 11, 5})
}

/*
RL-Rotation: Rotation root being left child
-------------------------------------------

. Insert: 14, 3, 16, 2, 11, 15, . Now insert: 4 (6 / 8 / 10)    . Rotate with: root 3, pivot 7  .
.         17, 1, 7, 12, 18, 5,  .                               .                               .
.         9, 13                 .                               .                               .
.                               .                               .                               .
.         14                    .             14                .                 14            .
.      +-/ \-+                  .         +--/ \--+             .            +---/ \---+        .
.     3       16                .        3         16           .           7           16      .
.    / \     / \                .      +/ \+      / \           .      +---/ \---+     / \      .
.   2   11  15  17              .     2     11   15  17         .     3           11  15  17    .
.  /   / \       \              .    /    +/ \+       \         .    / \         / \       \    .
. 1   7   12      18            .   1    7     12      18       .   2   5       9   12      18  .
.    / \   \                    .     +-/ \-+   \               .  /   / \     / \   \          .
.   5   9   13                  .    5       9   13             . 1  (4) [6] [8] [10] 13        .
.                               .   / \     / \                 .                               .
.                               . (4) [6] [8] [10]              .                               .
.                               .                               .                               .
.                               .  4: (0,0)                     .  4: (0,0)                     .
.  5: (0,0)                     .  5: (0,0) -> (1,0)            .  5: (0,0) -> (1,0)            .
.  7: (1,1)                     .  7: (1,1) -> (2,1)            .  7: (1,1) -> (2,1) -> (3,3)   .
. 11: (2,2)                     . 11: (2,2) -> (3,2)            . 11: (2,2) -> (3,2) -> (1,2)   .
.  3: (2,3)                     .  3: (2,3) -> (2,4) X          .  3: (2,3) -> (2,4) -> (2,2)   .
. 14: (4,3)                     . 14: (4,3)                     . 14: (4,3)                     .
*/

/*
RL-Rotation: Rotation root being right child
--------------------------------------------

. Insert: 5, 3, 8, 2, 4, 7, 16, . Now insert: 9 (11 / 13 / 15)  . Rotate with: root 8, pivot 12 .
.         1, 6, 12, 17, 10, 14, .                               .                               .
.         18                    .                               .                               .
.                               .                               .                               .
.         5                     .          5                    .           5                   .
.      +-/ \-+                  .      +--/ \--+                .      +---/ \---+              .
.     3       8                 .     3         8               .     3           12            .
.    / \     / \                .    / \      +/ \+             .    / \     +---/ \---+        .
.   2   4   7   16              .   2   4    7     16           .   2   4   8           16      .
.  /       /   / \              .  /        /    +/ \+          .  /       / \         / \      .
. 1       6   12  17            . 1        6    12    17        . 1       7   10      14  17    .
.            / \   \            .            +-/ \-+   \        .        /   / \     / \   \    .
.           10  14  18          .           10      14  18      .       6  (9) [11][13][15] 18  .
.                               .          / \     / \          .                               .
.                               .        (9) [11][13][15]       .                               .
.                               .                               .                               .
.                               .  9: (0,0)                     .  9: (0,0)                     .
. 10: (0,0)                     . 10: (0,0) -> (1,0)            . 10: (0,0) -> (1,0)            .
. 12: (1,1)                     . 12: (1,1) -> (2,1)            . 12: (1,1) -> (2,1) -> (3,3)    .
. 16: (2,2)                     . 16: (2,2) -> (3,2)            . 16: (2,2) -> (3,2) -> (1,2)   .
.  8: (2,3)                     .  8: (2,3) -> (2,4) X          .  8: (2,3) -> (2,4) -> (2,2)   .
.  5: (3,4)                     .  5: (3,4)                     .  5: (3,4)                     .
*/
