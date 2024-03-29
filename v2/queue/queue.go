package queue

import "github.com/maestre3d/ds/v2"

// Type kind of Queue (FIFO or LIFO).
type Type uint8

const (
	_ Type = iota
	// LIFO represents a Queue with Last-In First-Out (LIFO) item dequeue technique. Also known as Stack.
	LIFO
	// FIFO represents a Queue with First-In First-Out (FIFO) item dequeue technique.
	FIFO
)

// A Queue is a logical data structure which holds items to be later popped out from an internal physical data structure
// (array/linked list). This last operation depends on the Type (FIFO or LIFO).
type Queue[T any] interface {
	ds.Iterable[T]
	// Enqueue appends an item into the Queue.
	Enqueue(in T)
	// EnqueueAll appends a list of items into the Queue.
	EnqueueAll(in ...T)
	// Peek retrieves an item from the queue. Depending on the Queue Type, this will be the first item or last item
	// of the Queue.
	Peek() (T, bool)
	// Dequeue retrieves and removes an item from the queue. Depending on the Queue Type, this will be the first
	// item or last item of the Queue.
	Dequeue() (T, bool)
	// Clear removes all items from the queue. Depending on the concrete implementation, this may NOT reallocate
	// the underlying physical data structure, increasing performance.
	Clear()
	// Len retrieves total count of remaining items.
	Len() int
}
