package ds

type SliceIterator[T any] struct {
	src       []T
	iterPos   int
	totalIter int
	iterType  IterationType
}

var _ Iterator[any] = &SliceIterator[any]{}

func NewSliceIterator[T any](src []T, iterationType IterationType) *SliceIterator[T] {
	iterPos := 0
	if iterationType == BackwardIteration {
		iterPos = len(src) - 1
	}

	return &SliceIterator[T]{
		src:      src,
		iterPos:  iterPos,
		iterType: iterationType,
	}
}

func (s *SliceIterator[T]) HasNext() bool {
	return len(s.src) > s.totalIter
}

func (s *SliceIterator[T]) Next() T {
	defer func() {
		s.totalIter++
		if s.iterType == ForwardIteration {
			s.iterPos++
			return
		}
		s.iterPos--
	}()

	return s.src[s.iterPos]
}
