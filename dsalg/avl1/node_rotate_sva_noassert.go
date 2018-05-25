// Single-Valued Assignment

// +build !mva
// +build !assert

package avl

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
	panic("Not implemented")
}

//   N
//  / \              P
// T1  Q           /   \
//    / \   =>    N     Q
//   P  T4       / \   / \
//  / \         T1 T2 T3 T4
// T2 T3
func (n *IntTreeSetNode) rotateRightLeft(ptrN **IntTreeSetNode) {
	panic("Not implemented")
}
