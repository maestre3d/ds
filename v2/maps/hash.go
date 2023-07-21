package maps

import "github.com/maestre3d/ds/v2"

type HashMap[K comparable, V any] map[K]V

var _ Map[int, any] = HashMap[int, any]{}

func (h HashMap[K, V]) NewIterator(_ ds.IterationType) ds.Iterator[V] {
	cpMap := make(map[K]V)
	for k, v := range h {
		cpMap[k] = v
	}

	return HashIterator[K, V]{
		buf: cpMap,
	}
}

func (h HashMap[K, V]) Get(key K) V {
	return h[key]
}

func (h HashMap[K, V]) Contains(key K) bool {
	_, ok := h[key]
	return ok
}

func (h HashMap[K, V]) Put(key K, val V) {
	h[key] = val
}

func (h HashMap[K, V]) PutAll(keys []K, values []V) {
	if len(keys) != len(values) {
		return
	}

	for i, key := range keys {
		h.Put(key, values[i])
	}
}

func (h HashMap[K, V]) Remove(key K) {
	delete(h, key)
}

func (h HashMap[K, V]) RemoveAll() {
	for k := range h {
		delete(h, k)
	}
}
