package dsalg

import (
	"fmt"
)

type stringDequeNode struct {
	next  *stringDequeNode
	prev  *stringDequeNode
	value string
}

func (n *stringDequeNode) Reattach(next, prev *stringDequeNode, value string) {
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

type StringLinkedDeque struct {
	fnBack      func() string
	fnFront     func() string
	fnPushBack  func(string)
	fnPushFront func(string)
	fnPopBack   func()
	fnPopFront  func()
	back        *stringDequeNode
	front       *stringDequeNode
	length      int
	capacity    int
}

func NewStringLinkedDeque() *StringLinkedDeque {
	node := &stringDequeNode{}
	q := &StringLinkedDeque{
		back:     node,
		front:    node,
		capacity: 1,
	}
	q.fnBack, q.fnFront = q.backImpl, q.frontImpl
	q.fnPushBack, q.fnPushFront = q.pushBackImpl, q.pushFrontImpl
	q.fnPopBack, q.fnPopFront = q.popBackImpl, q.popFrontImpl
	return q
}

// Warning: Does not check emptiness!
// Invoking this method on an empty queue leak to undefined behavior.
func (q *StringLinkedDeque) backImpl() string {
	return q.back.value
}

// Warning: Does not check emptiness
// Invoking this method on an empty queue leak to undefined behavior.
func (q *StringLinkedDeque) frontImpl() string {
	return q.front.value
}

func (q *StringLinkedDeque) pushBackImpl(x string) {
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
		newBack = &stringDequeNode{nil, q.back, x}
		q.capacity++
	} else {
		newBack.Reattach(nil, q.back, x)
	}
	// Update Back Node
	q.back.next = newBack
	q.back = newBack
	q.length++
}

func (q *StringLinkedDeque) pushFrontImpl(x string) {
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
		newFront = &stringDequeNode{q.front, nil, x}
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
func (q *StringLinkedDeque) popBackImpl() {
	// More-Than-One-Element Queue
	if q.back != q.front {
		q.back = q.back.prev
	}
	// Non-Empty Queue
	q.length--
}

// Warning: Does not check emptiness
// Invoking this method on an empty queue leak to undefined behavior.
func (q *StringLinkedDeque) popFrontImpl() {
	// More-Than-One-Element Queue
	if q.front != q.back {
		q.front = q.front.next
	}
	// Non-Empty Queue
	q.length--
}

func (q *StringLinkedDeque) Back() string {
	return q.fnBack()
}

func (q *StringLinkedDeque) Front() string {
	return q.fnFront()
}

func (q *StringLinkedDeque) PushBack(x string) {
	q.fnPushBack(x)
}

func (q *StringLinkedDeque) PushFront(x string) {
	q.fnPushFront(x)
}

func (q *StringLinkedDeque) PopBack() {
	q.fnPopBack()
}

func (q *StringLinkedDeque) PopFront() {
	q.fnPopFront()
}

func (q *StringLinkedDeque) Reverse() {
	q.fnBack, q.fnFront = q.fnFront, q.fnBack
	q.fnPushBack, q.fnPushFront = q.fnPushFront, q.fnPushBack
	q.fnPopBack, q.fnPopFront = q.fnPopFront, q.fnPopBack
}

func (q *StringLinkedDeque) Len() int {
	return q.length
}

func (q *StringLinkedDeque) IsEmpty() bool {
	return q.length == 0
}

func (q *StringLinkedDeque) Cap() int {
	return q.capacity
}

func (q *StringLinkedDeque) Print() {
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
