package ds_test

import (
	"github.com/maestre3d/ds"
	"testing"
)

func TestNewSliceQueue(t *testing.T) {
	q := ds.NewSliceQueue[int](10)

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)

	a, b := q.Dequeue(), q.Dequeue()
	res := a * b
	t.Logf("%d * %d = %d", a, b, res)

	t.Log(q.Dequeue())
	t.Log(q.Dequeue())

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
