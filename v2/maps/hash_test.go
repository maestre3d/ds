package maps_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/maestre3d/ds/v2/maps"
	"github.com/maestre3d/ds/v2/set"
)

func TestHashMap(t *testing.T) {
	s := maps.HashMap[int, string]{}
	s.PutAll([]int{0, 1, 2, 3}, []string{"0", "1", "2", "3"})
	assert.True(t, s.Contains(3))
	s.Remove(3)
	assert.False(t, s.Contains(3))
	s.RemoveAll()
	assert.Equal(t, 0, len(s))
}

func TestHashMap_NewIterator(t *testing.T) {
	s := maps.HashMap[int, string]{}
	s.PutAll([]int{0, 1, 2, 3, 4, 5}, []string{"0", "1", "2", "3", "4", "5"})

	expSet := set.HashSet[string]{}
	iter := s.NewIterator(0)
	for iter.HasNext() {
		expSet.Add(iter.Next())
	}

	for _, v := range s {
		assert.True(t, expSet.Contains(v))
	}
	assert.Len(t, s, 6)
}
