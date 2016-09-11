package dsalg

import (
	"testing"
)

func checkLC(t *testing.T, q *LinkedDeque, expectedL, expectedC int) {
    l := q.Len()
	if l != expectedL {
		t.Errorf("q.Len() == %d; expected %d", l, expectedL)
	}
    c := q.Cap()
	if c != expectedC {
		t.Errorf("q.Cap() == %d; expected %d", c, expectedC)
	}
}

func checkFBLC(t *testing.T, q *LinkedDeque, expectedF, expectedB, expectedL, expectedC int) {
    f := q.Front().(int)
	if f != expectedF {
		t.Errorf("q.Front() == %d; expected %d", f, expectedF)
	}
	b := q.Back().(int)
	if b != expectedB {
		t.Errorf("q.Back() == %d; expected %d", b, expectedB)
	}
    checkLC(t, q, expectedL, expectedC)
}

func TestNewLinkedDeque(t *testing.T) {
	q := NewLinkedDeque()

	if q == nil {
		t.Errorf("Returned nil")
	}
    checkLC(t, q, 0, 1)
	isEmpty := q.IsEmpty()
	if !isEmpty {
		t.Errorf("q.IsEmpty() == false; expected true")
	}
}

func TestPushBack(t *testing.T) {
	q := NewLinkedDeque()

	q.PushBack(42)
    checkFBLC(t, q, 42, 42, 1, 1)

	q.PushBack(14)
    checkFBLC(t, q, 42, 14, 2, 2)
}

func TestPushFront(t *testing.T) {
	q := NewLinkedDeque()

	q.PushFront(42)
    checkFBLC(t, q, 42, 42, 1, 1)

	q.PushFront(14)
    checkFBLC(t, q, 14, 42, 2, 2)
}

func TestPushBackAndPushFront0(t *testing.T) {
	q := NewLinkedDeque()

	q.PushBack(42)
	q.PushFront(14)
    checkFBLC(t, q, 14, 42, 2, 2)

    q.PushBack(49)
    checkFBLC(t, q, 14, 49, 3, 3)
}

func TestPushBackAndPushFront1(t *testing.T) {
	q := NewLinkedDeque()

	q.PushFront(42)
	q.PushBack(14)
    checkFBLC(t, q, 42, 14, 2, 2)

    q.PushFront(49)
    checkFBLC(t, q, 49, 14, 3, 3)
}

func TestPopBack(t *testing.T) {
	q := NewLinkedDeque()

	q.PushBack(91)
	q.PopBack()
    checkLC(t, q, 0, 1)

	q.PushBack(92)
    checkFBLC(t, q, 92, 92, 1, 1)

    q.PushBack(93)
    checkFBLC(t, q, 92, 93, 2, 2)

	q.PopBack()
    checkFBLC(t, q, 92, 92, 1, 2)

    q.PopBack()
    checkLC(t, q, 0, 2)

    q.PushBack(94)
    checkFBLC(t, q, 94, 94, 1, 2)

    q.PushBack(95)
    checkFBLC(t, q, 94, 95, 2, 2)

    q.PushBack(96)
    checkFBLC(t, q, 94, 96, 3, 3)

    q.PopBack()
    checkFBLC(t, q, 94, 95, 2, 3)

    q.PopBack()
    checkFBLC(t, q, 94, 94, 1, 3)

    q.PopBack()
    checkLC(t, q, 0, 3)
}

func TestPopFront(t *testing.T) {
	q := NewLinkedDeque()

	q.PushFront(91)
	q.PopFront()
    checkLC(t, q, 0, 1)

	q.PushFront(92)
    checkFBLC(t, q, 92, 92, 1, 1)

    q.PushFront(93)
    checkFBLC(t, q, 93, 92, 2, 2)

	q.PopFront()
    checkFBLC(t, q, 92, 92, 1, 2)

    q.PopFront()
    checkLC(t, q, 0, 2)

    q.PushFront(94)
    checkFBLC(t, q, 94, 94, 1, 2)

    q.PushFront(95)
    checkFBLC(t, q, 95, 94, 2, 2)

    q.PushFront(96)
    checkFBLC(t, q, 96, 94, 3, 3)

    q.PopFront()
    checkFBLC(t, q, 95, 94, 2, 3)

    q.PopFront()
    checkFBLC(t, q, 94, 94, 1, 3)

    q.PopFront()
    checkLC(t, q, 0, 3)
}

func TestPopBackAndPopFront0(t *testing.T) {
    q := NewLinkedDeque()

	q.PushBack(91)
    q.PushBack(92)
    q.PushBack(93)
    q.PopFront()
	checkFBLC(t, q, 92, 93, 2, 3)

    q.PushBack(94)
	checkFBLC(t, q, 92, 94, 3, 3)

    q.PopBack()
	checkFBLC(t, q, 92, 93, 2, 3)

    q.PushFront(95)
	checkFBLC(t, q, 95, 93, 3, 3)
}

func TestPopBackAndPopFront1(t *testing.T) {
    q := NewLinkedDeque()

	q.PushBack(91)
    q.PushBack(92)
    q.PushBack(93)
    q.PopFront()
    q.PopFront()
	checkFBLC(t, q, 93, 93, 1, 3)

    q.PushBack(94)
	checkFBLC(t, q, 93, 94, 2, 3)

    q.PushBack(95)
	checkFBLC(t, q, 93, 95, 3, 3)

    q.PopBack()
    q.PopBack()
    q.PushFront(96)
	checkFBLC(t, q, 96, 93, 2, 3)

    q.PushFront(97)
	checkFBLC(t, q, 97, 93, 3, 3)
}

func TestPopBackAndPopFront2(t *testing.T) {
    q := NewLinkedDeque()

	q.PushBack(91)
    q.PushBack(92)
    q.PushBack(93)
    q.PopFront()
    q.PopFront()
    q.PopFront()
	checkLC(t, q, 0, 3)

    q.PushBack(94)
	checkFBLC(t, q, 94, 94, 1, 3)

    q.PushBack(95)
	checkFBLC(t, q, 94, 95, 2, 3)

    q.PushBack(96)
	checkFBLC(t, q, 94, 96, 3, 3)

    q.PopBack()
    q.PopBack()
    q.PopBack()
    q.PushFront(97)
	checkFBLC(t, q, 97, 97, 1, 3)

    q.PushFront(98)
	checkFBLC(t, q, 98, 97, 2, 3)

    q.PushFront(99)
	checkFBLC(t, q, 99, 97, 3, 3)
}
