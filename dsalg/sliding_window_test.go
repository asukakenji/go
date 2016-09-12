package dsalg

import (
	"reflect"
	"testing"
)

func TestSlidingWindowSum(t *testing.T) {
	cases := []struct {
		values     []int
		windowSize int
		expected   []int
	}{
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 1, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 2, []int{1, 3, 5, 7, 9, 11, 13, 15, 17}},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 3, []int{3, 6, 9, 12, 15, 18, 21, 24}},
	}
	for _, c := range cases {
		got := SlidingWindowSum(c.values, c.windowSize)
		if !reflect.DeepEqual(got, c.expected) {
			t.Errorf("SlidingWindowSum(%v, %d) == %v; expected %v", c.values, c.windowSize, got, c.expected)
		}
	}
}

func TestSlidingWindowMaximum(t *testing.T) {
	cases := []struct {
		values     []int
		windowSize int
		expected   []int
	}{
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 1, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 2, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 3, []int{2, 3, 4, 5, 6, 7, 8, 9}},
		{[]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, 1, []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}},
		{[]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, 2, []int{9, 8, 7, 6, 5, 4, 3, 2, 1}},
		{[]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, 3, []int{9, 8, 7, 6, 5, 4, 3, 2}},
		{[]int{0, 9, 1, 8, 2, 7, 3, 6, 4, 5}, 1, []int{0, 9, 1, 8, 2, 7, 3, 6, 4, 5}},
		{[]int{0, 9, 1, 8, 2, 7, 3, 6, 4, 5}, 2, []int{9, 9, 8, 8, 7, 7, 6, 6, 5}},
		{[]int{0, 9, 1, 8, 2, 7, 3, 6, 4, 5}, 3, []int{9, 9, 8, 8, 7, 7, 6, 6}},
		{[]int{9, 0, 8, 1, 7, 2, 6, 3, 5, 4}, 1, []int{9, 0, 8, 1, 7, 2, 6, 3, 5, 4}},
		{[]int{9, 0, 8, 1, 7, 2, 6, 3, 5, 4}, 2, []int{9, 8, 8, 7, 7, 6, 6, 5, 5}},
		{[]int{9, 0, 8, 1, 7, 2, 6, 3, 5, 4}, 3, []int{9, 8, 8, 7, 7, 6, 6, 5}},
		{[]int{0, 2, 4, 6, 8, 1, 3, 5, 7, 9}, 3, []int{4, 6, 8, 8, 8, 5, 7, 9}},
		{[]int{9, 7, 5, 3, 1, 8, 6, 4, 2, 0}, 3, []int{9, 7, 5, 8, 8, 8, 6, 4}},
	}
	for _, c := range cases {
		got := SlidingWindowMaximum(c.values, c.windowSize)
		if !reflect.DeepEqual(got, c.expected) {
			t.Errorf("SlidingWindowMaximum(%v, %d) == %v; expected %v", c.values, c.windowSize, got, c.expected)
		}
	}
}
