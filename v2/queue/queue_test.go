package queue_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/maestre3d/ds/v2/linkedlist"
	"github.com/maestre3d/ds/v2/queue"
	"github.com/maestre3d/ds/v2/slice"
)

func TestNewSliceQueue(t *testing.T) {
	tests := []struct {
		name      string
		q         queue.Queue[string]
		queueType queue.Type
		in        []string
		exp       []string
	}{
		{
			name:      "unknown slice",
			q:         queue.NewSliceQueue[string](0, 0),
			queueType: 0,
			in:        []string{"0", "1", "2", "3"},
			exp:       []string{"0", "1", "2", "3"},
		},
		{
			name:      "fifo slice",
			q:         queue.NewSliceQueue[string](queue.FIFO, 0),
			queueType: queue.FIFO,
			in:        []string{"0", "1", "2", "3"},
			exp:       []string{"0", "1", "2", "3"},
		},
		{
			name:      "lifo slice",
			q:         queue.NewSliceQueue[string](queue.LIFO, 0),
			queueType: queue.LIFO,
			in:        []string{"0", "1", "2", "3"},
			exp:       []string{"3", "2", "1", "0"},
		},
		{
			name:      "unknown list",
			q:         queue.NewListQueue[string](queue.FIFO, slice.NewList[string](0)),
			queueType: 0,
			in:        []string{"0", "1", "2", "3"},
			exp:       []string{"0", "1", "2", "3"},
		},
		{
			name:      "fifo list",
			q:         queue.NewListQueue[string](queue.FIFO, slice.NewList[string](0)),
			queueType: queue.FIFO,
			in:        []string{"0", "1", "2", "3"},
			exp:       []string{"0", "1", "2", "3"},
		},
		{
			name:      "lifo list",
			q:         queue.NewListQueue[string](queue.LIFO, slice.NewList[string](0)),
			queueType: queue.LIFO,
			in:        []string{"0", "1", "2", "3"},
			exp:       []string{"3", "2", "1", "0"},
		},
		{
			name:      "unknown list linked list",
			q:         queue.NewListQueue[string](queue.FIFO, linkedlist.NewList[string]()),
			queueType: 0,
			in:        []string{"0", "1", "2", "3"},
			exp:       []string{"0", "1", "2", "3"},
		},
		{
			name:      "fifo list linked list",
			q:         queue.NewListQueue[string](queue.FIFO, linkedlist.NewList[string]()),
			queueType: queue.FIFO,
			in:        []string{"0", "1", "2", "3"},
			exp:       []string{"0", "1", "2", "3"},
		},
		{
			name:      "lifo list linked list",
			q:         queue.NewListQueue[string](queue.LIFO, linkedlist.NewList[string]()),
			queueType: queue.LIFO,
			in:        []string{"0", "1", "2", "3"},
			exp:       []string{"3", "2", "1", "0"},
		},
		{
			name:      "unknown linked list",
			q:         queue.NewLinkedListQueue[string](0),
			queueType: 0,
			in:        []string{"0", "1", "2", "3"},
			exp:       []string{"0", "1", "2", "3"},
		},
		{
			name:      "fifo linked list",
			q:         queue.NewLinkedListQueue[string](queue.FIFO),
			queueType: queue.FIFO,
			in:        []string{"0", "1", "2", "3"},
			exp:       []string{"0", "1", "2", "3"},
		},
		{
			name:      "lifo linked list",
			q:         queue.NewLinkedListQueue[string](queue.LIFO),
			queueType: queue.LIFO,
			in:        []string{"0", "1", "2", "3"},
			exp:       []string{"3", "2", "1", "0"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, ok := tt.q.Peek()
			assert.False(t, ok)
			_, ok = tt.q.Dequeue()
			assert.False(t, ok)

			for _, item := range tt.in {
				tt.q.Enqueue(item)
			}

			for _, item := range tt.exp {
				itemQueue, okQueue := tt.q.Peek()
				assert.True(t, okQueue)
				assert.Equal(t, item, itemQueue)

				itemQueue, okQueue = tt.q.Dequeue()
				assert.True(t, okQueue)
				assert.Equal(t, item, itemQueue)
			}

			// validate buffer out-of-bounds error
			item, ok := tt.q.Dequeue()
			assert.False(t, ok)
			assert.Zero(t, item)
		})
	}
}

func TestSliceQueue_Clear(t *testing.T) {
	tests := []struct {
		name string
		q    queue.Queue[int]
		in   []int
	}{
		{
			name: "nil slice",
			q:    queue.NewSliceQueue[int](queue.FIFO, 0),
			in:   nil,
		},
		{
			name: "empty slice",
			q:    queue.NewSliceQueue[int](queue.FIFO, 0),
			in:   []int{},
		},
		{
			name: "populated slice",
			q:    queue.NewSliceQueue[int](queue.FIFO, 0),
			in:   []int{0, 1, 2},
		},
		{
			name: "nil list",
			q:    queue.NewListQueue[int](queue.FIFO, slice.NewList[int](0)),
			in:   nil,
		},
		{
			name: "empty list",
			q:    queue.NewListQueue[int](queue.FIFO, slice.NewList[int](0)),
			in:   []int{},
		},
		{
			name: "populated list",
			q:    queue.NewListQueue[int](queue.FIFO, slice.NewList[int](0)),
			in:   []int{0, 1, 2},
		},
		{
			name: "nil list linked list",
			q:    queue.NewListQueue[int](queue.FIFO, linkedlist.NewList[int]()),
			in:   nil,
		},
		{
			name: "empty list linked list",
			q:    queue.NewListQueue[int](queue.FIFO, linkedlist.NewList[int]()),
			in:   []int{},
		},
		{
			name: "populated list linked list",
			q:    queue.NewListQueue[int](queue.FIFO, linkedlist.NewList[int]()),
			in:   []int{0, 1, 2},
		},
		{
			name: "nil linked list",
			q:    queue.NewLinkedListQueue[int](queue.FIFO),
			in:   nil,
		},
		{
			name: "empty linked list",
			q:    queue.NewLinkedListQueue[int](queue.FIFO),
			in:   []int{},
		},
		{
			name: "populated linked list",
			q:    queue.NewLinkedListQueue[int](queue.FIFO),
			in:   []int{0, 1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.q.EnqueueAll(tt.in...)
			assert.Equal(t, len(tt.in), tt.q.Len())

			tt.q.Clear()
			assert.Equal(t, 0, tt.q.Len())
		})
	}
}

func BenchmarkNewSliceQueue(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		q := queue.NewSliceQueue[string](queue.LIFO, 4)
		q.Enqueue("some element 0")
		q.Enqueue("some element 1")
		q.Enqueue("some element 2")
		q.Enqueue("some element 3")
		q.Dequeue()
		q.Dequeue()
		q.Dequeue()
		q.Dequeue()
		q.Dequeue()
		q.Clear()
	}
}
