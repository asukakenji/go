package generator_test

import (
	"fmt"

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
	g := xrange(10, 20)

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
	// 10,11,12,13,14,15,16,17,18,19,
}
