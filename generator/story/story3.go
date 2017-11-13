// story3.go

// +build ignore

package main

import (
	"fmt"
)

func xrange(begin, end int) func(func(int)) {
	return func(consumer func(int)) {
		for i := begin; i < end; i++ {
			consumer(i)
		}
	}
}

func main() {
	xrange(10, 20)(func(x int) {
		fmt.Println(x)
	})
}
