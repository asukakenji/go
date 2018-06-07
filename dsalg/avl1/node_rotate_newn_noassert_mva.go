// - New node:
//   O [avl_newn] Return Value: newN *Node (*)
//   - [avl_ptrn] Parameter: ptrN **Node
// - Assertion:
//   O [avl_noassert] Without assertion (*)
//   - [avl_assert] With assertion
// - Assignment:
//   O [avl_mva] Multi-Valued Assignment (*)
//   - [avl_sva] Single-Valued Assignment

// +build !avl_ptrn
// +build !avl_assert
// +build !avl_sva

package avl

//   N              P
//  / \            / \
// T1  P    =>    N  T4
//    / \        / \
//  T23 T4      T1 T23
func (n *Node) rotateLeft() *Node {
	p := n.childR
	if p.balanceFactor == 0 {
		n.balanceFactor, p.balanceFactor = 1, -1
	} else {
		n.balanceFactor, p.balanceFactor = 0, 0
	}
	p.childL, n.childR = n, p.childL
	return p
}

//     N          P
//    / \        / \
//   P  T4  =>  T1  N
//  / \            / \
// T1 T23        T23 T4
func (n *Node) rotateRight() *Node {
	p := n.childL
	if p.balanceFactor == 0 {
		n.balanceFactor, p.balanceFactor = -1, 1
	} else {
		n.balanceFactor, p.balanceFactor = 0, 0
	}
	p.childR, n.childL = n, p.childR
	return p
}

//     N
//    / \            Q
//   P  T4         /   \
//  / \     =>    P     N
// T1  Q         / \   / \
//    / \       T1 T2 T3 T4
//   T2 T3
func (n *Node) rotateLeftRight() *Node {
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
	q.childR, q.childL, n.childL, p.childR = n, p, q.childR, q.childL
	return q
}

//   N
//  / \              Q
// T1  P           /   \
//    / \   =>    N     P
//   Q  T4       / \   / \
//  / \         T1 T2 T3 T4
// T2 T3
func (n *Node) rotateRightLeft() *Node {
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
	q.childL, q.childR, n.childR, p.childL = n, p, q.childL, q.childR
	return q
}
