package dsalg

import (
	"fmt"
)

type ArrayDeque struct {
	fnBack      func() interface{}
	fnFront     func() interface{}
	fnPushBack  func(interface{})
	fnPushFront func(interface{})
	fnPopBack   func()
	fnPopFront  func()
	values      []interface{}
	back        int
	front       int
	length      int
	capacity    int
}

// Warning: Does not check emptiness
// Invoking this method on an empty queue leak to undefined behavior.
func (q *ArrayDeque) backImpl() interface{} {
	return q.values[q.back]
}

// Warning: Does not check emptiness
// Invoking this method on an empty queue leak to undefined behavior.
func (q *ArrayDeque) frontImpl() interface{} {
	return q.values[q.front]
}

func (q *ArrayDeque) pushBackImpl(x interface{}) {
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

func (q *ArrayDeque) pushFrontImpl(x interface{}) {
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

func (q *ArrayDeque) popBackImpl() {
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

func (q *ArrayDeque) popFrontImpl() {
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

func NewArrayDeque(size int) *ArrayDeque {
	q := &ArrayDeque{
		values:   make([]interface{}, size),
		capacity: size,
	}
	q.fnBack, q.fnFront = q.backImpl, q.frontImpl
	q.fnPushBack, q.fnPushFront = q.pushBackImpl, q.pushFrontImpl
	q.fnPopBack, q.fnPopFront = q.popBackImpl, q.popFrontImpl
	return q
}

func (q *ArrayDeque) Back() interface{} {
	return q.fnBack()
}

func (q *ArrayDeque) Front() interface{} {
	return q.fnFront()
}

func (q *ArrayDeque) PushBack(x interface{}) {
	q.fnPushBack(x)
}

func (q *ArrayDeque) PushFront(x interface{}) {
	q.fnPushFront(x)
}

func (q *ArrayDeque) PopBack() {
	q.fnPopBack()
}

func (q *ArrayDeque) PopFront() {
	q.fnPopFront()
}

func (q *ArrayDeque) Reverse() {
	q.fnBack, q.fnFront = q.fnFront, q.fnBack
	q.fnPushBack, q.fnPushFront = q.fnPushFront, q.fnPushBack
	q.fnPopBack, q.fnPopFront = q.fnPopFront, q.fnPopBack
}

func (q *ArrayDeque) Len() int {
	return q.length
}

func (q *ArrayDeque) Cap() int {
	return q.capacity
}

func (q *ArrayDeque) IsEmpty() bool {
	return q.length == 0
}

func (q *ArrayDeque) IsFull() bool {
	return q.length == q.capacity
}

func (q *ArrayDeque) Print() {
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
