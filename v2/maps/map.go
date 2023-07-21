package maps

import "github.com/maestre3d/ds/v2"

type KeyValueItem[K comparable, V any] struct {
	Key   K
	Value V
}

func (v KeyValueItem[K, V]) IsEmpty() bool {
	var keyZeroVal K
	return v.Key == keyZeroVal
}

type Map[K comparable, V any] interface {
	ds.Iterable[KeyValueItem[K, V]]
	ds.Slicer[KeyValueItem[K, V]]
	Put(key K, val V)
	PutAll(keys []K, values []V)
	Remove(key K)
	RemoveAll()
	Get(key K) V
	Contains(key K) bool
	ToSliceKeys() []K
	ToSliceValues() []V
	Len() int
}
