package dsalg

import (
	"testing"
)

type Deque interface {
	BackValue() interface{}
	FrontValue() interface{}
	PushBackValue(v interface{})
	PushFrontValue(v interface{})
	PopBack() interface{}
	PopFront() interface{}
	Reverse()
	Len() int
	Cap() int
	IsEmpty() bool
	IsFull() bool
	String() string
}

func checkLC(t *testing.T, q Deque, expectedL, expectedC int, isCapChecked bool) {
	l := q.Len()
	if l != expectedL {
		t.Errorf("q.Len() == %d; expected %d", l, expectedL)
	}
	if isCapChecked {
		c := q.Cap()
		if c != expectedC {
			t.Errorf("q.Cap() == %d; expected %d", c, expectedC)
		}
	}
}

func checkFBLC(t *testing.T, q Deque, expectedF, expectedB, expectedL, expectedC int, isCapChecked bool) {
	f := q.FrontValue()
	if f != expectedF {
		t.Errorf("q.Front() == %d; expected %d", f, expectedF)
	}
	b := q.BackValue()
	if b != expectedB {
		t.Errorf("q.Back() == %d; expected %d", b, expectedB)
	}
	checkLC(t, q, expectedL, expectedC, isCapChecked)
}

func newArrayDeque() Deque {
	return NewArrayDeque(3)
}

func newLinkedDeque() Deque {
	return NewLinkedDeque()
}

func testNewDeque(t *testing.T, newDeque func() Deque, isCapChecked bool) {
	q := newDeque()

	if q == nil {
		t.Errorf("Returned nil")
	}
	checkLC(t, q, 0, 1, isCapChecked)
	isEmpty := q.IsEmpty()
	if !isEmpty {
		t.Errorf("q.IsEmpty() == false; expected true")
	}
	isFull := q.IsFull()
	if isFull {
		t.Errorf("q.IsFull() == true; expected false")
	}
}

func testPushBack(t *testing.T, newDeque func() Deque, isCapChecked bool) {
	q := newDeque()

	q.PushBackValue(42)
	checkFBLC(t, q, 42, 42, 1, 1, isCapChecked)

	q.PushBackValue(14)
	checkFBLC(t, q, 42, 14, 2, 2, isCapChecked)
}

func testPushFront(t *testing.T, newDeque func() Deque, isCapChecked bool) {
	q := newDeque()

	q.PushFrontValue(42)
	checkFBLC(t, q, 42, 42, 1, 1, isCapChecked)

	q.PushFrontValue(14)
	checkFBLC(t, q, 14, 42, 2, 2, isCapChecked)
}

func testPushBackAndPushFront0(t *testing.T, newDeque func() Deque, isCapChecked bool) {
	q := newDeque()

	q.PushBackValue(42)
	q.PushFrontValue(14)
	checkFBLC(t, q, 14, 42, 2, 2, isCapChecked)

	q.PushBackValue(49)
	checkFBLC(t, q, 14, 49, 3, 3, isCapChecked)
}

func testPushBackAndPushFront1(t *testing.T, newDeque func() Deque, isCapChecked bool) {
	q := newDeque()

	q.PushFrontValue(42)
	q.PushBackValue(14)
	checkFBLC(t, q, 42, 14, 2, 2, isCapChecked)

	q.PushFrontValue(49)
	checkFBLC(t, q, 49, 14, 3, 3, isCapChecked)
}

func testPopBack(t *testing.T, newDeque func() Deque, isCapChecked bool) {
	q := newDeque()

	q.PushBackValue(91)
	q.PopBack()
	checkLC(t, q, 0, 1, isCapChecked)

	q.PushBackValue(92)
	checkFBLC(t, q, 92, 92, 1, 1, isCapChecked)

	q.PushBackValue(93)
	checkFBLC(t, q, 92, 93, 2, 2, isCapChecked)

	q.PopBack()
	checkFBLC(t, q, 92, 92, 1, 2, isCapChecked)

	q.PopBack()
	checkLC(t, q, 0, 2, isCapChecked)

	q.PushBackValue(94)
	checkFBLC(t, q, 94, 94, 1, 2, isCapChecked)

	q.PushBackValue(95)
	checkFBLC(t, q, 94, 95, 2, 2, isCapChecked)

	q.PushBackValue(96)
	checkFBLC(t, q, 94, 96, 3, 3, isCapChecked)

	q.PopBack()
	checkFBLC(t, q, 94, 95, 2, 3, isCapChecked)

	q.PopBack()
	checkFBLC(t, q, 94, 94, 1, 3, isCapChecked)

	q.PopBack()
	checkLC(t, q, 0, 3, isCapChecked)
}

