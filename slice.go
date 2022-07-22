package ds

import "encoding/json"

// SortedSlice a kind of slice which is always sorted.
// The sort mechanism happens at appends operations using binary search to avoid O(n log n) time
// complexity of a sorting algorithm.
type SortedSlice[T Sortable] struct {
	buf []T
}

var _ json.Marshaler = SortedSlice[string]{}

var _ json.Unmarshaler = &SortedSlice[string]{}

// NewSortedSlice allocates a SortedSlice with l as length and c as capacity for the internal buffer.
func NewSortedSlice[T Sortable](l, c int) *SortedSlice[T] {
	return &SortedSlice[T]{
		buf: make([]T, l, c),
	}
}

// MarshalJSON encodes underlying slice as JSON data using DefaultJSONMarshaler.
func (s SortedSlice[T]) MarshalJSON() ([]byte, error) {
	return DefaultJSONMarshaler(s.buf)
}

// UnmarshalJSON decodes underlying slice from JSON data  using DefaultJSONUnmarshaler.
func (s *SortedSlice[T]) UnmarshalJSON(bytes []byte) error {
	if err := DefaultJSONUnmarshaler(bytes, &s.buf); err != nil {
		return err
	}
	return nil
}

func (s SortedSlice[T]) binarySearch(v T) (int, int, int) {
	low := 0
	upper := len(s.buf) - 1

	for low <= upper {
		pivot := (upper + low) / 2
		if v == s.buf[pivot] {
			return pivot, low, upper
		}

		if v > s.buf[pivot] {
			low = pivot + 1
		} else {
			upper = pivot - 1
		}
	}
	return -1, low, upper
}

func (s SortedSlice[T]) Search(v T) int {
	pos, _, _ := s.binarySearch(v)
	return pos
}

func (s *SortedSlice[T]) append(v T) {
	if len(s.buf) == 0 {
		s.buf = append(s.buf, v)
		return
	}

	pivot := 0
	pos, low, _ := s.binarySearch(v)
	if pos != -1 {
		pivot = pos
	} else {
		pivot = low
	}

	s.buf = append(s.buf[:pivot], append([]T{v}, s.buf[pivot:]...)...)
}

func (s *SortedSlice[T]) Append(v ...T) {
	for _, item := range v {
		s.append(item)
	}
}

func (s SortedSlice[T]) Slice() []T {
	return s.buf
}

func (s SortedSlice[T]) Len() int {
	return len(s.buf)
}

func (s SortedSlice[T]) Cap() int {
	return cap(s.buf)
}
