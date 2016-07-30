package stream

type IntStream interface {
	Stream
	FilterInt(predicate func(int) bool) IntStream
	MapInt(mapper func(int) interface{}) Stream
	ForEachInt(consumer func(int))
}
