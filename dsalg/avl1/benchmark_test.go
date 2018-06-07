package avl_test

import (
	"math/rand"
	"os"
	"strconv"
	"testing"

	avl "github.com/asukakenji/go/dsalg/avl1"
)

func BenchmarkTreeInsert(b *testing.B) {
	seedString := os.Getenv("SEED")
	seed, _ := strconv.ParseInt(seedString, 10, 64)
	rand.Seed(seed)
	values := rand.Perm(65535)
	b.ResetTimer()

	for i, count := 0, b.N; i < count; i++ {
		tree := avl.NewTree(avl.IntLess)
		for _, v := range values {
			tree.Insert(v)
		}
	}
}
