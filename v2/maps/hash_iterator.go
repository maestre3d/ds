package maps

import "github.com/maestre3d/ds/v2"

type HashIterator[K comparable, V any] struct {
	buf       map[K]V
	iterCount int
}

var _ ds.Iterator[KeyValueItem[int, any]] = HashIterator[int, any]{}

func (h HashIterator[K, V]) HasNext() bool {
	return h.iterCount < len(h.buf)
}

func (h HashIterator[K, V]) Next() KeyValueItem[K, V] {
	for k, v := range h.buf {
		delete(h.buf, k)
		h.iterCount++
		return KeyValueItem[K, V]{
			Key:   k,
			Value: v,
		}
	}

	return KeyValueItem[K, V]{}
}
