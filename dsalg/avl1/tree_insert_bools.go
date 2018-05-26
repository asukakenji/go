package avl

// Insert inserts v into the tree.
func (t *IntTreeSet) Insert(v int) {
	if isInserted, _ := t.root.insert(v, &t.root); isInserted {
		t.len++
	}
}
