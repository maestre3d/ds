package ds

import "encoding/json"

// Set logical data structure which holds non-repeated values.
type Set[T comparable] interface {
	// Append adds a value if not exists.
	Append(v ...T)
	// Exists checks if value exists.
	Exists(v T) bool
	// Slice retrieves a slice of T.
	Slice() []T
	// Map retrieves a map of T.
	Map() map[T]struct{}
}

// HashSet implementation of Set using a map as underlying physical data structure.
type HashSet[T comparable] map[T]struct{}

var _ Set[string] = HashSet[string]{}

// Append adds a value if not exists.
func (s HashSet[T]) Append(v ...T) {
	for _, item := range v {
		if !s.Exists(item) {
			s[item] = struct{}{}
		}
	}
}

// Exists checks if value exists.
func (s HashSet[T]) Exists(v T) bool {
	_, ok := s[v]
	return ok
}

// Map retrieves the underlying map.
func (s HashSet[T]) Map() map[T]struct{} {
	return s
}

// Slice retrieves a slice of T formed from the underlying map.
func (s HashSet[T]) Slice() []T {
	if len(s) == 0 {
		return nil
	}

	sl := make([]T, 0, len(s))
	for k := range s {
		sl = append(sl, k)
	}
	return sl
}

// SliceSet implementation of Set using a slice as underlying physical data structure.
type SliceSet[T comparable] struct {
	buf        []T
	totalItems uint
}

var _ Set[string] = &SliceSet[string]{}

var _ json.Marshaler = SliceSet[string]{}

var _ json.Unmarshaler = &SliceSet[string]{}

// NewSliceSet allocates a SliceSet using l as length and c as capacity for the underlying buffer.
func NewSliceSet[T comparable](l, c int) *SliceSet[T] {
	return &SliceSet[T]{
		buf: make([]T, l, c),
	}
}

// MarshalJSON encodes underlying slice as JSON data using DefaultJSONMarshaler.
func (s SliceSet[T]) MarshalJSON() ([]byte, error) {
	return DefaultJSONMarshaler(s.buf)
}

// UnmarshalJSON decodes underlying slice from JSON data  using DefaultJSONUnmarshaler.
func (s *SliceSet[T]) UnmarshalJSON(bytes []byte) error {
	if err := DefaultJSONUnmarshaler(bytes, &s.buf); err != nil {
		return err
	}
	return nil
}

// Exists checks if value exists.
func (s SliceSet[T]) Exists(v T) bool {
	for _, item := range s.buf {
		if item == v {
			return true
		}
	}
	return false
}

// Append adds a value if not exists.
func (s *SliceSet[T]) Append(v ...T) {
	// Time Complex: O(mn) where n = len(s.buf) & m = len(v)
	for _, item := range v {
		if !s.Exists(item) {
			s.buf = append(s.buf, item)
			s.totalItems++
		}
	}
}

// Slice retrieves the underlying slice.
func (s SliceSet[T]) Slice() []T {
	return s.buf
}

// Len retrieves the length of the underlying slice.
func (s SliceSet[T]) Len() int {
	return len(s.buf)
}

// Cap retrieves the capacity of the underlying slice.
func (s SliceSet[T]) Cap() int {
	return cap(s.buf)
}

// Map retrieves a map of T formed from the underlying slice.
func (s SliceSet[T]) Map() map[T]struct{} {
	if s.totalItems == 0 {
		return nil
	}
	m := make(map[T]struct{}, len(s.buf))
	for _, item := range s.buf {
		m[item] = struct{}{}
	}
	return m
}
