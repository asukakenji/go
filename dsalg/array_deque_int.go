package dsalg

import (
	"fmt"
)

type IntArrayDeque struct {
	fnBack      func() int
	fnFront     func() int
	fnPushBack  func(int)
	fnPushFront func(int)
	fnPopBack   func()
	fnPopFront  func()
	values      []int
	back        int
	front       int
	length      int
	capacity    int
}

// Warning: Does not check emptiness
// Invoking this method on an empty queue leak to undefined behavior.
func (q *IntArrayDeque) backImpl() int {
	return q.values[q.back]
}

// Warning: Does not check emptiness
// Invoking this method on an empty queue leak to undefined behavior.
func (q *IntArrayDeque) frontImpl() int {
	return q.values[q.front]
}

func (q *IntArrayDeque) pushBackImpl(x int) {
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

func (q *IntArrayDeque) pushFrontImpl(x int) {
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

func (q *IntArrayDeque) popBackImpl() {
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

func (q *IntArrayDeque) popFrontImpl() {
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

func NewIntArrayDeque(size int) *IntArrayDeque {
	q := &IntArrayDeque{
		values:   make([]int, size),
		capacity: size,
	}
	q.fnBack, q.fnFront = q.backImpl, q.frontImpl
	q.fnPushBack, q.fnPushFront = q.pushBackImpl, q.pushFrontImpl
	q.fnPopBack, q.fnPopFront = q.popBackImpl, q.popFrontImpl
	return q
}

func (q *IntArrayDeque) Back() int {
	return q.fnBack()
}

func (q *IntArrayDeque) Front() int {
	return q.fnFront()
}

func (q *IntArrayDeque) PushBack(x int) {
	q.fnPushBack(x)
}

func (q *IntArrayDeque) PushFront(x int) {
	q.fnPushFront(x)
}

func (q *IntArrayDeque) PopBack() {
	q.fnPopBack()
}

func (q *IntArrayDeque) PopFront() {
	q.fnPopFront()
}

func (q *IntArrayDeque) Reverse() {
	q.fnBack, q.fnFront = q.fnFront, q.fnBack
	q.fnPushBack, q.fnPushFront = q.fnPushFront, q.fnPushBack
	q.fnPopBack, q.fnPopFront = q.fnPopFront, q.fnPopBack
}

func (q *IntArrayDeque) Len() int {
	return q.length
}

func (q *IntArrayDeque) Cap() int {
	return q.capacity
}

func (q *IntArrayDeque) IsEmpty() bool {
	return q.length == 0
}

func (q *IntArrayDeque) IsFull() bool {
	return q.length == q.capacity
}

func (q *IntArrayDeque) Print() {
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
