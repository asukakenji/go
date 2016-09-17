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
	q := NewArrayDeque(w)
	for i := 0; i < w; i++ {
		for !q.IsEmpty() && values[i] >= values[q.BackValue().(int)] {
			q.PopBack()
		}
		q.PushBack(i)
	}
	for i := w; i < n; i++ {
		maxima[i-w] = values[q.FrontValue().(int)]
		for !q.IsEmpty() && values[i] >= values[q.BackValue().(int)] {
			q.PopBack()
		}
		for !q.IsEmpty() && q.FrontValue().(int) <= i-w {
			q.PopFront()
		}
		q.PushBack(i)
	}
	maxima[n-w] = values[q.FrontValue().(int)]
	return maxima
}

// Time Complexity: O(n)
// Space Complexity: O(n)
func SlidingWindowMinimum(values []int, w int) []int {
	if w <= 0 {
		panic(fmt.Sprintf("Invalid window size: %d", w))
	}
	n := len(values)
	if w > n {
		w = n
	}
	minima := make([]int, n-w+1)
	q := NewArrayDeque(w)
	for i := 0; i < w; i++ {
		for !q.IsEmpty() && values[i] <= values[q.BackValue().(int)] {
			q.PopBack()
		}
		q.PushBack(i)
	}
	for i := w; i < n; i++ {
		minima[i-w] = values[q.FrontValue().(int)]
		for !q.IsEmpty() && values[i] <= values[q.BackValue().(int)] {
			q.PopBack()
		}
		for !q.IsEmpty() && q.FrontValue().(int) <= i-w {
			q.PopFront()
		}
		q.PushBack(i)
	}
	minima[n-w] = values[q.FrontValue().(int)]
	return minima
}

// Time Complexity: O(n)
// Space Complexity: O(n)
func SlidingWindowMaxDiff(values []int, w int) []int {
	if w <= 0 {
		panic(fmt.Sprintf("Invalid window size: %d", w))
	}
	n := len(values)
	if w > n {
		w = n
	}
	maxDiffs := make([]int, n-w+1)
	qMax := NewArrayDeque(w)
	qMin := NewArrayDeque(w)
	for i := 0; i < w; i++ {
		for !qMax.IsEmpty() && values[i] >= values[qMax.BackValue().(int)] {
			qMax.PopBack()
		}
		qMax.PushBack(i)
		for !qMin.IsEmpty() && values[i] <= values[qMin.BackValue().(int)] {
			qMin.PopBack()
		}
		qMin.PushBack(i)
	}
	for i := w; i < n; i++ {
		maxDiffs[i-w] = values[qMax.FrontValue().(int)] - values[qMin.FrontValue().(int)]
		for !qMax.IsEmpty() && values[i] >= values[qMax.BackValue().(int)] {
			qMax.PopBack()
		}
		for !qMax.IsEmpty() && qMax.FrontValue().(int) <= i-w {
			qMax.PopFront()
		}
		qMax.PushBack(i)
		for !qMin.IsEmpty() && values[i] <= values[qMin.BackValue().(int)] {
			qMin.PopBack()
		}
		for !qMin.IsEmpty() && qMin.FrontValue().(int) <= i-w {
			qMin.PopFront()
		}
		qMin.PushBack(i)
	}
	maxDiffs[n-w] = values[qMax.FrontValue().(int)] - values[qMin.FrontValue().(int)]
	return maxDiffs
}
