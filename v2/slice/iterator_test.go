package slice_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/maestre3d/ds/v2"
	"github.com/maestre3d/ds/v2/slice"
)

func TestIterator(t *testing.T) {
	in := []int{0, 1, 2, 3, 4, 5, 6}
	iter := slice.NewIterator(in, ds.ForwardIteration)
	i := 0
	for iter.HasNext() {
		assert.Equal(t, in[i], iter.Next())
		i++
	}

	iter = slice.NewIterator(in, ds.BackwardIteration)
	i = len(in) - 1
	for iter.HasNext() {
		assert.Equal(t, in[i], iter.Next())
		i--
	}
}
