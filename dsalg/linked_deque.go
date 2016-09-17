package dsalg

import (
	"bytes"
	"fmt"
)

type DequeNode struct {
	next, prev *DequeNode
	Value      interface{}
}

func (n *DequeNode) reattach(next, prev *DequeNode, v interface{}) {
	// Detach from Non-Nil Next Node
	if n.next != nil {
		n.next.prev = n.prev
	}
	// Detach from Non-Nil Prev Node
	if n.prev != nil {
		n.prev.next = n.next
	}
	// Update Current Node
	n.next, n.prev, n.Value = next, prev, v
}

type LinkedDeque struct {
	fnBack      func() *DequeNode
	fnFront     func() *DequeNode
	fnPushBack  func(interface{}) *DequeNode
	fnPushFront func(interface{}) *DequeNode
	fnPopBack   func() interface{}
	fnPopFront  func() interface{}
	back        *DequeNode
	front       *DequeNode
	len         int
	cap         int
}

func (q *LinkedDeque) backImpl() *DequeNode {
	if q.len == 0 {
		return nil
	}
	return q.back
}

func (q *LinkedDeque) frontImpl() *DequeNode {
	if q.len == 0 {
		return nil
	}
	return q.front
}

func (q *LinkedDeque) pushBackImpl(v interface{}) *DequeNode {
	// Empty Queue
	if q.len == 0 {
		q.back.Value = v
		q.len++
		return q.back
	}
	// Back Does Not Point to Edge Node
	if q.back.next != nil {
		q.back = q.back.next
		q.back.Value = v
		q.len++
		return q.back
	}
	// Back Points to Edge Node
	// Reuse a node from before the front or create a new one
	newBack := q.front.prev
	if newBack == nil {
		newBack = &DequeNode{nil, q.back, v}
		q.cap++
	} else {
		newBack.reattach(nil, q.back, v)
	}
	// Update Back Node
	q.back.next = newBack
	q.back = newBack
	q.len++
	return q.back
}

func (q *LinkedDeque) pushFrontImpl(v interface{}) *DequeNode {
	// Empty Queue
	if q.len == 0 {
		q.front.Value = v
		q.len++
		return q.front
	}
	// Front Does Not Point to Edge Node
	if q.front.prev != nil {
		q.front = q.front.prev
		q.front.Value = v
		q.len++
		return q.front
	}
	// Front Points to Edge Node
	// Reuse a node from after the back or create a new one
	newFront := q.back.next
	if newFront == nil {
		newFront = &DequeNode{q.front, nil, v}
		q.cap++
	} else {
		newFront.reattach(q.front, nil, v)
	}
	// Update Front Node
	q.front.prev = newFront
	q.front = newFront
	q.len++
	return q.front
}

func (q *LinkedDeque) popBackImpl() interface{} {
	// Empty Queue
	if q.len == 0 {
		panic("Empty queue")
	}
	n := q.back
	// More-Than-One-Element Queue
	if q.len > 1 {
		q.back = q.back.prev
	}
	// Non-Empty Queue
	q.len--
	return n.Value
}

func (q *LinkedDeque) popFrontImpl() interface{} {
	// Empty Queue
	if q.len == 0 {
		panic("Empty queue")
	}
	n := q.front
	// More-Than-One-Element Queue
	if q.len > 1 {
		q.front = q.front.next
	}
	// Non-Empty Queue
	q.len--
	return n.Value
}

func NewLinkedDeque() *LinkedDeque {
	n := &DequeNode{}
	q := &LinkedDeque{
		back:  n,
		front: n,
		cap:   1,
	}
	q.fnBack, q.fnFront = q.backImpl, q.frontImpl
	q.fnPushBack, q.fnPushFront = q.pushBackImpl, q.pushFrontImpl
	q.fnPopBack, q.fnPopFront = q.popBackImpl, q.popFrontImpl
	return q
}

func (q *LinkedDeque) Back() *DequeNode {
	return q.fnBack()
}

func (q *LinkedDeque) Front() *DequeNode {
	return q.fnFront()
}

func (q *LinkedDeque) PushBack(v interface{}) *DequeNode {
	return q.fnPushBack(v)
}

func (q *LinkedDeque) PushFront(v interface{}) *DequeNode {
	return q.fnPushFront(v)
}

func (q *LinkedDeque) PopBack() interface{} {
	return q.fnPopBack()
}

func (q *LinkedDeque) PopFront() interface{} {
	return q.fnPopFront()
}

func (q *LinkedDeque) BackValue() interface{} {
	return q.Back().Value
}

func (q *LinkedDeque) FrontValue() interface{} {
	return q.Front().Value
}

func (q *LinkedDeque) PushBackValue(v interface{}) {
	q.PushBack(v)
}

func (q *LinkedDeque) PushFrontValue(v interface{}) {
	q.PushFront(v)
}

func (q *LinkedDeque) Reverse() {
	q.fnBack, q.fnFront = q.fnFront, q.fnBack
	q.fnPushBack, q.fnPushFront = q.fnPushFront, q.fnPushBack
	q.fnPopBack, q.fnPopFront = q.fnPopFront, q.fnPopBack
}

func (q *LinkedDeque) Len() int {
	return q.len
}

func (q *LinkedDeque) Cap() int {
	return q.cap
}

func (q *LinkedDeque) IsEmpty() bool {
	return q.len == 0
}

func (q *LinkedDeque) IsFull() bool {
	return false
}

func (q *LinkedDeque) String() string {
	var buffer bytes.Buffer
	ptr := q.front
	if ptr != nil {
		for ptr.prev != nil {
			ptr = ptr.prev
		}
	}
	buffer.WriteString("<- ")
	for ptr != nil {
		buffer.WriteString(fmt.Sprintf("%v", ptr.Value))
		if ptr == q.front {
			buffer.WriteString("(F)")
		}
		if ptr == q.back {
			buffer.WriteString("(B)")
		}
		if ptr.next != nil {
			buffer.WriteString(" <-> ")
		}
		ptr = ptr.next
	}
	buffer.WriteString(" ->")
	return buffer.String()
}
