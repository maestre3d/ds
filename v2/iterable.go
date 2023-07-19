package ds

// IterationType kind of iteration used by Iterator to traverse data structures.
type IterationType uint8

const (
	_                 IterationType = iota
	ForwardIteration                // Iterates from first item to last.
	BackwardIteration               // Iterates from last item to first.
)

// An Iterable represents a data structure which can be iterated by an Iterator instance.
type Iterable[T any] interface {
	// NewIterator allocates a new Iterator instance. Accepts IterationType to set traversing mechanisms (ForwardIteration,
	// BackwardIteration).
	NewIterator(iterationType IterationType) Iterator[T]
}
