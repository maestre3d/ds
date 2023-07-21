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

func (h HashMap[K, V]) ToSlice() []KeyValueItem[K, V] {
	if len(h) == 0 {
		return nil
	}

	buf := make([]KeyValueItem[K, V], 0, len(h))
	for k, v := range h {
		buf = append(buf, KeyValueItem[K, V]{
			Key:   k,
			Value: v,
		})
	}
	return buf
}

func (h HashMap[K, V]) ToSliceKeys() []K {
	if len(h) == 0 {
		return nil
	}

	buf := make([]K, 0, len(h))
	for k := range h {
		buf = append(buf, k)
	}
	return buf
}

func (h HashMap[K, V]) ToSliceValues() []V {
	if len(h) == 0 {
		return nil
	}

	buf := make([]V, 0, len(h))
	for _, v := range h {
		buf = append(buf, v)
	}
	return buf
}
