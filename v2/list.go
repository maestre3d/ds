package ds

// A List is a sequential set of items stored in a physical data structure (array or linked list).
type List[T any] interface {
	Iterable[T]
	// Len retrieves total number of items a List holds.
	Len() int
	// GetAt retrieves an item at the given position.
	GetAt(pos int) T
	// Append adds a new item at the end of the List.
	Append(v T)
	// AppendAll adds a set of new items at the end of the List.
	AppendAll(items ...T)
	// InsertAt inserts a new item at the given position. Item shifting might be performed.
	InsertAt(pos int, val T)
	// ReplaceAt replaces an existing item at the given position.
	ReplaceAt(pos int, val T)
	// Remove deletes an item at the end of the List.
	Remove()
	// RemoveAt deletes an item at a specific position.
	RemoveAt(pos int)
	// RemoveAll deletes every item from the List. Might reallocate.
	RemoveAll()
}
