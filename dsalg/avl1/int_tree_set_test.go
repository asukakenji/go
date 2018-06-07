package avl_test

import (
	"math/rand"
	"testing"

	avl "github.com/asukakenji/go/dsalg/avl1"
)

func TestTreeLenCapIsEmptyIsFull(t *testing.T) {
	values := rand.Perm(255)
	length := 0

	tree := new(avl.IntTreeSet)

	check := func() {
		expectedLen := length
		gotLen := tree.Len()
		if gotLen != expectedLen {
			t.Errorf("tree.Len() = %d, Expected: %d", gotLen, expectedLen)
		}

		expectedCap := length
		gotCap := tree.Cap()
		if gotCap != expectedCap {
			t.Errorf("tree.Cap() = %d, Expected: %d", gotCap, expectedCap)
		}

		expectedIsEmpty := (length == 0)
		gotIsEmpty := tree.IsEmpty()
		if gotIsEmpty != expectedIsEmpty {
			t.Errorf("tree.IsEmpty() = %t, Expected: %t", gotIsEmpty, expectedIsEmpty)
		}

		if tree.IsFull() {
			t.Errorf("tree.IsFull() = true, Expected: false")
		}
	}

	check()
	for _, value := range values {
		tree.Insert(value)
		length++
		check()
	}
}

func TestString(t *testing.T) {
	cases := []struct {
		value int
		s     string
	}{
		{1, "(1 / /)"},
		{2, "(1 / (2 / /))"},
		{0, "(1 (0 / /) (2 / /))"},
		{9, "(1 (0 / /) (2 / (9 / /)))"},
		{8, "(1 (0 / /) (8 (2 / /) (9 / /)))"},
		{3, "(2 (1 (0 / /) /) (8 (3 / /) (9 / /)))"},
		{4, "(2 (1 (0 / /) /) (8 (3 / (4 / /)) (9 / /)))"},
		{7, "(2 (1 (0 / /) /) (8 (4 (3 / /) (7 / /)) (9 / /)))"},
		{5, "(2 (1 (0 / /) /) (7 (4 (3 / /) (5 / /)) (8 / (9 / /))))"},
		{6, "(4 (2 (1 (0 / /) /) (3 / /)) (7 (5 / (6 / /)) (8 / (9 / /))))"},
		{-1, "(4 (2 (0 (-1 / /) (1 / /)) (3 / /)) (7 (5 / (6 / /)) (8 / (9 / /))))"},
	}

	tree := new(avl.IntTreeSet)
	for _, c := range cases {
		tree.Insert(c.value)
	}
}

func TestTreePrint(t *testing.T) {
	numbers := []int{7, 6, 4, 8, 9, 0, 1, 3, 5, 2}
	tree := new(avl.IntTreeSet)
	for _, number := range numbers {
		tree.Insert(number)
	}
	// t.Log(tree.Print("    "))
	tree.Print("    ")
}

func TestTreeContains(t *testing.T) {
	numbers := rand.Perm(255)

	tree := new(avl.IntTreeSet)
	for index1, number1 := range numbers {
		// t.Logf("Inserting %d", number1)
		tree.Insert(number1)
		for index2, number2 := range numbers {
			contains := tree.Contains(number2)
			if index2 <= index1 {
				if !contains {
					t.Errorf("tree does not contain %d", number2)
				}
			} else {
				if contains {
					t.Errorf("tree should not contain %d", number2)
				}
			}
		}
	}
}

type ValueAndBalanceFactor struct {
	value         interface{}
	balanceFactor int
}

type CheckTreeParams struct {
	t     *testing.T
	cid   int
	vabf  []*ValueAndBalanceFactor
	index int
}

func CheckTree(n *avl.IntTreeSetNode, acc0 interface{}) (bool, interface{}) {
	params := acc0.(*CheckTreeParams)
	var (
		expectedValue         interface{}
		expectedBalanceFactor int
		gotValue              interface{}
		gotBalanceFactor      int
	)
	expectedN := params.vabf[params.index]
	if expectedN != nil {
		expectedValue = expectedN.value
		expectedBalanceFactor = expectedN.balanceFactor
	}
	if n != nil {
		gotValue = n.Value()
		gotBalanceFactor = n.BalanceFactor()
	}
	if expectedN == nil && n != nil {
		params.t.Errorf("Case %d: Index: %d: n = (%v, %d), Expected: nil", params.cid, params.index, gotValue, gotBalanceFactor)
		return false, nil
	}
	if expectedN != nil && n == nil {
		params.t.Errorf("Case %d: Index: %d: n == nil, Expected: (%v, %d)", params.cid, params.index, expectedValue, expectedBalanceFactor)
		return false, nil
	}
	if gotValue != expectedValue || gotBalanceFactor != expectedBalanceFactor {
		params.t.Errorf("Case %d: Index: %d: n = (%v, %d), Expected: (%v, %d)", params.cid, params.index, gotValue, gotBalanceFactor, expectedValue, expectedBalanceFactor)
		return false, nil
	}
	params.index++
	return true, params
}

