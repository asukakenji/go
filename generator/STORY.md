# Implementing Generators (`yield`) in Go (Golang)

## Introduction

**TODO**

Say something about:

- Generator
- Python's `yield`
- ECMAScript 6's `yield`,
- Control Flow
- Infinite sequence should be supported.

## Background

This library originates from writing an answer to this question on StackOverflow:

[golang implementing generator / yield with channels: odd channel behavior][so1]

This reminds me of a similar question I asked when I began developing in Go (Golang):

[The idiomatic way to implement generators (yield) in Golang for recursive functions][so2]

And this question asked by another developer:

[Python-style generators in Go][so3]

Since none of the answers to the above questions fulfill my requirements,
I decided to roll my own implementation.

[so1]: http://stackoverflow.com/q/37648288/142239
[so2]: http://stackoverflow.com/q/34464146/142239
[so3]: http://stackoverflow.com/q/11385556/142239

## Requirements

1. The underlying mechanism of the generator should be encapsulated by the library.
   The [RAII][raii] principle should be followed.
2. The generator should return a value assignable to a variable.
   If the value is not safe to be discarded because some resources are acquired,
   the library should provide an API to release the resources.
3. The library should provide an API to stop the generator.
   The client should be able to stop the generation if no more values are needed.
4. There shouldn't be any resource leak (primarily unterminated goroutines,
   secondarily unclosed channels) provided that the client follows the usage guidelines.
5. The generator should be generic.
   Existing algorithms should be able to fit into the library with minimum effort.
   The library should provide an API for the algorithm developers to control the flow.

[raii]: https://en.wikipedia.org/wiki/Resource_Acquisition_Is_Initialization

## Python Code

Since the final version may be very complicated at first sight,
let's begin with implementing Python's `xrange` (or `range` in Python 3),
and build up the concepts step-by-step.

This demonstrates the basic requirement:

```python
# xrange1.py
for x in xrange(10, 20):
	print x
```

This demonstrates "Requirement 2" above --- the result is assigned to a variable (`r`):

```python
# xrange2.py
r = xrange(10, 20)
for x in r:
	print x
```

This demonstrates "Requirement 3" above --- the generation is stoppable:

```python
# xrange3.py
for x in xrange(10, 20):
	if x % 8 == 0:
		break
	print x
```

## Go Code

### 1. Naïve Implementation (Many Tutorials Use This!)

Let's start with a straight forward implementation
(<kbd>[view][story1.go]</kbd>):

```go
// story1.go
func xrange(begin, end int) <-chan int {
	ch := make(chan int)
	go func() {
		fmt.Println("Goroutine Started")
		defer fmt.Println("Goroutine Terminated")
		defer close(ch)
		for i := begin; i < end; i++ {
			ch <- i
		}
	}()
	return ch
}

func main() {
	for x := range xrange(10, 20) {
		fmt.Println(x)
	}
	time.Sleep(1 * time.Second)
}
```

Here is the output:

```
Goroutine Started
10
11
12
13
14
15
16
17
18
19
Goroutine Terminated
```

In the code above,
the creation of the channel and the goroutine is handled by the library
(fulfills "Requirement 1").
The generator returns a value (the channel) assignable to a variable
(fulfills "Requirement 2") like this
(<kbd>[view][story1a.go]</kbd>):

```go
// story1a.go
func main() {
	c := xrange(10, 20)
	for x := range c {
		fmt.Println(x)
	}
	time.Sleep(1 * time.Second)
}
```

However, there is a serious problem --- resource leak.
The goroutine is started in the body of `xrange()`.
If the client doesn't read all values from the returned channel,
the goroutine blocks at `ch <- i` and never terminates.
Since we don't have a way to "destroy" a goroutine in Go,
and it will not be garbage collected,
there is a resource leak (breaks "Requirement 4").
The channel won't be closed properly too,
since `close(ch)` (in `defer close(ch)`) is only invoked
immediately before the surrounding function (the goroutine) returns.

The following code demonstrates the resource leak
(<kbd>[view][story1b.go]</kbd>):

```go
// story1b.go
func main() {
	for x := range xrange(10, 20) {
		if x%8 == 0 {
			break
		}
		fmt.Println(x)
	}
	time.Sleep(1 * time.Second)
}
```

Here is the output:

```
Goroutine Started
10
11
12
13
14
15
```

### 2. An Alternative Approach - Consumer (Callback)

Before digging deeper into the "goroutine-channel" implementation,
let's take a look at a simpler alternative first.
Since this approach doesn't make use of goroutines or channels,
there is no resource leak, and thus "Requirement 4" is automatically fulfilled.

