// - New node:
//   O [avl_newn] Return Value: newN *IntTreeSetNode (*)
//   - [avl_ptrn] Parameter: ptrN **IntTreeSetNode
// - Flags:
//   O [avl_flags_bools] Flags using bools (*)
//   O [avl_flags_bools_2] Flags using bools
//   - [avl_flags_int] Flags using int
//   - [avl_flags_char] Flags using char

// +build !avl_ptrn
// +build !avl_flags_int,!avl_flags_char

package avl

// Insert inserts v into the tree.
func (t *IntTreeSet) Insert(v int) {
	var isInserted bool
	if t.root, isInserted, _ = t.root.insert(v); isInserted {
		t.len++
	}
}
