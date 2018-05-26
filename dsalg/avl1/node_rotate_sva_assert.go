// Single-Valued Assignment, with Assertion

// +build sva
// +build assert

package avl

import "fmt"

//   N              P
//  / \            / \
// T1  P    =>    N  T4
//    / \        / \
//  T23 T4      T1 T23
func (n *IntTreeSetNode) rotateLeft(ptrN **IntTreeSetNode) {
	p := n.childR
	n.childR = p.childL
	p.childL = n
	if p.balanceFactor == 0 {
		// Example:
		//   n: height = 7, balanceFactor = 2
		//   p: height = 6, balanceFactor = 0
		//  t1: height = 4
		// t23: height = 5
		//  t4: height = 5
		n.balanceFactor = 1
		p.balanceFactor = -1
	} else {
		// Example:
		//   n: height = 7, balanceFactor = 2
		//   p: height = 6, balanceFactor = 1
		//  t1: height = 4
		// t23: height = 4
		//  t4: height = 5
		assert(p.balanceFactor == 1, fmt.Sprintf("p.balanceFactor = %d, expected 1", p.balanceFactor))
		n.balanceFactor = 0
		p.balanceFactor = 0
	}
	*ptrN = p
}

//     N          P
//    / \        / \
//   P  T4  =>  T1  N
//  / \            / \
// T1 T23        T23 T4
func (n *IntTreeSetNode) rotateRight(ptrN **IntTreeSetNode) {
	p := n.childL
	n.childL = p.childR
	p.childR = n
	if p.balanceFactor == 0 {
		// Example:
		//   n: height = 7, balanceFactor = -2
		//   p: height = 6, balanceFactor = 0
		//  t1: height = 5
		// t23: height = 5
		//  t4: height = 4
		n.balanceFactor = -1
		p.balanceFactor = 1
	} else {
		// Example:
		//   n: height = 7, balanceFactor = -2
		//   p: height = 6, balanceFactor = -1
		//  t1: height = 5
		// t23: height = 4
		//  t4: height = 4
		assert(p.balanceFactor == -1, fmt.Sprintf("p.balanceFactor = %d, expected -1", p.balanceFactor))
		n.balanceFactor = 0
		p.balanceFactor = 0
	}
	*ptrN = p
}

//     N
//    / \            Q
//   P  T4         /   \
//  / \     =>    P     N
// T1  Q         / \   / \
//    / \       T1 T2 T3 T4
//   T2 T3
func (n *IntTreeSetNode) rotateLeftRight(ptrN **IntTreeSetNode) {
	p := n.childL
	q := p.childR
	p.childR = q.childL
	n.childL = q.childR
	q.childL = p
	q.childR = n
	if q.balanceFactor < 0 {
		n.balanceFactor = 1
		p.balanceFactor = 0
	} else if q.balanceFactor > 0 {
		n.balanceFactor = 0
		p.balanceFactor = -1
	} else {
		n.balanceFactor = 0
		p.balanceFactor = 0
	}
	q.balanceFactor = 0
	*ptrN = q
}

//   N
//  / \              Q
// T1  P           /   \
//    / \   =>    N     P
//   Q  T4       / \   / \
//  / \         T1 T2 T3 T4
// T2 T3
func (n *IntTreeSetNode) rotateRightLeft(ptrN **IntTreeSetNode) {
	p := n.childR
	q := p.childL
	p.childL = q.childR
	n.childR = q.childL
	q.childR = p
	q.childL = n
	if q.balanceFactor < 0 {
		n.balanceFactor = 0
		p.balanceFactor = 1
	} else if q.balanceFactor > 0 {
		n.balanceFactor = -1
		p.balanceFactor = 0
	} else {
		n.balanceFactor = 0
		p.balanceFactor = 0
	}
	q.balanceFactor = 0
	*ptrN = q
}
