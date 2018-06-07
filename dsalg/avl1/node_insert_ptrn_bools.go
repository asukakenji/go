// - New node:
//   - [avl_newn] Return Value: newN *Node (*)
//   O [avl_ptrn] Parameter: ptrN **Node
// - Flags:
//   O [avl_flags_bools] Flags using bools (*)
//   - [avl_flags_bools_2] Flags using bools
//   - [avl_flags_int] Flags using int
//   - [avl_flags_char] Flags using char

// +build !avl_newn,avl_ptrn
// +build !avl_flags_bools_2,!avl_flags_int,!avl_flags_char

package avl

// insert inserts v into the subtree rooted at n.
func (n *Node) insert(v interface{}, ptrN **Node, lessFunc func(interface{}, interface{}) bool) (bool, bool) {
	if n == nil {
		*ptrN = &Node{value: v}
		return true, true
	}
	if lessFunc(v, n.value) {
		if isAdded, needsPropagation := n.childL.insert(v, &n.childL, lessFunc); needsPropagation {
			if n.balanceFactor > 0 {
				n.balanceFactor--
				return true, false
			} else if n.balanceFactor == 0 {
				n.balanceFactor--
				return true, true
			} else {
				// Rotate
				if lessFunc(n.childL.value, v) {
					// LR Case
					n.rotateLeftRight(ptrN)
				} else {
					// LL Case
					n.rotateRight(ptrN)
				}
				return true, false
			}
		} else {
			return isAdded, false
		}
	}
	if lessFunc(n.value, v) {
		if isAdded, needsPropagation := n.childR.insert(v, &n.childR, lessFunc); needsPropagation {
			if n.balanceFactor < 0 {
				n.balanceFactor++
				return true, false
			} else if n.balanceFactor == 0 {
				n.balanceFactor++
				return true, true
			} else {
				// Rotate
				if lessFunc(v, n.childR.value) {
					// RL Case
					n.rotateRightLeft(ptrN)
				} else {
					// RR Case
					n.rotateLeft(ptrN)
				}
				return true, false
			}
		} else {
			return isAdded, false
		}
	}
	// v already existed, nothing to be done
	// TODO: Assign again?
	return false, false
}
