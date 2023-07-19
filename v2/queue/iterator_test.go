package queue_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/maestre3d/ds/v2/queue"
)

func TestIterator(t *testing.T) {
	tests := []struct {
		name   string
		in     []string
		inType queue.Type
		exp    []string
	}{
		{
			name:   "fifo",
			in:     []string{"0", "1", "2"},
			inType: queue.FIFO,
			exp:    []string{"0", "1", "2"},
		},
		{
			name:   "lifo",
			in:     []string{"0", "1", "2"},
			inType: queue.LIFO,
			exp:    []string{"2", "1", "0"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := queue.NewSliceQueue[string](tt.inType, 0)
			q.EnqueueAll(tt.in...)
			iter := q.NewIterator(0)
			i := 0
			for iter.HasNext() {
				assert.Equal(t, tt.exp[i], iter.Next())
				i++
			}
		})
	}
}
