// story5.go

// +build ignore

package main

import (
	"fmt"
	"time"
)

func xrange(begin, end int) func() <-chan int {
	return func() <-chan int {
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
}

func main() {
	for x := range xrange(10, 20)() {
		fmt.Println(x)
	}
	time.Sleep(1 * time.Second)
}
