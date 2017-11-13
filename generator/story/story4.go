// story4.go

// +build ignore

package main

import (
	"fmt"
	"time"
)

func xrange(begin, end int) chan int {
	ch := make(chan int)
	go func() {
		fmt.Println("Goroutine Started")
		defer fmt.Println("Goroutine Terminated")
		for i := begin; i < end; i++ {
			ch <- i
		}
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
	time.Sleep(1 * time.Second)
}
