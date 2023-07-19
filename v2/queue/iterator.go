package queue

import "github.com/maestre3d/ds/v2"

type Iterator[T any] struct {
	queue     Queue[T]
	initLen   int
	totalIter int
}

var _ ds.Iterator[any] = &Iterator[any]{}

func newIterator[T any](q Queue[T]) *Iterator[T] {
	return &Iterator[T]{
		queue:     q,
		initLen:   q.Len(),
		totalIter: 0,
	}
}

func (i *Iterator[T]) HasNext() bool {
	return i.initLen > i.totalIter
}

func (i *Iterator[T]) Next() T {
	defer func() {
		i.totalIter++
	}()
	val, _ := i.queue.Dequeue()
	return val
}
