// story1a.go

// +build ignore

package main

import (
	"fmt"
	"time"
)

func xrange(begin, end int) <-chan int {
	ch := make(chan int)
	go func() {
		fmt.Println("Goroutine Started")
		defer fmt.Println("Goroutine Terminated")
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
	time.Sleep(1 * time.Second)
}
