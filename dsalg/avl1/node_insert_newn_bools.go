// - New node:
//   O [avl_newn] Return Value: newN *IntTreeSetNode (*)
//   - [avl_ptrn] Parameter: ptrN **IntTreeSetNode
// - Insert returns
//   O [avl_bools] 2 bools (*)
//   - [avl_flags_int] Flags using int
//   - [avl_flags_char] Flags using char

// +build !avl_ptrn
// +build !avl_flags_int,!avl_flags_char

package avl

// insert inserts v into the subtree rooted at n.
func (n *IntTreeSetNode) insert(v int) (*IntTreeSetNode, bool, bool) {
	if n == nil {
		return &IntTreeSetNode{value: v}, true, true
	}
	var isAdded bool
	var needsPropagation bool
	if v < n.value {
		if n.childL, isAdded, needsPropagation = n.childL.insert(v); needsPropagation {
			if n.balanceFactor > 0 {
				n.balanceFactor--
				return n, true, false
			} else if n.balanceFactor == 0 {
				n.balanceFactor--
				return n, true, true
			} else {
				// Rotate
				if v > n.childL.value {
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
	if v > n.value {
		if n.childR, isAdded, needsPropagation = n.childR.insert(v); needsPropagation {
			if n.balanceFactor < 0 {
				n.balanceFactor++
				return n, true, false
			} else if n.balanceFactor == 0 {
				n.balanceFactor++
				return n, true, true
			} else {
				// Rotate
				if v < n.childR.value {
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
	return n, false, false
}
