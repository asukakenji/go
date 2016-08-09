// story5a.go

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
	f := xrange(10, 20)
	for x := range f() {
		fmt.Println(x)
	}
}
