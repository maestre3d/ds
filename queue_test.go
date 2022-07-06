package ds_test

import (
	"github.com/maestre3d/ds"
	"testing"
)

func TestNewSliceQueue(t *testing.T) {
	q := ds.NewSliceQueue[int](10)

	q.Push(1)
	q.Push(2)
	q.Push(3)
	q.Push(4)

	a, b := q.Pop(), q.Pop()
	res := a * b
	t.Logf("%d * %d = %d", a, b, res)

	t.Log(q.Pop())
	t.Log(q.Pop())

	tests := []struct {
		name string
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

		})
	}
}
