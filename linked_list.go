package ds

type linkedListItem[T any] struct {
	next *linkedListItem[T]
	val  T
}

type LinkedList[T any] struct {
	length int
	head   *linkedListItem[T]
}

func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{
		length: 0,
		head:   nil,
	}
}

func (l *LinkedList[T]) Append(v T) {
	l.length++
	if l.head == nil {
		l.head = &linkedListItem[T]{val: v}
		return
	}
	popItem := l.head
	l.head = &linkedListItem[T]{val: v, next: popItem}
}

func (l *LinkedList[T]) AppendAt(n int, v T) {
	if n > l.length {
		return
	} else if n == 0 {
		l.head.val = v
	}

	count := 0
	tmp := l.head
	for count <= n {
		if count == n {
			tmp.val = v
			break
		}
		tmp = tmp.next
		count++
	}
}

func (l *LinkedList[T]) AppendTail(v T) {
	l.AppendAt(l.length-1, v)
}

func (l *LinkedList[T]) GetAt(n int) T {
	var res T
	if n > l.length {
		return res
	} else if n == 0 {
		return l.head.val
	}
	count := 0
	tmp := l.head
	for count < n {
		tmp = tmp.next
		count++
	}
	if tmp != nil {
		res = tmp.val
	}
	return res
}

func (l *LinkedList[T]) GetTail() T {
	var res T // zero value
	if l.head == nil {
		return res
	}
	return l.head.val
}
