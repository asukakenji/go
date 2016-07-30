package stream

type Stream interface {
	Filter(predicate func(interface{}) bool) Stream
	Map(mapper func(interface{}) interface{}) Stream
	//FlatMap(mapper func(interface{}) Stream) Stream
	//Distinct() Stream
	ForEach(consumer func(interface{}))
}
