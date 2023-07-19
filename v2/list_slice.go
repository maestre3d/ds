package ds

// A SliceList is a concrete implementation of List using a Go slice as underlying data structure.
type SliceList[T any] struct {
	buf []T
}

var _ List[any] = &SliceList[any]{}

// NewSliceList allocates a new SliceList. Accepts an initial capacity for the underlying slice.
// If capacity is < 0, capacity will fallback to 0.
func NewSliceList[T any](initialCap int) *SliceList[T] {
	if initialCap < 0 {
		initialCap = 0
	}

	return &SliceList[T]{
		buf: make([]T, 0, initialCap),
	}
}

func (s SliceList[T]) NewIterator(iterationType IterationType) Iterator[T] {
	return NewSliceIterator[T](s.buf, iterationType)
}

func (s SliceList[T]) Len() int {
	return len(s.buf)
}

func (s SliceList[T]) GetAt(pos int) T {
	return s.buf[pos]
}

func (s SliceList[T]) Append(v T) {
	s.buf = append(s.buf, v)
}

func (s SliceList[T]) AppendAll(items ...T) {
	s.buf = append(s.buf, items...)
}

func (s SliceList[T]) InsertAt(pos int, val T) {
	s.buf[pos] = val
}

func (s SliceList[T]) Remove() {
	s.buf = s.buf[:s.Len()-1]
}

func (s SliceList[T]) RemoveAt(pos int) {
	if pos < len(s.buf)-1 {
		copy(s.buf[pos:], s.buf[pos+1:])
	}
	var zeroVal T
	s.buf[len(s.buf)-1] = zeroVal
	s.buf = s.buf[:len(s.buf)-1]
}
