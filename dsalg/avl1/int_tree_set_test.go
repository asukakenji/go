package avl_test

import (
	"math/rand"
	"testing"

	avl "github.com/asukakenji/go/dsalg/avl1"
)

func TestTreePrint(t *testing.T) {
	numbers := []int{7, 6, 4, 8, 9, 0, 1, 3, 5, 2}
	tree := new(avl.IntTreeSet)
	for _, number := range numbers {
		tree.Add(number)
	}
	// t.Log(tree.Print("    "))
	tree.Print("    ")
}

func TestTreeContains(t *testing.T) {
	numbers := rand.Perm(255)

	tree := new(avl.IntTreeSet)
	for index1, number1 := range numbers {
		// t.Logf("Adding %d", number1)
		tree.Add(number1)
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

func TestTreeAdd0(t *testing.T) {
	numbers := []int{1, 2, 0, 9, 8, 3, 4, 7, 5, 6, -1}
	strings := []string{
		"(1 / /)",                                                              // 1
		"(1 / (2 / /))",                                                        // 2
		"(1 (0 / /) (2 / /))",                                                  // 0
		"(1 (0 / /) (2 / (9 / /)))",                                            // 9
		"(1 (0 / /) (8 (2 / /) (9 / /)))",                                      // 8
		"(2 (1 (0 / /) /) (8 (3 / /) (9 / /)))",                                // 3
		"(2 (1 (0 / /) /) (8 (3 / (4 / /)) (9 / /)))",                          // 4
		"(2 (1 (0 / /) /) (8 (4 (3 / /) (7 / /)) (9 / /)))",                    // 7
		"(2 (1 (0 / /) /) (7 (4 (3 / /) (5 / /)) (8 / (9 / /))))",              // 5
		"(4 (2 (1 (0 / /) /) (3 / /)) (7 (5 / (6 / /)) (8 / (9 / /))))",        // 6
		"(4 (2 (0 (-1 / /) (1 / /)) (3 / /)) (7 (5 / (6 / /)) (8 / (9 / /))))", // -1
	}

	tree := new(avl.IntTreeSet)
	for index, number := range numbers {
		expected := strings[index]
		tree.Add(number)
		got := tree.String()
		if got != expected {
			t.Errorf("Expected: %s; Got: %s", expected, got)
		}
	}
}

type TreeChecker struct {
	t   *testing.T
	cid int
	n   []*struct {
		value         interface{}
		balanceFactor int
	}
	index int
}

func (tc *TreeChecker) Consume(n *avl.IntTreeSetNode) bool {
	if n == nil {
		if tc.n[tc.index] != nil {
			expectedValue := tc.n[tc.index].value
			expectedBalanceFactor := tc.n[tc.index].balanceFactor
			tc.t.Errorf("Case %d: n == nil, Expected: (%v, %d)", tc.cid, expectedValue, expectedBalanceFactor)
			return false
		}
	} else {
		gotValue := n.Value()
		expectedValue := tc.n[tc.index].value
		if gotValue != expectedValue {
			tc.t.Errorf("Case %d: n.Value() = %v, Expected: %v", tc.cid, gotValue, expectedValue)
			return false
		}
		gotBalanceFactor := n.BalanceFactor()
		expectedBalanceFactor := tc.n[tc.index].balanceFactor
		if gotBalanceFactor != expectedBalanceFactor {
			tc.t.Errorf("Case %d: n.Value() = %v: n.BalanceFactor() = %d, Expected: %d", tc.cid, gotValue, gotBalanceFactor, expectedBalanceFactor)
			return false
		}
	}
	tc.index++
	return true
}

func TestTreeAdd(t *testing.T) {
	cases := []struct {
		value int
		n     []*struct {
			value         interface{}
			balanceFactor int
		}
	}{
		{1, []*struct {
			value         interface{}
			balanceFactor int
		}{{1, 0}, nil, nil}},
		{2, []*struct {
			value         interface{}
			balanceFactor int
		}{{1, 1}, nil, {2, 0}, nil, nil}},
		{0, []*struct {
			value         interface{}
			balanceFactor int
		}{{1, 0}, {0, 0}, nil, nil, {2, 0}, nil, nil}},
		{9, []*struct {
			value         interface{}
			balanceFactor int
		}{{1, 1}, {0, 0}, nil, nil, {2, 1}, nil, {9, 0}, nil, nil}},
		{8, []*struct {
			value         interface{}
			balanceFactor int
		}{{1, 1}, {0, 0}, nil, nil, {8, 0}, {2, 0}, nil, nil, {9, 0}, nil, nil}},
		{3, []*struct {
			value         interface{}
			balanceFactor int
		}{{2, 0}, {1, -1}, {0, 0}, nil, nil, nil, {8, 0}, {3, 0}, nil, nil, {9, 0}, nil, nil}},
	}

	tree := new(avl.IntTreeSet)
	for cid, c := range cases {
		tree.Add(c.value)
		tc := &TreeChecker{t, cid, c.n, 0}
		tree.TraversePreOrder(tc.Consume)
	}
}

func TestTreeAdd2(t *testing.T) {
	numbers := []int{7, 6, 4, 8, 9, 0, 1, 3, 5, 2}
	strings := []string{
		"(7 / /)",
		"(7 (6 / /) /)",
		"(6 (4 / /) (7 / /))",
		"(6 (4 / /) (7 / (8 / /)))",
		"(6 (4 / /) (8 (7 / /) (9 / /)))",
		"(6 (4 (0 / /) /) (8 (7 / /) (9 / /)))",
		"(6 (1 (0 / /) (4 / /)) (8 (7 / /) (9 / /)))",
		"(6 (1 (0 / /) (4 (3 / /) /)) (8 (7 / /) (9 / /)))",
		"(6 (1 (0 / /) (4 (3 / /) (5 / /))) (8 (7 / /) (9 / /)))",
		"(6 (3 (1 (0 / /) (2 / /)) (4 / (5 / /))) (8 (7 / /) (9 / /)))",
	}

	tree := new(avl.IntTreeSet)
	for index, number := range numbers {
		expected := strings[index]
		tree.Add(number)
		got := tree.String()
		if got != expected {
			t.Errorf("Expected: %s; Got: %s", expected, got)
		}
	}
}

func TestTreeRemove(t *testing.T) {
	t.Skip("Not implemented")
}
