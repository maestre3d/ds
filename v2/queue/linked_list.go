package queue

import (
	"github.com/maestre3d/ds/v2"
	"github.com/maestre3d/ds/v2/linkedlist"
)

// A LinkedListQueue is a concrete implementation of a Queue using linkedlist.List underlying data structure.
//
// Can be LIFO (stack) or FIFO.
type LinkedListQueue[T any] struct {
	list      *linkedlist.List[T]
	queueType Type
}

var _ Queue[any] = LinkedListQueue[any]{}

// NewLinkedListQueue allocates a new LinkedListQueue Queue instance.
//
// Default Type is FIFO.
func NewLinkedListQueue[T any](t Type) LinkedListQueue[T] {
	return LinkedListQueue[T]{
		list:      linkedlist.NewList[T](),
		queueType: t,
	}
}

// NewIterator allocates a new Iterator instance.
func (l LinkedListQueue[T]) NewIterator(_ ds.IterationType) ds.Iterator[T] {
	return newIterator[T](l)
}

// Enqueue appends an item into the Queue.
func (l LinkedListQueue[T]) Enqueue(in T) {
	l.list.Append(in)
}

// EnqueueAll appends a list of items into the Queue.
func (l LinkedListQueue[T]) EnqueueAll(in ...T) {
	l.list.AppendAll(in...)
}

// Peek retrieves an item from the queue. Depending on the Queue Type, this will be the first item or last item
// of the Queue.
func (l LinkedListQueue[T]) Peek() (T, bool) {
	var getNodeFunc func() *linkedlist.Node[T]
	if l.queueType == LIFO {
		getNodeFunc = l.list.GetTail
	} else {
		getNodeFunc = l.list.GetHead
	}

	if node := getNodeFunc(); node != nil {
		return node.Value, true
	}

	var zeroVal T
	return zeroVal, false
}

// Dequeue retrieves and removes an item from the queue. Depending on the Queue Type, this will be the first
// item or last item of the Queue.
func (l LinkedListQueue[T]) Dequeue() (T, bool) {
	var getNodeFunc func() *linkedlist.Node[T]
	pos := 0
	if l.queueType == LIFO {
		getNodeFunc = l.list.GetTail
		pos = l.list.Len() - 1
	} else {
		getNodeFunc = l.list.GetHead
	}

	if node := getNodeFunc(); node != nil {
		l.list.RemoveAt(pos)
		return node.Value, true
	}

	var zeroVal T
	return zeroVal, false
}

// Clear removes all items from the queue.
func (l LinkedListQueue[T]) Clear() {
	l.list.RemoveAll()
}

// Len retrieves the length of the internal linkedlist.List.
func (l LinkedListQueue[T]) Len() int {
	return l.list.Len()
}
