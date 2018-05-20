package avl

import (
	"reflect"
	"testing"
)

type Tree interface {
	Len() int
}

func checkTraverse(t *testing.T, tree Tree, traverseFunc func(func(interface{})), expectedStack []interface{}) {
	stack := make([]interface{}, 0, tree.Len())
	consumer := func(v interface{}) {
		stack = append(stack, v)
	}
	traverseFunc(consumer)
	if !reflect.DeepEqual(stack, expectedStack) {
		t.Errorf("tree.TraverseXXX(consumer) == %v; expected %v", stack, expectedStack)
		panic("")
	}
}

func checkTraverseInt(t *testing.T, tree Tree, traverseFunc func(func(interface{})), expectedStack []int) {
	stack := make([]int, 0, tree.Len())
	consumer := func(v interface{}) {
		stack = append(stack, v.(int))
	}
	traverseFunc(consumer)
	if !reflect.DeepEqual(stack, expectedStack) {
		t.Errorf("tree.TraverseXXX(consumer) == %v; expected %v", stack, expectedStack)
		panic("")
	}
}
