package hash_map

import (
	"github.com/99_problems/go/src/data_structures"
	"github.com/99_problems/go/src/data_structures/list"
	"math"
)

type Map interface {
	Put(key data_structures.Any, value data_structures.Any)
	Get(key data_structures.Any) data_structures.Any
	Remove(key data_structures.Any) data_structures.Any
	Size() int
	Capacity() int
	IsEmpty() bool
	BucketDistribution() list.List
	Ratio() float64
}

const mapInitialCapacity = 10
const lowerThreshold = 0.3
const upperThreshold = 0.7

type MapEntry struct {
	Key   data_structures.Any
	Value data_structures.Any
}

func CopyMapEntry(entry MapEntry) *MapEntry {
	return &MapEntry{Key: entry.Key, Value: entry.Value}
}

type HashMap struct {
	buckets []*list.LinkedList
	size    int
}

func New() *HashMap {
	buckets := initBuckets(mapInitialCapacity)
	return &HashMap{buckets, 0}
}

func initBuckets(size int) []*list.LinkedList {
	buckets := make([]*list.LinkedList, size)
	for i, _ := range buckets {
		buckets[i] = list.NewLinkedList()
	}
	return buckets
}

func (m HashMap) Get(key data_structures.Any) data_structures.Any {
	bucket := m.getBucket(key)
	if bucket.IsEmpty() {
		return nil
	}

	iterator := bucket.Iterator()

	for iterator.HasNext() {
		entry := iterator.Next().(*MapEntry)
		if entry.Key == key {
			return entry.Value
		}
	}

	return nil
}

func (m *HashMap) Remove(key data_structures.Any) data_structures.Any {
	bucket := m.getBucket(key)
	if bucket.IsEmpty() {
		return nil
	}

	foundEntry := bucket.FindNode(func(value data_structures.Any) bool {
		if value.(*MapEntry).Key == key {
			return true
		}
		return false
	})

	if foundEntry == nil {
		return nil
	}

	value := foundEntry.Value
	bucket.RemoveNode(foundEntry)
	m.size -= 1

	m.ensureTresholdMet()

	return value
}

func (m *HashMap) Put(key data_structures.Any, value data_structures.Any) {
	bucket := m.getBucket(key)
	bucket.Add(&MapEntry{key, value})
	m.size += 1
	m.ensureTresholdMet()
}

func (m *HashMap) ensureTresholdMet() {
	ratio := m.Ratio()

	if ratio > upperThreshold {
		m.increaseSizeAndRehash()
		return
	}

	if m.Size() > mapInitialCapacity && ratio < lowerThreshold {
		m.decreaseSizeAndRehash()
	}
}

func (m *HashMap) Ratio() float64 {
	return float64(m.Size()) / float64(m.Capacity())
}

func (m *HashMap) increaseSizeAndRehash() {
	newBuckets := initBuckets(m.Capacity() * 2)
	rehash(m.buckets, newBuckets)
	m.buckets = newBuckets
}

func (m *HashMap) decreaseSizeAndRehash() {
	newSize := m.Capacity() / 2
	if newSize < mapInitialCapacity {
		newSize = mapInitialCapacity
	}

	newBuckets := initBuckets(newSize)
	rehash(m.buckets, newBuckets)
	m.buckets = newBuckets
}

func rehash(oldBuckets []*list.LinkedList, newBuckets []*list.LinkedList) {
	for _, bucket := range oldBuckets {
		iterator := bucket.Iterator()
		for iterator.HasNext() {
			next := iterator.Next()
			entry := next.(*MapEntry)
			newIndex := getBucketIndex(entry.Key, len(newBuckets))
			newBuckets[newIndex].Add(CopyMapEntry(*entry))
		}
	}
}

func (m HashMap) Size() int {
	return m.size
}

func (m HashMap) Capacity() int {
	return len(m.buckets)
}

func (m HashMap) IsEmpty() bool {
	return m.Size() == 0
}

func (m HashMap) getBucket(key data_structures.Any) *list.LinkedList {
	bucketIndex := getBucketIndex(key, m.Capacity())
	return m.buckets[bucketIndex]
}

func getBucketIndex(key data_structures.Any, capacity int) int {
	keyHash := data_structures.HashObject(key)
	return int(math.Abs(float64(keyHash % capacity)))
}

func (m HashMap) BucketDistribution() list.List {
	result := list.NewArrayList()
	for _, bucket := range m.buckets {
		result.Add(bucket.Size())
	}
	return result
}
