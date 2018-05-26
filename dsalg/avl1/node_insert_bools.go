package avl

// Add adds v to the subtree rooted at n.
func (n *IntTreeSetNode) Add(v int, ptrN **IntTreeSetNode) (bool, bool) {
	if n == nil {
		*ptrN = &IntTreeSetNode{value: v}
		return true, true
	}
	if v < n.value {
		if isAdded, needsPropagation := n.childL.Add(v, &n.childL); needsPropagation {
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
		if isAdded, needsPropagation := n.childR.Add(v, &n.childR); needsPropagation {
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
