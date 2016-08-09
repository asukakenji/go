// story5.go

// +build ignore

package main

import (
	"fmt"
)

func xrange(begin, end int) func() <-chan int {
	return func() <-chan int {
		ch := make(chan int)
		go func() {
			defer close(ch)
			for i := begin; i < end; i++ {
				ch <- i
			}
		}()
		return ch
	}
}

func main() {
	for x := range xrange(10, 20)() {
		fmt.Println(x)
	}
}
