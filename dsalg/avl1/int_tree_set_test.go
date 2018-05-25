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

type ValueAndBalanceFactor struct {
	value         interface{}
	balanceFactor int
}

type TreeChecker struct {
	t     *testing.T
	cid   int
	vabf  []*ValueAndBalanceFactor
	index int
}

func NewTreeChecker(t *testing.T, cid int, vabf []*ValueAndBalanceFactor) *TreeChecker {
	return &TreeChecker{t, cid, vabf, 0}
}

func (tc *TreeChecker) Consume(n *avl.IntTreeSetNode) bool {
	var (
		expectedValue         interface{}
		expectedBalanceFactor int
		gotValue              interface{}
		gotBalanceFactor      int
	)
	expectedN := tc.vabf[tc.index]
	if expectedN != nil {
		expectedValue = expectedN.value
		expectedBalanceFactor = expectedN.balanceFactor
	}
	if n != nil {
		gotValue = n.Value()
		gotBalanceFactor = n.BalanceFactor()
	}
	if expectedN == nil && n != nil {
		tc.t.Errorf("Case %d: n = (%v, %d), Expected: nil", tc.cid, gotValue, gotBalanceFactor)
		return false
	}
	if expectedN != nil && n == nil {
		tc.t.Errorf("Case %d: n == nil, Expected: (%v, %d)", tc.cid, expectedValue, expectedBalanceFactor)
		return false
	}
	if gotValue != expectedValue || gotBalanceFactor != expectedBalanceFactor {
		tc.t.Errorf("Case %d: n = (%v, %d), Expected: (%v, %d)", tc.cid, gotValue, gotBalanceFactor, expectedValue, expectedBalanceFactor)
		return false
	}
	tc.index++
	return true
}

func TestTreeAdd(t *testing.T) {
	cases := []struct {
		value int
		vabf  []*ValueAndBalanceFactor
	}{
		{1, []*ValueAndBalanceFactor{{1, 0}, nil, nil}},
		{2, []*ValueAndBalanceFactor{{1, 1}, nil, {2, 0}, nil, nil}},
		{0, []*ValueAndBalanceFactor{{1, 0}, {0, 0}, nil, nil, {2, 0}, nil, nil}},
		{9, []*ValueAndBalanceFactor{{1, 1}, {0, 0}, nil, nil, {2, 1}, nil, {9, 0}, nil, nil}},
		{8, []*ValueAndBalanceFactor{{1, 1}, {0, 0}, nil, nil, {8, 0}, {2, 0}, nil, nil, {9, 0}, nil, nil}},
		{3, []*ValueAndBalanceFactor{{2, 0}, {1, -1}, {0, 0}, nil, nil, nil, {8, 0}, {3, 0}, nil, nil, {9, 0}, nil, nil}},
	}

	tree := new(avl.IntTreeSet)
	for cid, c := range cases {
		tree.Add(c.value)
		tree.TraversePreOrder(NewTreeChecker(t, cid, c.vabf).Consume)
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
