package avl_test

import (
	"math/rand"
	"testing"

	avl "github.com/asukakenji/go/dsalg/avl1"
)

func TestTreePrint(t *testing.T) {
	numbers := []int{1, 2, 0, 9, 8, 3, 4, 7, 5, 6, -1}
	tree := new(avl.IntTreeSet)
	for _, number := range numbers {
		tree.Add(number)
	}
	tree.Print("   ")
}

func TestTreeContains(t *testing.T) {
	numbers := rand.Perm(255)

	tree := new(avl.IntTreeSet)
	for index1, number1 := range numbers {
		t.Logf("Adding %d", number1)
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

func TestTreeAdd(t *testing.T) {
	numbers := []int{1, 2, 0, 9, 8, 3, 4, 7, 5, 6, -1}
	strings := []string{
		"(1 / /)",
		"(1 / (2 / /))",
		"(1 (0 / /) (2 / /))",
		"(1 (0 / /) (2 / (9 / /)))",
		"(1 (0 / /) (8 (2 / /) (9 / /)))",
		"(2 (1 (0 / /) /) (8 (3 / /) (9 / /)))",
		"(2 (1 (0 / /) /) (8 (3 / (4 / /)) (9 / /)))",
		"(2 (1 (0 / /) /) (8 (4 (3 / /) (7 / /)) (9 / /)))",
		"(2 (1 (0 / /) /) (7 (4 (3 / /) (5 / /)) (8 / (9 / /))))",
		"(4 (2 (1 (0 / /) /) (3 / /)) (7 (5 / (6 / /)) (8 / (9 / /))))",
		"(4 (2 (0 (-1 / /) (1 / /)) (3 / /)) (7 (5 / (6 / /)) (8 / (9 / /))))",
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
