package stream

type BasicStream struct {
	openUpstreamFunc func(<-chan struct{}) <-chan interface{}
}

func NewBasicStream(openUpstreamFunc func(<-chan struct{}) <-chan interface{}) BasicStream {
	return BasicStream{openUpstreamFunc}
}

func (s BasicStream) Filter(predicate func(interface{}) bool) Stream {
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

func (s BasicStream) Map(mapper func(interface{}) interface{}) Stream {
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

/*
func (s BasicStream) MapToInt(mapper func(interface{}) int) IntStream {
	out := make(chan int)
	result := BasicIntStream{out}
	go func() {
		defer close(out)
		for v := range s.in {
			out <- mapper(v)
		}
	}()
	return result
}

// TODO: Deal with remaining streams when somebody calls closeUpstream on downstream.
func (s BasicStream) FlatMap(mapper func(interface{}) Stream) Stream {
	out := chan interface{}(nil)
	s2 := chan interface{}(nil)
	newOpenUpstream := func() <-chan interface{} {
		in := openUpstream()
		out = make(chan interface{})
		go func() {
			defer close(out)
			for v := range in {
				s2 := mapper(v)
				s2.ForEach(func (v2 interface{}) {
					out <- v2
				})
			}
		}()
	}
	newCloseUpstream := func() {
		s2.closeUpstream()
		close(out)
	}
	return BasicStream{newOpenUpstream, newCloseUpstream}
}

func (s BasicStream) FlatMapToInt(mapper func(interface{}) IntStream) IntStream {
	out := make(chan int)
	result := BasicIntStream{out}
	go func() {
		defer close(out)
		for v := range s.in {
			mapper(v).ForEach(func (v2 int) {
				out <- v2
			})
		}
	}()
	return result
}

func (s BasicStream) Distinct() Stream {
	out := chan interface{}(nil)
	newOpenUpstream := func() <-chan interface{} {
		in := openStream()
		out = make(chan interface{})
		go func() {
			defer close(out)
			seenValues := make([]interface{}, 0)
			for v := range in {
				isFound := false
				for _, v2 := range seenValues {
					if reflect.DeepEqual(v, v2) {
						isFound = true
						break
					}
				}
				if !isFound {
					seenValues = append(seenValues, v)
					out <- v
				}
			}
		}()
		return out
	}
	newCloseUpstream := func() {
		close(out)
	}
	return BasicStream{newOpenUpstream, newCloseUpstream}
}

func (s BasicStream) Sorted(less func(interface{}, interface{}) bool) Stream {
	out := chan interface{}(nil)
	newOpenUpstream := func() <-chan interface{} {
		in := openStream()
		out = make(chan interface{})
		go func() {
			defer close(out)
			values := make([]interface{}, 0)
			for v := range in {
				values = append(values, v)
			}
			sort.Stable(sortableSlice{values, less})
			for _, v := range values {
				out <- v
			}
		}
		return out
	}
	newCloseUpstream := func() {
		close(out)
	}
	return BasicStream{newOpenUpstream, newCloseUpstream}
}

type sortableSlice struct {
	slice []interface{}
	less func(interface{}, interface{}) bool
}

func (ss sortableSlice) Len() int {
	return len(ss.slice)
}

func (ss sortableSlice) Less(i, j int) bool {
	return less(ss.slice[i], ss.slice[j])
}

func (ss sortableSlice) Swap(i, j int) {
	ss.slice[i], ss.slice[j] = ss.slice[j], ss.slice[i]
}

func (s BasicStream) ForEach(consumer func(interface{})) {
	in := openUpstream()
	for v := range in {
		consumer(v)
	}
}
*/

func (s BasicStream) ForEach(consumer func(interface{})) {
	signal := make(chan struct{})
	in := s.openUpstreamFunc(signal)
	defer close(signal)
	for v := range in {
		consumer(v)
	}
}
