// - New node:
//   O [avl_newn] Return Value: newN *Node (*)
//   - [avl_ptrn] Parameter: ptrN **Node
// - Flags:
//   - [avl_flags_bools] Flags using bools (*)
//   O [avl_flags_bools_2] Flags using bools
//   - [avl_flags_int] Flags using int
//   - [avl_flags_char] Flags using char

// +build !avl_ptrn
// +build !avl_flags_bools,avl_flags_bools_2,!avl_flags_int,!avl_flags_char

package avl

// insert inserts v into the subtree rooted at n.
func (n *Node) insert(v interface{}, lessFunc func(interface{}, interface{}) bool) (*Node, bool, bool) {
	if n == nil {
		return &Node{value: v}, true, true
	}
	var isAdded bool
	var needsPropagation bool
	if lessFunc(v, n.value) {
		if n.childL, isAdded, needsPropagation = n.childL.insert(v, lessFunc); needsPropagation {
			if n.balanceFactor >= 0 {
				n.balanceFactor--
				return n, true, n.balanceFactor != 0
			} else {
				// Rotate
				if lessFunc(n.childL.value, v) {
					// LR Case
					n = n.rotateLeftRight()
				} else {
					// LL Case
					n = n.rotateRight()
				}
				return n, true, false
			}
		} else {
			return n, isAdded, false
		}
	}
	if lessFunc(n.value, v) {
		if n.childR, isAdded, needsPropagation = n.childR.insert(v, lessFunc); needsPropagation {
			if n.balanceFactor <= 0 {
				n.balanceFactor++
				return n, true, n.balanceFactor != 0
			} else {
				// Rotate
				if lessFunc(v, n.childR.value) {
					// RL Case
					n = n.rotateRightLeft()
				} else {
					// RR Case
					n = n.rotateLeft()
				}
				return n, true, false
			}
		} else {
			return n, isAdded, false
		}
	}
	// v already existed, nothing to be done
	// TODO: Assign again?
	return n, false, false
}
