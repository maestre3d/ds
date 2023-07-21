package queue

import "github.com/maestre3d/ds/v2"

// A SliceQueue is a concrete implementation of a Queue using a Go slice as underlying data structure.
//
// Can be LIFO (stack) or FIFO.
type SliceQueue[T any] struct {
	buf []T
	t   Type
}

var _ Queue[string] = &SliceQueue[string]{}

// NewSliceQueue allocates a new SliceQueue Queue instance. Accepts an initial capacity for the underlying slice.
// If capacity is < 0, capacity will fallback to 0.
//
// Default Type is FIFO.
func NewSliceQueue[T any](t Type, initCap int) *SliceQueue[T] {
	if initCap < 0 {
		initCap = 0
	}

	if t != FIFO && t != LIFO {
		t = FIFO
	}

	return &SliceQueue[T]{
		buf: make([]T, 0, initCap),
		t:   t,
	}
}

// NewIterator allocates a new Iterator instance.
func (s *SliceQueue[T]) NewIterator(_ ds.IterationType) ds.Iterator[T] {
	return newIterator[T](s)
}

// Enqueue appends an item into the Queue.
func (s *SliceQueue[T]) Enqueue(in T) {
	s.buf = append(s.buf, in)
}

// EnqueueAll appends a list of items into the Queue.
func (s *SliceQueue[T]) EnqueueAll(in ...T) {
	s.buf = append(s.buf, in...)
}

// Peek retrieves an item from the queue. Depending on the Queue Type, this will be the first item or last item
// of the Queue.
func (s *SliceQueue[T]) Peek() (T, bool) {
	if len(s.buf) == 0 {
		var zeroVal T
		return zeroVal, false
	}

	i := 0
	if s.t == LIFO {
		i = len(s.buf) - 1
	}
	return s.buf[i], true
}

// Dequeue retrieves and removes an item from the queue. Depending on the Queue Type, this will be the first
// item or last item of the Queue.
func (s *SliceQueue[T]) Dequeue() (T, bool) {
	if len(s.buf) == 0 {
		var zeroVal T
		return zeroVal, false
	}

	i := 0
	if s.t == LIFO {
		i = len(s.buf) - 1
	}
	snapshot := s.buf[i]

	if s.t == FIFO {
		s.buf = s.buf[1:]
	} else {
		s.buf = s.buf[:len(s.buf)-1]
	}
	return snapshot, true
}

// Len retrieves the length of the internal slice buffer.
func (s *SliceQueue[T]) Len() int {
	return len(s.buf)
}

// Cap retrieves the capacity of the internal slice buffer.
func (s *SliceQueue[T]) Cap() int {
	return cap(s.buf)
}

// Clear removes all items from the queue. This operation DOES NOT reallocate
// the underlying slice buffer, increasing performance.
func (s *SliceQueue[T]) Clear() {
	s.buf = s.buf[:0]
}
