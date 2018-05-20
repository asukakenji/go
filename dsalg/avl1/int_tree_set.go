package avl

// IntTreeSet implements a tree set based on AVL tree.
type IntTreeSet struct {
	root *intTreeSetNode
}

func (t *IntTreeSet) String() string {
	return t.root.String()
}

// Print prints the set.
func (t *IntTreeSet) Print(indentString string) {
	t.root.Print(indentString, 0)
}

// Contains returns whether the set contains v.
func (t *IntTreeSet) Contains(v int) bool {
	return t.root.Contains(v)
}

// Add adds v to the set.
func (t *IntTreeSet) Add(v int) {
	t.root.Add(v, &t.root)
}

// Remove removes v from the set.
func (t *IntTreeSet) Remove(v int) {
	t.root.Remove(v, &t.root)
}
