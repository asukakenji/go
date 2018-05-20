// Single-Valued Assignment

// +build !mva
// +build !assert

package avl

import "fmt"

//   N              Q
//  / \            / \
// T1  Q    =>    N  T4
//    / \        / \
//  T23 T4      T1 T23
func (n *IntTreeSetNode) rotateLeft(ptrN **IntTreeSetNode) {
	q := n.childR
	n.childR = q.childL
	q.childL = n
	if q.balanceFactor == 0 {
		// Example:
		//   n: height = 7, balanceFactor = 2
		//   q: height = 6, balanceFactor = 0
		//  t1: height = 4
		// t23: height = 5
		//  t4: height = 5
		n.balanceFactor = 1
		q.balanceFactor = -1
	} else {
		// Example:
		//   n: height = 7, balanceFactor = 2
		//   q: height = 6, balanceFactor = 1
		//  t1: height = 4
		// t23: height = 4
		//  t4: height = 5
		n.balanceFactor = 0
		q.balanceFactor = 0
	}
	*ptrN = q
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
	fmt.Printf("LR0: n: %d\n", n.balanceFactor)
	fmt.Printf("LR0: p: %d\n", p.balanceFactor)
	fmt.Printf("LR0: q: %d\n", q.balanceFactor)
	n.balanceFactor = n.childR.height() - n.childL.height()
	p.balanceFactor = p.childR.height() - p.childL.height()
	q.balanceFactor = q.childR.height() - q.childL.height()
	fmt.Printf("LR1: n: %d\n", n.balanceFactor)
	fmt.Printf("LR1: p: %d\n", p.balanceFactor)
	fmt.Printf("LR1: q: %d\n", q.balanceFactor)
	if n.balanceFactor < -1 || n.balanceFactor > 1 || p.balanceFactor < -1 || p.balanceFactor > 1 || q.balanceFactor < -1 || q.balanceFactor > 1 {
		n.Print("    ", 0)
	}
	*ptrN = q
}

//   N
//  / \              P
// T1  Q           /   \
//    / \   =>    N     Q
//   P  T4       / \   / \
//  / \         T1 T2 T3 T4
// T2 T3
func (n *IntTreeSetNode) rotateRightLeft(ptrN **IntTreeSetNode) {
	q := n.childR
	p := q.childL
	n.childR = p.childL
	q.childL = p.childR
	p.childL = n
	p.childR = q
	// if p.balanceFactor == 0 {
	// 	// Example:
	// 	//  n: height = 7, balanceFactor = 2
	// 	//  q: height = 6, balanceFactor = -1
	// 	//  p: height = 5, balanceFactor = 0
	// 	// t1: height = 4
	// 	// t2: height = 4
	// 	// t3: height = 4
	// 	// t4: height = 4
	// }
	fmt.Printf("RL0: n: %d\n", n.balanceFactor)
	fmt.Printf("RL0: q: %d\n", q.balanceFactor)
	fmt.Printf("RL0: p: %d\n", p.balanceFactor)
	n.balanceFactor = n.childR.height() - n.childL.height()
	q.balanceFactor = q.childR.height() - q.childL.height()
	p.balanceFactor = p.childR.height() - p.childL.height()
	fmt.Printf("RL1: n: %d\n", n.balanceFactor)
	fmt.Printf("RL1: q: %d\n", q.balanceFactor)
	fmt.Printf("RL1: p: %d\n", p.balanceFactor)
	if n.balanceFactor < -1 || n.balanceFactor > 1 || q.balanceFactor < -1 || q.balanceFactor > 1 || p.balanceFactor < -1 || p.balanceFactor > 1 {
		n.Print("    ", 0)
	}
	*ptrN = p
}
