package avl

// TraversePreOrder TODO: Write this comment!
func (t *IntTreeSet) TraversePreOrder(consumer func(int)) {
	t.root.TraversePreOrder(consumer)
}

// TraverseInOrder TODO: Write this comment!
func (t *IntTreeSet) TraverseInOrder(consumer func(int)) {
	t.root.TraverseInOrder(consumer)
}

// TraversePostOrder TODO: Write this comment!
func (t *IntTreeSet) TraversePostOrder(consumer func(int)) {
	t.root.TraversePostOrder(consumer)
}

// Fold TODO: Write this comment!
func Fold(
	traverseFunc func(func(int)),
	f func(int, interface{}) interface{},
	acc0 interface{},
) interface{} {
	acc := acc0
	traverseFunc(func(v int) {
		acc = f(v, acc)
	})
	return acc
}

// ConditionalTraversePreOrder TODO: Write this comment!
func (t *IntTreeSet) ConditionalTraversePreOrder(predicate func(int) bool) bool {
	return t.root.ConditionalTraversePreOrder(predicate)
}

// ConditionalTraverseInOrder TODO: Write this comment!
func (t *IntTreeSet) ConditionalTraverseInOrder(predicate func(int) bool) bool {
	return t.root.ConditionalTraverseInOrder(predicate)
}

// ConditionalTraversePostOrder TODO: Write this comment!
func (t *IntTreeSet) ConditionalTraversePostOrder(predicate func(int) bool) bool {
	return t.root.ConditionalTraversePostOrder(predicate)
}

// ConditionalFold TODO: Write this comment!
func ConditionalFold(
	conditionalTraverseFunc func(func(int) bool) bool,
	f func(int, interface{}) (bool, interface{}),
	acc0 interface{},
) (bool, interface{}) {
	var t bool
	acc := acc0
	conditionalTraverseFunc(func(v int) bool {
		t, acc = f(v, acc)
		return t
	})
	return t, acc
}

// TraverseNodePreOrder TODO: Write this comment!
func (t *IntTreeSet) TraverseNodePreOrder(consumer func(*IntTreeSetNode)) {
	t.root.TraverseNodePreOrder(consumer)
}

// TraverseNodeInOrder TODO: Write this comment!
func (t *IntTreeSet) TraverseNodeInOrder(consumer func(*IntTreeSetNode)) {
	t.root.TraverseNodeInOrder(consumer)
}

// TraverseNodePostOrder TODO: Write this comment!
func (t *IntTreeSet) TraverseNodePostOrder(consumer func(*IntTreeSetNode)) {
	t.root.TraverseNodePostOrder(consumer)
}

// FoldNode TODO: Write this comment!
func FoldNode(
	traverseNodeFunc func(func(*IntTreeSetNode)),
	f func(*IntTreeSetNode, interface{}) interface{},
	acc0 interface{},
) interface{} {
	acc := acc0
	traverseNodeFunc(func(n *IntTreeSetNode) {
		acc = f(n, acc)
	})
	return acc
}

// ConditionalNodeTraversePreOrder TODO: Write this comment!
func (t *IntTreeSet) ConditionalTraverseNodePreOrder(predicate func(*IntTreeSetNode) bool) bool {
	return t.root.ConditionalTraverseNodePreOrder(predicate)
}

// ConditionalTraverseNodeInOrder TODO: Write this comment!
func (t *IntTreeSet) ConditionalTraverseNodeInOrder(predicate func(*IntTreeSetNode) bool) bool {
	return t.root.ConditionalTraverseNodeInOrder(predicate)
}

// ConditionalTraverseNodePostOrder TODO: Write this comment!
func (t *IntTreeSet) ConditionalTraverseNodePostOrder(predicate func(*IntTreeSetNode) bool) bool {
	return t.root.ConditionalTraverseNodePostOrder(predicate)
}

// ConditionalFoldNode TODO: Write this comment!
func ConditionalFoldNode(
	conditionalTraverseNodeFunc func(func(*IntTreeSetNode) bool) bool,
	f func(*IntTreeSetNode, interface{}) (bool, interface{}),
	acc0 interface{},
) (bool, interface{}) {
	var t bool
	acc := acc0
	conditionalTraverseNodeFunc(func(n *IntTreeSetNode) bool {
		t, acc = f(n, acc)
		return t
	})
	return t, acc
}
