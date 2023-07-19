package linkedlist

import "github.com/maestre3d/ds/v2"

type Iterator[T any] struct {
	list          *List[T]
	currentNode   *Node[T]
	totalIterated int

	iterType ds.IterationType
}

var _ ds.Iterator[any] = &Iterator[any]{}

func (i *Iterator[T]) HasNext() bool {
	return i.list.Len() > i.totalIterated
}

func (i *Iterator[T]) Next() T {
	defer func() {
		i.totalIterated++
	}()
	if i.iterType == ds.BackwardIteration {
		return i.nextReverse()
	}

	return i.next()
}

func (i *Iterator[T]) next() T {
	if i.totalIterated == 0 {
		i.currentNode = i.list.head
		return i.currentNode.Value
	}
	i.currentNode = i.currentNode.Next
	return i.currentNode.Value
}

func (i *Iterator[T]) nextReverse() T {
	if i.totalIterated == 0 {
		i.currentNode = i.list.tail
		return i.currentNode.Value
	}
	i.currentNode = i.currentNode.Previous
	return i.currentNode.Value
}
