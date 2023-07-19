package linkedlist_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/maestre3d/ds/v2"
	"github.com/maestre3d/ds/v2/linkedlist"
)

func TestNewList(t *testing.T) {
	list := linkedlist.NewList[string]()
	list.AppendAll("0", "1", "2")
	list.InsertAt(0, "-1")
	list.InsertAt(3, "1.5")
	list.Append("4")
	list.InsertAt(list.Len()-1, "3.5")
	list.ReplaceAt(list.Len()-1, "5")
	list.ReplaceAt(list.Len()-2, "4")

	iterator := list.NewIterator(ds.BackwardIteration)
	i := 0
	exp := []string{"5", "4", "2", "1.5", "1", "0", "-1"}
	for iterator.HasNext() {
		assert.Equal(t, exp[i], iterator.Next())
		i++
	}

	iterator = list.NewIterator(ds.ForwardIteration)
	i = 0
	exp = []string{"-1", "0", "1", "1.5", "2", "4", "5"}
	for iterator.HasNext() {
		assert.Equal(t, exp[i], iterator.Next())
		i++
	}

	assert.Equal(t, "-1", list.GetHead().Value)
	assert.Equal(t, "5", list.GetTail().Value)
}
