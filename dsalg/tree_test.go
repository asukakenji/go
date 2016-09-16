package dsalg

import (
	"reflect"
	"testing"
)

type Tree interface {
	Len() int
}

func checkTraverse(t *testing.T, tree Tree, traverseFunc func(func(int)), expectedStack []int) {
	stack := make([]int, 0, tree.Len())
	consumer := func(x int) {
		stack = append(stack, x)
	}
	traverseFunc(consumer)
	if !reflect.DeepEqual(stack, expectedStack) {
		t.Errorf("tree.TraverseXXX(consumer) == %v; expected %v", stack, expectedStack)
		panic("")
	}
}
