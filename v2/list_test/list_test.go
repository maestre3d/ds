package list_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/maestre3d/ds/v2"
	"github.com/maestre3d/ds/v2/linkedlist"
	"github.com/maestre3d/ds/v2/slice"
)

func TestNewList(t *testing.T) {
	tests := []struct {
		name string
		list ds.List[int]
	}{
		{
			name: "slice",
			list: slice.NewList[int](0),
		},
		{
			name: "linked list",
			list: linkedlist.NewList[int](),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, 0, tt.list.Len())
			tt.list.AppendAll(1, 2, 3)
			assert.Equal(t, 3, tt.list.Len())
			assert.Equal(t, 3, tt.list.GetAt(tt.list.Len()-1))

			tt.list.Append(4)
			assert.Equal(t, 4, tt.list.Len())

			tt.list.Remove()
			assert.Equal(t, 3, tt.list.Len())
			assert.Equal(t, 3, tt.list.GetAt(tt.list.Len()-1))

			tt.list.RemoveAt(1)
			assert.Equal(t, 2, tt.list.Len())
			assert.Equal(t, 3, tt.list.GetAt(1))

			tt.list.ReplaceAt(0, -1)
			assert.Equal(t, 2, tt.list.Len())
			assert.Equal(t, -1, tt.list.GetAt(0))

			tt.list.InsertAt(0, -2)
			assert.Equal(t, 3, tt.list.Len())
			assert.Equal(t, -1, tt.list.GetAt(1))

			tt.list.RemoveAll()
			assert.Equal(t, 0, tt.list.Len())

			// test nil
			tt.list.Remove()
			assert.Zero(t, tt.list.GetAt(1))
			tt.list.InsertAt(1, 1) // out of bounds
			assert.Equal(t, 0, tt.list.Len())
			tt.list.InsertAt(0, 1) // first item always inserts
			assert.Equal(t, 1, tt.list.Len())
		})
	}
}
