package main

import (
	"fmt"
	"math/rand"

	avl "github.com/asukakenji/go/dsalg/avl1"
)

func main() {
	numbers := rand.Perm(255)

	t := new(avl.IntTreeSet)
	for _, number := range numbers {
		fmt.Printf("Adding %d...\n", number)
		t.Add(number)
		fmt.Println()
	}
}
