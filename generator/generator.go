package generator

type Yield func(interface{}) bool

type Generator struct {
	impl func(Yield)
	stop func()
}

func NewGenerator(impl func(Yield)) *Generator {
	return &Generator{
		impl: impl,
	}
}

func (g *Generator) Start() <-chan interface{} {
	chOut := make(chan interface{})
	chStop := make(chan struct{})
	g.stop = func() {
		close(chStop)
	}
	go func() {
		defer close(chOut)
		g.impl(Yield(func(msg interface{}) bool {
			select {
			case chOut <- msg:
				return true
			case <-chStop:
				return false
			}
		}))
	}()
	return chOut
}

func (g *Generator) Stop() {
	g.stop()
}
