// story8a.go

// +build ignore

package main

import (
	"fmt"
	"time"
)

type IntGenerator struct {
	impl func(chan<- int, <-chan struct{})
	stop func()
}

func NewIntGenerator(impl func(chan<- int, <-chan struct{})) *IntGenerator {
	return &IntGenerator{
		impl: impl,
	}
}

func (g *IntGenerator) Start() <-chan int {
	chOut := make(chan int)
	chStop := make(chan struct{})
	g.stop = func() {
		close(chStop)
	}
	go func() {
		fmt.Println("Goroutine Started")
		defer fmt.Println("Goroutine Terminated")
		defer close(chOut)
		g.impl(chOut, chStop)
	}()
	return chOut
}

func (g *IntGenerator) Stop() {
	g.stop()
}

func XRange(begin, end int) *IntGenerator {
	return NewIntGenerator(func(chOut chan<- int, chStop <-chan struct{}) {
		for i := begin; i < end; i++ {
			select {
			case chOut <- i:
				// Nothing to be done
			case <-chStop:
				break
			}
		}
	})
}

func main() {
	g := XRange(10, 20)
	c := g.Start()
	for x := range c {
		if x%8 == 0 {
			break
		}
		fmt.Println(x)
	}
	g.Stop()
	time.Sleep(1 * time.Second)
}
