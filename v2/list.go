package ds

// A List is a sequential set of items stored in a physical data structure (array or linked list).
type List[T any] interface {
	Iterable[T]
	// Len retrieves total number of items a List holds.
	Len() int
	// GetAt retrieves an item at the given position.
	GetAt(pos int) T
	// Append adds a new item at the end of the list.
	Append(v T)
	// AppendAll adds a set of new items at the end of the list.
	AppendAll(items ...T)
	// InsertAt inserts a new item at the given position.
	InsertAt(pos int, val T)
	// Remove deletes an item at the end of the list.
	Remove()
	// RemoveAt deletes an item at a specific position.
	RemoveAt(pos int)
}
