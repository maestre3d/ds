package set

import "github.com/maestre3d/ds/v2"

type Set[K comparable] interface {
	ds.Iterable[K]
	ds.Slicer[K]
	Contains(key K) bool
	Add(key K)
	AddAll(keys ...K)
	Remove(key K)
	RemoveAll()
}
