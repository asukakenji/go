package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	seedString := os.Getenv("SEED")
	// if seedString == "" {
	// 	panic("No seed")
	// }
	seed, err := strconv.ParseInt(seedString, 10, 64)
	if err != nil {
		panic("Invalid seed")
	}
	rand.Seed(seed)
	numbers := rand.Perm(255)
	fmt.Println(numbers[0])
	fmt.Println(numbers[1])
	fmt.Println(numbers[2])
	fmt.Println(numbers[3])
}
