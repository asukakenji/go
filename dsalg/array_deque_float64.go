package dsalg

import (
	"fmt"
)

type Float64ArrayDeque struct {
	fnBack      func() float64
	fnFront     func() float64
	fnPushBack  func(float64)
	fnPushFront func(float64)
	fnPopBack   func()
	fnPopFront  func()
	values      []float64
	back        int
	front       int
	length      int
	capacity    int
}

// Warning: Does not check emptiness
// Invoking this method on an empty queue leak to undefined behavior.
func (q *Float64ArrayDeque) backImpl() float64 {
	return q.values[q.back]
}

// Warning: Does not check emptiness
// Invoking this method on an empty queue leak to undefined behavior.
func (q *Float64ArrayDeque) frontImpl() float64 {
	return q.values[q.front]
}

func (q *Float64ArrayDeque) pushBackImpl(x float64) {
	if q.length == q.capacity {
		panic("Full queue")
	}
	if q.length != 0 {
		q.back++
		if q.back == q.capacity {
			q.back = 0
		}
	}
	q.values[q.back] = x
	q.length++
}

func (q *Float64ArrayDeque) pushFrontImpl(x float64) {
	if q.length == q.capacity {
		panic("Full queue")
	}
	if q.length != 0 {
		q.front--
		if q.front == -1 {
			q.front = q.capacity - 1
		}
	}
	q.values[q.front] = x
	q.length++
}

func (q *Float64ArrayDeque) popBackImpl() {
	if q.length == 0 {
		panic("Empty queue")
	}
	// More-Than-One-Element Queue
	if q.length > 1 {
		q.back--
		if q.back == -1 {
			q.back = q.capacity - 1
		}
	}
	// Non-Empty Queue
	q.length--
}

func (q *Float64ArrayDeque) popFrontImpl() {
	if q.length == 0 {
		panic("Empty queue")
	}
	// More-Than-One-Element Queue
	if q.length > 1 {
		q.front++
		if q.front == q.capacity {
			q.front = 0
		}
	}
	// Non-Empty Queue
	q.length--
}

func NewFloat64ArrayDeque(size int) *Float64ArrayDeque {
	q := &Float64ArrayDeque{
		values:   make([]float64, size),
		capacity: size,
	}
	q.fnBack, q.fnFront = q.backImpl, q.frontImpl
	q.fnPushBack, q.fnPushFront = q.pushBackImpl, q.pushFrontImpl
	q.fnPopBack, q.fnPopFront = q.popBackImpl, q.popFrontImpl
	return q
}

func (q *Float64ArrayDeque) Back() float64 {
	return q.fnBack()
}

func (q *Float64ArrayDeque) Front() float64 {
	return q.fnFront()
}

func (q *Float64ArrayDeque) PushBack(x float64) {
	q.fnPushBack(x)
}

func (q *Float64ArrayDeque) PushFront(x float64) {
	q.fnPushFront(x)
}

func (q *Float64ArrayDeque) PopBack() {
	q.fnPopBack()
}

func (q *Float64ArrayDeque) PopFront() {
	q.fnPopFront()
}

func (q *Float64ArrayDeque) Reverse() {
	q.fnBack, q.fnFront = q.fnFront, q.fnBack
	q.fnPushBack, q.fnPushFront = q.fnPushFront, q.fnPushBack
	q.fnPopBack, q.fnPopFront = q.fnPopFront, q.fnPopBack
}

func (q *Float64ArrayDeque) Len() int {
	return q.length
}

func (q *Float64ArrayDeque) Cap() int {
	return q.capacity
}

func (q *Float64ArrayDeque) IsEmpty() bool {
	return q.length == 0
}

func (q *Float64ArrayDeque) IsFull() bool {
	return q.length == q.capacity
}

func (q *Float64ArrayDeque) Print() {
	sep := ""
	fmt.Print("[")
	for i := range q.values {
		fmt.Print(sep, q.values[i])
		if i == q.front {
			fmt.Print("(F)")
		}
		if i == q.back {
			fmt.Print("(B)")
		}
		sep = " "
	}
	fmt.Println("]")
}
