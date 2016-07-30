package stream

type IntStream interface {
	Filter(predicate func(int) bool) IntStream
	Map(mapper func(int) interface{}) Stream
	ForEach(consumer func(int))
}
