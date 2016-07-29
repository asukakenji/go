package stream

type IntStream interface {
	ForEach(consumer func(int))
	Filter(predicate func(int) bool) IntStream
	Map(mapper func(int) interface{}) Stream
}





type BasicIntStream struct {
	in <- chan int
}

func (s BasicIntStream) ForEach(consumer func(int)) {
	for {
		v, ok := <- s.in
		if !ok {
			break
		}
		consumer(v)
	}
}

func (s BasicIntStream) Filter(predicate func(int) bool) IntStream {
	out := make(chan int)
	result := BasicIntStream{out}
	go func() {
		defer close(out)
		for {
			v, ok := <- s.in
			if !ok {
				break
			}
			if predicate(v) {
				out <- v
			}
		}
	}()
	return result
}

func (s BasicIntStream) Map(mapper func(int) interface{}) Stream {
	out := make(chan interface{})
	result := BasicStream{out}
	go func() {
		defer close(out)
		for {
			v, ok := <- s.in
			if !ok {
				break
			}
			out <- mapper(v)
		}
	}()
	return result
}





type RangeIntStream struct {
	from int
	to int
}

func (s RangeIntStream) ForEach(consumer func(int)) {
	for v := s.from; v <= s.to; v++ {
		consumer(v)
	}
}

func (s RangeIntStream) Filter(predicate func(int) bool) IntStream {
	out := make(chan int)
	result := BasicIntStream{out}
	go func() {
		defer close(out)
		for v := s.from; v <= s.to; v++ {
			if predicate(v) {
				out <- v
			}
		}
	}()
	return result
}

func (s RangeIntStream) Map(mapper func(int) interface{}) Stream {
	out := make(chan interface{})
	result := BasicStream{out}
	go func() {
		defer close(out)
		for v := s.from; v <= s.to; v++ {
			out <- mapper(v)
		}
	}()
	return result
}





type SliceIntStream struct {
	slice []int
}

func (s SliceIntStream) ForEach(consumer func(int)) {
	for _, v := range s.slice {
		consumer(v)
	}
}

func (s SliceIntStream) Filter(predicate func(int) bool) IntStream {
	out := make(chan int)
	result := BasicIntStream{out}
	go func() {
		defer close(out)
		for _, v := range s.slice {
			if predicate(v) {
				out <- v
			}
		}
	}()
	return result
}

func (s SliceIntStream) Map(mapper func(int) interface{}) Stream {
	out := make(chan interface{})
	result := BasicStream{out}
	go func() {
		defer close(out)
		for _, v := range s.slice {
			out <- mapper(v)
		}
	}()
	return result
}
