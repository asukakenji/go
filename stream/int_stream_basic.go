package stream

type BasicIntStream struct {
	openUpstreamFunc func(<-chan struct{}) <-chan int
}

func NewBasicIntStream(openUpstreamFunc func(<-chan struct{}) <-chan int) BasicIntStream {
	return BasicIntStream{openUpstreamFunc}
}

func (s BasicIntStream) FilterInt(predicate func(int) bool) IntStream {
	return NewBasicIntStream(func(downstreamSignal <-chan struct{}) <-chan int {
		signal, in, inBackup := s.prepareInChannels()
		out, outBackup, outForReturn := prepareIntOutChannels()
		go func() {
			defer close(signal)
			defer close(outBackup)
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
					in, out = nil, outBackup
				case out <- e:
					in, out = inBackup, nil
				case _, ok := <-downstreamSignal:
					if !ok {
						break outer
					}
				}
			}
		}()
		return outForReturn
	})
}

func (s BasicIntStream) MapInt(mapper func(int) interface{}) Stream {
	return NewBasicStream(func(downstreamSignal <-chan struct{}) <-chan interface{} {
		signal, in, inBackup := s.prepareInChannels()
		out, outBackup, outForReturn := prepareOutChannels()
		go func() {
			defer close(signal)
			defer close(outBackup)
			var e interface{}
		outer:
			for {
				select {
				case v, ok := <-in:
					if !ok {
						break outer
					}
					e = mapper(v)
					in, out = nil, outBackup
				case out <- e:
					in, out = inBackup, nil
				case _, ok := <-downstreamSignal:
					if !ok {
						break outer
					}
				}
			}
		}()
		return outForReturn
	})
}

func (s BasicIntStream) ForEachInt(consumer func(int)) {
	signal, in, _ := s.prepareInChannels()
	defer close(signal)
	for v := range in {
		consumer(v)
	}
}

// Adapter Methods

func (s BasicIntStream) Filter(predicate func(interface{}) bool) Stream {
	return s.FilterInt(func(x int) bool {
		return predicate(x)
	})
}

func (s BasicIntStream) Map(mapper func(interface{}) interface{}) Stream {
	return s.MapInt(func(x int) interface{} {
		return mapper(x)
	})
}

func (s BasicIntStream) ForEach(consumer func(interface{})) {
	s.ForEachInt(func(x int) {
		consumer(x)
	})
}

// Helper Methods

func (s BasicIntStream) prepareInChannels() (signal chan<- struct{}, in, inBackup <-chan int) {
	signalBackup := make(chan struct{})
	inCommon := s.openUpstreamFunc(signalBackup)
	return signalBackup, inCommon, inCommon
}
