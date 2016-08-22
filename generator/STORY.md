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

This reminds me of this question I asked when I began developing in Go (Golang):

[The idiomatic way to implement generators (yield) in Golang for recursive functions][so2]

And this question asked by another developer:

[Python-style generators in Go][so3]

Since none of the answers to the above questions fulfill my requirements,
I decided roll my own implementation.

[so1]: http://stackoverflow.com/q/37648288/142239
[so2]: http://stackoverflow.com/q/34464146/142239
[so3]: http://stackoverflow.com/q/11385556/142239

## Requirements

1. The underlying mechanism of the generator should be encapsulated by the library.
   The [RAII][raii] principle should be followed.
2. The generator should return a value assignable to a variable.
   If the value is not safe to be discarded because resource is acquired,
   the library should provide an API the release the resource.
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
the creation of the channel and the goroutine are handled by the library
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
Since we don't have a way to "destroy" the goroutine,
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
(which could also be a closure, or a function reference) as parameter.
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

At this point, our requirements are all fulfilled. Are we done? No, of course!
What we just developed is a function that accepts a callback,
which is very different from what a generator is.

The callback approach hard wires the control flow for you.
Your callback functions fill in the holes in the library code.
This style is more popular in (and appropriate for) things like
"providing a comparison callback to a sorting function",
or "providing a event handler callback to a event-driven framework".

On the other hand, the generator approach let you control the flow.
You may treat the generator is a source of objects,
an object is taken from the source each time it is invoked.

### 4. Goroutine-Channel Approach - Wrong Solution

**TODO: Write this!**

### 5. Preventing Resource Leak (start only when I ask so)

**TODO: Write this!**

### 6. Preventing Resource Leak (stop when I say so)

**TODO: Write this!**

### 7. Preventing Resource Leak (start + stop)

**TODO: Write this!**

### 8. Packaging and Generalization



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
