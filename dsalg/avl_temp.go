package avl

import "fmt"

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type AVLNode struct {
	value         int
	leftChild     *AVLNode
	rightChild    *AVLNode
	balanceFactor int
}

//   X          Z
//  / \        / \
// T1  Z  =>  X  T4
//    / \    / \
//  T23 T4  T1 T23
func (x *AVLNode) leftRotate(ptrX **AVLNode) {
	z := x.rightChild
	x.rightChild = z.leftChild // T23
	z.leftChild = x
	if z.balanceFactor == 0 {
		// Example:
		// T1.height() = 4
		// Z.height() = 6
		// T23.height() = 5
		// T4.height() = 5
		x.balanceFactor = 1
		z.balanceFactor = -1
	} else if z.balanceFactor == 1 {
		// Case: z.balanceFactor == 1
		// Example:
		// T1.height() = 4
		// Z.height() = 6
		// T23.height() = 4
		// T4.height() = 5
		x.balanceFactor = 0
		z.balanceFactor = 0
	} else {
		panic(fmt.Errorf("z.balanceFactor == %d", z.balanceFactor))
	}
	*ptrX = z
}

//     Z          X
//    / \        / \
//   X  T4  =>  T1  Z
//  / \            / \
// T1 T23        T23 T4
func (z *AVLNode) rightRotate(ptrZ **AVLNode) {
	x := z.leftChild
	z.leftChild = x.rightChild // T23
	x.rightChild = z
	if x.balanceFactor == 0 {
		// Example:
		// X.height() = 6
		// T1.height() = 5
		// T23.height() = 5
		// T4.height() = 4
		z.balanceFactor = -1
		x.balanceFactor = 1
	} else if x.balanceFactor == -1 {
		// Case: x.balanceFactor == -1
		// Example:
		// X.height() = 6
		// T1.height() = 5
		// T23.height() = 4
		// T4.height() = 4
		z.balanceFactor = 0
		x.balanceFactor = 0
	} else {
		panic(fmt.Errorf("x.balanceFactor == %d", x.balanceFactor))
	}
	*ptrZ = x
}

func (n *AVLNode) height() int {
	if n == nil {
		return -1
	}
	leftHeight := n.leftChild.height()
	rightHeight := n.rightChild.height()
	return max(leftHeight, rightHeight) + 1
}

func (n *AVLNode) Contains(v int) bool {
	if n == nil {
		return false
	}
	if v < n.value {
		return n.leftChild.Contains(v)
	}
	if v > n.value {
		return n.rightChild.Contains(v)
	}
	return true
}

func (n *AVLNode) Add(v int, ptrN **AVLNode) bool {
	if n == nil {
		*ptrN = &AVLNode{value: v}
		return true
	}
	if v < n.value {
		if n.leftChild.Add(v, &n.leftChild) {
			n.balanceFactor--
			if n.balanceFactor > 1 {
				// Rotation
			}
		}
	}
	if v > n.value {
		if n.rightChild.Add(v, &n.rightChild) {
			n.balanceFactor++
			if n.balanceFactor < -1 {
				// Rotation
			}
		}
	}
	// v is already in the tree
	return false
}

type AVLTree struct {
	root *AVLNode
}

func (t *AVLTree) Contains(v int) bool {
	return t.root.Contains(v)
}

func (t *AVLTree) Add(v int) {
	t.root.Add(v, &t.root)
}

func (t *AVLTree) AddAll(vs []int) {
	for v := range vs {
		t.Add(v)
	}
}

func main() {
}
