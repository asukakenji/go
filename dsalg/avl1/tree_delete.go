package avl

// Delete deletes v from the tree.
func (t *IntTreeSet) Delete(v int) {
	if t.root.delete(v, &t.root) {
		t.len--
	}
}
