package ds

import (
	"container/list"
	"errors"
)

type DoublyLinkedList[T any] struct {
	l          *list.List
	iterBuffer *list.Element
	iterCount  int
}

var _ List[int] = &DoublyLinkedList[int]{}

func NewDoublyLinkedList[T any]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{
		l: list.New(),
	}
}

func traverseDoublyLinkedList(n *list.Element, stopIndex int, isReverse bool) *list.Element {
	if isReverse {
		if n.Prev() == nil || stopIndex == 0 {
			return n
		}
		if stopIndex >= 0 {
			stopIndex--
		}

		return traverseDoublyLinkedList(n.Prev(), stopIndex, isReverse)
	}
	if n.Next() == nil || stopIndex == 0 {
		return n
	}
	if stopIndex >= 0 {
		stopIndex--
	}

	return traverseDoublyLinkedList(n.Next(), stopIndex, isReverse)
}

func (d DoublyLinkedList[T]) HasNext() bool {
	return (d.l.Len() > 0 && d.iterCount == 0) || d.iterBuffer != nil
}

func (d *DoublyLinkedList[T]) Next() T {
	if d.iterBuffer == nil && d.iterCount == 0 {
		d.iterBuffer = d.l.Front()
	}
	res := d.iterBuffer.Value.(T)
	d.iterBuffer = d.iterBuffer.Next()
	d.iterCount++
	return res
}

func (d DoublyLinkedList[T]) Len() int {
	return d.l.Len()
}

func (d DoublyLinkedList[T]) Append(v T) {
	_ = d.l.PushBack(v)
}

func (d DoublyLinkedList[T]) Remove() {
	d.l.Remove(d.l.Back())
}

func (d DoublyLinkedList[T]) search(i int) *list.Element {
	if i+1 > d.l.Len() {
		return nil
	}

	reverseTraversal := false
	traversingNode := d.l.Front()
	half := d.l.Len() / 2
	if i > half {
		traversingNode = d.l.Back()
		reverseTraversal = true
	}
	return traverseDoublyLinkedList(traversingNode, i, reverseTraversal)
}

func (d DoublyLinkedList[T]) InsertAt(i int, v T) {
	if node := d.search(i); node != nil {
		d.l.InsertBefore(v, node)
	}
}

func (d DoublyLinkedList[T]) GetAt(i int) (res T) {
	if node := d.search(i); node != nil {
		return node.Value.(T)
	}
	return
}

func (d DoublyLinkedList[T]) RemoveAt(i int) {
	if node := d.search(i); node != nil {
		d.l.Remove(node)
	}
}

type linkedListNode[T any] struct {
	val  T
	next *linkedListNode[T]
}

type SinglyLinkedList[T any] struct {
	length int
	head   *linkedListNode[T]

	iterStarted bool
	iterBuffer  *linkedListNode[T]
}

var _ List[int] = &SinglyLinkedList[int]{}

func NewSinglyLinkedList[T any]() *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{
		length: 0,
		head:   nil,
	}
}

func (l SinglyLinkedList[T]) Len() int {
	return l.length
}

func traverseLinkedList[T any](n *linkedListNode[T], stopIndex int) *linkedListNode[T] {
	if n.next == nil || stopIndex == 0 {
		return n
	}
	if stopIndex >= 0 {
		stopIndex--
	}

	return traverseLinkedList(n.next, stopIndex)
}

func (l *SinglyLinkedList[T]) Append(v T) {
	l.length++
	node := &linkedListNode[T]{
		val: v,
	}
	if l.head == nil {
		l.head = node
		return
	}

	if tail := traverseLinkedList(l.head, -1); tail != nil {
		tail.next = node
	}
}

func (l *SinglyLinkedList[T]) Remove() {
	tmpNode := l.head
	for tmpNode != nil {
		if tmpNode.next == nil {
			tmpNode = nil
			break
		}
		tmpNode = tmpNode.next
	}
}

func (l SinglyLinkedList[T]) GetAt(i int) (res T) {
	if i >= l.length {
		return
	} else if node := traverseLinkedList(l.head, i); node != nil {
		return node.val
	}
	return
}

func (l *SinglyLinkedList[T]) InsertAt(i int, v T) {
	if i >= l.length {
		panic(errors.New("linked_list: insertion out of range"))
	} else if i == 0 {
		l.head.val = v
		return
	}

	if node := traverseLinkedList(l.head, i); node != nil {
		node.val = v
	}
}

func (l *SinglyLinkedList[T]) RemoveAt(i int) {
	if i >= l.length {
		panic(errors.New("linked_list: removal out of range"))
	}

	l.length--
	if i == 0 {
		tmp := *l.head.next
		l.head = &tmp
		return
	}

	if node := traverseLinkedList[T](l.head, i); node != nil {
	}
}

func (l SinglyLinkedList[T]) HasNext() bool {
	if l.length == 0 {
		return false
	} else if l.iterStarted {
		return l.iterBuffer != nil
	}
	return true
}

func (l *SinglyLinkedList[T]) Next() T {
	if !l.iterStarted {
		l.iterStarted = true
		l.iterBuffer = l.head
	}
	res := l.iterBuffer.val
	l.iterBuffer = l.iterBuffer.next
	return res
}
