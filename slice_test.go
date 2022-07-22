package ds_test

import (
	"testing"

	"github.com/maestre3d/ds"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewSortedSlice(t *testing.T) {
	l, c := 10, 15
	s := ds.NewSortedSlice[uint8](l, c)
	assert.Equal(t, l, s.Len())
	assert.Equal(t, l, len(s.Slice()))
	assert.Equal(t, c, s.Cap())
	assert.Equal(t, c, cap(s.Slice()))
}

func TestSortedSlice_Append(t *testing.T) {
	tests := []struct {
		name           string
		l, c           int
		values         []string
		lastSortedItem string
	}{
		{
			name:           "Nil",
			l:              0,
			c:              3,
			values:         nil,
			lastSortedItem: "",
		},
		{
			name:           "Unsorted",
			l:              0,
			c:              3,
			values:         []string{"b", "c", "a"},
			lastSortedItem: "c",
		},
		{
			name:           "Sorted",
			l:              0,
			c:              3,
			values:         []string{"aa", "ab", "ac"},
			lastSortedItem: "ac",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := ds.NewSortedSlice[string](tt.l, tt.c)
			s.Append(tt.values...)
			if len(tt.values) == 0 {
				return
			}
			assert.Equal(t, tt.lastSortedItem, s.Slice()[len(tt.values)-1])
		})
	}
}

func TestSortedSlice_Search(t *testing.T) {
	tests := []struct {
		name      string
		l, c      int
		values    []string
		searchVal string
		expPos    int
	}{
		{
			name:      "Nil",
			l:         0,
			c:         3,
			values:    nil,
			searchVal: "",
			expPos:    -1,
		},
		{
			name:      "Unsorted",
			l:         0,
			c:         4,
			values:    []string{"b", "f", "d", "a"},
			searchVal: "f",
			expPos:    3,
		},
		{
			name:      "Sorted",
			l:         0,
			c:         3,
			values:    []string{"aa", "ab", "ac"},
			searchVal: "ac",
			expPos:    2,
		},
		{
			name:      "Not existent",
			l:         0,
			c:         3,
			values:    []string{"aa", "ab", "ac"},
			searchVal: "abc",
			expPos:    -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := ds.NewSortedSlice[string](tt.l, tt.c)
			s.Append(tt.values...)
			assert.Equal(t, tt.expPos, s.Search(tt.searchVal))
		})
	}
}

func TestSortedSlice_MarshalJSON(t *testing.T) {
	set := ds.NewSortedSlice[string](0, 5)
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
		Set *ds.SortedSlice[string]
	}
	f := Foo{
		Bar: "baz",
		Set: set,
	}

	out, err = ds.DefaultJSONMarshaler(f)
	require.NoError(t, err)
	require.NotNil(t, out)
}

func TestSortedSlice_UnmarshalJSON(t *testing.T) {
	set := ds.NewSortedSlice[string](0, 5)
	set.Append("a")
	set.Append("b")
	set.Append("foo")
	set.Append("bar")
	set.Append("baz")
	out, err := set.MarshalJSON()
	require.NoError(t, err)
	require.NotNil(t, out)

	var setB ds.SortedSlice[string]
	err = setB.UnmarshalJSON([]byte("foo bar"))
	require.Error(t, err)
	err = setB.UnmarshalJSON(out)
	require.NoError(t, err)
}
