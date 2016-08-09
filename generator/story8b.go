// story8b.go

// +build ignore

package main

import (
	"fmt"
)

type IntYielder func(int) bool

type IntGenerator struct {
	impl func(IntYielder)
	stop func()
}

func NewIntGenerator(impl func(IntYielder)) *IntGenerator {
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
		defer close(chOut)
		g.impl(IntYielder(func(msg int) bool {
			select {
			case chOut <- msg:
				return true
			case <-chStop:
				return false
			}
		}))
	}()
	return chOut
}

func (g *IntGenerator) Stop() {
	g.stop()
}

func XRange(begin, end int) *IntGenerator {
	return NewIntGenerator(func(yield IntYielder) {
		for i := begin; i < end; i++ {
			if !yield(i) {
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
}
