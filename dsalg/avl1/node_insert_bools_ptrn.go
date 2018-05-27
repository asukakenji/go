// - Insert returns
//   O [avl_bools] 2 bools (*)
//   - [avl_flags_int] Flags using int
//   - [avl_flags_char] Flags using char
// - New node:
//   - [avl_newn] Return Value: newN *IntTreeSetNode (*)
//   O [avl_ptrn] Parameter: ptrN **IntTreeSetNode

// +build !avl_flags_int,!avl_flags_char
// +build avl_ptrn

package avl

// insert inserts v into the subtree rooted at n.
func (n *IntTreeSetNode) insert(v int, ptrN **IntTreeSetNode) (bool, bool) {
	if n == nil {
		*ptrN = &IntTreeSetNode{value: v}
		return true, true
	}
	if v < n.value {
		if isAdded, needsPropagation := n.childL.insert(v, &n.childL); needsPropagation {
			if n.balanceFactor > 0 {
				n.balanceFactor--
				return true, false
			} else if n.balanceFactor == 0 {
				n.balanceFactor--
				return true, true
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
			if n.balanceFactor < 0 {
				n.balanceFactor++
				return true, false
			} else if n.balanceFactor == 0 {
				n.balanceFactor++
				return true, true
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
