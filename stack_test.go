package ds_test

import (
	"fmt"
	"testing"

	"github.com/maestre3d/ds"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewSliceStack(t *testing.T) {
	type testTable struct {
		name   string
		inCap  int
		inData []string
	}
	test := []testTable{
		{
			name:   "empty",
			inCap:  0,
			inData: nil,
		},
		{
			name:  "with single data",
			inCap: 1,
			inData: []string{
				"foo",
			},
		},
		{
			name:  "with data",
			inCap: 5,
			inData: []string{
				"foo",
				"bar",
				"baz",
				"foobar",
				"bazinga",
			},
		},
	}

	testFunc := func(t *testing.T, stack *ds.SliceStack[string], tt testTable) {
		assert.Equal(t, 0, stack.Len())
		assert.Equal(t, tt.inCap, stack.Cap())
		assert.Equal(t, "", stack.Peek())
		assert.Equal(t, "", stack.Pop())

		for _, d := range tt.inData {
			stack.Push(d)
		}

		assert.Equal(t, len(tt.inData), stack.Len())
		if len(tt.inData) == 0 {
			return
		}

		assert.Equal(t, tt.inData[len(tt.inData)-1], stack.Peek())
		assert.Equal(t, len(tt.inData), stack.Len())
		for i := len(tt.inData) - 1; i >= 0; i-- {
			assert.Equal(t, tt.inData[i], stack.Pop())
		}
		assert.Equal(t, 0, stack.Len())
	}

	for _, tt := range test {
		t.Run(fmt.Sprintf("stack_fixed %s", tt.name), func(t *testing.T) {
			stack := ds.NewSliceStackFixed[string](tt.inCap)
			testFunc(t, stack, tt)
		})
		t.Run(fmt.Sprintf("stack_dynamic %s", tt.name), func(t *testing.T) {
			stack := ds.NewSliceStack[string]()
			tt.inCap = 0
			testFunc(t, stack, tt)
		})
	}
}

func TestNewLinkedList(t *testing.T) {
	test := []struct {
		name   string
		inData []string
	}{
		{
			name:   "empty",
			inData: nil,
		},
		{
			name: "with single data",
			inData: []string{
				"foo",
			},
		},
		{
			name: "with data",
			inData: []string{
				"foo",
				"bar",
				"baz",
				"foobar",
				"bazinga",
			},
		},
	}

	for _, tt := range test {
		t.Run(fmt.Sprintf(tt.name), func(t *testing.T) {
			stack := ds.NewLinkedStack[string]()
			assert.Equal(t, 0, stack.Len())
			assert.Equal(t, "", stack.Peek())
			assert.Equal(t, "", stack.Pop())

			for _, d := range tt.inData {
				stack.Push(d)
			}

			assert.Equal(t, len(tt.inData), stack.Len())
			if len(tt.inData) == 0 {
				return
			}

			assert.Equal(t, tt.inData[len(tt.inData)-1], stack.Peek())
			assert.Equal(t, len(tt.inData), stack.Len())
			for i := len(tt.inData) - 1; i >= 0; i-- {
				assert.Equal(t, tt.inData[i], stack.Pop())
			}
			assert.Equal(t, 0, stack.Len())
		})
	}
}

func TestSliceStack_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		inItems []any
		exp     []byte
		err     error
	}{
		{
			name:    "nil",
			inItems: nil,
			exp:     []byte("[]"),
		},
		{
			name:    "empty",
			inItems: []any{},
			exp:     []byte("[]"),
		},
		{
			name:    "with single data int",
			inItems: []any{1},
			exp:     []byte("[1]"),
		},
		{
			name:    "with data int",
			inItems: []any{1, 2, 3, 10, 58},
			exp:     []byte("[1,2,3,10,58]"),
		},
		{
			name:    "with data string",
			inItems: []any{"foo", "bar", "baz"},
			exp:     []byte("[\"foo\",\"bar\",\"baz\"]"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := ds.NewSliceStack[any]()
			for _, d := range tt.inItems {
				stack.Push(d)
			}

			out, err := stack.MarshalJSON()
			assert.Equal(t, tt.err, err)
			if err != nil {
				return
			}
			assert.EqualValues(t, tt.exp, out)
			require.NoError(t, stack.UnmarshalJSON(out))
		})
	}
}
