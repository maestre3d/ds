package ds_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/maestre3d/ds/v2"
)

func TestSliceIterator(t *testing.T) {
	in := []int{0, 1, 2, 3, 4, 5, 6}
	iter := ds.NewSliceIterator(in, ds.ForwardIteration)
	i := 0
	for iter.HasNext() {
		assert.Equal(t, in[i], iter.Next())
		i++
	}

	iter = ds.NewSliceIterator(in, ds.BackwardIteration)
	i = len(in) - 1
	for iter.HasNext() {
		assert.Equal(t, in[i], iter.Next())
		i--
	}
}
