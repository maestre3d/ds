package linkedlist_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/maestre3d/ds/v2"
	"github.com/maestre3d/ds/v2/linkedlist"
)

func TestIterator_Next(t *testing.T) {
	tests := []struct {
		name     string
		in       []int
		iterType ds.IterationType
		exp      []int
	}{
		{
			name:     "nil",
			in:       nil,
			iterType: 0,
			exp:      []int{},
		},
		{
			name:     "empty",
			in:       []int{},
			iterType: 0,
			exp:      []int{},
		},
		{
			name:     "single",
			in:       []int{0},
			iterType: 0,
			exp:      []int{0},
		},
		{
			name:     "multi",
			in:       []int{0, 1, 2, 3},
			iterType: 0,
			exp:      []int{0, 1, 2, 3},
		},
		{
			name:     "single forward",
			in:       []int{0},
			iterType: ds.ForwardIteration,
			exp:      []int{0},
		},
		{
			name:     "multi forward",
			in:       []int{0, 1, 2, 3},
			iterType: ds.ForwardIteration,
			exp:      []int{0, 1, 2, 3},
		},
		{
			name:     "single backwards",
			in:       []int{0},
			iterType: ds.BackwardIteration,
			exp:      []int{0},
		},
		{
			name:     "multi backwards",
			in:       []int{0, 1, 2, 3},
			iterType: ds.BackwardIteration,
			exp:      []int{3, 2, 1, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := linkedlist.NewList[int]()
			list.AppendAll(tt.in...)

			iter := list.NewIterator(tt.iterType)

			bufOut := make([]int, 0, list.Len())
			for iter.HasNext() {
				bufOut = append(bufOut, iter.Next())
			}

			assert.Equal(t, tt.exp, bufOut)
		})
	}
}
