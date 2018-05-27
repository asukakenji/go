// - Insert returns
//   O [avl_bools] 2 bools (*)
//   - [avl_flags_int] Flags using int
//   - [avl_flags_char] Flags using char
// - New node:
//   O [avl_newn] Return Value: newN *IntTreeSetNode (*)
//   - [avl_ptrn] Parameter: ptrN **IntTreeSetNode

// +build !avl_flags_int,!avl_flags_char
// +build !avl_ptrn

package avl

// Insert inserts v into the tree.
func (t *IntTreeSet) Insert(v int) {
	var isInserted bool
	if t.root, isInserted, _ = t.root.insert(v); isInserted {
		t.len++
	}
}
