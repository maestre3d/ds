package ds

// Queue logical data structure which may be implemented as LIFO (stack) or FIFO (default queue). Depending on the
// implementation, it will read data from the tail or head of the underlying data structure which may be
// a heap (i.e. priority queues), slice or linked list.
type Queue[T any] interface {
	SerializableJSON
	Len() int
	Enqueue(T)
	Dequeue() T
	Peek() T
}

// SliceQueue also known as FIFO queue. SliceQueue is a slice-backed logical data structure which is only* able
// to read data from the first element from its underlying slice.
//
// *The structure contains a PeekAt method to read any index from the underlying slice.
type SliceQueue[T any] struct {
	buf []T
}

var _ Queue[any] = &SliceQueue[any]{}

func NewSliceQueue[T any](s int) *SliceQueue[T] {
	return &SliceQueue[T]{
		buf: make([]T, 0, s),
	}
}

func (s SliceQueue[T]) Len() int {
	return len(s.buf)
}

func (s SliceQueue[T]) Cap() int {
	return cap(s.buf)
}

func (s *SliceQueue[T]) Enqueue(v T) {
	s.buf = append(s.buf, v)
}

func (s *SliceQueue[T]) Dequeue() T {
	var res T
	if len(s.buf) == 0 {
		return res
	}
	res = s.buf[0]
	s.buf = s.buf[1:]
	return res
}

func (s SliceQueue[T]) Peek() T {
	var zeroVal T
	if len(s.buf) == 0 {
		return zeroVal
	}
	return s.buf[0]
}

func (s SliceQueue[T]) MarshalJSON() ([]byte, error) {
	return DefaultJSONMarshaler(s.buf)
}

func (s *SliceQueue[T]) UnmarshalJSON(bytes []byte) error {
	return DefaultJSONUnmarshaler(bytes, &s.buf)
}
