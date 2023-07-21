package set

import "github.com/maestre3d/ds/v2"

type HashIterator[K comparable] struct {
	buf       map[K]struct{}
	iterCount int
}

var _ ds.Iterator[int] = HashIterator[int]{}

func (h HashIterator[K]) HasNext() bool {
	return h.iterCount < len(h.buf)
}

func (h HashIterator[K]) Next() K {
	for k := range h.buf {
		delete(h.buf, k)
		h.iterCount++
		return k
	}
	var zeroVal K
	return zeroVal
}
