// - New node:
//   - [avl_newn] Return Value: newN *Node (*)
//   O [avl_ptrn] Parameter: ptrN **Node
// - Flags:
//   O [avl_flags_bools] Flags using bools (*)
//   O [avl_flags_bools_2] Flags using bools
//   - [avl_flags_int] Flags using int
//   - [avl_flags_char] Flags using char

// +build !avl_newn,avl_ptrn
// +build !avl_flags_int,!avl_flags_char

package avl

// Insert inserts v into t.
func (t *Tree) Insert(v interface{}) {
	if isInserted, _ := t.root.insert(v, &t.root, t.lessFunc); isInserted {
		t.len++
	}
}
