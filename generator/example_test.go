package generator_test

import (
	"fmt"

	"github.com/asukakenji/go/generator"
)

func ExampleNewGenerator() {
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

	g := xrange(10, 20)

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
