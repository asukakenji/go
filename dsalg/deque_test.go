package dsalg

import (
	"testing"
)

type queue interface {
	Back() interface{}
	Front() interface{}
	PushBack(x interface{})
	PushFront(x interface{})
	PopBack()
	PopFront()
	Reverse()
	Len() int
	Cap() int
	IsEmpty() bool
	IsFull() bool
	Print()
}

func checkLC(t *testing.T, q queue, expectedL, expectedC int, isCapChecked bool) {
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

func checkFBLC(t *testing.T, q queue, expectedF, expectedB, expectedL, expectedC int, isCapChecked bool) {
	f := q.Front().(int)
	if f != expectedF {
		t.Errorf("q.Front() == %d; expected %d", f, expectedF)
	}
	b := q.Back().(int)
	if b != expectedB {
		t.Errorf("q.Back() == %d; expected %d", b, expectedB)
	}
	checkLC(t, q, expectedL, expectedC, isCapChecked)
}

func newArrayDeque() queue {
	return NewArrayDeque(3)
}

func newLinkedDeque() queue {
	return NewLinkedDeque()
}

func testNewDeque(t *testing.T, newQueue func() queue, isCapChecked bool) {
	q := newQueue()

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

func testPushBack(t *testing.T, newQueue func() queue, isCapChecked bool) {
	q := newQueue()

	q.PushBack(42)
	checkFBLC(t, q, 42, 42, 1, 1, isCapChecked)

	q.PushBack(14)
	checkFBLC(t, q, 42, 14, 2, 2, isCapChecked)
}

func testPushFront(t *testing.T, newQueue func() queue, isCapChecked bool) {
	q := newQueue()

	q.PushFront(42)
	checkFBLC(t, q, 42, 42, 1, 1, isCapChecked)

	q.PushFront(14)
	checkFBLC(t, q, 14, 42, 2, 2, isCapChecked)
}

func testPushBackAndPushFront0(t *testing.T, newQueue func() queue, isCapChecked bool) {
	q := newQueue()

	q.PushBack(42)
	q.PushFront(14)
	checkFBLC(t, q, 14, 42, 2, 2, isCapChecked)

	q.PushBack(49)
	checkFBLC(t, q, 14, 49, 3, 3, isCapChecked)
}

func testPushBackAndPushFront1(t *testing.T, newQueue func() queue, isCapChecked bool) {
	q := newQueue()

	q.PushFront(42)
	q.PushBack(14)
	checkFBLC(t, q, 42, 14, 2, 2, isCapChecked)

	q.PushFront(49)
	checkFBLC(t, q, 49, 14, 3, 3, isCapChecked)
}

func testPopBack(t *testing.T, newQueue func() queue, isCapChecked bool) {
	q := newQueue()

	q.PushBack(91)
	q.PopBack()
	checkLC(t, q, 0, 1, isCapChecked)

	q.PushBack(92)
	checkFBLC(t, q, 92, 92, 1, 1, isCapChecked)

	q.PushBack(93)
	checkFBLC(t, q, 92, 93, 2, 2, isCapChecked)

	q.PopBack()
	checkFBLC(t, q, 92, 92, 1, 2, isCapChecked)

	q.PopBack()
	checkLC(t, q, 0, 2, isCapChecked)

	q.PushBack(94)
	checkFBLC(t, q, 94, 94, 1, 2, isCapChecked)

	q.PushBack(95)
	checkFBLC(t, q, 94, 95, 2, 2, isCapChecked)

	q.PushBack(96)
	checkFBLC(t, q, 94, 96, 3, 3, isCapChecked)

	q.PopBack()
	checkFBLC(t, q, 94, 95, 2, 3, isCapChecked)

	q.PopBack()
	checkFBLC(t, q, 94, 94, 1, 3, isCapChecked)

	q.PopBack()
	checkLC(t, q, 0, 3, isCapChecked)
}

func testPopFront(t *testing.T, newQueue func() queue, isCapChecked bool) {
	q := newQueue()

	q.PushFront(91)
	q.PopFront()
	checkLC(t, q, 0, 1, isCapChecked)

	q.PushFront(92)
	checkFBLC(t, q, 92, 92, 1, 1, isCapChecked)

	q.PushFront(93)
	checkFBLC(t, q, 93, 92, 2, 2, isCapChecked)

	q.PopFront()
	checkFBLC(t, q, 92, 92, 1, 2, isCapChecked)

	q.PopFront()
	checkLC(t, q, 0, 2, isCapChecked)

	q.PushFront(94)
	checkFBLC(t, q, 94, 94, 1, 2, isCapChecked)

	q.PushFront(95)
	checkFBLC(t, q, 95, 94, 2, 2, isCapChecked)

	q.PushFront(96)
	checkFBLC(t, q, 96, 94, 3, 3, isCapChecked)

	q.PopFront()
	checkFBLC(t, q, 95, 94, 2, 3, isCapChecked)

	q.PopFront()
	checkFBLC(t, q, 94, 94, 1, 3, isCapChecked)

	q.PopFront()
	checkLC(t, q, 0, 3, isCapChecked)
}

func testPopBackAndPopFront0(t *testing.T, newQueue func() queue, isCapChecked bool) {
	q := newQueue()

	q.PushBack(91)
	q.PushBack(92)
	q.PushBack(93)
	q.PopFront()
	checkFBLC(t, q, 92, 93, 2, 3, isCapChecked)

	q.PushBack(94)
	checkFBLC(t, q, 92, 94, 3, 3, isCapChecked)

	q.PopBack()
	checkFBLC(t, q, 92, 93, 2, 3, isCapChecked)

	q.PushFront(95)
	checkFBLC(t, q, 95, 93, 3, 3, isCapChecked)
}

func testPopBackAndPopFront1(t *testing.T, newQueue func() queue, isCapChecked bool) {
	q := newQueue()

	q.PushBack(91)
	q.PushBack(92)
	q.PushBack(93)
	q.PopFront()
	q.PopFront()
	checkFBLC(t, q, 93, 93, 1, 3, isCapChecked)

	q.PushBack(94)
	checkFBLC(t, q, 93, 94, 2, 3, isCapChecked)

	q.PushBack(95)
	checkFBLC(t, q, 93, 95, 3, 3, isCapChecked)

	q.PopBack()
	q.PopBack()
	q.PushFront(96)
	checkFBLC(t, q, 96, 93, 2, 3, isCapChecked)

	q.PushFront(97)
	checkFBLC(t, q, 97, 93, 3, 3, isCapChecked)
}

func testPopBackAndPopFront2(t *testing.T, newQueue func() queue, isCapChecked bool) {
	q := newQueue()

	q.PushBack(91)
	q.PushBack(92)
	q.PushBack(93)
	q.PopFront()
	q.PopFront()
	q.PopFront()
	checkLC(t, q, 0, 3, isCapChecked)

	q.PushBack(94)
	checkFBLC(t, q, 94, 94, 1, 3, isCapChecked)

	q.PushBack(95)
	checkFBLC(t, q, 94, 95, 2, 3, isCapChecked)

	q.PushBack(96)
	checkFBLC(t, q, 94, 96, 3, 3, isCapChecked)

	q.PopBack()
	q.PopBack()
	q.PopBack()
	q.PushFront(97)
	checkFBLC(t, q, 97, 97, 1, 3, isCapChecked)

	q.PushFront(98)
	checkFBLC(t, q, 98, 97, 2, 3, isCapChecked)

	q.PushFront(99)
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
