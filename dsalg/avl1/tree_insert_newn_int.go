// - New node:
//   O [avl_newn] Return Value: newN *Node (*)
//   - [avl_ptrn] Parameter: ptrN **Node
// - Flags:
//   - [avl_flags_bools] Flags using bools (*)
//   - [avl_flags_bools_2] Flags using bools
//   O [avl_flags_int] Flags using int
//   - [avl_flags_char] Flags using char

// +build !avl_ptrn
// +build !avl_flags_bools,!avl_flags_bools_2,avl_flags_int,!avl_flags_char

package avl

type Flags int

const (
	FLAGS_NONE              Flags = 0x0
	FLAGS_NEEDS_PROPAGATION Flags = 0x1
	FLAGS_IS_INSERTED       Flags = 0x10
)

// Insert inserts v into t.
func (t *Tree) Insert(v interface{}) {
	var flags Flags
	if t.root, flags = t.root.insert(v, t.lessFunc); flags&FLAGS_IS_INSERTED != FLAGS_NONE {
		t.len++
	}
}
