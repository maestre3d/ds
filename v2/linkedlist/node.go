package linkedlist

// A Node is an item on a List. Holds a sentinel value T and two pointers, next and previous nodes.
type Node[T any] struct {
	Value    T
	Next     *Node[T]
	Previous *Node[T]
}
