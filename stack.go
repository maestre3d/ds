package ds

// SliceStack LIFO queue.
type SliceStack[T any] struct {
	buf []T
}

func NewSliceStack[T any](s int) *SliceStack[T] {
	return &SliceStack[T]{
		buf: make([]T, 0, s),
	}
}

func (s *SliceStack[T]) Len() int {
	return len(s.buf)
}

func (s *SliceStack[T]) Cap() int {
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
	tmp := s.buf[len(s.buf)-1]
	s.buf = s.buf[:len(s.buf)-1]
	return tmp
}
