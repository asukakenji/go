// - New node:
//   - [avl_newn] Return Value: newN *Node (*)
//   O [avl_ptrn] Parameter: ptrN **Node
// - Assertion:
//   O [avl_noassert] Without assertion (*)
//   - [avl_assert] With assertion
// - Assignment:
//   O [avl_mva] Multi-Valued Assignment (*)
//   - [avl_sva] Single-Valued Assignment

// +build !avl_newn,avl_ptrn
// +build !avl_assert
// +build !avl_sva

package avl

//   N              P
//  / \            / \
// T1  P    =>    N  T4
//    / \        / \
//  T23 T4      T1 T23
func (n *Node) rotateLeft(ptrN **Node) {
	p := n.childR
	if p.balanceFactor == 0 {
		n.balanceFactor, p.balanceFactor = 1, -1
	} else {
		n.balanceFactor, p.balanceFactor = 0, 0
	}
	p.childL, *ptrN, n.childR = n, p, p.childL
}

//     N          P
//    / \        / \
//   P  T4  =>  T1  N
//  / \            / \
// T1 T23        T23 T4
func (n *Node) rotateRight(ptrN **Node) {
	p := n.childL
	if p.balanceFactor == 0 {
		n.balanceFactor, p.balanceFactor = -1, 1
	} else {
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
func (n *Node) rotateLeftRight(ptrN **Node) {
	p := n.childL
	q := p.childR
	if q.balanceFactor < 0 {
		n.balanceFactor, p.balanceFactor = 1, 0
	} else if q.balanceFactor > 0 {
		n.balanceFactor, p.balanceFactor = 0, -1
	} else {
		n.balanceFactor, p.balanceFactor = 0, 0
	}
	q.balanceFactor = 0
	q.childR, q.childL, *ptrN, n.childL, p.childR = n, p, q, q.childR, q.childL
}

//   N
//  / \              Q
// T1  P           /   \
//    / \   =>    N     P
//   Q  T4       / \   / \
//  / \         T1 T2 T3 T4
// T2 T3
func (n *Node) rotateRightLeft(ptrN **Node) {
	p := n.childR
	q := p.childL
	if q.balanceFactor < 0 {
		n.balanceFactor, p.balanceFactor = 0, 1
	} else if q.balanceFactor > 0 {
		n.balanceFactor, p.balanceFactor = -1, 0
	} else {
		n.balanceFactor, p.balanceFactor = 0, 0
	}
	q.balanceFactor = 0
	q.childL, q.childR, *ptrN, n.childR, p.childL = n, p, q, q.childL, q.childR
}
