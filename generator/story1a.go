// story1a.go

// +build ignore

package main

import (
	"fmt"
)

func xrange(begin, end int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := begin; i < end; i++ {
			ch <- i
		}
	}()
	return ch
}

func main() {
	c := xrange(10, 20)
	for x := range c {
		fmt.Println(x)
	}
}
