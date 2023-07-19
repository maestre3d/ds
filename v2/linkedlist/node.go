package linkedlist

// A Node is an item on a List.
type Node[T any] struct {
	Value    T
	Next     *Node[T]
	Previous *Node[T]
}
