package generator

// Generator is a type helps implementing Python-style "yield" generator.
type Generator struct {
	// start is the implementation of the Start method.
	//
	// It is assigned by the NewGenerator function and the Stop method,
	// and unassigned by itself.
	// Therefore, if Start is called more than once without Stop,
	// the program panics.
	start func() <-chan interface{}

	// stop is the implementation of the Stop method.
	//
	// It is assigned by the Start method.
	// Therefore, if Stop is called prior to Start,
	// the program panics.
	stop func()
}

// makeStart creates the start method of a Generator.
func makeStart(impl func(yield func(v interface{}) bool), g *Generator) func() <-chan interface{} {
	return func() <-chan interface{} {
		g.start = nil
		chOut := make(chan interface{})
		chStop := make(chan struct{})
		g.stop = makeStop(impl, g, chStop)
		go func() {
			defer close(chOut)
			yield := func(v interface{}) bool {
				select {
				case chOut <- v:
					return true
				case <-chStop:
					return false
				}
			}
			impl(yield)
		}()
		return chOut
	}
}

// makeStop creates the stop method of a Generator.
func makeStop(impl func(yield func(v interface{}) bool), g *Generator, chStop chan struct{}) func() {
	return func() {
		close(chStop)
		g.start = makeStart(impl, g)
	}
}

// NewGenerator creates a new Generator.
func NewGenerator(impl func(yield func(v interface{}) bool)) *Generator {
	g := &Generator{}
	g.start = makeStart(impl, g)
	return g
}

// Start starts the generator.
//
// If Start is called more than once without Stop,
// the program panics.
//
func (g *Generator) Start() <-chan interface{} {
	return g.start()
}

// Stop stops the generator.
//
// If Stop is called prior to Start, the program panics.
//
func (g *Generator) Stop() {
	g.stop()
}
