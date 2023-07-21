package slice

import "github.com/maestre3d/ds/v2"

type Iterator[T any] struct {
	src       []T
	iterPos   int
	totalIter int
	iterType  ds.IterationType
}

var _ ds.Iterator[any] = &Iterator[any]{}

func NewIterator[T any](src []T, iterationType ds.IterationType) *Iterator[T] {
	iterPos := 0
	if iterationType == ds.BackwardIteration {
		iterPos = len(src) - 1
	}

	return &Iterator[T]{
		src:      src,
		iterPos:  iterPos,
		iterType: iterationType,
	}
}

func (s *Iterator[T]) HasNext() bool {
	return len(s.src) > s.totalIter
}

func (s *Iterator[T]) Next() T {
	defer func() {
		s.totalIter++
		if s.iterType == ds.ForwardIteration {
			s.iterPos++
			return
		}
		s.iterPos--
	}()

	return s.src[s.iterPos]
}
