package ds_test

import (
	"testing"

	"github.com/maestre3d/ds"
)

func TestNewSliceStack(t *testing.T) {
	var queue ds.SerializableQueue[string]
	queue = ds.NewSliceStack[string](5)
	queue.Push("foo")
	queue.Push("bar")
	jsonData, _ := queue.MarshalJSON()
	t.Log(string(jsonData))
	t.Log(queue.Len())
	t.Log(queue.Peek())
	t.Log(queue.Pop())
	t.Log(queue.Pop())
	t.Log(queue.Len())
	t.Log(queue.Pop())

	queue = ds.NewSliceQueue[string](5)
	queue.Push("foo")
	queue.Push("bar")

	type Foo struct {
		Bar      string                       `json:"bar"`
		TodoList ds.SerializableQueue[string] `json:"todo_list"`
	}
	f := Foo{
		Bar:      "the quick brown fox",
		TodoList: queue,
	}

	jsonData, _ = ds.DefaultJSONMarshaler(f)
	t.Log(string(jsonData))
}
