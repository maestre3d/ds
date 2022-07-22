package ds

// Sortable primitive built-in type which can be compared using greater or less than (>, <, >=, <=) operators.
type Sortable interface {
	string | int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | uintptr | float32 | float64
}
