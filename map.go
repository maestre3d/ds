package ds

type HashBag[T comparable] map[T]T

func (h HashBag[T]) Set(k, v T) {
	h[k] = v
	h[v] = k
}

func (h HashBag[T]) Sets(v ...T) {
	if len(v)%2 != 0 {
		return
	}

	for i := 0; i < len(v); i++ {
		h.Set(v[i], v[i+1])
		i++
	}
}

func (h HashBag[T]) Get(k T) T {
	return h[k]
}
