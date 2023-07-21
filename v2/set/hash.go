package set

import "github.com/maestre3d/ds/v2"

type HashSet[K comparable] map[K]struct{}

var _ Set[int] = HashSet[int]{}

func (h HashSet[K]) NewIterator(_ ds.IterationType) ds.Iterator[K] {
	cpMap := make(map[K]struct{})
	for k := range h {
		cpMap[k] = struct{}{}
	}
	return HashIterator[K]{
		buf: cpMap,
	}
}

func (h HashSet[K]) Contains(key K) bool {
	_, ok := h[key]
	return ok
}

func (h HashSet[K]) Add(key K) {
	h[key] = struct{}{}
}

func (h HashSet[K]) AddAll(keys ...K) {
	for _, key := range keys {
		h.Add(key)
	}
}

func (h HashSet[K]) Remove(key K) {
	delete(h, key)
}

func (h HashSet[K]) RemoveAll() {
	for k := range h {
		delete(h, k)
	}
}

func (h HashSet[K]) ToSlice() []K {
	if len(h) == 0 {
		return nil
	}

	buf := make([]K, 0, len(h))
	for k := range h {
		buf = append(buf, k)
	}
	return buf
}
