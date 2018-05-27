// - New node:
//   - [avl_newn] Return Value: newN *IntTreeSetNode (*)
//   O [avl_ptrn] Parameter: ptrN **IntTreeSetNode
// - Insert returns
//   O [avl_bools] 2 bools (*)
//   - [avl_flags_int] Flags using int
//   - [avl_flags_char] Flags using char

// +build avl_ptrn
// +build !avl_flags_int,!avl_flags_char

package avl

// Insert inserts v into the tree.
func (t *IntTreeSet) Insert(v int) {
	if isInserted, _ := t.root.insert(v, &t.root); isInserted {
		t.len++
	}
}
