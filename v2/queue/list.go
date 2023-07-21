package queue

import (
	"github.com/maestre3d/ds/v2"
)

type ListQueue[T any] struct {
	list      ds.List[T]
	queueType Type
}

var _ Queue[any] = ListQueue[any]{}

func NewListQueue[T any](t Type, l ds.List[T]) ListQueue[T] {
	return ListQueue[T]{
		list:      l,
		queueType: t,
	}
}

func (l ListQueue[T]) NewIterator(_ ds.IterationType) ds.Iterator[T] {
	return newIterator[T](l)
}

func (l ListQueue[T]) Enqueue(in T) {
	l.list.Append(in)
}

func (l ListQueue[T]) EnqueueAll(in ...T) {
	l.list.AppendAll(in...)
}

func (l ListQueue[T]) Peek() (T, bool) {
	if l.list.Len() == 0 {
		var zeroVal T
		return zeroVal, false
	}

	pos := 0
	if l.queueType == LIFO {
		pos = l.list.Len() - 1
	}

	val := l.list.GetAt(pos)
	return val, true
}

func (l ListQueue[T]) Dequeue() (T, bool) {
	if l.list.Len() == 0 {
		var zeroVal T
		return zeroVal, false
	}

	pos := 0
	if l.queueType == LIFO {
		pos = l.list.Len() - 1
	}

	val := l.list.GetAt(pos)
	l.list.RemoveAt(pos)
	return val, true
}

func (l ListQueue[T]) Clear() {
	l.list.RemoveAll()
}

func (l ListQueue[T]) Len() int {
	return l.list.Len()
}
