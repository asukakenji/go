// story2a.go

// +build ignore

package main

import (
	"fmt"
)

func xrange(begin, end int, predicate func(int) bool) {
	for i := begin; i < end; i++ {
		if !predicate(i) {
			break
		}
	}
}

func main() {
	xrange(10, 20, func(x int) bool {
		if x%8 == 0 {
			return false
		}
		fmt.Println(x)
		return true
	})
}
