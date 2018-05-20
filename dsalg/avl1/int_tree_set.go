package avl

// IntTreeSet implements a tree set based on AVL tree.
type IntTreeSet struct {
	root *IntTreeSetNode
	len  int
}

func (t *IntTreeSet) TraversePreOrder(consumer func(*IntTreeSetNode)) {
	t.root.TraversePreOrder(consumer)
}

func (t *IntTreeSet) TraverseInOrder(consumer func(*IntTreeSetNode)) {
	t.root.TraverseInOrder(consumer)
}

func (t *IntTreeSet) TraversePostOrder(consumer func(*IntTreeSetNode)) {
	t.root.TraversePostOrder(consumer)
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
	if t.root.Add(v, &t.root) {
		t.len++
	}
}

// Remove removes v from the set.
func (t *IntTreeSet) Remove(v int) {
	if t.root.Remove(v, &t.root) {
		t.len--
	}
}

func (t *IntTreeSet) Len() int {
	return t.len
}

func (t *IntTreeSet) Cap() int {
	return t.len
}

func (t *IntTreeSet) IsEmpty() bool {
	return t.len == 0
}

func (t *IntTreeSet) IsFull() bool {
	return false
}
