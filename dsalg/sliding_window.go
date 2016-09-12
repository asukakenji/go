package dsalg

import (
	"fmt"
)

// Time Complexity: O(n)
// Space Complexity: O(n)
func SlidingWindowSum(values []int, w int) []int {
	if w <= 0 {
		panic(fmt.Sprintf("Invalid window size: %d", w))
	}
	n := len(values)
	if w > n {
		w = n
	}
	sums := make([]int, n-w+1)
	for i := 0; i < w; i++ {
		sums[0] += values[i]
	}
	for i := w; i < n; i++ {
		sums[i-w+1] = sums[i-w] + (values[i] - values[i-w])
	}
	return sums
}

// Time Complexity: O(n)
// Space Complexity: O(n)
// Reference: http://articles.leetcode.com/sliding-window-maximum/
func SlidingWindowMaximum(values []int, w int) []int {
	if w <= 0 {
		panic(fmt.Sprintf("Invalid window size: %d", w))
	}
	n := len(values)
	if w > n {
		w = n
	}
	maxima := make([]int, n-w+1)
	q := NewLinkedDeque()
	for i := 0; i < w; i++ {
		for !q.IsEmpty() && values[i] >= values[q.Back().(int)] {
			q.PopBack()
		}
		q.PushBack(i)
	}
	for i := w; i < n; i++ {
		maxima[i-w] = values[q.Front().(int)]
		for !q.IsEmpty() && values[i] >= values[q.Back().(int)] {
			q.PopBack()
		}
		for !q.IsEmpty() && q.Front().(int) <= i-w {
			q.PopFront()
		}
		q.PushBack(i)
	}
	maxima[n-w] = values[q.Front().(int)]
	return maxima
}
