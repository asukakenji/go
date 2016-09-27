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

// Reference: sort.IntSlice.Less
func IntLess(a, b interface{}) bool {
	return a.(int) < b.(int)
}

// Reference: sort.Float64Slice.Less
func Float64Less(a, b interface{}) bool {
	return a.(float64) < b.(float64) || isNaN(a.(float64)) && !isNaN(b.(float64))
}

// Reference: sort.StringSlice.Less
func StringLess(a, b interface{}) bool {
	return a.(string) < b.(string)
}