Here is the basic block
(<kbd>[view][story2.go]</kbd>):

```go
// story2.go
func xrange(begin, end int, consumer func(int)) {
	for i := begin; i < end; i++ {
		consumer(i)
	}
}

func main() {
	xrange(10, 20, func(x int) {
		fmt.Println(x)
	})
}
```

Since the generation is not stoppable (breaks "Requirement 3"),
let's modify it like this
(<kbd>[view][story2a.go]</kbd>):

```go
// story2a.go
func xrange(begin, end int, predicate func(int) bool) {
	for i := begin; i < end; i++ {
		if !predicate(i) {
			break
		}
	}
}

func main() {
	xrange(10, 20, func(x int) bool {
		if x%8 == 0 {
			return false
		}
		fmt.Println(x)
		return true
	})
}
```

This looks good,
but the generator doesn't return a value assignable to a variable (breaks "Requirement 2").
To achieve this,
the generator should not take the callback (`consumer` or `predicate`) as a parameter,
it should instead *return* a closure that takes the callback as a parameter.

### 3. Consumer (Callback) + Closure

Let's start with the "`consumer` version" first
(<kbd>[view][story3.go]</kbd>):

```go
// story3.go
func xrange(begin, end int) func(func(int)) {
	return func(consumer func(int)) {
		for i := begin; i < end; i++ {
			consumer(i)
		}
	}
}

func main() {
	xrange(10, 20)(func(x int) {
		fmt.Println(x)
	})
}
```

The `xrange()` function returns a closure that takes a `consumer`
(which could also be a closure or a function reference) as a parameter.
The `for`-loop is "hard-wired" for every invocation
(`begin` and `end` could be different, though),
but what to do with the result `i` is determined by the `consumer` supplied.
The returned closure could be carried around like this
(<kbd>[view][story3a.go]</kbd>):

```go
// story3a.go
func main() {
	f := xrange(10, 20)
	f(func(x int) {
		fmt.Println(x)
	})
}
```

The variable `f`, containing the closure,
can be passed around or discarded without any resource leak (fulfills "Requirement 2").
Its invocation could be deferred until a suitable `consumer` is ready,
which tells it what to do with each number it generates.
In the code above, the `consumer` contains only one line: `fmt.Println(x)`.

Now, let's modify it so that it is stoppable
(<kbd>[view][story3b.go]</kbd>):

```go
// story3b.go
func xrange(begin, end int) func(func(int) bool) {
	return func(predicate func(int) bool) {
		for i := begin; i < end; i++ {
			if !predicate(i) {
				break
			}
		}
	}
}

func main() {
	f := xrange(10, 20)
	f(func(x int) bool {
		if x%8 == 0 {
			return false
		}
		fmt.Println(x)
		return true
	})
}
```

At this point, our requirements are all fulfilled. Are we done? Of course not!
What we just developed is a function that accepts a callback,
which is very different from what a generator is.

The callback approach hard wires the control flow for the client.
The client callback functions fill in the holes in the library code.
It is more popular in (and more appropriate for) situations like
"providing a comparison callback to a sorting function",
or "providing an event handler callback to an event-driven framework".

On the other hand, the generator approach lets the client control the flow.
You may treat the generator as a stream of objects,
and one object is taken away from the stream each time it is invoked.

### 4. Goroutine-Channel Approach - Wrong Solution

Before starting our first step into the correct solution,
let's see an incorrect one first
(<kbd>[view][story4.go]</kbd>):

```go
// story4.go
func xrange(begin, end int) chan int {
	ch := make(chan int)
	go func() {
		fmt.Println("Goroutine Started")
		defer fmt.Println("Goroutine Terminated")
		for i := begin; i < end; i++ {
			ch <- i
		}
	}()
	return ch
}

func main() {
	c := xrange(10, 20)
	for x := range c {
		if x%8 == 0 {
			break
		}
		fmt.Println(x)
	}
	close(c)
	time.Sleep(1 * time.Second)
}
```

The code above modified the generator to return a `chan int` instead of a `<-chan int`,
so that the client code has a chance to `close()` the channel.

However, there are 2 major problems.
Firstly, as already stated in Requirement 4,
the primary resource leak is unterminated goroutines, not unclosed channels.
Closing the channel does not automatically terminate the goroutine.
Secondly, since the goroutine continues to run after the channel is closed,
it panics when it reaches `ch <- i`, and the program crashes.

### 5. Preventing Resource Leak (Start Only When I Ask So)

Let's make our first step into the correct solution!

The following is what Requirement 2 states:

> The generator should return a value assignable to a variable.
> If the value is not safe to be discarded because some resources are acquired,
> the library should provide an API to release the resources.

