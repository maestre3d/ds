package maps

import "github.com/maestre3d/ds/v2"

type Map[K comparable, V any] interface {
	ds.Iterable[V]
	Put(key K, val V)
	PutAll(keys []K, values []V)
	Remove(key K)
	RemoveAll()
	Get(key K) V
	Contains(key K) bool
}
