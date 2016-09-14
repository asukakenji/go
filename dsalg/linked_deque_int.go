package dsalg

import (
	"fmt"
)

type intDequeNode struct {
	next  *intDequeNode
	prev  *intDequeNode
	value int
}

func (n *intDequeNode) Reattach(next, prev *intDequeNode, value int) {
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

type IntLinkedDeque struct {
	fnBack      func() int
	fnFront     func() int
	fnPushBack  func(int)
	fnPushFront func(int)
	fnPopBack   func()
	fnPopFront  func()
	back        *intDequeNode
	front       *intDequeNode
	length      int
	capacity    int
}

// Warning: Does not check emptiness!
// Invoking this method on an empty queue leak to undefined behavior.
func (q *IntLinkedDeque) backImpl() int {
	return q.back.value
}

// Warning: Does not check emptiness
// Invoking this method on an empty queue leak to undefined behavior.
func (q *IntLinkedDeque) frontImpl() int {
	return q.front.value
}

func (q *IntLinkedDeque) pushBackImpl(x int) {
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
		newBack = &intDequeNode{nil, q.back, x}
		q.capacity++
	} else {
		newBack.Reattach(nil, q.back, x)
	}
	// Update Back Node
	q.back.next = newBack
	q.back = newBack
	q.length++
}

func (q *IntLinkedDeque) pushFrontImpl(x int) {
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
		newFront = &intDequeNode{q.front, nil, x}
		q.capacity++
	} else {
		newFront.Reattach(q.front, nil, x)
	}
	// Update Front Node
	q.front.prev = newFront
	q.front = newFront
	q.length++
}

func (q *IntLinkedDeque) popBackImpl() {
	// Empty Queue
	if q.length == 0 {
		panic("Empty queue")
	}
	// More-Than-One-Element Queue
	if q.length > 1 {
		q.back = q.back.prev
	}
	// Non-Empty Queue
	q.length--
}

func (q *IntLinkedDeque) popFrontImpl() {
	// Empty Queue
	if q.length == 0 {
		panic("Empty queue")
	}
	// More-Than-One-Element Queue
	if q.length > 1 {
		q.front = q.front.next
	}
	// Non-Empty Queue
	q.length--
}

func NewIntLinkedDeque() *IntLinkedDeque {
	node := &intDequeNode{}
	q := &IntLinkedDeque{
		back:     node,
		front:    node,
		capacity: 1,
	}
	q.fnBack, q.fnFront = q.backImpl, q.frontImpl
	q.fnPushBack, q.fnPushFront = q.pushBackImpl, q.pushFrontImpl
	q.fnPopBack, q.fnPopFront = q.popBackImpl, q.popFrontImpl
	return q
}

func (q *IntLinkedDeque) Back() int {
	return q.fnBack()
}

func (q *IntLinkedDeque) Front() int {
	return q.fnFront()
}

func (q *IntLinkedDeque) PushBack(x int) {
	q.fnPushBack(x)
}

func (q *IntLinkedDeque) PushFront(x int) {
	q.fnPushFront(x)
}

func (q *IntLinkedDeque) PopBack() {
	q.fnPopBack()
}

func (q *IntLinkedDeque) PopFront() {
	q.fnPopFront()
}

func (q *IntLinkedDeque) Reverse() {
	q.fnBack, q.fnFront = q.fnFront, q.fnBack
	q.fnPushBack, q.fnPushFront = q.fnPushFront, q.fnPushBack
	q.fnPopBack, q.fnPopFront = q.fnPopFront, q.fnPopBack
}

func (q *IntLinkedDeque) Len() int {
	return q.length
}

func (q *IntLinkedDeque) Cap() int {
	return q.capacity
}

func (q *IntLinkedDeque) IsEmpty() bool {
	return q.length == 0
}

func (q *IntLinkedDeque) IsFull() bool {
	return false
}

func (q *IntLinkedDeque) Print() {
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