func testPopFront(t *testing.T, newDeque func() Deque, isCapChecked bool) {
	q := newDeque()

	q.PushFrontValue(91)
	q.PopFront()
	checkLC(t, q, 0, 1, isCapChecked)

	q.PushFrontValue(92)
	checkFBLC(t, q, 92, 92, 1, 1, isCapChecked)

	q.PushFrontValue(93)
	checkFBLC(t, q, 93, 92, 2, 2, isCapChecked)

	q.PopFront()
	checkFBLC(t, q, 92, 92, 1, 2, isCapChecked)

	q.PopFront()
	checkLC(t, q, 0, 2, isCapChecked)

	q.PushFrontValue(94)
	checkFBLC(t, q, 94, 94, 1, 2, isCapChecked)

	q.PushFrontValue(95)
	checkFBLC(t, q, 95, 94, 2, 2, isCapChecked)

	q.PushFrontValue(96)
	checkFBLC(t, q, 96, 94, 3, 3, isCapChecked)

	q.PopFront()
	checkFBLC(t, q, 95, 94, 2, 3, isCapChecked)

	q.PopFront()
	checkFBLC(t, q, 94, 94, 1, 3, isCapChecked)

	q.PopFront()
	checkLC(t, q, 0, 3, isCapChecked)
}

func testPopBackAndPopFront0(t *testing.T, newDeque func() Deque, isCapChecked bool) {
	q := newDeque()

	q.PushBackValue(91)
	q.PushBackValue(92)
	q.PushBackValue(93)
	q.PopFront()
	checkFBLC(t, q, 92, 93, 2, 3, isCapChecked)

	q.PushBackValue(94)
	checkFBLC(t, q, 92, 94, 3, 3, isCapChecked)

	q.PopBack()
	checkFBLC(t, q, 92, 93, 2, 3, isCapChecked)

	q.PushFrontValue(95)
	checkFBLC(t, q, 95, 93, 3, 3, isCapChecked)
}

func testPopBackAndPopFront1(t *testing.T, newDeque func() Deque, isCapChecked bool) {
	q := newDeque()

	q.PushBackValue(91)
	q.PushBackValue(92)
	q.PushBackValue(93)
	q.PopFront()
	q.PopFront()
	checkFBLC(t, q, 93, 93, 1, 3, isCapChecked)

	q.PushBackValue(94)
	checkFBLC(t, q, 93, 94, 2, 3, isCapChecked)

	q.PushBackValue(95)
	checkFBLC(t, q, 93, 95, 3, 3, isCapChecked)

	q.PopBack()
	q.PopBack()
	q.PushFrontValue(96)
	checkFBLC(t, q, 96, 93, 2, 3, isCapChecked)

	q.PushFrontValue(97)
	checkFBLC(t, q, 97, 93, 3, 3, isCapChecked)
}

func testPopBackAndPopFront2(t *testing.T, newDeque func() Deque, isCapChecked bool) {
	q := newDeque()

	q.PushBackValue(91)
	q.PushBackValue(92)
	q.PushBackValue(93)
	q.PopFront()
	q.PopFront()
	q.PopFront()
	checkLC(t, q, 0, 3, isCapChecked)

	q.PushBackValue(94)
	checkFBLC(t, q, 94, 94, 1, 3, isCapChecked)

	q.PushBackValue(95)
	checkFBLC(t, q, 94, 95, 2, 3, isCapChecked)

	q.PushBackValue(96)
	checkFBLC(t, q, 94, 96, 3, 3, isCapChecked)

	q.PopBack()
	q.PopBack()
	q.PopBack()
	q.PushFrontValue(97)
	checkFBLC(t, q, 97, 97, 1, 3, isCapChecked)

	q.PushFrontValue(98)
	checkFBLC(t, q, 98, 97, 2, 3, isCapChecked)

	q.PushFrontValue(99)
	checkFBLC(t, q, 99, 97, 3, 3, isCapChecked)
}

func TestNewDeque(t *testing.T) {
	testNewDeque(t, newArrayDeque, false)
	testNewDeque(t, newLinkedDeque, true)
}

func TestPushBack(t *testing.T) {
	testPushBack(t, newArrayDeque, false)
	testPushBack(t, newLinkedDeque, true)
}

func TestPushFront(t *testing.T) {
	testPushFront(t, newArrayDeque, false)
	testPushFront(t, newLinkedDeque, true)
}

func TestPushBackAndPushFront0(t *testing.T) {
	testPushBackAndPushFront0(t, newArrayDeque, false)
	testPushBackAndPushFront0(t, newLinkedDeque, true)
}

func TestPushBackAndPushFront1(t *testing.T) {
	testPushBackAndPushFront1(t, newArrayDeque, false)
	testPushBackAndPushFront1(t, newLinkedDeque, true)
}

func TestPopBack(t *testing.T) {
	testPopBack(t, newArrayDeque, false)
	testPopBack(t, newLinkedDeque, true)
}

func TestPopFront(t *testing.T) {
	testPopFront(t, newArrayDeque, false)
	testPopFront(t, newLinkedDeque, true)
}

func TestPopBackAndPopFront0(t *testing.T) {
	testPopBackAndPopFront0(t, newArrayDeque, false)
	testPopBackAndPopFront0(t, newLinkedDeque, true)
}

func TestPopBackAndPopFront1(t *testing.T) {
	testPopBackAndPopFront1(t, newArrayDeque, false)
	testPopBackAndPopFront1(t, newLinkedDeque, true)
}

func TestPopBackAndPopFront2(t *testing.T) {
	testPopBackAndPopFront2(t, newArrayDeque, false)
	testPopBackAndPopFront2(t, newLinkedDeque, true)
}
