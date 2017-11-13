package generator_test

import (
	"fmt"

	"github.com/asukakenji/go/generator"
)

func ExampleNewGenerator() {
	xrange := func(begin, end int) *generator.Generator {
		return generator.NewGenerator(func(yield func(interface{}) bool) {
			for i := begin; i < end; i++ {
				if !yield(i) {
					break
				}
			}
		})
	}

	g := xrange(10, 20)
	c := g.Start()
	for x := range c {
		if x.(int)%8 == 0 {
			break
		}
		fmt.Println(x)
	}
	g.Stop()
	// Output: 10
	// 11
	// 12
	// 13
	// 14
	// 15
}