This could easily be achieved by returning "a function that returns a channel",
instead of "returning a channel" as in Step 1
(<kbd>[view][story5.go]</kbd>):

```go
// story5.go
func xrange(begin, end int) func() <-chan int {
	return func() <-chan int {
		ch := make(chan int)
		go func() {
			fmt.Println("Goroutine Started")
			defer fmt.Println("Goroutine Terminated")
			defer close(ch)
			for i := begin; i < end; i++ {
				ch <- i
			}
		}()
		return ch
	}
}

func main() {
	for x := range xrange(10, 20)() {
		fmt.Println(x)
	}
	time.Sleep(1 * time.Second)
}
```

There are 2 points worth to notice:
1. Since a closure or a function reference does not involve resource management,
   there is no resource leak no matter how the returned value is manipulated.
2. The major reason of why we had resource leak so far is that
   the goroutine started in the body of the generator.
   Now, the goroutine is started in the body of the returned function,
   so there is no resource leak if we simply discard it without invoking it.

In the client code, `xrange(10, 20)` returns "a function that returns a channel",
`xrange(10, 20)()` invokes the returned function,
and evaluates to the returned channel.
It is then fed to `range`, turning the line into a `for`-`range`-loop over a channel.

The following is a variant of the client code shown above
(<kbd>[view][story5a.go]</kbd>):

```go
// story5a.go
func main() {
	f := xrange(10, 20)
	for x := range f() {
		fmt.Println(x)
	}
	time.Sleep(1 * time.Second)
}
```

You can see that the value returned by `xrange()` is assigned to `f`
(fulfills "Requirement 2").

However, there is still resource leak if not all the values are received from the channel:

<kbd>[view][story5b.go]</kbd>:

```go
// story5b.go
func main() {
	f := xrange(10, 20)
	for x := range f() {
		if x%8 == 0 {
			break
		}
		fmt.Println(x)
	}
	time.Sleep(1 * time.Second)
}
```

### 6. Preventing Resource Leak (Stop When I Say So)

**TODO: Write this!**

<kbd>[view][story6.go]</kbd>:

```go
// story6.go
func xrange(begin, end int) (<-chan int, chan<- struct{}) {
	chOut := make(chan int)
	chStop := make(chan struct{})
	go func() {
		fmt.Println("Goroutine Started")
		defer fmt.Println("Goroutine Terminated")
		defer close(chOut)
		for i := begin; i < end; i++ {
			select {
			case chOut <- i:
				// Nothing to be done
			case <-chStop:
				break
			}
		}
	}()
	return chOut, chStop
}

func main() {
	chOut, chStop := xrange(10, 20)
	for x := range chOut {
		fmt.Println(x)
	}
	close(chStop)
	time.Sleep(1 * time.Second)
}
```

<kbd>[view][story6a.go]</kbd>:

```go
// story6a.go
func main() {
	chOut, chStop := xrange(10, 20)
	for x := range chOut {
		if x%8 == 0 {
			break
		}
		fmt.Println(x)
	}
	close(chStop)
	time.Sleep(1 * time.Second)
}
```

### 7. Preventing Resource Leak (Start + Stop)

**TODO: Write this!**

<kbd>[view][story7.go]</kbd>:

```go
// story7.go
func xrange(begin, end int) func() (<-chan int, chan<- struct{}) {
	return func() (<-chan int, chan<- struct{}) {
		chOut := make(chan int)
		chStop := make(chan struct{})
		go func() {
			fmt.Println("Goroutine Started")
			defer fmt.Println("Goroutine Terminated")
			defer close(chOut)
			for i := begin; i < end; i++ {
				select {
				case chOut <- i:
					// Nothing to be done
				case <-chStop:
					break
				}
			}
		}()
		return chOut, chStop
	}
}

func main() {
	f := xrange(10, 20)
	chOut, chStop := f()
	for x := range chOut {
		if x%8 == 0 {
			break
		}
		fmt.Println(x)
	}
	close(chStop)
	time.Sleep(1 * time.Second)
}
```

<kbd>[view][story7a.go]</kbd>:

```go
// story7a.go
func xrange(begin, end int) func() (<-chan int, func()) {
	return func() (<-chan int, func()) {
		chOut := make(chan int)
		chStop := make(chan struct{})
		fnStop := func() {
			close(chStop)
		}
		go func() {
			fmt.Println("Goroutine Started")
			defer fmt.Println("Goroutine Terminated")
			defer close(chOut)
			for i := begin; i < end; i++ {
				select {
				case chOut <- i:
					// Nothing to be done
				case <-chStop:
					break
				}
			}
		}()
		return chOut, fnStop
	}
}

func main() {
	fnStart := xrange(10, 20)
	chOut, fnStop := fnStart()
	for x := range chOut {
		if x%8 == 0 {
			break
		}
		fmt.Println(x)
	}
	fnStop()
	time.Sleep(1 * time.Second)
}
```

