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
	q := NewIntLinkedDeque()
	for i := 0; i < w; i++ {
		for !q.IsEmpty() && values[i] >= values[q.Back()] {
			q.PopBack()
		}
		q.PushBack(i)
	}
	for i := w; i < n; i++ {
		maxima[i-w] = values[q.Front()]
		for !q.IsEmpty() && values[i] >= values[q.Back()] {
			q.PopBack()
		}
		for !q.IsEmpty() && q.Front() <= i-w {
			q.PopFront()
		}
		q.PushBack(i)
	}
	maxima[n-w] = values[q.Front()]
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
	q := NewIntLinkedDeque()
	for i := 0; i < w; i++ {
		for !q.IsEmpty() && values[i] <= values[q.Back()] {
			q.PopBack()
		}
		q.PushBack(i)
	}
	for i := w; i < n; i++ {
		minima[i-w] = values[q.Front()]
		for !q.IsEmpty() && values[i] <= values[q.Back()] {
			q.PopBack()
		}
		for !q.IsEmpty() && q.Front() <= i-w {
			q.PopFront()
		}
		q.PushBack(i)
	}
	minima[n-w] = values[q.Front()]
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
	qMax := NewIntLinkedDeque()
	qMin := NewIntLinkedDeque()
	for i := 0; i < w; i++ {
		for !qMax.IsEmpty() && values[i] >= values[qMax.Back()] {
			qMax.PopBack()
		}
		qMax.PushBack(i)
		for !qMin.IsEmpty() && values[i] <= values[qMin.Back()] {
			qMin.PopBack()
		}
		qMin.PushBack(i)
	}
	for i := w; i < n; i++ {
		maxDiffs[i-w] = values[qMax.Front()] - values[qMin.Front()]
		for !qMax.IsEmpty() && values[i] >= values[qMax.Back()] {
			qMax.PopBack()
		}
		for !qMax.IsEmpty() && qMax.Front() <= i-w {
			qMax.PopFront()
		}
		qMax.PushBack(i)
		for !qMin.IsEmpty() && values[i] <= values[qMin.Back()] {
			qMin.PopBack()
		}
		for !qMin.IsEmpty() && qMin.Front() <= i-w {
			qMin.PopFront()
		}
		qMin.PushBack(i)
	}
	maxDiffs[n-w] = values[qMax.Front()] - values[qMin.Front()]
	return maxDiffs
}
