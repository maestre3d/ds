package linkedlist

import "github.com/maestre3d/ds/v2"

// A List is a physical data structure with dynamic memory allocation, using Node to hold sentinel values and pointer
// references to items from the list. This List in particular, implements a double linked-list.
//
// Finally, List implements ds.List.
type List[T any] struct {
	head   *Node[T]
	tail   *Node[T]
	length int
}

var _ ds.List[any] = &List[any]{}

// NewList allocates a new List.
func NewList[T any]() *List[T] {
	return &List[T]{
		head: nil,
	}
}

// NewIterator allocates a new Iterator instance.
func (l *List[T]) NewIterator(iterationType ds.IterationType) ds.Iterator[T] {
	return &Iterator[T]{
		list:     l,
		iterType: iterationType,
	}
}

// Len retrieves total items on the list.
func (l *List[T]) Len() int {
	return l.length
}

// GetHead retrieves the head Node of the List.
func (l *List[T]) GetHead() *Node[T] {
	return l.head
}

// GetTail retrieves the tail Node of the List.
func (l *List[T]) GetTail() *Node[T] {
	return l.tail
}

// GetNodeAt retrieves a Node at a given position.
func (l *List[T]) GetNodeAt(pos int) *Node[T] {
	if pos == 0 {
		return l.head
	}

	current := l.head
	for i := 0; i < pos; i++ {
		current = current.Next
	}
	return current
}

// GetAt retrieves the sentinel value from a Node at a given position.
func (l *List[T]) GetAt(pos int) T {
	if pos == 0 {
		return l.head.Value
	}

	current := l.head
	for i := 0; i < pos; i++ {
		current = current.Next
	}
	return current.Value
}

// Append attaches a Node at the tail of the list (i.e. last element).
func (l *List[T]) Append(item T) {
	defer func() {
		l.length++
	}()

	if l.head == nil {
		l.head = &Node[T]{
			Value: item,
		}
		l.tail = l.head
		return
	}

	current := l.head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = &Node[T]{
		Value:    item,
		Previous: current,
	}
	l.tail = current.Next
}

// AppendAll attaches a sequence of Node(s) at the tail of the list (i.e. last element).
func (l *List[T]) AppendAll(items ...T) {
	for _, item := range items {
		l.Append(item)
	}
}

// InsertAt inserts a Node at a given position, shifting Node pointers (Node.Next & Node.Previous) if required.
func (l *List[T]) InsertAt(pos int, item T) {
	defer func() {
		l.length++
	}()

	if pos == 0 {
		node := &Node[T]{
			Value: item,
			Next:  l.head,
		}
		l.head.Previous = node
		l.head = node
		return
	}

	parentNode := l.GetNodeAt(pos - 1)
	currentNode := l.GetNodeAt(pos)
	node := &Node[T]{
		Value:    item,
		Next:     currentNode,
		Previous: currentNode.Previous,
	}
	parentNode.Next = node
	currentNode.Previous = node
}

// ReplaceAt inserts given value into the Node at the given position.
func (l *List[T]) ReplaceAt(pos int, val T) {
	node := l.GetNodeAt(pos)
	node.Value = val
}

func (l *List[T]) Remove() {
	//node := l.GetNodeAt(l.Len() - 1)
	//node = nil
}

func (l *List[T]) RemoveAt(pos int) {
}
