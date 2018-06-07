// - New node:
//   O [avl_newn] Return Value: newN *IntTreeSetNode (*)
//   - [avl_ptrn] Parameter: ptrN **IntTreeSetNode
// - Flags:
//   - [avl_flags_bools] Flags using bools (*)
//   - [avl_flags_bools_2] Flags using bools
//   O [avl_flags_int] Flags using int
//   - [avl_flags_char] Flags using char

// +build !avl_ptrn
// +build !avl_flags_bools,!avl_flags_bools_2,avl_flags_int,!avl_flags_char

package avl

// insert inserts v into the subtree rooted at n.
func (n *IntTreeSetNode) insert(v int) (*IntTreeSetNode, Flags) {
	if n == nil {
		return &IntTreeSetNode{value: v}, FLAGS_IS_INSERTED | FLAGS_NEEDS_PROPAGATION
	}
	var flags Flags
	if v < n.value {
		if n.childL, flags = n.childL.insert(v); flags&FLAGS_NEEDS_PROPAGATION != FLAGS_NONE {
			if n.balanceFactor > 0 {
				n.balanceFactor--
				return n, FLAGS_IS_INSERTED
			} else if n.balanceFactor == 0 {
				n.balanceFactor--
				return n, FLAGS_IS_INSERTED | FLAGS_NEEDS_PROPAGATION
			} else {
				// Rotate
				if v > n.childL.value {
					// LR Case
					n = n.rotateLeftRight()
				} else {
					// LL Case
					n = n.rotateRight()
				}
				return n, FLAGS_IS_INSERTED
			}
		} else {
			return n, flags & FLAGS_IS_INSERTED
		}
	}
	if v > n.value {
		if n.childR, flags = n.childR.insert(v); flags&FLAGS_NEEDS_PROPAGATION != FLAGS_NONE {
			if n.balanceFactor < 0 {
				n.balanceFactor++
				return n, FLAGS_IS_INSERTED
			} else if n.balanceFactor == 0 {
				n.balanceFactor++
				return n, FLAGS_IS_INSERTED | FLAGS_NEEDS_PROPAGATION
			} else {
				// Rotate
				if v < n.childR.value {
					// RL Case
					n = n.rotateRightLeft()
				} else {
					// RR Case
					n = n.rotateLeft()
				}
				return n, FLAGS_IS_INSERTED
			}
		} else {
			return n, flags & FLAGS_IS_INSERTED
		}
	}
	// v already existed, nothing to be done
	return n, FLAGS_NONE
}
