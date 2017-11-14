package generator_test

import (
	"fmt"
	"sync"

	"github.com/asukakenji/go/generator"
)

func ExampleGenerator_fibonacci() {
	// fib is a function literal (anonymous function / closure) here because
	// of the limitations of Go's "Example" mechanism. It is more preferable to
	// write it as a function with the following signature:
	//
	//     func Fib() *generator.Generator
	//
	fib := func() *generator.Generator {
		impl := func(yield func(v interface{}) bool) {
			a, b := 0, 1
			for {
				if !yield(a) {
					return
				}
				a, b = b, a+b
			}
		}
		return generator.NewGenerator(impl)
	}

	// Create a new Generator.
	// fib does not require any parameter and g always begins with 0, 1, 1...
	g := fib()

	// g.Start() is called to start the generator and receive the channel.
	// Note that g.Stop() is essential no matter ch is exhausted or not.
	// It should be called via defer as shown.
	func() {
		ch := g.Start()
		defer g.Stop()
		for v := range ch {
			if v.(int) > 20 {
				break
			}
			fmt.Printf("%d,", v)
		}
		fmt.Println()
	}()
	// Output: 0,1,1,2,3,5,8,13,
}

func ExampleGenerator_xRange() {
	// xrange is a function literal (anonymous function / closure) here because
	// of the limitations of Go's "Example" mechanism. It is more preferable to
	// write it as a function with the following signature:
	//
	//     func XRange(begin, end int) *generator.Generator
	//
	xrange := func(begin, end int) *generator.Generator {
		impl := func(yield func(v interface{}) bool) {
			for i := begin; i < end; i++ {
				if !yield(i) {
					return
				}
			}
		}
		return generator.NewGenerator(impl)
	}

	// Create a new Generator.
	// xrange requires 2 parameters: begin and end.
	// When g starts, it always begins with 10, 11, 12...
	g := xrange(10, 30)

	// g.Start() is called to start the generator and receive the channel.
	// Note that g.Stop() is essential no matter ch is exhausted or not.
	// It should be called via defer as shown.
	func() {
		ch := g.Start()
		defer g.Stop()
		for v := range ch {
			if v.(int)%8 == 0 {
				break
			}
			fmt.Printf("%d,", v)
		}
		fmt.Println()
	}()

	// After g.Stop() is called, g.Start() could be called to start the
	// generator from the beginning again.
	func() {
		ch := g.Start()
		defer g.Stop()
		for v := range ch {
			fmt.Printf("%d,", v)
		}
		fmt.Println()
	}()
	// Output: 10,11,12,13,14,15,
	// 10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,
}

func ExampleGenerator_xRangeGoroutines() {
	// The same as the XRange example above.
	xrange := func(begin, end int) *generator.Generator {
		impl := func(yield func(v interface{}) bool) {
			for i := begin; i < end; i++ {
				if !yield(i) {
					return
				}
			}
		}
		return generator.NewGenerator(impl)
	}

	// The same as the XRange example above.
	g := xrange(10, 30)

	// ch, the channel returned from g.Start(), could be shared between
	// goroutines. Each yield from the generator implementation is delivered
	// (to one of the goroutines) only once. Therefore, the following code
	// prints each number in 10 ~ 29 ONCE in no particular order.
	//
	// This behavior is different from using multithreading in Python,
	// which results in printing 10 ~ 29 in EACH thread (TWICE for 2 threads,
	// 3 TIMES for 3 threads).
	func() {
		ch := g.Start()
		defer g.Stop()

		var wg sync.WaitGroup
		for i := 0; i < 3; i++ {
			wg.Add(1)
			go func(ch <-chan interface{}) {
				defer wg.Done()
				for v := range ch {
					fmt.Println(v)
				}
			}(ch)
		}
		wg.Wait()
	}()
	// Unordered output:
	// 10
	// 12
	// 11
	// 13
	// 16
	// 17
	// 18
	// 19
	// 20
	// 21
	// 22
	// 15
	// 23
	// 24
	// 14
	// 27
	// 28
	// 26
	// 25
	// 29
}
