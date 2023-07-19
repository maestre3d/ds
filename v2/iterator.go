package ds

// An Iterator traverses an underlying data structure.
type Iterator[T any] interface {
	// HasNext checks if there is any item left to be iterated.
	HasNext() bool
	// Next retrieves the value from the current item.
	Next() T
}
