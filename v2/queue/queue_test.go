package queue_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/maestre3d/ds/v2/queue"
)

func TestNewSliceQueue(t *testing.T) {
	tests := []struct {
		name      string
		queueType queue.Type
		initCap   int
		exp       []string
	}{
		{
			name:      "unknown",
			queueType: 0,
			initCap:   0,
			exp:       []string{"0", "1", "2", "3"},
		},
		{
			name:      "fifo",
			queueType: queue.FIFO,
			initCap:   0,
			exp:       []string{"0", "1", "2", "3"},
		},
		{
			name:      "lifo",
			queueType: queue.LIFO,
			initCap:   0,
			exp:       []string{"0", "1", "2", "3"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var q queue.Queue[string] = queue.NewSliceQueue[string](tt.queueType, tt.initCap)
			_, ok := q.Peek()
			assert.False(t, ok)
			_, ok = q.Dequeue()
			assert.False(t, ok)

			for _, item := range tt.exp {
				q.Enqueue(item)
			}

			if tt.queueType == queue.FIFO || tt.queueType == 0 {
				for _, item := range tt.exp {
					itemQueue, okQueue := q.Peek()
					assert.True(t, okQueue)
					assert.Equal(t, item, itemQueue)

					itemQueue, okQueue = q.Dequeue()
					assert.True(t, okQueue)
					assert.Equal(t, item, itemQueue)
				}
			} else if tt.queueType == queue.LIFO {
				for i := len(tt.exp) - 1; i >= 0; i-- {
					itemQueue, okQueue := q.Peek()
					assert.True(t, okQueue)
					assert.Equal(t, tt.exp[i], itemQueue)

					itemQueue, okQueue = q.Dequeue()
					assert.True(t, okQueue)
					assert.Equal(t, tt.exp[i], itemQueue)
				}
			}

			// validate buffer out-of-bounds error
			item, ok := q.Dequeue()
			assert.False(t, ok)
			assert.Zero(t, item)
		})
	}
}

func TestSliceQueue_Clear(t *testing.T) {
	q := queue.NewSliceQueue[string](queue.FIFO, 10)
	q.EnqueueAll("0", "1", "2")
	assert.Equal(t, 3, q.Len())
	assert.Equal(t, 10, q.Cap())

	q.Clear()
	assert.Equal(t, 0, q.Len())
	assert.Equal(t, 10, q.Cap())
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
