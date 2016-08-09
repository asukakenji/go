// story4.go

// +build ignore

package main

import (
	"fmt"
)

func xrange(begin, end int) chan int {
	ch := make(chan int)
	go func() {
		for i := begin; i < end; i++ {
			ch <- i
		}
		fmt.Println("X")
	}()
	return ch
}

func main() {
	c := xrange(10, 20)
	for x := range c {
		if x%8 == 0 {
			break
		}
		fmt.Println(x)
	}
	close(c)
}
