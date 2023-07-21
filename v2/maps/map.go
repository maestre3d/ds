package maps

import "github.com/maestre3d/ds/v2"

type KeyValueItem[K comparable, V any] struct {
	Key   K
	Value V
}

type Map[K comparable, V any] interface {
	ds.Iterable[V]
	ds.Slicer[KeyValueItem[K, V]]
	Put(key K, val V)
	PutAll(keys []K, values []V)
	Remove(key K)
	RemoveAll()
	Get(key K) V
	Contains(key K) bool
	ToSliceKeys() []K
	ToSliceValues() []V
}
