package set_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/maestre3d/ds/v2/set"
)

func TestHashSet(t *testing.T) {
	s := set.HashSet[int]{}
	s.AddAll(0, 1, 2, 3)
	assert.True(t, s.Contains(3))
	assert.Len(t, s.ToSlice(), 4)
	s.Remove(3)
	assert.False(t, s.Contains(3))
	s.RemoveAll()
	assert.Equal(t, 0, len(s))
	assert.Len(t, s.ToSlice(), 0)
}

func TestHashSet_NewIterator(t *testing.T) {
	s := set.HashSet[int]{}
	s.AddAll(0, 1, 2, 3, 4, 5)

	expSet := set.HashSet[int]{}
	iter := s.NewIterator(0)
	for iter.HasNext() {
		expSet.Add(iter.Next())
	}

	for k := range s {
		assert.True(t, expSet.Contains(k))
	}
	assert.Len(t, s, 6)
}
