package generator

type Yielder func(interface{}) bool

type Generator struct {
	impl func(Yielder)
	stop func()
}

func NewGenerator(impl func(Yielder)) *Generator {
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
		g.impl(Yielder(func(msg interface{}) bool {
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
