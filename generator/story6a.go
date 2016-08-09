// story6a.go

// +build ignore

package main

import (
	"fmt"
)

func xrange(begin, end int) (<-chan int, chan<- struct{}) {
	chOut := make(chan int)
	chStop := make(chan struct{})
	go func() {
		defer close(chOut)
		for i := begin; i < end; i++ {
			select {
			case chOut <- i:
				// Nothing to be done
			case <-chStop:
				break
			}
		}
	}()
	return chOut, chStop
}

func main() {
	chOut, chStop := xrange(10, 20)
	for x := range chOut {
		if x%8 == 0 {
			break
		}
		fmt.Println(x)
	}
	close(chStop)
}
