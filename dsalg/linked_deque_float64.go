package dsalg

import (
	"fmt"
)

type float64DequeNode struct {
	next  *float64DequeNode
	prev  *float64DequeNode
	value float64
}

func (n *float64DequeNode) Reattach(next, prev *float64DequeNode, value float64) {
	// Detach from Non-Nil Next Node
	if n.next != nil {
		n.next.prev = n.prev
	}
	// Detach from Non-Nil Prev Node
	if n.prev != nil {
		n.prev.next = n.next
	}
	// Update Current Node
	n.next, n.prev, n.value = next, prev, value
}

type Float64LinkedDeque struct {
	fnBack      func() float64
	fnFront     func() float64
	fnPushBack  func(float64)
	fnPushFront func(float64)
	fnPopBack   func()
	fnPopFront  func()
	back        *float64DequeNode
	front       *float64DequeNode
	length      int
	capacity    int
}

// Warning: Does not check emptiness!
// Invoking this method on an empty queue leak to undefined behavior.
func (q *Float64LinkedDeque) backImpl() float64 {
	return q.back.value
}

// Warning: Does not check emptiness
// Invoking this method on an empty queue leak to undefined behavior.
func (q *Float64LinkedDeque) frontImpl() float64 {
	return q.front.value
}

func (q *Float64LinkedDeque) pushBackImpl(x float64) {
	// Empty Queue
	if q.length == 0 {
		q.back.value = x
		q.length++
		return
	}
	// Back Does Not Point to Edge Node
	if q.back.next != nil {
		q.back = q.back.next
		q.back.value = x
		q.length++
		return
	}
	// Back Points to Edge Node
	// Reuse a node from before the front or create a new one
	newBack := q.front.prev
	if newBack == nil {
		newBack = &float64DequeNode{nil, q.back, x}
		q.capacity++
	} else {
		newBack.Reattach(nil, q.back, x)
	}
	// Update Back Node
	q.back.next = newBack
	q.back = newBack
	q.length++
}

func (q *Float64LinkedDeque) pushFrontImpl(x float64) {
	// Empty Queue
	if q.length == 0 {
		q.front.value = x
		q.length++
		return
	}
	// Front Does Not Point to Edge Node
	if q.front.prev != nil {
		q.front = q.front.prev
		q.front.value = x
		q.length++
		return
	}
	// Front Points to Edge Node
	// Reuse a node from after the back or create a new one
	newFront := q.back.next
	if newFront == nil {
		newFront = &float64DequeNode{q.front, nil, x}
		q.capacity++
	} else {
		newFront.Reattach(q.front, nil, x)
	}
	// Update Front Node
	q.front.prev = newFront
	q.front = newFront
	q.length++
}

// Warning: Does not check emptiness
// Invoking this method on an empty queue leak to undefined behavior.
func (q *Float64LinkedDeque) popBackImpl() {
	// More-Than-One-Element Queue
	if q.back != q.front {
		q.back = q.back.prev
	}
	// Non-Empty Queue
	q.length--
}

// Warning: Does not check emptiness
// Invoking this method on an empty queue leak to undefined behavior.
func (q *Float64LinkedDeque) popFrontImpl() {
	// More-Than-One-Element Queue
	if q.front != q.back {
		q.front = q.front.next
	}
	// Non-Empty Queue
	q.length--
}

func NewFloat64LinkedDeque() *Float64LinkedDeque {
	node := &float64DequeNode{}
	q := &Float64LinkedDeque{
		back:     node,
		front:    node,
		capacity: 1,
	}
	q.fnBack, q.fnFront = q.backImpl, q.frontImpl
	q.fnPushBack, q.fnPushFront = q.pushBackImpl, q.pushFrontImpl
	q.fnPopBack, q.fnPopFront = q.popBackImpl, q.popFrontImpl
	return q
}

func (q *Float64LinkedDeque) Back() float64 {
	return q.fnBack()
}

func (q *Float64LinkedDeque) Front() float64 {
	return q.fnFront()
}

func (q *Float64LinkedDeque) PushBack(x float64) {
	q.fnPushBack(x)
}

func (q *Float64LinkedDeque) PushFront(x float64) {
	q.fnPushFront(x)
}

func (q *Float64LinkedDeque) PopBack() {
	q.fnPopBack()
}

func (q *Float64LinkedDeque) PopFront() {
	q.fnPopFront()
}

func (q *Float64LinkedDeque) Reverse() {
	q.fnBack, q.fnFront = q.fnFront, q.fnBack
	q.fnPushBack, q.fnPushFront = q.fnPushFront, q.fnPushBack
	q.fnPopBack, q.fnPopFront = q.fnPopFront, q.fnPopBack
}

func (q *Float64LinkedDeque) Len() int {
	return q.length
}

func (q *Float64LinkedDeque) Cap() int {
	return q.capacity
}

func (q *Float64LinkedDeque) IsEmpty() bool {
	return q.length == 0
}

func (q *Float64LinkedDeque) IsFull() bool {
	return false
}

func (q *Float64LinkedDeque) Print() {
	ptr := q.front
	if ptr != nil {
		for ptr.prev != nil {
			ptr = ptr.prev
		}
	}
	fmt.Print("<- ")
	for ptr != nil {
		fmt.Print(ptr.value)
		if ptr == q.front {
			fmt.Print("(F)")
		}
		if ptr == q.back {
			fmt.Print("(B)")
		}
		if ptr.next != nil {
			fmt.Print(" <-> ")
		}
		ptr = ptr.next
	}
	fmt.Println(" ->")
}
