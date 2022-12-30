package ds

// Stack also known as Last-In-First-Out (LIFO) queue. Stack is a logical data structure which is only able
// to read data from the last element from its underlying storage. An item is also removed from the storage
// when Pop() is called.
type Stack[T any] interface {
	SerializableJSON
	// Len retrieves the number of items stored.
	Len() int
	// Push inserts an item into the stack.
	Push(T)
	// Pop retrieves and removes an item from the stack.
	Pop() T
	// Peek retrieves an item from the stack (no removal).
	Peek() T
}

// SliceStack also known as LIFO queue. SliceStack is a slice-backed logical data structure which is only able
// to read data from the last element from its underlying slice.
type SliceStack[T any] struct {
	buf       []T
	isDynamic bool
	pivot     int
}

var _ Stack[any] = &SliceStack[any]{}

// NewSliceStack allocates a new dynamic-sized Stack using a slice as underlying physical data structure.
// Internal slice will grow when fulfilling previous capacity using a rounding double-like mechanism.
func NewSliceStack[T any]() *SliceStack[T] {
	return &SliceStack[T]{
		buf:       make([]T, 0),
		isDynamic: true,
	}
}

// NewSliceStackFixed allocates a new fixed-size Stack using a slice as underlying physical data structure.
// As the underlying slice is fixed, the slice will not grow and thus new items pushed will be ignored after reaching the slice
// maximum capacity.
func NewSliceStackFixed[T any](c int) *SliceStack[T] {
	return &SliceStack[T]{
		buf: make([]T, c),
	}
}

// Len retrieves the number of items stored.
func (s SliceStack[T]) Len() int {
	return s.pivot
}

// Cap retrieves the current capacity of the underlying slice.
func (s SliceStack[T]) Cap() int {
	return cap(s.buf)
}

// Push inserts an item into the stack.
func (s *SliceStack[T]) Push(v T) {
	if s.isDynamic {
		s.buf = append(s.buf, v)
		s.pivot++
		return
	} else if (s.pivot) > cap(s.buf) {
		return
	}
	s.buf[s.pivot] = v
	s.pivot++
}

// Pop retrieves and removes an item from the stack.
func (s *SliceStack[T]) Pop() (res T) {
	if s.pivot == 0 {
		return res
	}
	tmpZeroVal := res
	s.pivot--
	res = s.buf[s.pivot]
	s.buf[s.pivot] = tmpZeroVal
	return
}

// Peek retrieves an item from the stack (no removal).
func (s SliceStack[T]) Peek() T {
	var zeroVal T
	if s.pivot == 0 {
		return zeroVal
	}
	return s.buf[s.pivot-1]
}

// MarshalJSON converts internal slice into a JSON array.
func (s SliceStack[T]) MarshalJSON() ([]byte, error) {
	return DefaultJSONMarshaler(s.buf)
}

// UnmarshalJSON converts and stores a JSON array into internal slice.
func (s *SliceStack[T]) UnmarshalJSON(bytes []byte) error {
	return DefaultJSONUnmarshaler(bytes, &s.buf)
}

type LinkedStack[T any] struct {
	headNode *linkedListNode[T]
	length   int
}

func NewLinkedStack[T any]() LinkedStack[T] {
	return LinkedStack[T]{}
}

var _ Stack[int] = &LinkedStack[int]{}

func (l LinkedStack[T]) Len() int {
	return l.length
}

func (l *LinkedStack[T]) Push(t T) {
	node := &linkedListNode[T]{
		val: t,
	}
	l.length++
	if l.headNode == nil {
		l.headNode = node
		return
	}
	tmpNode := *l.headNode
	l.headNode = node
	node.next = &tmpNode
}

func (l *LinkedStack[T]) Pop() (res T) {
	if l.headNode == nil {
		return
	}
	l.length--
	res = l.headNode.val
	l.headNode = l.headNode.next
	return
}

func (l LinkedStack[T]) Peek() T {
	var res T
	if l.headNode == nil {
		return res
	}
	return l.headNode.val
}

func (l LinkedStack[T]) MarshalJSON() ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (l LinkedStack[T]) UnmarshalJSON(bytes []byte) error {
	//TODO implement me
	panic("implement me")
}
