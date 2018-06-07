package avl

// Delete deletes v from t.
func (t *Tree) Delete(v interface{}) {
	if t.root.delete(v, &t.root) {
		t.len--
	}
}
