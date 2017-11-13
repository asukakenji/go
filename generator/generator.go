package generator

// Generator is a type helps implementing Python-style "yield" generator.
type Generator struct {
	// impl is the implementation of the generator.
	impl func(func(interface{}) bool)

	// stop is the implementation of the Stop method.
	//
	// It is assigned by the Start method.
	// Therefore, if Stop is called prior to Start,
	// the program panics.
	stop func()
}

// NewGenerator creates a new Generator.
func NewGenerator(impl func(func(interface{}) bool)) *Generator {
	return &Generator{
		impl: impl,
	}
}

// Start starts the generator.
func (g *Generator) Start() <-chan interface{} {
	chOut := make(chan interface{})
	chStop := make(chan struct{})
	g.stop = func() {
		close(chStop)
	}
	go func() {
		defer close(chOut)
		g.impl(func(v interface{}) bool {
			select {
			case chOut <- v:
				return true
			case <-chStop:
				return false
			}
		})
	}()
	return chOut
}

// Stop stops the generator.
//
// If Stop is called prior to Start, the program panics.
//
func (g *Generator) Stop() {
	g.stop()
}
