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

func (s SliceIntStream) Filter(predicate func(int) bool) IntStream {
	return NewBasicIntStream(func (downstreamSignal <-chan struct{}) <-chan int {
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

func (s SliceIntStream) Map(mapper func(int) interface{}) Stream {
	return NewBasicStream(func (downstreamSignal <-chan struct{}) <-chan interface{} {
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

func (s SliceIntStream) ForEach(consumer func(int)) {
	for _, e := range s.slice {
		consumer(e)
	}
}
