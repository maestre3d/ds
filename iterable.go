package ds

type Iterable[T any] interface {
	HasNext() bool
	Next() T
}
