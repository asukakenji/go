// Multi-Valued Assignment

// +build mva

package avl

import "fmt"

//   N              Q
//  / \            / \
// T1  Q    =>    N  T4
//    / \        / \
//  T23 T4      T1 T23
func (n *IntTreeSetNode) rotateLeft(ptrN **IntTreeSetNode) {
	q := n.childR
	if q.balanceFactor == 0 {
		// Example:
		//   n: height = 7, balanceFactor = 2
		//   q: height = 6, balanceFactor = 0
		//  t1: height = 4
		// t23: height = 5
		//  t4: height = 5
		n.balanceFactor, q.balanceFactor = 1, -1
	} else {
		// Example:
		//   n: height = 7, balanceFactor = 2
		//   q: height = 6, balanceFactor = 1
		//  t1: height = 4
		// t23: height = 4
		//  t4: height = 5
		assert(q.balanceFactor == 1, fmt.Sprintf("q.balanceFactor = %d, expected 1", q.balanceFactor))
		n.balanceFactor, q.balanceFactor = 0, 0
	}
	q.childL, *ptrN, n.childR = n, q, q.childL
}

//     N          P
//    / \        / \
//   P  T4  =>  T1  N
//  / \            / \
// T1 T23        T23 T4
func (n *IntTreeSetNode) rotateRight(ptrN **IntTreeSetNode) {
	p := n.childL
	if p.balanceFactor == 0 {
		// Example:
		//   n: height = 7, balanceFactor = -2
		//   p: height = 6, balanceFactor = 0
		//  t1: height = 5
		// t23: height = 5
		//  t4: height = 4
		n.balanceFactor, p.balanceFactor = -1, 1
	} else {
		// Example:
		//   n: height = 7, balanceFactor = -2
		//   p: height = 6, balanceFactor = -1
		//  t1: height = 5
		// t23: height = 4
		//  t4: height = 4
		assert(p.balanceFactor == -1, fmt.Sprintf("p.balanceFactor = %d, expected -1", p.balanceFactor))
		n.balanceFactor, p.balanceFactor = 0, 0
	}
	p.childR, *ptrN, n.childL = n, p, p.childR
}

//     N
//    / \            Q
//   P  T4         /   \
//  / \     =>    P     N
// T1  Q         / \   / \
//    / \       T1 T2 T3 T4
//   T2 T3
func (n *IntTreeSetNode) rotateLeftRight(ptrN **IntTreeSetNode) {
	// TODO: Write this!
}

//   N
//  / \              P
// T1  Q           /   \
//    / \   =>    N     Q
//   P  T4       / \   / \
//  / \         T1 T2 T3 T4
// T2 T3
func (n *IntTreeSetNode) rotateRightLeft(ptrN **IntTreeSetNode) {
	// TODO: Write this!
}
