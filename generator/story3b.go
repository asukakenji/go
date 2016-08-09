// story3b.go

// +build ignore

package main

import (
	"fmt"
)

func xrange(begin, end int) func(func(int) bool) {
	return func(predicate func(int) bool) {
		for i := begin; i < end; i++ {
			if !predicate(i) {
				break
			}
		}
	}
}

func main() {
	f := xrange(10, 20)
	f(func(x int) bool {
		if x%8 == 0 {
			return false
		}
		fmt.Println(x)
		return true
	})
}
