// - New node:
//   - [avl_newn] Return Value: newN *IntTreeSetNode (*)
//   O [avl_ptrn] Parameter: ptrN **IntTreeSetNode
// - Flags:
//   - [avl_flags_bools] Flags using bools (*)
//   O [avl_flags_bools_2] Flags using bools
//   - [avl_flags_int] Flags using int
//   - [avl_flags_char] Flags using char

// +build !avl_newn,avl_ptrn
// +build !avl_flags_bools,avl_flags_bools_2,!avl_flags_int,!avl_flags_char

package avl

// insert inserts v into the subtree rooted at n.
func (n *IntTreeSetNode) insert(v int, ptrN **IntTreeSetNode) (bool, bool) {
	if n == nil {
		*ptrN = &IntTreeSetNode{value: v}
		return true, true
	}
	if v < n.value {
		if isAdded, needsPropagation := n.childL.insert(v, &n.childL); needsPropagation {
			if n.balanceFactor >= 0 {
				n.balanceFactor--
				return true, n.balanceFactor != 0
			} else {
				// Rotate
				if v > n.childL.value {
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
	if v > n.value {
		if isAdded, needsPropagation := n.childR.insert(v, &n.childR); needsPropagation {
			if n.balanceFactor <= 0 {
				n.balanceFactor++
				return true, n.balanceFactor != 0
			} else {
				// Rotate
				if v < n.childR.value {
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
	return false, false
}
