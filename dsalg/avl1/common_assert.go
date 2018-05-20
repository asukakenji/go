// +build assert

package avl

func assert(condition bool, message interface{}) {
	if !condition {
		panic(message)
	}
}
