package avl

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func IntLess(a interface{}, b interface{}) bool {
	return a.(int) < b.(int)
}

func Float64Less(a interface{}, b interface{}) bool {
	return a.(float64) < b.(float64)
}

func StringLess(a interface{}, b interface{}) bool {
	return a.(string) < b.(string)
}
