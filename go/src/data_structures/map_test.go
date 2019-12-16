package data_structures_test

import (
	"github.com/99_problems/go/src/data_structures"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestHashMap(t *testing.T) {
	hashMap := data_structures.NewHashMap()
	hashMap.Put("key", 123)
	assert.Equal(t, 123, hashMap.Get("key"))
	assert.Nil(t, hashMap.Get("not_exists"))
	assert.Equal(t, 1, hashMap.Size())
	assert.False(t, hashMap.IsEmpty())
}

func TestHashMapDistribution(t *testing.T) {
	hashMap := data_structures.NewHashMap()
	put100(hashMap)
	distribution := hashMap.BucketDistribution()

	var emptyBucket int
	for i := 0; i < distribution.Size(); i++ {
		if distribution.Get(i) == 0 {
			emptyBucket += 1
		}
	}

	expectedEmptyBuckets := float64(hashMap.Capacity()) * (1 - hashMap.Ratio()) * 2

	assert.True(t, float64(emptyBucket) < expectedEmptyBuckets)
}

func TestHashMapRemove(t *testing.T) {
	hashMap := data_structures.NewHashMap()
	hashMap.Put("key", 123)
	assert.Equal(t, 123, hashMap.Get("key"))
	assert.Nil(t, hashMap.Get("not_exists"))
	assert.Equal(t, 1, hashMap.Size())
	assert.False(t, hashMap.IsEmpty())
}

func TestHashMapRemoveMany(t *testing.T) {
	hashMap := data_structures.NewHashMap()
	put100(hashMap)
	remove100(t, hashMap)
	assert.Equal(t, 0, hashMap.Size())
}

func TestHashMapCapacity(t *testing.T) {
	hashMap := data_structures.NewHashMap()
	assert.Equal(t, 10, hashMap.Capacity())

	put100(hashMap)

	assert.Equal(t, 100, hashMap.Size())
	assert.NotEqual(t, 10, hashMap.Capacity())

	remove100(t, hashMap)

	assert.Equal(t, 0, hashMap.Size())
	assert.Equal(t, 20, hashMap.Capacity())
}

func remove100(t *testing.T, hashMap *data_structures.HashMap) {
	for i := 0; i < 100; i++ {
		assert.NotNil(t, hashMap.Remove("key "+strconv.Itoa(i)))
	}
}

func put100(hashMap *data_structures.HashMap) {
	for i := 0; i < 100; i++ {
		hashMap.Put("key "+strconv.Itoa(i), i)
	}
}
