package ds

// Slicer turns a data structure into a Go slice.
type Slicer[T any] interface {
	// ToSlice retrieves items as a Go slice.
	ToSlice() []T
}
