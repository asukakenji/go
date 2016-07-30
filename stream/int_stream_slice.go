package stream

type SliceIntStream struct {
	slice []int
}

func IntStreamOf(values ...int) SliceIntStream {
	return SliceIntStream{values}
}

func (s SliceIntStream) Len() int {
	return len(s.slice)
}

func (s SliceIntStream) Cap() int {
	return cap(s.slice)
}

func (s SliceIntStream) Get(index int) int {
	return s.slice[index]
}

func (s SliceIntStream) FilterInt(predicate func(int) bool) IntStream {
	return NewBasicIntStream(func(downstreamSignal <-chan struct{}) <-chan int {
		out := make(chan int)
		go func() {
			defer close(out)
		outer:
			for _, e := range s.slice {
				if !predicate(e) {
					continue
				}
				select {
				case out <- e:
					// Sent, nothing to be done
				case _, ok := <-downstreamSignal:
					if !ok {
						break outer
					}
				}
			}
		}()
		return out
	})
}

func (s SliceIntStream) MapInt(mapper func(int) interface{}) Stream {
	return NewBasicStream(func(downstreamSignal <-chan struct{}) <-chan interface{} {
		out := make(chan interface{})
		go func() {
			defer close(out)
		outer:
			for _, e := range s.slice {
				select {
				case out <- mapper(e):
					// Sent, nothing to be done
				case _, ok := <-downstreamSignal:
					if !ok {
						break outer
					}
				}
			}
		}()
		return out
	})
}

func (s SliceIntStream) ForEachInt(consumer func(int)) {
	for _, e := range s.slice {
		consumer(e)
	}
}

// Adapter Methods

func (s SliceIntStream) Filter(predicate func(interface{}) bool) Stream {
	return s.FilterInt(func(x int) bool {
		return predicate(x)
	})
}

func (s SliceIntStream) Map(mapper func(interface{}) interface{}) Stream {
	return s.MapInt(func(x int) interface{} {
		return mapper(x)
	})
}

func (s SliceIntStream) ForEach(consumer func(interface{})) {
	s.ForEachInt(func(x int) {
		consumer(x)
	})
}
