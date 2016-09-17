package dsalg

import (
	"bytes"
	"fmt"
)

type ArrayDeque struct {
	fnBack      func() int
	fnFront     func() int
	fnPushBack  func(interface{}) int
	fnPushFront func(interface{}) int
	fnPopBack   func() interface{}
	fnPopFront  func() interface{}
	values      []interface{}
	back        int
	front       int
	len         int
	cap         int
}

func (q *ArrayDeque) backImpl() int {
	if q.len == 0 {
		return -1
	}
	return q.back
}

func (q *ArrayDeque) frontImpl() int {
	if q.len == 0 {
		return -1
	}
	return q.front
}

func (q *ArrayDeque) pushBackImpl(v interface{}) int {
	// Full Queue
	if q.len == q.cap {
		panic("Full queue")
	}
	// Non-Empty Queue
	if q.len != 0 {
		q.back++
		if q.back == q.cap {
			q.back = 0
		}
	}
	// Non-Full Queue
	q.values[q.back] = v
	q.len++
	return q.back
}

func (q *ArrayDeque) pushFrontImpl(v interface{}) int {
	// Full Queue
	if q.len == q.cap {
		panic("Full queue")
	}
	// Non-Empty Queue
	if q.len != 0 {
		q.front--
		if q.front == -1 {
			q.front = q.cap - 1
		}
	}
	// Non-Full Queue
	q.values[q.front] = v
	q.len++
	return q.front
}

func (q *ArrayDeque) popBackImpl() interface{} {
	// Empty Queue
	if q.len == 0 {
		panic("Empty queue")
	}
	// More-Than-One-Element Queue
	i := q.back
	if q.len > 1 {
		q.back--
		if q.back == -1 {
			q.back = q.cap - 1
		}
	}
	// Non-Empty Queue
	q.len--
	return q.values[i]
}

func (q *ArrayDeque) popFrontImpl() interface{} {
	// Empty Queue
	if q.len == 0 {
		panic("Empty queue")
	}
	// More-Than-One-Element Queue
	i := q.front
	if q.len > 1 {
		q.front++
		if q.front == q.cap {
			q.front = 0
		}
	}
	// Non-Empty Queue
	q.len--
	return q.values[i]
}

func NewArrayDeque(size int) *ArrayDeque {
	q := &ArrayDeque{
		values: make([]interface{}, size),
		cap:    size,
	}
	q.fnBack, q.fnFront = q.backImpl, q.frontImpl
	q.fnPushBack, q.fnPushFront = q.pushBackImpl, q.pushFrontImpl
	q.fnPopBack, q.fnPopFront = q.popBackImpl, q.popFrontImpl
	return q
}

func (q *ArrayDeque) Back() int {
	return q.fnBack()
}

func (q *ArrayDeque) Front() int {
	return q.fnFront()
}

func (q *ArrayDeque) PushBack(v interface{}) int {
	return q.fnPushBack(v)
}

func (q *ArrayDeque) PushFront(v interface{}) int {
	return q.fnPushFront(v)
}

func (q *ArrayDeque) PopBack() interface{} {
	return q.fnPopBack()
}

func (q *ArrayDeque) PopFront() interface{} {
	return q.fnPopFront()
}

func (q *ArrayDeque) BackValue() interface{} {
	return q.values[q.Back()]
}

func (q *ArrayDeque) FrontValue() interface{} {
	return q.values[q.Front()]
}

func (q *ArrayDeque) PushBackValue(v interface{}) {
	q.PushBack(v)
}

func (q *ArrayDeque) PushFrontValue(v interface{}) {
	q.PushFront(v)
}

func (q *ArrayDeque) Reverse() {
	q.fnBack, q.fnFront = q.fnFront, q.fnBack
	q.fnPushBack, q.fnPushFront = q.fnPushFront, q.fnPushBack
	q.fnPopBack, q.fnPopFront = q.fnPopFront, q.fnPopBack
}

func (q *ArrayDeque) Len() int {
	return q.len
}

func (q *ArrayDeque) Cap() int {
	return q.cap
}

func (q *ArrayDeque) IsEmpty() bool {
	return q.len == 0
}

func (q *ArrayDeque) IsFull() bool {
	return q.len == q.cap
}

func (q *ArrayDeque) String() string {
	var buffer bytes.Buffer
	sep := ""
	buffer.WriteString("[")
	for i := range q.values {
		buffer.WriteString(sep)
		buffer.WriteString(fmt.Sprintf("%v", q.values[i]))
		if i == q.front {
			buffer.WriteString("(F)")
		}
		if i == q.back {
			buffer.WriteString("(B)")
		}
		sep = " "
	}
	buffer.WriteString("]")
	return buffer.String()
}
