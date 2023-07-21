package slice

import "github.com/maestre3d/ds/v2"

// A List is a concrete implementation of ds.List using a Go slice as underlying data structure.
type List[T any] struct {
	buf []T
}

var _ ds.List[any] = &List[any]{}

// NewList allocates a new List. Accepts an initial capacity for the underlying slice.
// If capacity is < 0, capacity will fallback to 0.
func NewList[T any](initialCap int) *List[T] {
	if initialCap < 0 {
		initialCap = 0
	}

	return &List[T]{
		buf: make([]T, 0, initialCap),
	}
}

// NewIterator allocates a new Iterator instance.
func (s *List[T]) NewIterator(iterationType ds.IterationType) ds.Iterator[T] {
	return NewIterator[T](s.buf, iterationType)
}

// Len retrieves total number of items the internal slice holds.
func (s *List[T]) Len() int {
	return len(s.buf)
}

// Cap retrieves the total capacity the internal slice holds.
func (s *List[T]) Cap() int {
	return cap(s.buf)
}

// GetAt retrieves an item at the given position.
func (s *List[T]) GetAt(pos int) T {
	if pos >= len(s.buf) {
		var zeroVal T
		return zeroVal
	}
	return s.buf[pos]
}

// Append adds a new item at the end of the list.
func (s *List[T]) Append(v T) {
	s.buf = append(s.buf, v)
}

// AppendAll adds a set of new items at the end of the list.
func (s *List[T]) AppendAll(items ...T) {
	s.buf = append(s.buf, items...)
}

// InsertAt inserts a new item at the given position. Item shifting is performed.
func (s *List[T]) InsertAt(pos int, val T) {
	if pos >= len(s.buf) {
		if cap(s.buf)-1 > pos && pos == 0 {
			s.buf = append(s.buf, val)
		}
		return
	}
	s.buf = append(s.buf[:pos], append([]T{val}, s.buf[pos:]...)...)
}

// ReplaceAt replaces an existing item at the given position.
func (s *List[T]) ReplaceAt(pos int, val T) {
	s.buf[pos] = val
}

// Remove deletes an item at the end of the list.
func (s *List[T]) Remove() {
	if len(s.buf) == 0 {
		return
	}

	s.buf = s.buf[:s.Len()-1]
}

// RemoveAt deletes an item at a specific position.
func (s *List[T]) RemoveAt(pos int) {
	if len(s.buf) == 0 {
		return
	}

	if pos < len(s.buf)-1 {
		copy(s.buf[pos:], s.buf[pos+1:])
	}
	var zeroVal T
	s.buf[len(s.buf)-1] = zeroVal
	s.buf = s.buf[:len(s.buf)-1]
}

// RemoveAll deletes every item from the list. Keeps internal slice capacity.
func (s *List[T]) RemoveAll() {
	s.buf = s.buf[:0]
}
