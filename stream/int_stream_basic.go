package stream

type BasicIntStream struct {
	openUpstreamFunc func(<-chan struct{}) <-chan int
}

func NewBasicIntStream(openUpstreamFunc func(<-chan struct{}) <-chan int) BasicIntStream {
	return BasicIntStream{openUpstreamFunc}
}

func (s BasicIntStream) Filter(predicate func(int) bool) IntStream {
	return NewBasicIntStream(func (downstreamSignal <-chan struct{}) <-chan int {
		signal := make(chan struct{})
		in := s.openUpstreamFunc(signal)
		in2 := in
		out := (chan int)(nil)
		out2 := make(chan int)
		go func() {
			defer close(signal)
			defer close(out2)
			var e int
		outer:
			for {
				select {
				case v, ok := <-in:
					if !ok {
						break outer
					}
					if !predicate(v) {
						continue
					}
					e = v
					in, out = nil, out2
				case out <- e:
					in, out = in2, nil
				case _, ok := <-downstreamSignal:
					if !ok {
						break outer
					}
				}
			}
		}()
		return out2
	})
}

func (s BasicIntStream) Map(mapper func(int) interface{}) Stream {
	return NewBasicStream(func (downstreamSignal <-chan struct{}) <-chan interface{} {
		signal := make(chan struct{})
		in := s.openUpstreamFunc(signal)
		in2 := in
		out := (chan interface{})(nil)
		out2 := make(chan interface{})
		go func() {
			defer close(signal)
			defer close(out2)
			var e interface{}
		outer:
			for {
				select {
				case v, ok := <-in:
					if !ok {
						break outer
					}
					e = mapper(v)
					in, out = nil, out2
				case out <- e:
					in, out = in2, nil
				case _, ok := <-downstreamSignal:
					if !ok {
						break outer
					}
				}
			}
		}()
		return out2
	})
}

func (s BasicIntStream) ForEach(consumer func(int)) {
	signal := make(chan struct{})
	in := s.openUpstreamFunc(signal)
	defer close(signal)
	for v := range in {
		consumer(v)
	}
}
