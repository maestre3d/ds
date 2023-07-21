package slice_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/maestre3d/ds/v2/slice"
)

func TestList_Append(t *testing.T) {
	list := slice.NewList[int](0)
	list.Append(0)
	list.Append(1)
	list.Append(2)
	assert.Equal(t, 3, list.Len())

	list.AppendAll(3, 4, 5)
	assert.Equal(t, 6, list.Len())
}

func TestList_InsertAt(t *testing.T) {
	list := slice.NewList[int](0)
	list.AppendAll(-1, -2, -3)
	assert.Equal(t, -1, list.GetAt(0))
	assert.Equal(t, 3, list.Len())

	list.InsertAt(0, 1)
	assert.Equal(t, 1, list.GetAt(0))
	assert.Equal(t, 4, list.Len())

	list.InsertAt(1, 2)
	assert.Equal(t, 2, list.GetAt(1))
	assert.Equal(t, 5, list.Len())

	list.InsertAt(2, 3)
	assert.Equal(t, 3, list.GetAt(2))
	assert.Equal(t, 6, list.Len())
}

func TestList_ReplaceAt(t *testing.T) {
	list := slice.NewList[int](3)
	list.AppendAll(-1, -2, -3)
	assert.Equal(t, -1, list.GetAt(0))

	list.ReplaceAt(0, 1)
	assert.Equal(t, 1, list.GetAt(0))

	list.ReplaceAt(1, 2)
	assert.Equal(t, 2, list.GetAt(1))

	list.ReplaceAt(2, 3)
	assert.Equal(t, 3, list.GetAt(2))
	assert.Equal(t, 3, list.Cap())
}

func TestList_Remove(t *testing.T) {
	list := slice.NewList[int](0)
	list.AppendAll(1, 2, 3)

	list.Remove()
	assert.Equal(t, 2, list.GetAt(list.Len()-1))

	list.RemoveAt(0)
	assert.Equal(t, 2, list.GetAt(0))

	list.Remove()
	assert.Equal(t, 0, list.Len())

	// test nil values
	list.Remove()
	list.RemoveAt(1)

	// test out of bounds
	list.AppendAll(1)
	list.RemoveAt(3)
}

func TestList_RemoveAll(t *testing.T) {
	list := slice.NewList[int](5)
	list.AppendAll(1, 2, 3, 4, 5)
	assert.Equal(t, 5, list.Len())
	assert.Equal(t, 5, list.Cap())
	list.RemoveAll()
	assert.Equal(t, 0, list.Len())
	assert.Equal(t, 5, list.Cap())
}
