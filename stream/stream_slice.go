package stream

type SliceStream struct {
	slice []interface{}
}

func StreamOf(values ...interface{}) SliceStream {
	return SliceStream{values}
}

func (s SliceStream) Len() int {
	return len(s.slice)
}

func (s SliceStream) Cap() int {
	return cap(s.slice)
}

func (s SliceStream) Get(index int) interface{} {
	return s.slice[index]
}

func (s SliceStream) Filter(predicate func(interface{}) bool) Stream {
	return NewBasicStream(func (downstreamSignal <-chan struct{}) <-chan interface{} {
		out := make(chan interface{})
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

func (s SliceStream) Map(mapper func(interface{}) interface{}) Stream {
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

func (s SliceStream) ForEach(consumer func(interface{})) {
	for _, e := range s.slice {
		consumer(e)
	}
}
