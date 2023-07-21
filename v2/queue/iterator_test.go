package queue_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/maestre3d/ds/v2/queue"
)

func TestIterator(t *testing.T) {
	tests := []struct {
		name string
		q    queue.Queue[string]
		in   []string
		exp  []string
	}{
		{
			name: "fifo slice",
			q:    queue.NewSliceQueue[string](queue.FIFO, 0),
			in:   []string{"0", "1", "2"},
			exp:  []string{"0", "1", "2"},
		},
		{
			name: "lifo slice",
			q:    queue.NewSliceQueue[string](queue.LIFO, 0),
			in:   []string{"0", "1", "2"},
			exp:  []string{"2", "1", "0"},
		},
		{
			name: "fifo linked list",
			q:    queue.NewLinkedListQueue[string](queue.FIFO),
			in:   []string{"0", "1", "2"},
			exp:  []string{"0", "1", "2"},
		},
		{
			name: "lifo linked list",
			q:    queue.NewLinkedListQueue[string](queue.LIFO),
			in:   []string{"0", "1", "2"},
			exp:  []string{"2", "1", "0"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.q.EnqueueAll(tt.in...)
			iter := tt.q.NewIterator(0)
			i := 0
			for iter.HasNext() {
				assert.Equal(t, tt.exp[i], iter.Next())
				i++
			}
		})
	}
}
