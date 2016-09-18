package dsalg

// min is a copy of strconv.min
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// max is a copy of strconv.max
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// isNaN is a copy of math.IsNaN to avoid a dependency on the math package.
func isNaN(f float64) bool {
	return f != f
}
