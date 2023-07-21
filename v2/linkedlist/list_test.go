package linkedlist_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/maestre3d/ds/v2"
	"github.com/maestre3d/ds/v2/linkedlist"
)

func TestNewList_Mut(t *testing.T) {
	tests := []struct {
		name    string
		in      []string
		mutFunc func(list *linkedlist.List[string])
		exp     []string
	}{
		{
			name: "mutability",
			in:   []string{"0", "1", "2"},
			mutFunc: func(list *linkedlist.List[string]) {
				list.InsertAt(0, "-1")
				list.InsertAt(3, "1.5")
				list.Append("4")
				list.InsertAt(list.Len()-1, "3.5")
				list.ReplaceAt(list.Len()-1, "5")
				list.ReplaceAt(list.Len()-2, "4")
			},
			exp: []string{"-1", "0", "1", "1.5", "2", "4", "5"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := linkedlist.NewList[string]()
			list.AppendAll(tt.in...)

			tt.mutFunc(list)

			bufOut := make([]string, 0, list.Len())
			iter := list.NewIterator(ds.ForwardIteration)
			for iter.HasNext() {
				bufOut = append(bufOut, iter.Next())
			}
			assert.Equal(t, tt.exp, bufOut)
		})
	}
}

func TestList_Remove(t *testing.T) {
	list := linkedlist.NewList[int]()
	list.AppendAll(1, 2, 3, 4)
	list.Remove()
	assert.Equal(t, 3, list.GetTail().Value)
	assert.Equal(t, 3, list.GetAt(list.Len()-1))
	assert.Equal(t, 3, list.Len())

	list.RemoveAt(0)
	assert.Equal(t, 2, list.GetHead().Value)
	assert.Equal(t, 2, list.GetAt(0))
	assert.Equal(t, 2, list.Len())
}
