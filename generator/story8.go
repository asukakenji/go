// story8.go

// +build ignore

package main

import (
	"fmt"
)

type XRange struct {
	begin int
	end   int
	stop  func()
}

func NewXRange(begin, end int) *XRange {
	return &XRange{
		begin: begin,
		end:   end,
	}
}

func (r *XRange) Start() <-chan int {
	chOut := make(chan int)
	chStop := make(chan struct{})
	r.stop = func() {
		close(chStop)
	}
	go func() {
		defer close(chOut)
		for i := r.begin; i < r.end; i++ {
			select {
			case chOut <- i:
				// Nothing to be done
			case <-chStop:
				break
			}
		}
	}()
	return chOut
}

func (r *XRange) Stop() {
	r.stop()
}

func main() {
	r := NewXRange(10, 20)
	c := r.Start()
	for x := range c {
		if x%8 == 0 {
			break
		}
		fmt.Println(x)
	}
	r.Stop()
}
