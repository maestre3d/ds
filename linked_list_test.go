package ds_test

import (
	"github.com/maestre3d/ds"
	"testing"
)

func TestLinkedList_Append(t *testing.T) {
	l := ds.NewLinkedList[string]()
	l.Append("foo")
	l.Append("bar")
	l.Append("baz")
	t.Log(l.GetAt(2))

	l.AppendTail("foobar")
	t.Log(l.GetAt(2))
}
