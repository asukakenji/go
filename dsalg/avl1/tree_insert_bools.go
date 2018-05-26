package avl

// Add adds v to the set.
func (t *IntTreeSet) Add(v int) {
	if isAdded, _ := t.root.Add(v, &t.root); isAdded {
		t.len++
	}
}
