package ds

type Map[K, V any] interface {
	Iterable[V]
	Put(key K, val V)
	Remove(key K)
	Get(key K) V
	Contains(key K) bool
}
