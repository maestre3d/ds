package ds

type List[T any] interface {
	Iterable[T]
	Len() int
	Append(v T)
	Remove()
	InsertAt(i int, v T)
	GetAt(i int) T
	RemoveAt(i int)
}
