package ds

// SliceQueue FIFO Queue.
type SliceQueue[T any] struct {
	buf []T
}

func NewSliceQueue[T any](s int) *SliceQueue[T] {
	return &SliceQueue[T]{
		buf: make([]T, 0, s),
	}
}

func (s *SliceQueue[T]) Len() int {
	return len(s.buf)
}

func (s *SliceQueue[T]) Cap() int {
	return cap(s.buf)
}

func (s *SliceQueue[T]) Push(v T) {
	s.buf = append(s.buf, v)
}

func (s *SliceQueue[T]) Pop() T {
	var res T
	if len(s.buf) == 0 {
		return res
	}
	tmp := s.buf[0]
	s.buf = s.buf[1:]
	return tmp
}