### 8. Packaging and Generalization

**TODO: Write this!**

<kbd>[view][story8.go]</kbd>:

```go
// story8.go
type XRange struct {
	begin int
	end   int
	stop  func()
}

func NewXRange(begin, end int) *XRange {
	return &XRange{
		begin: begin,
		end:   end,
	}
}

func (r *XRange) Start() <-chan int {
	chOut := make(chan int)
	chStop := make(chan struct{})
	r.stop = func() {
		close(chStop)
	}
	go func() {
		fmt.Println("Goroutine Started")
		defer fmt.Println("Goroutine Terminated")
		defer close(chOut)
		for i := r.begin; i < r.end; i++ {
			select {
			case chOut <- i:
				// Nothing to be done
			case <-chStop:
				break
			}
		}
	}()
	return chOut
}

func (r *XRange) Stop() {
	r.stop()
}

func main() {
	r := NewXRange(10, 20)
	c := r.Start()
	for x := range c {
		if x%8 == 0 {
			break
		}
		fmt.Println(x)
	}
	r.Stop()
	time.Sleep(1 * time.Second)
}
```

<kbd>[view][story8a.go]</kbd>:

```go
// story8a.go
type IntGenerator struct {
	impl func(chan<- int, <-chan struct{})
	stop func()
}

func NewIntGenerator(impl func(chan<- int, <-chan struct{})) *IntGenerator {
	return &IntGenerator{
		impl: impl,
	}
}

func (g *IntGenerator) Start() <-chan int {
	chOut := make(chan int)
	chStop := make(chan struct{})
	g.stop = func() {
		close(chStop)
	}
	go func() {
		fmt.Println("Goroutine Started")
		defer fmt.Println("Goroutine Terminated")
		defer close(chOut)
		g.impl(chOut, chStop)
	}()
	return chOut
}

func (g *IntGenerator) Stop() {
	g.stop()
}

func XRange(begin, end int) *IntGenerator {
	return NewIntGenerator(func(chOut chan<- int, chStop <-chan struct{}) {
		for i := begin; i < end; i++ {
			select {
			case chOut <- i:
				// Nothing to be done
			case <-chStop:
				break
			}
		}
	})
}

func main() {
	g := XRange(10, 20)
	c := g.Start()
	for x := range c {
		if x%8 == 0 {
			break
		}
		fmt.Println(x)
	}
	g.Stop()
	time.Sleep(1 * time.Second)
}
```

<kbd>[view][story8b.go]</kbd>:

```go
// story8b.go
type IntYielder func(int) bool

type IntGenerator struct {
	impl func(IntYielder)
	stop func()
}

func NewIntGenerator(impl func(IntYielder)) *IntGenerator {
	return &IntGenerator{
		impl: impl,
	}
}

func (g *IntGenerator) Start() <-chan int {
	chOut := make(chan int)
	chStop := make(chan struct{})
	g.stop = func() {
		close(chStop)
	}
	go func() {
		fmt.Println("Goroutine Started")
		defer fmt.Println("Goroutine Terminated")
		defer close(chOut)
		g.impl(IntYielder(func(msg int) bool {
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

func (g *IntGenerator) Stop() {
	g.stop()
}

func XRange(begin, end int) *IntGenerator {
	return NewIntGenerator(func(yield IntYielder) {
		for i := begin; i < end; i++ {
			if !yield(i) {
				break
			}
		}
	})
}

func main() {
	g := XRange(10, 20)
	c := g.Start()
	for x := range c {
		if x%8 == 0 {
			break
		}
		fmt.Println(x)
	}
	g.Stop()
	time.Sleep(1 * time.Second)
}
```




[story1.go]: story1.go
[story1a.go]: story1a.go
[story1b.go]: story1b.go
[story2.go]: story2.go
[story2a.go]: story2a.go
[story3.go]: story3.go
[story3a.go]: story3a.go
[story3b.go]: story3b.go
[story4.go]: story4.go
[story5.go]: story5.go
[story5a.go]: story5a.go
[story5b.go]: story5b.go
[story6.go]: story6.go
[story6a.go]: story6a.go
[story7.go]: story7.go
[story7a.go]: story7a.go
[story8.go]: story8.go
[story8a.go]: story8a.go
[story8b.go]: story8b.go