func TestTreeInsert_1(t *testing.T) {
	cases := []struct {
		value int
		s     string
		vabf  []*ValueAndBalanceFactor
	}{
		{
			1,
			"(1 / /)",
			[]*ValueAndBalanceFactor{{1, 0}, nil, nil},
		},
		{
			2,
			"(1 / (2 / /))",
			[]*ValueAndBalanceFactor{{1, 1}, nil, {2, 0}, nil, nil},
		},
		{
			0,
			"(1 (0 / /) (2 / /))",
			[]*ValueAndBalanceFactor{{1, 0}, {0, 0}, nil, nil, {2, 0}, nil, nil},
		},
		{
			9,
			"(1 (0 / /) (2 / (9 / /)))",
			[]*ValueAndBalanceFactor{{1, 1}, {0, 0}, nil, nil, {2, 1}, nil, {9, 0}, nil, nil},
		},
		{
			8,
			"(1 (0 / /) (8 (2 / /) (9 / /)))",
			[]*ValueAndBalanceFactor{{1, 1}, {0, 0}, nil, nil, {8, 0}, {2, 0}, nil, nil, {9, 0}, nil, nil},
		},
		{
			3,
			"(2 (1 (0 / /) /) (8 (3 / /) (9 / /)))",
			[]*ValueAndBalanceFactor{{2, 0}, {1, -1}, {0, 0}, nil, nil, nil, {8, 0}, {3, 0}, nil, nil, {9, 0}, nil, nil},
		},
		{
			4,
			"(2 (1 (0 / /) /) (8 (3 / (4 / /)) (9 / /)))",
			[]*ValueAndBalanceFactor{{2, 1}, {1, -1}, {0, 0}, nil, nil, nil, {8, -1}, {3, 1}, nil, {4, 0}, nil, nil, {9, 0}, nil, nil},
		},
		{
			7,
			"(2 (1 (0 / /) /) (8 (4 (3 / /) (7 / /)) (9 / /)))",
			[]*ValueAndBalanceFactor{{2, 1}, {1, -1}, {0, 0}, nil, nil, nil, {8, -1}, {4, 0}, {3, 0}, nil, nil, {7, 0}, nil, nil, {9, 0}, nil, nil},
		},
		{
			5,
			"(2 (1 (0 / /) /) (7 (4 (3 / /) (5 / /)) (8 / (9 / /))))",
			[]*ValueAndBalanceFactor{{2, 1}, {1, -1}, {0, 0}, nil, nil, nil, {7, 0}, {4, 0}, {3, 0}, nil, nil, {5, 0}, nil, nil, {8, 1}, nil, {9, 0}, nil, nil},
		},
		{
			6,
			"(4 (2 (1 (0 / /) /) (3 / /)) (7 (5 / (6 / /)) (8 / (9 / /))))",
			[]*ValueAndBalanceFactor{{4, 0}, {2, -1}, {1, -1}, {0, 0}, nil, nil, nil, {3, 0}, nil, nil, {7, 0}, {5, 1}, nil, {6, 0}, nil, nil, {8, 1}, nil, {9, 0}, nil, nil},
		},
		{
			-1,
			"(4 (2 (0 (-1 / /) (1 / /)) (3 / /)) (7 (5 / (6 / /)) (8 / (9 / /))))",
			[]*ValueAndBalanceFactor{{4, 0}, {2, -1}, {0, 0}, {-1, 0}, nil, nil, {1, 0}, nil, nil, {3, 0}, nil, nil, {7, 0}, {5, 1}, nil, {6, 0}, nil, nil, {8, 1}, nil, {9, 0}, nil, nil},
		},
	}

	tree := new(avl.IntTreeSet)
	for cid, c := range cases {
		tree.Insert(c.value)
		expectedString := c.s
		gotString := tree.String()
		if gotString != expectedString {
			t.Errorf("Case %d: tree.String() = %s, Expected: %s", cid, gotString, expectedString)
		}
		avl.ConditionalFoldNode(tree.ConditionalVisitNodePreOrder, CheckTree, &CheckTreeParams{t, cid, c.vabf, 0})
	}
}

func TestTreeInsert_2(t *testing.T) {
	cases := []struct {
		value int
		s     string
		// vabf  []*ValueAndBalanceFactor
	}{
		{7, "(7 / /)"},
		{6, "(7 (6 / /) /)"},
		{4, "(6 (4 / /) (7 / /))"},
		{8, "(6 (4 / /) (7 / (8 / /)))"},
		{9, "(6 (4 / /) (8 (7 / /) (9 / /)))"},
		{0, "(6 (4 (0 / /) /) (8 (7 / /) (9 / /)))"},
		{1, "(6 (1 (0 / /) (4 / /)) (8 (7 / /) (9 / /)))"},
		{3, "(6 (1 (0 / /) (4 (3 / /) /)) (8 (7 / /) (9 / /)))"},
		{5, "(6 (1 (0 / /) (4 (3 / /) (5 / /))) (8 (7 / /) (9 / /)))"},
		{2, "(6 (3 (1 (0 / /) (2 / /)) (4 / (5 / /))) (8 (7 / /) (9 / /)))"},
	}

	tree := new(avl.IntTreeSet)
	for cid, c := range cases {
		tree.Insert(c.value)
		expectedString := c.s
		gotString := tree.String()
		if gotString != expectedString {
			t.Errorf("Case %d: tree.String() = %s, Expected: %s", cid, gotString, expectedString)
		}
		// tree.VisitPreOrder(CheckTree, &CheckTreeParams{t, cid, c.vabf, 0})
	}
}

func TestTreeRemove(t *testing.T) {
	t.Skip("Not implemented")
}
