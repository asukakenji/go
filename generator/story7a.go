// story7a.go

// +build ignore

package main

import (
	"fmt"
)

func xrange(begin, end int) func() (<-chan int, func()) {
	return func() (<-chan int, func()) {
		chOut := make(chan int)
		chStop := make(chan struct{})
		fnStop := func() {
			close(chStop)
		}
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
		return chOut, fnStop
	}
}

func main() {
	fnStart := xrange(10, 20)
	chOut, fnStop := fnStart()
	for x := range chOut {
		if x%8 == 0 {
			break
		}
		fmt.Println(x)
	}
	fnStop()
}
