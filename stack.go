package ds

// SliceStack also known as LIFO queue. SliceStack is a slice-backed logical data structure which is only* able
// to read data from the last element from its underlying slice.
//
// *The structure contains a PeekAt method to read any index from the underlying slice.
type SliceStack[T any] struct {
	buf []T
}

var _ SerializableQueue[any] = &SliceStack[any]{}

func NewSliceStack[T any](c int) *SliceStack[T] {
	return &SliceStack[T]{
		buf: make([]T, 0, c),
	}
}

func (s SliceStack[T]) Len() int {
	return len(s.buf)
}

func (s SliceStack[T]) Cap() int {
	return cap(s.buf)
}

func (s *SliceStack[T]) Push(v T) {
	s.buf = append(s.buf, v)
}

func (s *SliceStack[T]) Pop() T {
	var res T
	if len(s.buf) == 0 {
		return res
	}
	res = s.buf[len(s.buf)-1]
	s.buf = s.buf[:len(s.buf)-1]
	return res
}

func (s SliceStack[T]) Peek() T {
	var zeroVal T
	if len(s.buf) == 0 {
		return zeroVal
	}
	return s.buf[len(s.buf)-1]
}

func (s SliceStack[T]) PeekAt(i int) T {
	var zeroVal T
	if len(s.buf) == 0 || i > len(s.buf)-1 {
		return zeroVal
	}
	return s.buf[i]
}

func (s SliceStack[T]) MarshalJSON() ([]byte, error) {
	return DefaultJSONMarshaler(s.buf)
}

func (s *SliceStack[T]) UnmarshalJSON(bytes []byte) error {
	return DefaultJSONUnmarshaler(bytes, &s.buf)
}
