// story2.go

// +build ignore

package main

import (
	"fmt"
)

func xrange(begin, end int, consumer func(int)) {
	for i := begin; i < end; i++ {
		consumer(i)
	}
}

func main() {
	xrange(10, 20, func(x int) {
		fmt.Println(x)
	})
}
