package ds

type Set[K comparable] interface {
	Iterable[K]
	Contains(key K) bool
	Add(key K)
	Remove(key K)
}
