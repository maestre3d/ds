package maps

import "github.com/maestre3d/ds/v2"

type HashIterator[K comparable, V any] struct {
	buf       map[K]V
	iterCount int
}

var _ ds.Iterator[any] = HashIterator[int, any]{}

func (h HashIterator[K, V]) HasNext() bool {
	return h.iterCount < len(h.buf)
}

func (h HashIterator[K, V]) Next() V {
	for k, v := range h.buf {
		delete(h.buf, k)
		h.iterCount++
		return v
	}
	var zeroVal V
	return zeroVal
}
