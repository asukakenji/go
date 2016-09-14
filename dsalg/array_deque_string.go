package dsalg

import (
	"fmt"
)

type StringArrayDeque struct {
	fnBack      func() string
	fnFront     func() string
	fnPushBack  func(string)
	fnPushFront func(string)
	fnPopBack   func()
	fnPopFront  func()
	values      []string
	back        int
	front       int
	length      int
	capacity    int
}

// Warning: Does not check emptiness
// Invoking this method on an empty queue leak to undefined behavior.
func (q *StringArrayDeque) backImpl() string {
	return q.values[q.back]
}

// Warning: Does not check emptiness
// Invoking this method on an empty queue leak to undefined behavior.
func (q *StringArrayDeque) frontImpl() string {
	return q.values[q.front]
}

func (q *StringArrayDeque) pushBackImpl(x string) {
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

func (q *StringArrayDeque) pushFrontImpl(x string) {
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

func (q *StringArrayDeque) popBackImpl() {
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

func (q *StringArrayDeque) popFrontImpl() {
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

func NewStringArrayDeque(size int) *StringArrayDeque {
	q := &StringArrayDeque{
		values:   make([]string, size),
		capacity: size,
	}
	q.fnBack, q.fnFront = q.backImpl, q.frontImpl
	q.fnPushBack, q.fnPushFront = q.pushBackImpl, q.pushFrontImpl
	q.fnPopBack, q.fnPopFront = q.popBackImpl, q.popFrontImpl
	return q
}

func (q *StringArrayDeque) Back() string {
	return q.fnBack()
}

func (q *StringArrayDeque) Front() string {
	return q.fnFront()
}

func (q *StringArrayDeque) PushBack(x string) {
	q.fnPushBack(x)
}

func (q *StringArrayDeque) PushFront(x string) {
	q.fnPushFront(x)
}

func (q *StringArrayDeque) PopBack() {
	q.fnPopBack()
}

func (q *StringArrayDeque) PopFront() {
	q.fnPopFront()
}

func (q *StringArrayDeque) Reverse() {
	q.fnBack, q.fnFront = q.fnFront, q.fnBack
	q.fnPushBack, q.fnPushFront = q.fnPushFront, q.fnPushBack
	q.fnPopBack, q.fnPopFront = q.fnPopFront, q.fnPopBack
}

func (q *StringArrayDeque) Len() int {
	return q.length
}

func (q *StringArrayDeque) Cap() int {
	return q.capacity
}

func (q *StringArrayDeque) IsEmpty() bool {
	return q.length == 0
}

func (q *StringArrayDeque) IsFull() bool {
	return q.length == q.capacity
}

func (q *StringArrayDeque) Print() {
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
