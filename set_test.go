package ds_test

import (
	"testing"

	"github.com/maestre3d/ds"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewHashSet(t *testing.T) {
	l := 5
	set := make(ds.HashSet[int], l)
	assert.Equal(t, 0, len(set))
	assert.Equal(t, 0, len(set.Slice()))
	assert.Equal(t, 0, cap(set.Slice()))
	assert.Nil(t, set.Slice())

	set.Append(8)
	set.Append(10)
	set.Append(5)
	set.Append(12)
	assert.Equal(t, 4, len(set))
	assert.Equal(t, 4, len(set.Slice()))
	assert.Equal(t, 4, cap(set.Slice()))
}

func TestHashSet_Exists(t *testing.T) {
	tests := []struct {
		name      string
		cap       int
		values    []string
		existsVal string
		exp       bool
	}{
		{
			name: "Empty",
		},
		{
			name:      "Exists",
			cap:       4,
			values:    []string{"10", "3", "4", "11"},
			existsVal: "11",
			exp:       true,
		},
		{
			name:      "Not exists",
			cap:       4,
			values:    []string{"10", "3", "4", "11"},
			existsVal: "12",
			exp:       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			set := ds.HashSet[string]{}
			set.Append(tt.values...)
			assert.Equal(t, tt.exp, set.Exists(tt.existsVal))
		})
	}
}

func TestNewSliceSet(t *testing.T) {
	l, c := 5, 10
	set := ds.NewSliceSet[int](l, c)
	assert.Equal(t, l, set.Len())
	assert.Equal(t, l, len(set.Slice()))
	assert.Equal(t, 0, len(set.Map()))
	assert.Equal(t, c, set.Cap())
	assert.Equal(t, c, cap(set.Slice()))
	set.Append(8)
	set.Append(10)
	set.Append(5)
	set.Append(12)
	assert.Equal(t, l, len(set.Map()))
}

func TestSliceSet_Exists(t *testing.T) {
	tests := []struct {
		name      string
		len, cap  int
		values    []int
		existsVal int
		exp       bool
	}{
		{
			name: "Empty",
		},
		{
			name:      "Exists",
			len:       0,
			cap:       4,
			values:    []int{10, 3, 4, 11},
			existsVal: 11,
			exp:       true,
		},
		{
			name:      "Not exists",
			len:       0,
			cap:       4,
			values:    []int{10, 3, 4, 11},
			existsVal: 12,
			exp:       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			set := ds.NewSliceSet[int](tt.len, tt.cap)
			set.Append(tt.values...)
			assert.Equal(t, tt.exp, set.Exists(tt.existsVal))
		})
	}
}

func TestSliceSet_MarshalJSON(t *testing.T) {
	set := ds.NewSliceSet[string](0, 5)
	set.Append("a")
	set.Append("b")
	set.Append("foo")
	set.Append("bar")
	set.Append("baz")
	out, err := set.MarshalJSON()
	require.NoError(t, err)
	require.NotNil(t, out)

	type Foo struct {
		Bar string
		Set *ds.SliceSet[string]
	}
	f := Foo{
		Bar: "baz",
		Set: set,
	}

	out, err = ds.DefaultJSONMarshaler(f)
	require.NoError(t, err)
	require.NotNil(t, out)
}

func TestSliceSet_UnmarshalJSON(t *testing.T) {
	set := ds.NewSliceSet[string](0, 5)
	set.Append("a")
	set.Append("b")
	set.Append("foo")
	set.Append("bar")
	set.Append("baz")
	out, err := set.MarshalJSON()
	require.NoError(t, err)
	require.NotNil(t, out)

	var setB ds.SliceSet[string]
	err = setB.UnmarshalJSON([]byte("foo bar"))
	require.Error(t, err)
	err = setB.UnmarshalJSON(out)
	require.NoError(t, err)
}
