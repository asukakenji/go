package dsalg

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
